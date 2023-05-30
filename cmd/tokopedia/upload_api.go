package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/v2_gots_sdk"
)

type UploadApi struct {
	upload *upload_app.UploadApp
}

func (api *UploadApi) Start(ctx *gin.Context) {
	api.upload.Start()
	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}
func (api *UploadApi) Stop(ctx *gin.Context) {
	api.upload.Cancel()
	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}
func (api *UploadApi) Status(ctx *gin.Context) {
	status, _ := api.upload.Status()

	ctx.JSON(http.StatusOK, status)
}

func RegisterCommand(g *v2_gots_sdk.SdkGroup, upload *upload_app.UploadApp) {

	api := UploadApi{
		upload: upload,
	}

	command := g.Group("upload")

	command.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "start",
		Response:     Response{},
	}, api.Start)

	command.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "stop",
		Response:     Response{},
	}, api.Stop)

	command.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "status",
		Response:     upload_app.UploadStatus{},
	}, api.Status)

}
