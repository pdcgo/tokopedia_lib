package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type CheckVerifApi struct {
	base repo.BaseInterface
}

type RunCheckVerifPayload struct {
	Fname string `json:"fname"`
	Akuns []*report.CekVerifReport
}

func (cekbot *CheckVerifApi) runBin(fname string) {
	cmd := exec.Command("bin/tokopedia.exe", "cv", "-fname", fname)
	cmd.Dir = cekbot.base.Path()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}

	cmd.Start()

}

func (cekbot *CheckVerifApi) RunCekverif(ctx *gin.Context) {
	var payload RunCheckbotPayload
	ctx.BindJSON(&payload)

	hasil := make([]*report.CekVerifReport, len(payload.Akuns))

	for ind, akun := range payload.Akuns {
		hasil[ind] = &report.CekVerifReport{
			DriverAccount: akun,
		}
	}
	fname := cekbot.base.Path(payload.Fname)

	report.SaveCekVerifReport(fname, hasil)

	cekbot.runBin(payload.Fname)
	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}

func RegisterCheckVerifApi(grp *v2_gots_sdk.SdkGroup, base repo.BaseInterface) {
	api := CheckVerifApi{
		base: base,
	}
	delgrp := grp.Group("check_verif")
	delgrp.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "run",
		Payload:      RunCheckVerifPayload{},
	}, api.RunCekverif)
}
