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
				Usage:    "The base URL of Landscape (can also be set via LANDSCAPE_BASE_URL env var)",
				Required: true,
				Sources:  cli.EnvVars("LANDSCAPE_BASE_URL"),
			},
			&cli.StringFlag{
				Name:     accessKeyFlag,
				Aliases:  []string{"ak", "access_key"},
				Usage:    "An access key for the Landscape API (can also be set via LANDSCAPE_ACCESS_KEY env var)",
				Required: true,
				Sources:  cli.EnvVars("LANDSCAPE_ACCESS_KEY"),
			},
			&cli.StringFlag{
				Name:     secretKeyFlag,
				Aliases:  []string{"sk", "secret_key"},
				Usage:    "An secret key for the Landscape API (can also be set via LANDSCAPE_SECRET_KEY env var)",
				Required: true,
				Sources:  cli.EnvVars("LANDSCAPE_SECRET_KEY"),
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			baseURL := c.String(baseURLFlag)
			accessKey := c.String(accessKeyFlag)
			secretKey := c.String(secretKeyFlag)

			lp := client.NewAccessKeyProvider(accessKey, secretKey)
			api, err := client.NewLandscapeAPIClient(baseURL, lp)
			if err != nil {
				return ctx, err
			}

			return context.WithValue(ctx, apiClientKey, api), nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
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
