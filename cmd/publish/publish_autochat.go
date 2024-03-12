package main

import (
	"log"

	"github.com/pdcgo/autoupdater"
	"github.com/pdcgo/tokopedia_lib/lib/app_builder"
	"github.com/urfave/cli/v2"
)

var autochatStorage = "tokopedia_chat_artifact"
var autochatVariant = "auto_chat"
var autochatEntryPoint = "./bin/autochat.exe"
var autochatVersion = "1.0.1"

func publishAutochat(ctx *cli.Context) error {

	up := autoupdater.Publiser{
		Version:       autochatVersion,
		Storage:       autochatStorage,
		Variant:       autochatVariant,
		OutputDir:     "dist",
		AppEntryPoint: autochatEntryPoint,
		BuildCmd: []autoupdater.BuildFunc{func(outputdir string) (string, error) {
			return app_builder.BuildBynaryCmd(outputdir, "./cmd/autochat", autochatEntryPoint)
		}},
	}

	mode := ctx.String("m")
	log.Println(mode)
	if mode == "release" {
		up.Run()
	} else {
		up.RunBuild()
	}

	return nil
}
