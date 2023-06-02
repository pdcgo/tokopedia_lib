package main

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
)

func BuildBynaryCmd(outputdir string, pathapp string, entryPoint string) (string, error) {
	log.Println("create aplication... " + pathapp)
	var outb, errb bytes.Buffer

	updatefname := filepath.Join(outputdir, entryPoint)
	flags := "-X 'main.DebugMode=false' -s -w"

	cmdBuild := exec.Command("go", "build", "-o", updatefname, "-ldflags", flags, pathapp)
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
