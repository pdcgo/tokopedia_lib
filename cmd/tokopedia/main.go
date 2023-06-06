package main

import (
	"log"
	"os"
	"time"

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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
