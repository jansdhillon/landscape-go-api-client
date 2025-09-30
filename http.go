package landscape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *LandscapeAPIClient) DoRequest(method, relativeURL string, body map[string]any, queryArgs, urlArgs map[string]any, target any) (any, error) {
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
		for key, raw := range queryArgs {
			switch v := raw.(type) {
			case []string:
				values.Del(key)
				for _, s := range v {
					values.Add(key, s)
				}
			case []any:
				values.Del(key)
				for _, item := range v {
					values.Add(key, fmt.Sprint(item))
				}
			default:
				values.Set(key, fmt.Sprint(v))
			}
		}
		fullURL.RawQuery = values.Encode()
	}

	req, err := http.NewRequest(method, fullURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if target != nil {
		if err := json.Unmarshal(bodyBytes, target); err != nil {
			return nil, err
		}
		return target, nil
	} else {
		return nil, fmt.Errorf("failed to unmarshal response into given struct %T", target)
	}
}
