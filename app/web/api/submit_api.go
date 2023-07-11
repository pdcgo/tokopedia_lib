package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/autosubmit"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
)

type SubmitApi struct {
	base repo.BaseInterface
}

func RegisterSubmitApi(grp *v2_gots_sdk.SdkGroup, base repo.BaseInterface) {
	submit := SubmitApi{
		base: base,
	}

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "run",
		Payload:      &autosubmit.AutoSubmit{},
	}, submit.Run)

}

func (sub *SubmitApi) Run(ctx *gin.Context) {
	var submit autosubmit.AutoSubmit

	err := ctx.BindJSON(&submit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "bad_request",
		})
		return
	}

	cmd := exec.Command("bin/tokopedia.exe", "submit_ktp")
	cmd.Dir = sub.base.Path()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}

	cmd.Start()
}
