package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/pdc_launcher/essential"
	"github.com/pdcgo/tokopedia_lib/app/autochat"
	"github.com/urfave/cli/v2"
)

var AppLogname = "golang_tokopedia_autochat"
var AppID = 2

func main() {
	sleepErrorhandler := func(err error) {
		time.Sleep(time.Hour)
	}

	app := &cli.App{
		Name:  "Tokopedia Server Tool",
		Usage: "Binary Tokopedia Server Tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "./",
			},
		},
		Action: func(ctx *cli.Context) error {

			app := pdc_application.PdcApplication{
				Base: &legacy_source.BaseConfig{
					BaseData: ctx.String("base"),
				},
				Version:       essential.Version,
				AppID:         essential.AppID,
				Credential:    essential.LoggerCredential,
				OnPanic:       []func(err error){sleepErrorhandler},
				OnError:       []func(err error){sleepErrorhandler},
				ReplaceLogger: true,
			}

			return app.RunWithLicenseFile(essential.UniversalConfigName, AppLogname, func(app *pdc_application.PdcApplication) error {

				autochatapp, err := autochat.InitApplication(app.Base)
				if err != nil {
					return err
				}

				fmt.Println("---------- PDC Tokopedia Autochat ----------")
				fmt.Println("Mode :")
				fmt.Println("1. Auto Send ( kirim pesan ke seller )")
				fmt.Println("2. Auto Reply ( balasan otomatis ke buyer )")
				fmt.Println("--------------------------------------------")

				var mode autochat.AutochatMode
				fmt.Print("Pilih mode : ")
				fmt.Scanln(&mode)

				return autochatapp.Run(mode)
			})
		},
	}

	if err := app.Run(os.Args); err != nil {
		pdc_common.ReportError(err)
		log.Fatal(err)
	}

}
