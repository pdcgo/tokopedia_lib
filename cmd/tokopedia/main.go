package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/tokopedia_lib/app/web"
	"github.com/pdcgo/tokopedia_lib/lib/datasource"
	"github.com/pdcgo/v2_gots_sdk"
)

func SetUpTokopediaRouter(r *gin.Engine, prefix string) {
	db := datasource.NewSqliteDatabase("tokopedia_data.db")

	app := upload_app.NewUploadApp()
	repo := upload_app.NewAkunRepo(db)

	sdk := v2_gots_sdk.NewApiSdk(r)
	save := sdk.GenerateSdkFunc("frontend/src/sdk.ts")

	g := sdk.Group("/tokopedia")
	RegisterAkunApi(g, db, repo)
	RegisterCommand(g, app)

	web.RegisterTokopediaFrontend(r, prefix)
	save()
}

func main() {
	r := gin.Default()
	SetUpTokopediaRouter(r, "tokopedia")
	r.Run("localhost:8080")
}
