package main

import (
	"log"
	"os"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/autosubmit"
	"github.com/pdcgo/tokopedia_lib/app/cek_bot"
	"github.com/pdcgo/tokopedia_lib/app/cek_verification"
	"github.com/urfave/cli/v2"
)

func catch() {
	if r := recover(); r != nil {
		log.Println(r)

		time.Sleep(time.Minute)
	}
}

func main() {
	defer catch()

	checkbotCommand := cek_bot.CreateCheckbotCommand()
	checkverifCommand := cek_verification.CreateCheckVerifCommand()
	deleteCommand := createDeleteCommand()
	submitCommand := autosubmit.CreateSubmitCommand()

	app := &cli.App{
		Name:  "Tokopedia Server Tool",
		Usage: "Binary Tokopedia Server Tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Value:   "../",
			},
		},
		Action: runWebServer,
		Commands: []*cli.Command{
			{
				Name:    "shopee_toped",
				Aliases: []string{"st"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "base",
						Aliases: []string{"b"},
						Value:   "../",
					},
				},
				Action: runUploadShopeeToped,
			},
			checkbotCommand,
			deleteCommand,
			checkverifCommand,
			submitCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		pdc_common.ReportError(err)
		log.Fatal(err)
	}

}
