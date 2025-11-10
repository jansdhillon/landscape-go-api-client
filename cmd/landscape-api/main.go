// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jansdhillon/landscape-go-api-client/client"
	"github.com/urfave/cli/v3"
)

type ctxKey string

const apiClientKey ctxKey = "landscape-api-client"

const (
	baseURLFlag   = "base-url"
	accessKeyFlag = "access-key"
	secretKeyFlag = "secret-key"
	emailFlag     = "email"
	passwordFlag  = "password"
	accountFlag   = "account"
)

func main() {
	cmd := &cli.Command{
		Name:  "landscape-api",
		Usage: "Interact with the Landscape API.",
		Commands: []*cli.Command{
			scriptCmd,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     baseURLFlag,
				Aliases:  []string{"u", "url", "base_url"},
				Usage:    "The base URL of Landscape (can also be set via LANDSCAPE_BASE_URL env var).",
				Required: true,
				Sources:  cli.EnvVars("LANDSCAPE_BASE_URL"),
			},
			&cli.StringFlag{
				Name:    accessKeyFlag,
				Aliases: []string{"ak", "access_key"},
				Usage:   "An access key for the Landscape API (can also be set via LANDSCAPE_ACCESS_KEY env var). If provided, you must also provide the -secret-key flag or set the LANDSCAPE_SECRET_KEY env var.",
				Sources: cli.EnvVars("LANDSCAPE_ACCESS_KEY"),
			},
			&cli.StringFlag{
				Name:    secretKeyFlag,
				Aliases: []string{"sk", "secret_key"},
				Usage:   "An secret key for the Landscape API (can also be set via LANDSCAPE_SECRET_KEY env var). If provided, you must also provide the -access-key flag or set the LANDSCAPE_ACCESS_KEY env var.",
				Sources: cli.EnvVars("LANDSCAPE_SECRET_KEY"),
			},
			&cli.StringFlag{
				Name:    emailFlag,
				Aliases: []string{"e"},
				Usage:   "An email to access the Landscape API (can also be set via LANDSCAPE_EMAIL env var). If provided, you must also provide the -password flag or set the LANDSCAPE_PASSWORD env var.",
				Sources: cli.EnvVars("LANDSCAPE_EMAIL"),
			},
			&cli.StringFlag{
				Name:    passwordFlag,
				Aliases: []string{"p"},
				Usage:   "A password to access the Landscape API (can also be set via LANDSCAPE_PASSWORD env var). If provided, you must also provide the -email flag or set the LANDSCAPE_EMAIL env var.",
				Sources: cli.EnvVars("LANDSCAPE_PASSWORD"),
			},
			&cli.StringFlag{
				Name:    accountFlag,
				Aliases: []string{"a"},
				Usage:   "An account to login into the Landscape API with (can also be set via LANDSCAPE_ACCOUNT env var). If provided, you must also provide the -email and -password flags or set the LANDSCAPE_EMAIL and LANDSCAPE_PASSWORD env vars.",
				Sources: cli.EnvVars("LANDSCAPE_ACCOUNT"),
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			baseURL := c.String(baseURLFlag)
			if baseURL == "" {
				return ctx, fmt.Errorf("base URL must be provided")
			}

			email := c.String(emailFlag)
			password := c.String(passwordFlag)
			account := c.String(accountFlag)

			var lp client.LoginProvider

			if email != "" && password != "" {
				if account != "" {
					lp = client.NewEmailPasswordProvider(email, password, &account)
				} else {
					lp = client.NewEmailPasswordProvider(email, password, nil)
				}

			} else {
				accessKey := c.String(accessKeyFlag)
				secretKey := c.String(secretKeyFlag)

				if accessKey == "" || secretKey == "" {
					return ctx, fmt.Errorf("must provide the -e & -p flags or the -ak & -sk flags, or set either the LANDSCAPE_EMAIL & LANDSCAPE_PASSWORD env vars or the LANDSCAPE_ACCESS_KEY & LANDSCAPE_SECRET_KEY env vars")
				}

				lp = client.NewAccessKeyProvider(accessKey, secretKey)
			}

			api, err := client.NewLandscapeAPIClient(baseURL, lp)
			if err != nil {
				return ctx, err
			}

			return context.WithValue(ctx, apiClientKey, api), nil
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}

func WriteResponseToRoot(_ context.Context, cmd *cli.Command, res *http.Response) error {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, body, "", "  "); err != nil {
		return err
	}

	out.WriteTo(cmd.Root().Writer)
	fmt.Fprintln(cmd.Root().Writer)
	return nil
}
