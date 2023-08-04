package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pdcgo/autoupdater"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var storage = "tokopedia_uploader_artifact"
var variant = "beta"
var entryPoint = "./bin/tokopedia.exe"

func createConfig(outputdir string) (string, error) {
	log.Println("create config lisensi...")
	config := pdc_common.PdcConfig{}

	fconfig := filepath.Join(outputdir, "config.yml.example")
	f, err := os.OpenFile(fconfig, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)

	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	yaml.NewEncoder(f).Encode(&config)

	return fconfig, nil
}

func buildBinary(outputdir string) (string, error) {
	log.Println("create aplication...")
	var outb, errb bytes.Buffer

	updatefname := filepath.Join(outputdir, entryPoint)
	flags := "-X 'main.DebugMode=false' -s -w"

	cmdBuild := exec.Command("go", "build", "-o", updatefname, "-ldflags", flags, "./cmd/cek_akun")
	cmdBuild.Stdout = &outb
	cmdBuild.Stderr = &errb

	err := cmdBuild.Run()

	if err != nil {
		return "", err
	}

	// compress bin
	log.Println("compressing application...")
	comfname, _ := filepath.Abs(updatefname)
	cmdBuild = exec.Command("upx", "--best", "--lzma", comfname)
	cmdBuild.Stdout = &outb
	cmdBuild.Stderr = &errb

	err = cmdBuild.Run()

	if err != nil {
		return "", err
	}

	return updatefname, nil
}

func publishCekAkun(ctx *cli.Context) error {

	up := autoupdater.Publiser{
		Version:       config.Version,
		Storage:       storage,
		Variant:       variant,
		OutputDir:     "dist",
		AppEntryPoint: entryPoint,
		BuildCmd:      []autoupdater.BuildFunc{buildBinary, createConfig},
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
