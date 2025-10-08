package landscape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LandscapeAPIClient struct {
	AccountName string
	RootURL     string
	AccessKey   string
	SecretKey   string
	HTTP        *http.Client
}

func NewLandscapeAPIClient(accountName, rootURL, accessKey, secretKey string) *LandscapeAPIClient {
	return &LandscapeAPIClient{
		AccountName: accountName,
		RootURL:     rootURL,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		HTTP:        &http.Client{},
	}
}

func (c *LandscapeAPIClient) DoRequest(method, relativeURL string, body any, queryArgs map[string]any) (*http.Response, error) {
	baseURL, err := url.Parse(c.RootURL)
	if err != nil {
		return nil, err
	}
	relative, err := url.Parse(relativeURL)
	if err != nil {
		return nil, err
	}
	fullURL := baseURL.ResolveReference(relative)

	if len(queryArgs) > 0 {
		values := fullURL.Query()
		for key, v := range queryArgs {
			switch t := v.(type) {
			case []string:
				values[key] = append([]string(nil), t...)
			case []any:
				for _, item := range t {
					values.Add(key, fmt.Sprint(item))
				}
			default:
				values.Set(key, fmt.Sprint(t))
			}
		}
		fullURL.RawQuery = values.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, fullURL.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type LoginResponse struct {
	Token          string           `json:"token"`
	Accounts       []map[string]any `json:"accounts"`
	CurrentAccount string           `json:"current_account"`
	Email          string           `json:"email"`
	Name           string           `json:"name"`
}

func (c *LandscapeAPIClient) Login() (*LoginResponse, error) {
	body := map[string]any{
		"access_key": c.AccessKey,
		"secret_key": c.SecretKey,
	}

	res, err := c.DoRequest(http.MethodPost, "login/access-key", body, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var loginRes LoginResponse
	if err := json.Unmarshal(bodyBytes, &loginRes); err != nil {
		return nil, err
	}

	return &loginRes, nil
}
