package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
)

func main() {
	conf := config.NewAppConfig(".\\..")
	if !conf.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	app, err := chat.InitApplication(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
