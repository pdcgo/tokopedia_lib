package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/v2_gots_sdk"
)

type EtalaseMapApi struct {
	service *services.EtalaseMapService
}

func NewEtalaseMapApi(
	service *services.EtalaseMapService,
) *EtalaseMapApi {
	api := EtalaseMapApi{
		service: service,
	}

	return &api

}

func (api *EtalaseMapApi) RegisterApi(gr *v2_gots_sdk.SdkGroup) {
	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list",
		Response:     &ListMapEtalaseRes{},
		Query:        ListMapEtalaseQuery{},
	}, api.ListMapEtalase)

	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodDelete,
		RelativePath: "delete",
		Query:        &DeleteEtalaseQuery{},
	}, api.DeleteEtalase)

	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "update",
		Response:     &Response{},
		Payload:      []*services.EtalaseMapItem{},
	}, api.UpdateMapEtalase)

	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list_etalase",
		Response:     []*services.EtalasePayload{},
	}, func(ctx *gin.Context) {
		data, _ := api.service.ListEtalase()
		ctx.JSON(http.StatusOK, &data)
	})
}

func (api *EtalaseMapApi) UpdateMapEtalase(ctx *gin.Context) {
	var payload []*services.EtalaseMapItem
	err := ctx.BindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Msg: err.Error(),
		})
		return
	}

	errs := api.service.UpdateBulkMap(payload)

	ctx.JSON(http.StatusOK, errs)
}

type ListMapEtalaseRes struct {
	Data []*services.ShopeeEtalaseMapItem `json:"data"`
}

type ListMapEtalaseQuery struct {
	Namespace string `json:"namespace" form:"namespace"`
}

func (api *EtalaseMapApi) ListMapEtalase(ctx *gin.Context) {
	query := ListMapEtalaseQuery{}
	ctx.BindQuery(&query)

	hasil := ListMapEtalaseRes{
		Data: []*services.ShopeeEtalaseMapItem{},
	}
	data, _ := api.service.GetListMap(query.Namespace)
	hasil.Data = data

	ctx.JSON(http.StatusOK, &hasil)
}

type DeleteEtalaseQuery struct {
	Name string `json:"name" form:"name"`
}

func (api *EtalaseMapApi) DeleteEtalase(ctx *gin.Context) {
	var query DeleteEtalaseQuery

	ctx.BindQuery(&query)

	api.service.DeleteEtalase(query.Name)
	ctx.JSON(http.StatusOK, &Response{})
}
