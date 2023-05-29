package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func buildFrontendAsset() {
	var outb, errb bytes.Buffer

	cmdBuild := exec.Command("npm", "run", "build", "--prefix", "./frontend")
	cmdBuild.Stdout = &outb
	cmdBuild.Stderr = &errb

	err := cmdBuild.Run()

	fmt.Println("out:", outb.String(), "err:", errb.String())
	if err != nil {
		log.Fatal(err)
	}
	source, _ := filepath.Abs("frontend/build")
	dest, _ := filepath.Abs("app/web/assets/frontend")
	err = os.RemoveAll(dest)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(source, dest)
	if err != nil {
		log.Fatal(err)
	}

}
