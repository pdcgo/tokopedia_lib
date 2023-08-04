package config

import (
	_ "embed"
	"os"
)

var Version = "0.0.1"

//go:embed ..\..\logger_credentials.json
var Cred []byte

var devmode = os.Getenv("DEV_MODE") != ""
var LogName = "golang_tokopedia_cekbot"

var fname = "data/config.json"
