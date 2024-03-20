package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type AutoReplyApi struct {
	BaseApi
	areplyConfig *config.AutoReplyConfig
}

func NewAutoReplyApi(areplyConfig *config.AutoReplyConfig) *AutoReplyApi {
	return &AutoReplyApi{
		areplyConfig: areplyConfig,
	}
}

func (api *AutoReplyApi) get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.areplyConfig.Data)
}

func (api *AutoReplyApi) update(ctx *gin.Context) {

	err := ctx.BindJSON(&api.areplyConfig.Data)
	if err != nil {
		api.BaseResponseBadRequest(err)
		return
	}

	err = api.areplyConfig.Save()
	if err != nil {
		api.BaseResponseInternalServerError(err)
		return
	}

	ctx.JSON(http.StatusOK, api.areplyConfig.Data)
}

func (api *AutoReplyApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Response:     config.AutoReplyConfigData{},
	}, api.get)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "",
		Payload:      config.AutoReplyConfigData{},
		Response:     config.AutoReplyConfigData{},
	}, api.update)
}
