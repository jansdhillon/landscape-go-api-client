// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
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
		baseURL,
		client.NewAccessKeyProvider(ak, sk),
	)

	if err != nil {
		log.Fatalf("failed to create Landscape API client: %v", err)
	}

	// Create a V1 script
	rawCode := "#!/bin/bash\n \"hello\" > /home/ubuntu/hello.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(rawCode))
	scriptType := "V1"
	createParams := &client.CreateScriptParams{
		Version:    "2011-08-01",
		Action:     "CreateScript",
		Title:      rand.Text(),
		Code:       enc,
		ScriptType: &scriptType,
	}
	createdScriptRes, err := landscapeAPIClient.CreateScriptWithResponse(ctx, createParams)

	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	log.Printf("raw create script response: %s", createdScriptRes.Body)
	if createdScriptRes.JSON200 == nil {
		log.Fatalf("error creating script: %s", createdScriptRes.Status())
	}

	script, err := client.ParseLegacyResponse[client.V1Script](createdScriptRes.Body)
	if err != nil {
		log.Fatalf("failed to parse script response: %v", err)
	}
	out, _ := json.MarshalIndent(script, "", "  ")
	log.Printf("created script:\n%s", out)

}
