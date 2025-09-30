package landscape

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDoRequest(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		query      map[string]any
		wantQuery  url.Values
		statusCode int
	}{
		{
			name:       "post request succeeds",
			method:     http.MethodPost,
			path:       "/api",
			statusCode: http.StatusOK,
			wantQuery:  url.Values{},
		},
		{
			name:   "get request includes query args",
			method: http.MethodGet,
			path:   "/resources",
			query: map[string]any{
				"status":  "active",
				"limit":   10,
				"verbose": true,
			},
			wantQuery: url.Values{
				"status":  []string{"active"},
				"limit":   []string{"10"},
				"verbose": []string{"true"},
			},
			statusCode: http.StatusAccepted,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method {
					t.Fatalf("expected method %s, got %s", tc.method, r.Method)
				}
				if r.URL.Path != tc.path {
					t.Fatalf("expected path %s, got %s", tc.path, r.URL.Path)
				}

				wantEncoded := ""
				if tc.wantQuery != nil {
					wantEncoded = tc.wantQuery.Encode()
				}
				if got := r.URL.Query().Encode(); got != wantEncoded {
					t.Fatalf("unexpected query string: want %q, got %q", wantEncoded, got)
				}

				w.WriteHeader(tc.statusCode)
			}))
			defer server.Close()

			client := &LandscapeAPIClient{
				AccountName: "test",
				RootURL:     server.URL,
				AccessKey:   "key",
				SecretKey:   "secret",
				HTTP:        server.Client(),
			}

			resp, err := client.DoRequest(tc.method, tc.path, nil, tc.query, nil, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := resp.(*http.Response)
			if got.StatusCode != tc.statusCode {
				t.Fatalf("expected status %d, got %d", tc.statusCode, got.StatusCode)
			}
		})
	}
}
