// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jansdhillon/landscape-go-api-client/client"
	"github.com/urfave/cli/v3"
)

const (
	nameFlag          = "name"
	materialFlag      = "material"
	materialFileFlag  = "file"
	accessGroupFlag   = "access-group"
	seriesFlag        = "series"
	distributionFlag  = "distribution"
	componentsFlag    = "components"
	architecturesFlag = "architectures"
	modeFlag          = "mode"
	gpgKeyFlag        = "gpg-key"
	mirrorURIFlag     = "mirror-uri"
	mirrorSuiteFlag   = "mirror-suite"
	mirrorGpgKeyFlag  = "mirror-gpg-key"
	mirrorSeriesFlag  = "mirror-series"
	originFlag        = "origin"
)

var gpgKeyCmd = &cli.Command{
	Name:  "gpg-key",
	Usage: "Manage GPG keys.",
	Commands: []*cli.Command{
		{
			Name:  "import",
			Usage: "Import a GPG key.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "Name of the GPG key. Must be unique within the account, start with an alphanumeric character, and only contain lowercase letters, numbers and - or + signs.",
					Required: true,
				},
				&cli.StringFlag{
					Name:    materialFileFlag,
					Aliases: []string{"f"},
					Usage:   "Path to the GPG key file (e.g. key.asc). Use this instead of --material to avoid shell word-splitting issues with armored keys.",
				},
				&cli.StringFlag{
					Name:    materialFlag,
					Aliases: []string{"m"},
					Usage:   "The text representation of the key (literal string). Prefer --file for armored keys to avoid shell expansion issues.",
				},
			},
			Action: importGPGKeyAction,
		},
	},
}

var distributionCmd = &cli.Command{
	Name:  "distribution",
	Usage: "Manage distributions.",
	Commands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new distribution.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "Name of the distribution. Must be unique within the account, start with an alphanumeric character, and only contain lowercase letters, numbers and - or + signs.",
					Required: true,
				},
				&cli.StringFlag{
					Name:    accessGroupFlag,
					Aliases: []string{"g"},
					Usage:   "Optional name of the access group to create the distribution into.",
				},
			},
			Action: createDistributionAction,
		},
	},
}

var pocketCmd = &cli.Command{
	Name:  "pocket",
	Usage: "Manage pockets.",
	Commands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new pocket in a series.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "Name of the pocket. Must be unique within the series, start with an alphanumeric character, and only contain lowercase letters, numbers and - or + signs.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     seriesFlag,
					Aliases:  []string{"s"},
					Usage:    "The name of the series to create the pocket in.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     distributionFlag,
					Aliases:  []string{"d"},
					Usage:    "The name of the distribution the series belongs to.",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     componentsFlag,
					Aliases:  []string{"c"},
					Usage:    "A list of components the pocket will handle. Can be specified multiple times.",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     architecturesFlag,
					Aliases:  []string{"a"},
					Usage:    "A list of architectures the pocket will handle. Can be specified multiple times.",
					Required: true,
				},
				&cli.StringFlag{
					Name:    modeFlag,
					Aliases: []string{"m"},
					Usage:   "The pocket mode. Can be 'pull', 'mirror', or 'upload'.",
					Value:   "upload",
				},
				&cli.StringFlag{
					Name:     gpgKeyFlag,
					Aliases:  []string{"k"},
					Usage:    "The name of the GPG key to use to sign package lists for this pocket.",
					Required: true,
				},
				&cli.StringFlag{
					Name:  mirrorURIFlag,
					Usage: "The URI to mirror for pockets in 'mirror' mode.",
				},
				&cli.StringFlag{
					Name:  mirrorSuiteFlag,
					Usage: "The repository entry under dists/ to mirror for pockets in 'mirror' mode.",
				},
				&cli.StringFlag{
					Name:  mirrorGpgKeyFlag,
					Usage: "The name of the GPG key to use to verify the mirrored archive signature.",
				},
				&cli.StringFlag{
					Name:  originFlag,
					Usage: "The origin of this pocket.",
				},
			},
			Action: createPocketAction,
		},
	},
}

var mirrorCmd = &cli.Command{
	Name:  "mirror",
	Usage: "Manage mirror pockets.",
	Commands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new mirror pocket in a series.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "Name of the mirror pocket.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     seriesFlag,
					Aliases:  []string{"s"},
					Usage:    "The name of the series to create the pocket in.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     distributionFlag,
					Aliases:  []string{"d"},
					Usage:    "The name of the distribution the series belongs to.",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     componentsFlag,
					Aliases:  []string{"c"},
					Usage:    "A list of components the pocket will handle. Can be specified multiple times.",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     architecturesFlag,
					Aliases:  []string{"a"},
					Usage:    "A list of architectures the pocket will handle. Can be specified multiple times.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     gpgKeyFlag,
					Aliases:  []string{"k"},
					Usage:    "The name of the GPG key to use to sign package lists for this pocket.",
					Required: true,
				},
				&cli.StringFlag{
					Name:  mirrorURIFlag,
					Usage: "The URI to mirror.",
				},
				&cli.StringFlag{
					Name:  mirrorSuiteFlag,
					Usage: "The repository entry under dists/ to mirror.",
				},
				&cli.StringFlag{
					Name:  mirrorGpgKeyFlag,
					Usage: "The name of the GPG key to use to verify the mirrored archive signature.",
				},
			},
			Action: createMirrorAction,
		},
		{
			Name:  "sync",
			Usage: "Synchronize a mirror pocket.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "The name of the pocket to synchronize.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     seriesFlag,
					Aliases:  []string{"s"},
					Usage:    "The name of the series.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     distributionFlag,
					Aliases:  []string{"d"},
					Usage:    "The name of the distribution.",
					Required: true,
				},
			},
			Action: syncMirrorAction,
		},
	},
}

func importGPGKeyAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	var material string
	if filePath := cmd.String(materialFileFlag); filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read GPG key file: %w", err)
		}
		material = string(data)
	} else if m := cmd.String(materialFlag); m != "" {
		material = m
	} else {
		return fmt.Errorf("one of --file or --material must be provided")
	}

	res, err := api.LegacyImportGPGKey(ctx, &client.LegacyImportGPGKeyParams{
		Name:     cmd.String(nameFlag),
		Material: material,
	})
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func createDistributionAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	params := &client.LegacyCreateDistributionParams{
		Name: cmd.String(nameFlag),
	}

	if ag := cmd.String(accessGroupFlag); ag != "" {
		params.AccessGroup = &ag
	}

	res, err := api.LegacyCreateDistribution(ctx, params)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func createPocketAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	params := &client.LegacyCreatePocketParams{
		Name:          cmd.String(nameFlag),
		Series:        cmd.String(seriesFlag),
		Distribution:  cmd.String(distributionFlag),
		Components:    cmd.StringSlice(componentsFlag),
		Architectures: cmd.StringSlice(architecturesFlag),
		Mode:          cmd.String(modeFlag),
		GpgKey:        cmd.String(gpgKeyFlag),
	}

	if v := cmd.String(mirrorURIFlag); v != "" {
		params.MirrorUri = &v
	}
	if v := cmd.String(mirrorSuiteFlag); v != "" {
		params.MirrorSuite = &v
	}
	if v := cmd.String(mirrorGpgKeyFlag); v != "" {
		params.MirrorGpgKey = &v
	}
	if v := cmd.String(originFlag); v != "" {
		params.Origin = &v
	}

	res, err := api.LegacyCreatePocket(ctx, params)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func createMirrorAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	params := &client.LegacyCreatePocketParams{
		Name:          cmd.String(nameFlag),
		Series:        cmd.String(seriesFlag),
		Distribution:  cmd.String(distributionFlag),
		Components:    cmd.StringSlice(componentsFlag),
		Architectures: cmd.StringSlice(architecturesFlag),
		Mode:          "mirror",
		GpgKey:        cmd.String(gpgKeyFlag),
	}

	if v := cmd.String(mirrorURIFlag); v != "" {
		params.MirrorUri = &v
	}
	if v := cmd.String(mirrorSuiteFlag); v != "" {
		params.MirrorSuite = &v
	}
	if v := cmd.String(mirrorGpgKeyFlag); v != "" {
		params.MirrorGpgKey = &v
	}

	res, err := api.LegacyCreatePocket(ctx, params)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

var seriesCmd = &cli.Command{
	Name:  "series",
	Usage: "Manage series.",
	Commands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create a new series in a distribution.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     nameFlag,
					Aliases:  []string{"n"},
					Usage:    "Name of the series. Must be unique within the distribution, start with an alphanumeric character, and only contain lowercase letters, numbers and - or + signs.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     distributionFlag,
					Aliases:  []string{"d"},
					Usage:    "The name of the distribution to create the series in.",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:    componentsFlag,
					Aliases: []string{"c"},
					Usage:   "List of components for the created pockets. Required when --pockets is specified.",
				},
				&cli.StringSliceFlag{
					Name:    architecturesFlag,
					Aliases: []string{"a"},
					Usage:   "List of architectures for the created pockets. Required when --pockets is specified.",
				},
				&cli.StringSliceFlag{
					Name:  "pockets",
					Usage: "Pockets to create in the series (mirror mode by default). Can be specified multiple times.",
				},
				&cli.StringFlag{
					Name:  gpgKeyFlag,
					Usage: "The name of the GPG key to use to sign package lists of the created pockets.",
				},
				&cli.StringFlag{
					Name:  mirrorURIFlag,
					Usage: "The URI to mirror for the created pockets.",
				},
				&cli.StringFlag{
					Name:  mirrorSeriesFlag,
					Usage: "The remote series to mirror. Defaults to the name of the series being created.",
				},
				&cli.StringFlag{
					Name:  mirrorGpgKeyFlag,
					Usage: "The name of the GPG key to use to verify the mirrored repositories for created pockets.",
				},
			},
			Action: createSeriesAction,
		},
	},
}

func createSeriesAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	params := &client.LegacyCreateSeriesParams{
		Name:         cmd.String(nameFlag),
		Distribution: cmd.String(distributionFlag),
	}

	if v := cmd.StringSlice(componentsFlag); len(v) > 0 {
		params.Components = &v
	}
	if v := cmd.StringSlice(architecturesFlag); len(v) > 0 {
		params.Architectures = &v
	}
	if v := cmd.StringSlice("pockets"); len(v) > 0 {
		params.Pockets = &v
	}
	if v := cmd.String(gpgKeyFlag); v != "" {
		params.GpgKey = &v
	}
	if v := cmd.String(mirrorURIFlag); v != "" {
		params.MirrorUri = &v
	}
	if v := cmd.String(mirrorSeriesFlag); v != "" {
		params.MirrorSeries = &v
	}
	if v := cmd.String(mirrorGpgKeyFlag); v != "" {
		params.MirrorGpgKey = &v
	}

	res, err := api.LegacyCreateSeries(ctx, params)
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}

func syncMirrorAction(ctx context.Context, cmd *cli.Command) error {
	api, ok := ctx.Value(apiClientKey).(*client.ClientWithResponses)
	if !ok || api == nil {
		return fmt.Errorf("api client not initialized")
	}

	res, err := api.LegacySyncMirrorPocket(ctx, &client.LegacySyncMirrorPocketParams{
		Name:         cmd.String(nameFlag),
		Series:       cmd.String(seriesFlag),
		Distribution: cmd.String(distributionFlag),
	})
	if err != nil {
		return err
	}
	return WriteResponseToRoot(ctx, cmd, res)
}
