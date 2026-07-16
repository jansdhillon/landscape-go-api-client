// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jansdhillon/landscape-go-api-client/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	baseURL := os.Getenv("LANDSCAPE_BASE_URL")
	if baseURL == "" {
		log.Fatalf("base URL not set")
	}
	ak := os.Getenv("LANDSCAPE_ACCESS_KEY")
	if ak == "" {
		log.Fatalf("access key not set")
	}

	sk := os.Getenv("LANDSCAPE_SECRET_KEY")
	if sk == "" {
		log.Fatalf("secret key not set")
	}

	landscapeAPIClient, err := client.NewLandscapeAPIClient(
		ctx,
		http.DefaultClient,
		baseURL,
		client.NewAccessKeyProvider(ak, sk),
	)

	if err != nil {
		log.Fatalf("failed to create Landscape API client: %v", err)
	}

	// Create a V1 script
	rawCode := "#!/bin/bash\n \"hello\" > /home/ubuntu/hello.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(rawCode))
	createdScriptRes, err := client.LegacyAPIRequestWithResponse[client.V1Script](ctx, landscapeAPIClient, "CreateScript", map[string]any{
		"title":       rand.Text(),
		"code":        enc,
		"script_type": "V1",
	})

	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	log.Printf("raw create script response: %s", createdScriptRes.Body)
	if createdScriptRes.JSON == nil {
		log.Fatalf("error creating script: %d", createdScriptRes.StatusCode())
	}

	out, _ := json.MarshalIndent(createdScriptRes.JSON, "", "  ")
	log.Printf("created script:\n%s", out)

}
