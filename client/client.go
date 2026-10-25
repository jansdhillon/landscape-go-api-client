// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// legacyAPIVersion is the fixed "version" query parameter required by every
// legacy (action-based) Landscape API request.
const legacyAPIVersion = "2011-08-01"

// ParseLegacyResponse decodes the raw body bytes from a legacy API action
// response into T. Use this instead of the JSON200 field on legacy action
// responses, which is an untyped map[string]any that cannot carry
// typed methods.
func ParseLegacyResponse[T any](body []byte) (T, error) {
	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("failed to parse legacy API response: %w", err)
	}
	return result, nil
}

// LegacyResponse is the parsed result of a legacy API action request, pairing
// the raw HTTP response with its typed body. It mirrors the *WithResponse
// return types generated for the OpenAPI-described (REST v2) endpoints, so
// that legacy and generated calls feel consistent to callers.
type LegacyResponse[T any] struct {
	HTTPResponse *http.Response
	Body         []byte
	JSON         *T
}

// StatusCode returns the HTTP status code of the underlying response, or 0 if
// there is no response.
func (r *LegacyResponse[T]) StatusCode() int {
	if r == nil || r.HTTPResponse == nil {
		return 0
	}
	return r.HTTPResponse.StatusCode
}

// LoginProvider is an interface that knows how to obtain a JWT token
// given a pre-configured API client (ClientWithResponses). Implementations
// can call the appropriate login endpoints (email/password or access key).
type LoginProvider interface {
	Login(ctx context.Context, client *ClientWithResponses) (string, error)
}

// EmailPasswordProvider logs in with an email/password pair and optionally
// specifies an account name.
type EmailPasswordProvider struct {
	Email    string
	Password string
	Account  *string
}

func NewEmailPasswordProvider(email, password string, account *string) *EmailPasswordProvider {
	return &EmailPasswordProvider{
		Email:    email,
		Password: password,
		Account:  account,
	}
}

// Login implements LoginProvider for EmailPasswordProvider.
func (p *EmailPasswordProvider) Login(ctx context.Context, c *ClientWithResponses) (string, error) {
	resp, err := c.LoginWithPasswordWithResponse(ctx, LoginWithPasswordJSONRequestBody{
		Email:    openapi_types.Email(p.Email),
		Password: p.Password,
	})
	if err != nil {
		return "", fmt.Errorf("login with password request failed: %w", err)
	}
	if resp == nil {
		return "", fmt.Errorf("nil response from login")
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("login failed with status: %d", resp.StatusCode())
	}

	return resp.JSON200.Token, nil
}

// AccessKeyProvider logs in with an access key/secret key pair.
type AccessKeyProvider struct {
	AccessKey string
	SecretKey string
}

func NewAccessKeyProvider(accessKey, secretKey string) *AccessKeyProvider {
	return &AccessKeyProvider{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

// Login implements LoginProvider for AccessKeyProvider.
func (p *AccessKeyProvider) Login(ctx context.Context, c *ClientWithResponses) (string, error) {
	resp, err := c.LoginWithAccessKeyWithResponse(ctx, LoginWithAccessKeyJSONRequestBody{
		AccessKey: p.AccessKey,
		SecretKey: p.SecretKey,
	})
	if err != nil {
		return "", fmt.Errorf("login with access key request failed: %w", err)
	}
	if resp == nil {
		return "", fmt.Errorf("nil response from login")
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("login failed with status: %d", resp.StatusCode())
	}

	return resp.JSON200.Token, nil
}

// LandscapeAPIClient is a Landscape API client that combines two things
// under one authenticated roof:
//
//   - The OpenAPI-generated *ClientWithResponses (embedded), which covers the
//     REST v2 endpoints described by the Landscape OpenAPI spec. Its methods
//     (e.g. GetScriptWithResponse) are promoted onto LandscapeAPIClient.
//   - LegacyAPIRequest, a manual helper for the legacy, action-based API
//     (?action=...&version=2011-08-01), which is not (and will not be)
//     described by the OpenAPI spec.
type LandscapeAPIClient struct {
	// Embedding promotes all REST v2 (OpenAPI-generated) methods.
	*ClientWithResponses

	// jwt is the bearer token obtained from the LoginProvider.
	jwt string

	// httpClient is used for legacy (non-generated) requests.
	httpClient HttpRequestDoer

	// serverURL is the parsed base URL of the Landscape server, used to build
	// legacy API request URLs.
	serverURL *url.URL
}

// NewLandscapeAPIClient creates a new Landscape API client configured with authentication
// provided by the given LoginProvider. The provider is used to obtain a JWT token which
// is then applied to subsequent requests as a Bearer token.
//
// opts are passed directly to the underlying generated client. Use
// WithHTTPClient to supply a custom *http.Client, ex. one configured with
// a custom TLS cert pool for self-signed server certificates.
func NewLandscapeAPIClient(ctx context.Context, hc HttpRequestDoer, baseURL string, loginProvider LoginProvider, opts ...ClientOption) (*LandscapeAPIClient, error) {
	// The login (temp) client must use the caller's HTTP client too, so that
	// options like a custom TLS cert pool apply to the login request as well
	// as to the authenticated client built below.
	loginOpts := make([]ClientOption, 0, len(opts)+1)
	loginOpts = append(loginOpts, opts...)
	loginOpts = append(loginOpts, WithHTTPClient(hc))

	tempClient, err := NewClientWithResponses(baseURL, loginOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp client: %w", err)
	}

	token, err := loginProvider.Login(ctx, tempClient)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	authEditor := func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}

	// Build a fresh option slice for the authenticated client instead of
	// mutating the caller's opts, so login requests aren't affected by
	// options meant only for the authenticated client, and so callers can
	// safely reuse the opts slice they passed in.
	authedOpts := make([]ClientOption, 0, len(opts)+2)
	authedOpts = append(authedOpts, opts...)
	authedOpts = append(authedOpts, WithRequestEditorFn(authEditor), WithHTTPClient(hc))

	genClient, err := NewClientWithResponses(baseURL, authedOpts...)
	if err != nil {
		return nil, fmt.Errorf("cannot create generated API client: %w", err)
	}

	parsed, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse base URL %q: %w", baseURL, err)
	}

	return &LandscapeAPIClient{
		ClientWithResponses: genClient,
		jwt:                 token,
		httpClient:          hc,
		serverURL:           parsed,
	}, nil
}

// EncodeQueryRequestEditor returns a RequestEditorFn that
// adds the given url.Values as query arguments in the request
// URL.
func EncodeQueryRequestEditor(values url.Values) RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		query := req.URL.Query()

		for k, v := range values {
			query.Del(k)

			for _, arg := range v {
				query.Add(k, arg)
			}
		}

		req.URL.RawQuery = query.Encode()

		return nil
	}
}

// LegacyAPIRequest issues a request against the legacy, action-based
// Landscape API (i.e. GET /api/?action=<action>&version=2011-08-01&...),
// which is not described by the OpenAPI spec. kwargs are encoded as query
// parameters using the same "form" style encoding oapi-codegen uses for
// generated endpoints, so slices, bools, numbers, etc. are all serialized
// consistently with the rest of the client.
//
// The caller is responsible for closing the returned response's Body.
func (c *LandscapeAPIClient) LegacyAPIRequest(ctx context.Context, action string, kwargs map[string]any, reqEditors ...RequestEditorFn) (*http.Response, error) {
	requestURL := url.URL{
		Scheme: c.serverURL.Scheme,
		Host:   c.serverURL.Host,
		Path:   "/api",
	}

	query := url.Values{}
	query.Set("version", legacyAPIVersion)
	query.Set("action", action)

	for k, v := range kwargs {
		if v == nil {
			continue
		}

		encoded, err := runtime.StyleParamWithLocation("form", true, k, runtime.ParamLocationQuery, v)
		if err != nil {
			return nil, fmt.Errorf("cannot encode legacy API parameter %q: %w", k, err)
		}

		parsed, err := url.ParseQuery(encoded)
		if err != nil {
			return nil, fmt.Errorf("cannot parse encoded legacy API parameter %q: %w", k, err)
		}

		for pk, pv := range parsed {
			for _, v2 := range pv {
				query.Add(pk, v2)
			}
		}
	}

	requestURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create new request: %w", err)
	}

	for _, editor := range reqEditors {
		if err := editor(ctx, req); err != nil {
			return nil, fmt.Errorf("cannot apply request editor: %w", err)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot make request: %w", err)
	}

	return resp, nil
}

// LegacyAPIRequestWithResponse calls LegacyAPIRequest and parses the response
// body as T, mirroring the *WithResponse helpers generated for REST v2
// endpoints. It always reads and closes the response body.
func LegacyAPIRequestWithResponse[T any](ctx context.Context, c *LandscapeAPIClient, action string, kwargs map[string]any, reqEditors ...RequestEditorFn) (*LegacyResponse[T], error) {
	httpResp, err := c.LegacyAPIRequest(ctx, action, kwargs, reqEditors...)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read legacy API response body: %w", err)
	}

	result := &LegacyResponse[T]{
		HTTPResponse: httpResp,
		Body:         body,
	}

	if httpResp.StatusCode == http.StatusOK {
		parsed, err := ParseLegacyResponse[T](body)
		if err != nil {
			return result, err
		}
		result.JSON = &parsed
	}

	return result, nil
}
