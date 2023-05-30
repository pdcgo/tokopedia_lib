package main

import (
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/urfave/cli/v2"
)

func runApplication(cCtx *cli.Context) error {
	base := cCtx.String("b")

	WalkPurnaSession(base, func(fname string, sess *PurnaSession) error {
		log.Println("opening session", fname)
		driver := &tokopedia_lib.DriverAccount{
			Username: "",
			Password: "",
			Secret:   "",
			Session:  sess,
		}
		driver.DevMode = os.Getenv("DEV_MODE") != ""

		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			chromedp.Run(dctx.Ctx, chromedp.Navigate("https://seller.tokopedia.com"))

			<-dctx.Ctx.Done()

			return nil
		})

		return nil
	})

	return nil

}

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "base",
				Aliases:  []string{"b"},
				Required: true,
			},
		},
		Name:   "Change email harum mumet",
		Usage:  "emailmumet",
		Action: runApplication,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
