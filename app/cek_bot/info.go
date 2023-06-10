package cek_bot

import (
	_ "embed"
	"os"
)

var Version = "0.0.1"

//go:embed ..\..\logger_credentials.json
var cred []byte

var devmode = os.Getenv("DEV_MODE") != ""
var LogName = "golang_tokopedia_cekbot"
