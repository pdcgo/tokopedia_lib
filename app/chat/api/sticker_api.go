package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/v2_gots_sdk"
)

type StickerApi struct {
	*BaseDriverApi
}

func NewStickerApi(driverApi *BaseDriverApi) *StickerApi {
	return &StickerApi{
		BaseDriverApi: driverApi,
	}
}

type StickerQuery struct {
	Shopid int `json:"shopid" schema:"shopid" form:"shopid"`
}

func (api *StickerApi) group(ctx *gin.Context) {

	query := StickerQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {

		res, err := driverApi.ChatGetGroupSticker(1)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

type StickerBundleQuery struct {
	Shopid int    `json:"shopid" schema:"shopid" form:"shopid"`
	Id     string `json:"id" schema:"id" form:"id"`
}

func (api *StickerApi) bundle(ctx *gin.Context) {

	query := StickerBundleQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {

		res, err := driverApi.ChatGetBundleSticker(&apimodel.ChatGetBundleStickerVar{
			ID:    query.Id,
			Limit: 8,
		})
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

func (api *StickerApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "group",
		Query:        StickerQuery{},
		Response:     apimodel.ChatGetGroupStickerResp{},
	}, api.group)

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "bundle",
		Query:        StickerBundleQuery{},
		Response:     apimodel.ChatGetBundleStickerResp{},
	}, api.bundle)
}
