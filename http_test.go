package landscape

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type CreateComputerRequest struct {
	Title     string `json:"title"`
	AccountID int    `json:"account_id"`
}

func TestDoRequest(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		query      map[string]any
		wantQuery  url.Values
		body       any
		wantBody   any
		statusCode int
	}{
		{
			name:       "post request with typed body succeeds",
			method:     http.MethodPost,
			path:       "/computers",
			body:       CreateComputerRequest{Title: "hello", AccountID: 1},
			wantBody:   CreateComputerRequest{Title: "hello", AccountID: 1},
			statusCode: http.StatusCreated,
			wantQuery:  url.Values{},
		},
		{
			name:   "get request includes query args",
			method: http.MethodGet,
			path:   "/computers",
			query: map[string]any{
				"limit": 10,
			},
			wantQuery:  url.Values{"limit": []string{"10"}},
			statusCode: http.StatusAccepted,
		},
		{
			name:   "post request with map body and query args",
			method: http.MethodPost,
			path:   "/computers",
			query: map[string]any{
				"all_computers": true,
			},
			body:       map[string]any{"title": "hello", "account_id": 1},
			wantBody:   map[string]any{"title": "hello", "account_id": 1},
			wantQuery:  url.Values{"all_computers": []string{"true"}},
			statusCode: http.StatusCreated,
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

				if !reflect.DeepEqual(r.URL.Query(), tc.wantQuery) {
					t.Fatalf("unexpected query: want %v, got %v", tc.wantQuery, r.URL.Query())
				}

				if tc.wantBody != nil {
					data, err := io.ReadAll(r.Body)
					if err != nil {
						t.Fatalf("failed to read body: %v", err)
					}
					defer r.Body.Close()

					switch want := tc.wantBody.(type) {
					case CreateComputerRequest:
						var got CreateComputerRequest
						if err := json.Unmarshal(data, &got); err != nil {
							t.Fatalf("failed to unmarshal body: %v", err)
						}
						if !reflect.DeepEqual(got, want) {
							t.Fatalf("unexpected body: want %+v, got %+v", want, got)
						}

					case map[string]any:
						var got map[string]any
						if err := json.Unmarshal(data, &got); err != nil {
							t.Fatalf("failed to unmarshal body: %v", err)
						}
						for k, v := range want {
							if iv, ok := v.(int); ok {
								want[k] = float64(iv)
							}
						}
						if !reflect.DeepEqual(got, want) {
							t.Fatalf("unexpected body: want %+v, got %+v", want, got)
						}

					default:
						t.Fatalf("unhandled body type %T", tc.wantBody)
					}
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

			resp, err := client.DoRequest(tc.method, tc.path, tc.body, tc.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if resp.StatusCode != tc.statusCode {
				t.Fatalf("expected status %d, got %d", tc.statusCode, resp.StatusCode)
			}
		})
	}
}
