// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

	// Create a V2 script
	rawCode := "#!/bin/bash\n \"hello\" > /home/ubuntu/hello.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(rawCode))
	createdScriptRes, err := client.LegacyAPIRequestWithResponse[client.V1Script](ctx, landscapeAPIClient, "CreateScript", map[string]any{
		"title":       rand.Text(),
		"code":        enc,
		"script_type": "V2",
	})
	if err != nil {
		log.Fatalf("failed to create script: %v", err)
	}

	log.Printf("raw create script response: %s", createdScriptRes.Body)
	if createdScriptRes.JSON == nil {
		log.Fatalf("error creating script: %d", createdScriptRes.StatusCode())
	}

	createdScript := createdScriptRes.JSON

	raw := "#!/bin/bash\necho \"newcode\" > /home/ubuntu/myscript.txt"
	enc = base64.StdEncoding.EncodeToString([]byte(raw))

	res, err := client.LegacyAPIRequestWithResponse[client.V2Script](ctx, landscapeAPIClient, "EditScript", map[string]any{
		"script_id": createdScript.Id,
		"username":  "jim",
		"code":      enc,
	})
	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	log.Printf("raw edit script response: %s", res.Body)
	if res.JSON == nil {
		log.Fatalf("failed to edit script: %d", res.StatusCode())
	}

	editedScript := res.JSON

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
