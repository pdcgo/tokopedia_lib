package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/pdcgo/v2_gots_sdk"
)

type CekbotApi struct {
	base repo.BaseInterface
}

const (
	CREATE_NEW_CONSOLE = 0x10
)

type CekbotAkun struct {
}

type RunCheckbotPayload struct {
	Fname string `json:"fname"`
	Akuns []*tokopedia_lib.DriverAccount
}

func (cekbot *CekbotApi) runBin(fname string) {
	cmd := exec.Command("bin/tokopedia.exe", "cekbot", "-fname", fname)
	cmd.Dir = cekbot.base.Path()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}

	cmd.Start()

}

func (cekbot *CekbotApi) RunCekbot(ctx *gin.Context) {
	var payload RunCheckbotPayload
	ctx.BindJSON(&payload)

	hasil := make([]*report.CekReport, len(payload.Akuns))

	for ind, akun := range payload.Akuns {
		hasil[ind] = &report.CekReport{
			DriverAccount: akun,
		}
	}
	fname := cekbot.base.Path(payload.Fname)
	report.SaveCekReport(fname, hasil)

	cekbot.runBin(payload.Fname)
	ctx.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}

func RegisterCekbotApi(grp *v2_gots_sdk.SdkGroup, base repo.BaseInterface) {
	api := CekbotApi{
		base: base,
	}

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "run",
		Payload:      RunCheckbotPayload{},
	}, api.RunCekbot)
}
