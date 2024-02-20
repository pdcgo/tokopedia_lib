package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type UploadApi struct {
	upload *upload_app.UploadApp
	Base   string
}

func (upapi *UploadApi) Start(ctx *gin.Context) {

	cmd := exec.Command("bin/tokopedia.exe", "shopee_toped", "-b", "./")
	cmd.Dir = upapi.Base
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}

	cmd.Start()

	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}
func (upapi *UploadApi) Stop(ctx *gin.Context) {
	upapi.upload.Cancel()
	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}
func (api *UploadApi) Status(ctx *gin.Context) {
	status, _ := api.upload.Status()

	ctx.JSON(http.StatusOK, status)
}

func RegisterCommand(g *v2_gots_sdk.SdkGroup, upload *upload_app.UploadApp, base string) {

	upapi := UploadApi{
		upload: upload,
		Base:   base,
	}

	command := g.Group("upload")

	command.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "start",
		Response:     Response{},
	}, upapi.Start)

	command.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "stop",
		Response:     Response{},
	}, upapi.Stop)

	command.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "status",
		Response:     upload_app.UploadAppStatus{},
	}, upapi.Status)

}
