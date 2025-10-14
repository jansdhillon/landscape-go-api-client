package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime/types"

	landscape "github.com/jansdhillon/landscape-go-client/client"
)

func main() {
	baseURL := "https://landscape-server-quickstart.com"

	email := os.Getenv("LANDSCAPE_EMAIL")
	password := os.Getenv("LANDSCAPE_PASSWORD")
	accessKey := os.Getenv("LANDSCAPE_ACCESS_KEY")
	secretKey := os.Getenv("LANDSCAPE_SECRET_KEY")
	if (email == "" || password == "") && (accessKey == "" || secretKey == "") {
		log.Fatal("set either LANDSCAPE_EMAIL and LANDSCAPE_PASSWORD or LANDSCAPE_ACCESS_KEY and LANDSCAPE_SECRET_KEY")
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
	}

	ctx := context.Background()

	token, err := login(ctx, httpClient, baseURL, email, password, accessKey, secretKey)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}

	query := url.Values{}
	query.Set("version", "2011-08-01")
	query.Set("action", "CreateScript")
	query.Set("title", "Example2")
	query.Set("code", "IyEvYmluL2Jhc2gKZWNobyAiSGVsbG8gZnJvbSBsYW5kc2NhcGUiCg==")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/api/?"+query.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response body: %v", err)
	}
	log.Printf("status: %d body: %s", resp.StatusCode, string(bodyBytes))

	var decoded interface{}
	if err := json.Unmarshal(bodyBytes, &decoded); err != nil {
		log.Printf("body (non-json): %s", string(bodyBytes))
		return
	}
	pretty, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		log.Printf("body: %s", string(bodyBytes))
		return
	}
	log.Printf("pretty body:\n%s", pretty)
}

func login(ctx context.Context, httpClient *http.Client, baseURL, email, password, accessKey, secretKey string) (string, error) {
	authClient, err := landscape.NewClientWithResponses(baseURL, landscape.WithHTTPClient(httpClient))
	if err != nil {
		return "", fmt.Errorf("init auth client: %w", err)
	}

	if accessKey != "" && secretKey != "" {
		resp, err := authClient.LoginWithAccessKeyWithResponse(ctx, landscape.AccessKeyLoginRequest{
			AccessKey: accessKey,
			SecretKey: secretKey,
		})
		if err != nil {
			return "", fmt.Errorf("access-key login request failed: %w", err)
		}
		if resp.JSON200 == nil {
			return "", fmt.Errorf("access-key login unexpected response: status %d body %s", resp.StatusCode(), strings.TrimSpace(string(resp.Body)))
		}
		return resp.JSON200.Token, nil
	}

	resp, err := authClient.LoginWithPasswordWithResponse(ctx, landscape.LoginRequest{
		Email:    types.Email(email),
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("password login request failed: %w", err)
	}
	if resp.JSON200 == nil {
		return "", fmt.Errorf("password login unexpected response: status %d body %s", resp.StatusCode(), strings.TrimSpace(string(resp.Body)))
	}
	return resp.JSON200.Token, nil
}
