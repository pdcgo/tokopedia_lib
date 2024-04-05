package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type StickerApi struct {
	BaseApi
	driverGroup *group.DriverGroup
}

func NewStickerApi(driverGroup *group.DriverGroup) *StickerApi {
	return &StickerApi{
		driverGroup: driverGroup,
	}
}

func (api *StickerApi) group(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.driverGroup.WithDriverApi(query.Shopid, func(dapi *group.DriverApi) error {
		res, err := dapi.Api.ChatGetGroupSticker(1)
		if err == nil {
			ctx.JSON(http.StatusOK, res)
		}
		return err
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

type StickerBundleQuery struct {
	*BaseQuery
	Id string `json:"id" schema:"id" form:"id"`
}

func (api *StickerApi) bundle(ctx *gin.Context) {

	query := StickerBundleQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.driverGroup.WithDriverApi(query.Shopid, func(dapi *group.DriverApi) error {
		res, err := dapi.Api.ChatGetBundleSticker(&apimodel.ChatGetBundleStickerVar{
			ID:    query.Id,
			Limit: 8,
		})
		if err == nil {
			ctx.JSON(http.StatusOK, res)
		}
		return err
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

func (api *StickerApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "group",
		Query:        BaseQuery{},
		Response:     apimodel.ChatGetGroupStickerResp{},
	}, api.group)

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "bundle",
		Query:        StickerBundleQuery{},
		Response:     apimodel.ChatGetBundleStickerResp{},
	}, api.bundle)
}
