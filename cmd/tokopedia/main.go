package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/web"
	"github.com/pdcgo/v2_gots_sdk"
)

func main() {
	r := gin.Default()

	sdk := v2_gots_sdk.NewApiSdk(r)
	save := sdk.GenerateSdkFunc("frontend/src/sdk.ts")

	sdk.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "/ping",
	},
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	web.RegisterTokopediaFrontend(r, "tokopedia/")

	save()

	r.Run("localhost:8080")
}
