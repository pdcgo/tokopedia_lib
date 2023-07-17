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
		Response:     &EtalaseListMapRes{},
	}, api.ListEtalase)

	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "add",
		Payload:      &services.EtalasePayload{},
		Response:     &Response{},
	}, api.AddEtalase)

	gr.Register(&v2_gots_sdk.Api{
		Method:       http.MethodDelete,
		RelativePath: "",
		Query:        &DeleteEtalaseQuery{},
	}, api.DeleteEtalase)

}

type EtalaseListMapRes struct {
	Data []*services.EtalasePayload `json:"data"`
}

func (api *EtalaseMapApi) ListEtalase(ctx *gin.Context) {
	hasil := EtalaseListMapRes{
		Data: []*services.EtalasePayload{},
	}
	data, _ := api.service.ListEtalase()
	hasil.Data = data

	ctx.JSON(http.StatusOK, &hasil)
}

func (api *EtalaseMapApi) AddEtalase(ctx *gin.Context) {
	var payload services.EtalasePayload

	ctx.BindJSON(&payload)

	api.service.AddMap(&payload)
	ctx.JSON(http.StatusOK, &Response{})
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
