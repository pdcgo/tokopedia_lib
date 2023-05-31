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

	app := upload_app.NewUploadApp(db, &upload_app.UploadConfig{
		Concurent: 5,
	})
	repo := upload_app.NewAkunRepo(db)

	sdk := v2_gots_sdk.NewApiSdk(r)
	save := sdk.GenerateSdkFunc("frontend/src/client/sdk_types.ts", true)

	g := sdk.Group("/tokopedia")
	RegisterAkunApi(g, db, repo)
	RegisterCommand(g, app)

	web.RegisterTokopediaFrontend(r, prefix)
	save()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	SetUpTokopediaRouter(r, "tokopedia")
	r.Run("localhost:8080")
}
