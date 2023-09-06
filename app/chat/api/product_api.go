package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/v2_gots_sdk"
)

type ProductApi struct {
	BaseApi
	publicApi *api_public.TokopediaApiPublic
}

func NewProductApi(publicApi *api_public.TokopediaApiPublic) *ProductApi {

	return &ProductApi{
		publicApi: publicApi,
	}
}

type ProductQuery struct {
	Shopid int `json:"shopid" schema:"shopid" form:"shopid"`
}

func (api *ProductApi) chatSearch(ctx *gin.Context) {

	query := ProductQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := model_public.SearchProductVar{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	res, err := api.publicApi.SearchProductQueryV4(&model_public.ParamsVar{
		Params: payload.GetQuery(),
	})

	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *ProductApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "chat_search",
		Query:        ProductQuery{},
		Payload:      model_public.SearchProductVar{},
		Response:     model_public.SearchProductQueryV4Resp{},
	}, api.chatSearch)
}
