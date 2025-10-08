package landscape

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetScript(t *testing.T) {
	wantScript := ScriptType{
		ID:            42,
		Title:         "Maintenance Script",
		VersionNumber: 3,
		CreatedBy:     User{ID: 1, Name: "Alice"},
		CreatedAt:     "2025-10-07T12:00:00Z",
		LastEditedBy:  User{ID: 2, Name: "Bob"},
		LastEditedAt:  "2025-10-07T14:00:00Z",
		ScriptProfiles: []any{
			map[string]any{"profile": "production"},
		},
		Status: "active",
		Attachments: []Attachment{
			{Filename: "deploy.sh", ID: 12},
		},
		Code:         "#!/bin/bash\necho 'Deploying...'\n",
		Interpreter:  "bash",
		AccessGroup:  "admins",
		TimeLimit:    300,
		Username:     "alice",
		IsRedactable: false,
		IsEditable:   true,
		IsExecutable: true,
	}

	t.Run("successfully fetches and unmarshals a script", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Fatalf("expected GET method, got %s", r.Method)
			}
			expectedPath := "/scripts/42"
			if r.URL.Path != expectedPath {
				t.Fatalf("expected path %s, got %s", expectedPath, r.URL.Path)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(wantScript)
		}))
		defer server.Close()

		client := &LandscapeAPIClient{
			AccountName: "test",
			RootURL:     server.URL + "/",
			AccessKey:   "key",
			SecretKey:   "secret",
			HTTP:        server.Client(),
		}

		got, err := client.GetScript(42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(*got, wantScript) {
			t.Fatalf("unexpected result:\nwant %+v\ngot  %+v", wantScript, *got)
		}
	})

	t.Run("returns error on non-200 response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "not found", http.StatusNotFound)
		}))
		defer server.Close()

		client := &LandscapeAPIClient{
			AccountName: "test",
			RootURL:     server.URL + "/",
			HTTP:        server.Client(),
		}

		server.CloseClientConnections()

		_, err := client.GetScript(42)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("returns error for invalid JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{invalid-json}`)
		}))
		defer server.Close()

		client := &LandscapeAPIClient{
			AccountName: "test",
			RootURL:     server.URL + "/",
			HTTP:        server.Client(),
		}

		_, err := client.GetScript(42)
		if err == nil {
			t.Fatal("expected JSON unmarshal error, got nil")
		}
	})
}
