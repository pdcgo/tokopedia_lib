package api

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
)

type DeleterApi struct {
	base repo.BaseInterface
}

type DeleteSettingRes struct {
	Data *deleter_product.DeleteConfig `json:"data"`
}

func (runapi *DeleterApi) GetSetting(c *gin.Context) {
	fname := runapi.base.Path("data/deleter_config.json")
	cfg, _ := deleter_product.NewDeleteConfig(fname)

	res := DeleteSettingRes{
		Data: cfg,
	}

	c.JSON(http.StatusOK, &res)

}

func (runapi *DeleterApi) UpdateSetting(c *gin.Context) {
	var payload deleter_product.DeleteConfig
	c.BindJSON(&payload)

	fname := runapi.base.Path("data/deleter_config.json")
	deleter_product.SaveDeleteConfig(fname, &payload)

	res := DeleteSettingRes{
		Data: &payload,
	}
	c.JSON(http.StatusOK, &res)
}

func (runapi *DeleterApi) RunDelete(c *gin.Context) {
	cmd := exec.Command("bin/tokopedia.exe", "delete_product", "-base", "./")
	cmd.Dir = runapi.base.Path()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    CREATE_NEW_CONSOLE,
		NoInheritHandles: true,
	}

	cmd.Start()

	c.JSON(http.StatusOK, Response{
		Msg: "success",
	})
}

func RegisterDeleterApi(grp *v2_gots_sdk.SdkGroup, base repo.BaseInterface) {
	deleter := DeleterApi{
		base: base,
	}
	delgrp := grp.Group("deleter")

	delgrp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "setting",
		Response:     DeleteSettingRes{},
	}, deleter.GetSetting)

	delgrp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "setting",
		Payload:      deleter_product.DeleteConfig{},
		Response:     DeleteSettingRes{},
	}, deleter.UpdateSetting)

	delgrp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "run_delete",
		Response:     Response{},
	}, deleter.RunDelete)
}
