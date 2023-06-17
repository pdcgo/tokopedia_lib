package main

import (
	"log"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/urfave/cli/v2"
)

func createDeleteCommand() *cli.Command {
	command := cli.Command{
		Name:    "delete_product",
		Aliases: []string{"delprod"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "../",
			},
		},
		Action: runDeleteCommand,
	}
	return &command

}

func runDeleteCommand(ctx *cli.Context) error {
	rootBase := ctx.String("b")
	baseData := &legacy_source.BaseConfig{
		BaseData: rootBase,
	}

	fname := baseData.Path("data/deleter_config.json")
	config, err := deleter_product.NewDeleteConfig(fname)

	if err != nil {
		return err
	}

	runner := deleter_product.NewDeleteRunner(config)

	log.Println("running deleter")
	runner.Run()

	return nil

}
