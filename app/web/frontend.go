package web

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var frontendAssets embed.FS

func RegisterTokopediaFrontend(r *gin.Engine, prefix string) {

	httpfs := http.FS(frontendAssets)

	fs.WalkDir(frontendAssets, "assets/frontend", func(file string, d fs.DirEntry, err error) error {
		uri := strings.ReplaceAll(file, "assets/frontend", "")

		if uri == "" {
			return nil
		}
		uri = prefix + uri
		r.StaticFileFS(uri, file, httpfs)

		return nil
	})

	r.GET("/"+prefix, func(ctx *gin.Context) {
		data, _ := frontendAssets.ReadFile("assets/frontend/index.html")
		ctx.Data(http.StatusOK, "text/html", data)
	})
}
