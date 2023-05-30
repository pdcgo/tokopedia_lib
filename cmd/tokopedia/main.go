package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/tokopedia_lib/app/web"
	"github.com/pdcgo/v2_gots_sdk"
)

func SetUpTokopediaRouter(r *gin.Engine, prefix string) {

	upload_app := upload_app.NewUploadApp()

	sdk := v2_gots_sdk.NewApiSdk(r)
	save := sdk.GenerateSdkFunc("frontend/src/sdk.ts")

	g := sdk.Group("/tokopedia")
	RegisterAkunApi(g)
	RegisterCommand(g, upload_app)

	web.RegisterTokopediaFrontend(r, prefix)
	save()
}

func main() {
	r := gin.Default()
	SetUpTokopediaRouter(r, "tokopedia")
	r.Run("localhost:8080")
}
