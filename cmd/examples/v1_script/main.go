// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/url"
	"os"
	"strconv"

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
	createParams := client.LegacyActionParams("CreateScript")
	rawCode := "#!/bin/bash\n \"hello\" > /home/ubuntu/hello.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(rawCode))
	queryArgsEditorFn := client.EncodeQueryRequestEditor(url.Values{
		"title":       []string{rand.Text()},
		"code":        []string{enc},
		"script_type": []string{"V1"},
	})
	createdScriptRes, err := landscapeAPIClient.InvokeLegacyActionWithResponse(ctx, createParams, queryArgsEditorFn)

	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	if createdScriptRes == nil {
		log.Fatalf("legacy action returned nil response")
	}

	if createdScriptRes.StatusCode() != 200 {
		log.Fatalf("legacy action failed: status=%d body=%s", createdScriptRes.StatusCode(), string(createdScriptRes.Body))
	}

	log.Printf("raw create script response: %s", createdScriptRes.Body)

	if createdScriptRes.JSON200 == nil {
		log.Fatalf("legacy action did not return a script object: %s", string(createdScriptRes.Body))
	}

	createdScript, err := createdScriptRes.JSON200.AsV2Script()
	if err != nil {
		log.Fatalf("failed to convert returned script into V1 Script type: %v", err)
	}

	editParams := client.LegacyActionParams("EditScript")
	raw := "#!/bin/bash\necho \"newcode\" > /home/ubuntu/goodbyeworld2.txt"
	enc = base64.StdEncoding.EncodeToString([]byte(raw))
	queryArgsEditorFn = client.EncodeQueryRequestEditor(url.Values{
		"script_id": []string{strconv.Itoa(createdScript.Id)},
		"username":  []string{"jim"},
		"code":      []string{enc},
	})

	editedScriptRes, err := landscapeAPIClient.InvokeLegacyActionWithResponse(ctx, editParams, queryArgsEditorFn)

	if err != nil {
		log.Fatalf("failed to invoke legacy action: %v", err)
	}

	if editedScriptRes == nil {
		log.Fatalf("legacy action returned nil response")
	}

	if editedScriptRes.StatusCode() != 200 {
		log.Fatalf("legacy action failed: status=%d body=%s", editedScriptRes.StatusCode(), string(editedScriptRes.Body))
	}

	log.Printf("raw edit script response: %s", editedScriptRes.Body)

	if editedScriptRes.JSON200 == nil {
		log.Fatalf("legacy action did not return a script object: %s", string(editedScriptRes.Body))
	}

	editedScript, err := editedScriptRes.JSON200.AsV1Script()
	if err != nil {
		log.Fatalf("failed to convert returned script into Script type: %v", err)
	}

	log.Printf("edited script title: %s", editedScript.Title)
	if editedScript.Attachments != nil {
		log.Printf("edited script attachments: %s", *editedScript.Attachments)
	}

	if editedScript.Creator != nil {
		log.Printf("edited creator email: %s", *editedScript.Creator.Email)
		log.Printf("edited creator id: %d", *editedScript.Creator.Id)
		log.Printf("edited creator name: %s", *editedScript.Creator.Name)
	}

}
