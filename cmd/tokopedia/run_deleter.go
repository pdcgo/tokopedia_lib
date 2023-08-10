package main

import (
	"log"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
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

	app := pdc_application.PdcApplication{
		Base: &legacy_source.BaseConfig{
			BaseData: ctx.String("base"),
		},
		Credential:    Cred,
		Version:       "development",
		ReplaceLogger: true,
		AppID:         AppID,
	}

	app.RunWithLicenseFile("data/config.json", "golang_tokopedia_delete", func(app *pdc_application.PdcApplication) {

		fname := app.Base.Path("data/deleter_config.json")
		config, err := deleter_product.NewDeleteConfig(fname)

		if err != nil {
			pdc_common.ReportError(err)
		}

		runner := deleter_product.NewDeleteRunner(config)

		log.Println("running deleter")
		runner.Run()

	})

	return nil

}
