package web

import (
	"embed"
	"io/fs"
	"net/http"
	"net/url"
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
		uri, _ = url.JoinPath("/", prefix, "./", uri)
		r.StaticFileFS(uri, file, httpfs)

		return nil
	})

	pathhome, _ := url.JoinPath("/", prefix, "./")

	r.GET(pathhome, func(ctx *gin.Context) {
		data, _ := frontendAssets.ReadFile("assets/frontend/index.html")
		ctx.Data(http.StatusOK, "text/html", data)
	})
}
