package landscape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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
