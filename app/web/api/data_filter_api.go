package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/v2_gots_sdk"
)

type baseInterface interface {
	Path(...string) string
}

type DataFilterApi struct {
	base baseInterface
}

func NewDataFilterApi(
	base baseInterface,

) *DataFilterApi {
	fapi := DataFilterApi{
		base: base,
	}

	return &fapi
}

func (fapi *DataFilterApi) RegisterApi(grp *v2_gots_sdk.SdkGroup) {

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "fcity",
		Response:     []*api_public.Fcity{},
	}, fapi.FcityData)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "categories",
		Response:     []*api_public.Fcity{},
	}, fapi.FcityData)
}

func (fapi *DataFilterApi) FcityData(ctx *gin.Context) {
	fname := fapi.base.Path("data", "fcity_tokopedia.json")
	cities, _ := api_public.GetFcity(fname)
	ctx.JSON(http.StatusOK, cities)
}
