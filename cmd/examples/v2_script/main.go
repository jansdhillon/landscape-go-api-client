// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

	// Create a V2 script
	rawCode := "#!/bin/bash\n \"hello\" > /home/ubuntu/hello.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(rawCode))
	scriptType := "V2"
	createdScriptRes, err := landscapeAPIClient.LegacyCreateScriptWithResponse(ctx, &client.LegacyCreateScriptParams{
		Title:      rand.Text(),
		Code:       enc,
		ScriptType: &scriptType,
	})
	if err != nil {
		log.Fatalf("failed to create script: %v", err)
	}

	log.Printf("raw create script response: %s", createdScriptRes.Body)
	if createdScriptRes.JSON200 == nil {
		log.Fatalf("error creating script: %s", createdScriptRes.Status())
	}

	createdScript, err := client.ParseLegacyResponse[client.V1Script](createdScriptRes.Body)
	if err != nil {
		log.Fatalf("failed to parse response as script: %v", err)
	}

	raw := "#!/bin/bash\necho \"newcode\" > /home/ubuntu/myscript.txt"
	enc = base64.StdEncoding.EncodeToString([]byte(raw))
	username := "jim"

	res, err := landscapeAPIClient.LegacyEditScriptWithResponse(ctx, &client.LegacyEditScriptParams{
		ScriptId: createdScript.Id,
		Username: &username,
		Code:     &enc,
	})
	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	log.Printf("raw edit script response: %s", res.Body)
	if res.JSON200 == nil {
		log.Fatalf("failed to edit script: %s", res.Status())
	}

	editedScript, err := client.ParseLegacyResponse[client.V2Script](res.Body)
	if err != nil {
		log.Fatalf("failed to parse response as V2 script: %s", err)
	}

	log.Printf("edited script title: %s", editedScript.Title)
	if editedScript.Attachments != nil {
		log.Printf("edited script attachments: %+v", *editedScript.Attachments)
	}
	if editedScript.Attachments != nil {
		log.Printf("edited script attachments count: %d", len(*editedScript.Attachments))
		for i, attachment := range *editedScript.Attachments {
			log.Printf("attachment %d: %+v", i, attachment)
		}
	}
	if editedScript.CreatedBy != nil {
		log.Printf("edited created by id: %d", *editedScript.CreatedBy.Id)
		log.Printf("edited created by name: %s", *editedScript.CreatedBy.Name)
	}

}
