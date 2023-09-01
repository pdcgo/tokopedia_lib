package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/pdcgo/v2_gots_sdk"
)

func CreateChatSdk(r *gin.Engine) *v2_gots_sdk.ApiSdk {
	sdk := v2_gots_sdk.NewApiSdk(r)

	// cors
	sdk.R.Use(func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")
		c.Header("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")
		c.Next()
	})

	return sdk
}
