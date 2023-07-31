package config

import "os"

var Version = "0.0.1"

var devmode = os.Getenv("DEV_MODE") != ""
var LogName = "golang_tokopedia_cekbot"

var fname = "data/config.json"
