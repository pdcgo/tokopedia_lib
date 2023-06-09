package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
)

type CategoryApi struct {
	repo *repo.CategoryRepo
}

func (catapi *CategoryApi) GetListCateg(c *gin.Context) {

	c.JSON(http.StatusOK, catapi.repo.Data)

}

func (catapi *CategoryApi) UpdateCateg(c *gin.Context) {
	var payload repo.UpdateTopedCategoryPayload
	c.BindJSON(&payload)
	go func() {
		err := catapi.repo.UpdateCateg(&payload)
		if err != nil {
			pdc_common.ReportError(err)
		}

	}()

	c.JSON(http.StatusOK, &Response{})

}

func RegisterCategoryApi(grp *v2_gots_sdk.SdkGroup, base repo.BaseInterface) {
	catapi := CategoryApi{
		repo: repo.NewCategoryRepo(base),
	}

	grp = grp.Group("category")

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list",
		Response:     api.CategoryAllListLiteRes{},
	}, catapi.GetListCateg)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "update_category",
		Payload:      repo.UpdateTopedCategoryPayload{},
		Response:     Response{},
	}, catapi.UpdateCateg)

}
