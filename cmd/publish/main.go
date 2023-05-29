package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Publisher",
		Usage: "Publisher data",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Pdc Publisher")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "publish",
				Aliases: []string{"p"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "m",
						Value: "release",
						Usage: "mode 'release' atau 'build'",
					},
				},
				Action: publishCekAkun,
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "build options",
						Value: string(AllBuildFlag),
						Usage: getUsageBuildFlagString(),
					},
				},
				Action: func(ctx *cli.Context) error {
					buildFrontendAsset()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
