// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jansdhillon/landscape-go-api-client/client"
	"github.com/urfave/cli/v3"
)

const (
	codeFlag       = "code"
	fileFlag       = "file"
	scriptTypeFlag = "script-type"
	titleFlag      = "title"
	scriptIdFlag   = "script-id"
)

var scriptCmd = &cli.Command{
	Name:  "script",
	Usage: "Manage and create Landscape scripts.",
	Commands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new script.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     titleFlag,
					Aliases:  []string{"t"},
					Required: true,
				},
				&cli.StringFlag{
					Name:     codeFlag,
					Aliases:  []string{"c"},
					Required: true,
				},
				&cli.StringFlag{
					Name:     scriptTypeFlag,
					Aliases:  []string{"st", "script_type"},
					Required: false,
					Value:    "V1",
				},
			},
			Action: createScriptAction,
		},
		{
			Name:      "edit",
			Usage:     "Edit an existing script.",
			ArgsUsage: "[script-id]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     titleFlag,
					Aliases:  []string{"t"},
					Required: false,
				},
				&cli.StringFlag{
					Name:     codeFlag,
					Aliases:  []string{"c"},
					Required: true,
				},
			},
			Action: editScriptAction,
		},
		{
			Name:      "get",
			Usage:     "Get an existing script.",
			ArgsUsage: "[script-id]",
			Action:    getScriptAction,
		},
		{
			Name:  "attachment",
			Usage: "Create or manage script attachments.",
			Commands: []*cli.Command{
				{
					Name:      "create",
					Usage:     "Create a script attachment.",
					ArgsUsage: "[script-id]",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     fileFlag,
							Aliases:  []string{"f"},
							Usage:    "The file you wish to use as an attachment. The format for this parameter is: <filename>$$<base64 encoded file contents>.",
							Required: true,
						},
					},
					Action: createScriptAttachmentAction,
				},
			},
		},
	},
}

func createScriptAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	title := cmd.String(titleFlag)
	code := cmd.String(codeFlag)
	scriptType := cmd.String(scriptTypeFlag)

	enc := base64.StdEncoding.EncodeToString([]byte(code))

	params := client.LegacyActionParams("CreateScript")
	edit := client.EncodeQueryRequestEditor(url.Values{
		"title":       []string{title},
		"code":        []string{enc},
		"script_type": []string{scriptType},
	})

	res, err := api.InvokeLegacyAction(ctx, params, edit)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func editScriptAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api not initialized")
	}

	noArgs := cmd.Args().Len() == 0
	scriptIDStr := cmd.Args().First()
	if noArgs || scriptIDStr == "" {
		return fmt.Errorf("script ID must be provided as the first argument")
	}

	scriptID, err := strconv.Atoi(scriptIDStr)
	if err != nil {
		return fmt.Errorf("couldn't convert script ID to string: %s", err)
	}

	title := cmd.String(titleFlag)
	code := cmd.String(codeFlag)

	enc := base64.StdEncoding.EncodeToString([]byte(code))

	params := client.LegacyActionParams("EditScript")
	edit := client.EncodeQueryRequestEditor(url.Values{
		"title":     []string{title},
		"code":      []string{enc},
		"script_id": []string{strconv.Itoa(scriptID)},
	})

	res, err := api.InvokeLegacyAction(ctx, params, edit)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func getScriptAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	noArgs := cmd.Args().Len() == 0
	scriptIDStr := cmd.Args().First()
	if noArgs || scriptIDStr == "" {
		return fmt.Errorf("script ID must be provided as the first argument")
	}

	scriptID, err := strconv.Atoi(scriptIDStr)
	if err != nil {
		return fmt.Errorf("couldn't convert script ID to string: %s", err)
	}

	res, err := api.GetScript(ctx, scriptID)
	if err != nil {
		return err
	}

	return WriteResponseToRoot(ctx, cmd, res)
}

func createScriptAttachmentAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	noArgs := cmd.Args().Len() == 0
	scriptIDStr := cmd.Args().First()
	if noArgs || scriptIDStr == "" {
		return fmt.Errorf("script ID must be provided as the first argument")
	}

	scriptID, err := strconv.Atoi(scriptIDStr)
	if err != nil {
		return fmt.Errorf("couldn't convert script ID to string: %s", err)
	}

	file := cmd.String(fileFlag)

	params := client.LegacyActionParams("CreateScriptAttachment")
	edit := client.EncodeQueryRequestEditor(url.Values{
		"script_id": []string{strconv.Itoa(scriptID)},
		"file":      []string{file},
	})

	res, err := api.InvokeLegacyAction(ctx, params, edit)
	if err != nil {
		return err
	}

	return WriteResponseToRoot(ctx, cmd, res)
}
