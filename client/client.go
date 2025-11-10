// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

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

// NewLandscapeAPIClient creates a new Landscape API client configured with authentication
// provided by the given LoginProvider. The provider is used to obtain a JWT token which
// is then applied to subsequent requests as a Bearer token.
func NewLandscapeAPIClient(baseURL string, loginProvider LoginProvider) (*ClientWithResponses, error) {
	tempClient, err := NewClientWithResponses(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp client: %w", err)
	}

	token, err := loginProvider.Login(context.Background(), tempClient)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	authEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}

	return NewClientWithResponses(baseURL, WithRequestEditorFn(authEditor))
}

// LegacyActionParams is a helper to call legacy API
// actions by creating a pointer to a InvokeLegacyActionParams with
// the hardcoded version, as well as the provided action.
func LegacyActionParams(action string) *InvokeLegacyActionParams {
	return &InvokeLegacyActionParams{
		Action:  action,
		Version: "2011-08-01",
	}
}

// EncodeQueryRequestEditor returns a RequestEditorFn that
// adds the given url.Values as query arguments in the request
// URL.
func EncodeQueryRequestEditor(values url.Values) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
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
