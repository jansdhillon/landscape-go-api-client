package main

import (
	"context"
	"encoding/base64"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/jansdhillon/landscape-go-api-client/client"
)

func main() {
	apiUrl := os.Getenv("LANDSCAPE_API_URL")
	if apiUrl == "" {
		log.Fatalf("API URL not set")
	}
	ak := os.Getenv("LANDSCAPE_API_ACCESS_KEY")
	if ak == "" {
		log.Fatalf("access key not set")
	}

	sk := os.Getenv("LANDSCAPE_API_SECRET_KEY")
	if sk == "" {
		log.Fatalf("secret key not set")
	}

	scriptId := os.Getenv("LANDSCAPE_SCRIPT_ID")

	scriptIdInt, err := strconv.Atoi(scriptId)

	if err != nil {
		log.Fatalf("failed to convert scriptId to integer: %s", err)
	}

	landscapeAPIClient, err := client.NewLandscapeAPIClient(
		apiUrl,
		client.NewAccessKeyProvider(ak, sk),
	)
	if err != nil {
		log.Fatalf("failed to create Landscape API client: %v", err)
	}

	ctx := context.Background()

	scriptRes, err := landscapeAPIClient.GetScriptWithResponse(ctx, scriptIdInt)

	if err != nil {
		log.Fatalf("failed to get script: %s", err)
	}

	if scriptRes.JSON200 == nil {
		log.Fatal("failed to get script (nil)")
	}

	log.Printf("raw script response: %s", scriptRes.Body)

	log.Printf("script title: %s", scriptRes.JSON200.Title)
	if scriptRes.JSON200.Code != nil {
		log.Printf("script code: %s", *scriptRes.JSON200.Code)
	}
	if scriptRes.JSON200.Attachments != nil {
		log.Printf("script attachments: %s", scriptRes.JSON200.Attachments)

	}

	editParams := client.LegacyActionParams("EditScript")
	raw := "#!/bin/bash\necho \"goodbyeworld2\" > /home/ubuntu/goodbyeworld2.txt"
	enc := base64.StdEncoding.EncodeToString([]byte(raw))
	queryArgsEditorFn := client.EncodeQueryRequestEditor(url.Values{
		"script_id": []string{scriptId},
		"title":     []string{"goodbyeworld5"},
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

	editedScript, err := editedScriptRes.JSON200.AsScript()
	if err != nil {
		log.Fatalf("failed to convert returned script into Script type: %v", err)
	}

	log.Printf("edited script title: %s", editedScript.Title)
	if editedScript.Code != nil {
		log.Printf("edited script code: %s", *editedScript.Code)
	}
	if editedScript.Attachments != nil {
		log.Printf("edited script attachments: %s", editedScript.Attachments)
	}

	if editedScript.CreatedBy != nil {
		log.Printf("edited CreatedBy name: %s", *editedScript.CreatedBy.Name)
		log.Printf("edited CreatedBy id: %d", *editedScript.CreatedBy.Id)
	}

	if editedScript.Creator != nil {
		log.Printf("edited creator email: %s", *editedScript.Creator.Email)
		log.Printf("edited creator id: %d", *editedScript.Creator.Id)
		log.Printf("edited creator name: %s", *editedScript.Creator.Name)
	}

}
