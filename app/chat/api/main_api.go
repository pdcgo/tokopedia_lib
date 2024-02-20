package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type MainApi struct{}

func NewMainApi() *MainApi {
	return &MainApi{}
}

type BotConfigRes struct {
	Version string `json:"version"`
}

func (api *MainApi) bot(ctx *gin.Context) {

	res := BotConfigRes{
		Version: config.Version,
	}
	ctx.JSON(http.StatusOK, res)
}

func (api *MainApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "bot",
		Response:     BotConfigRes{},
	}, api.bot)
}
