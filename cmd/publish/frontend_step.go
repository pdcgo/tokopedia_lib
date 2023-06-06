package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func installDependendAsset() {
	log.Println("installing dependencies frontend with npm..")
	var outb, errb bytes.Buffer

	cmdBuild := exec.Command("npm", "install", "--save-dev")
	cmdBuild.Dir = "./frontend"
	cmdBuild.Stdout = &outb
	cmdBuild.Stderr = &errb

	err := cmdBuild.Run()

	fmt.Println("out:", outb.String(), "err:", errb.String())
	if err != nil {
		log.Fatal(err)
	}
}

func buildFrontendAsset() {
	installDependendAsset()

	var outb, errb bytes.Buffer

	cmdBuild := exec.Command("npm", "run", "build", "--prefix", "./frontend")
	cmdBuild.Stdout = &outb
	cmdBuild.Stderr = &errb

	err := cmdBuild.Run()

	fmt.Println("out:", outb.String(), "err:", errb.String())
	if err != nil {
		log.Fatal(err)
	}
	source, _ := filepath.Abs("frontend/dist")
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
