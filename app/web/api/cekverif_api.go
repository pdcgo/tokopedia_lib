package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/cek_verification"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
)

type CheckVerifApi struct {
	base repo.BaseInterface
}

type RunCheckVerifPayload struct {
	Fname string `json:"fname"`
	Akuns []*cek_verification.VerifDriverAccount
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

	hasil := make([]*cek_verification.VerifDriverAccount, len(payload.Akuns))

	for ind, akun := range payload.Akuns {
		hasil[ind] = &cek_verification.VerifDriverAccount{
			DriverAccount: akun,
		}
	}
	fname := cekbot.base.Path(payload.Fname)

	cek_verification.SaveCekReport(fname, hasil)

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
	delgrp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "run",
		Payload:      RunCheckVerifPayload{},
	}, api.RunCekverif)
}
