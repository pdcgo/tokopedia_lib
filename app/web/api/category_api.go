package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/v2_gots_sdk"
)

type baseInterface interface {
	Path(item ...string) string
}

type CategoryApi struct {
	fname string
	base  baseInterface
}

func (catapi *CategoryApi) GetListCateg(c *gin.Context) {
	fname := catapi.base.Path(catapi.fname)
	var hasil api.CategoryAllListLiteRes
	data, _ := os.ReadFile(fname)

	json.Unmarshal(data, &hasil)

	c.JSON(http.StatusOK, &hasil)

}

type UpdateTopedCategoryPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Secret   string `json:"secret" form:"secret"`
}

func (catapi *CategoryApi) UpdateCateg(c *gin.Context) {
	var payload UpdateTopedCategoryPayload
	c.BindJSON(&payload)
	go func() {
		driver, _ := tokopedia_lib.NewDriverAccount(payload.Username, payload.Password, payload.Secret)
		sellerApi, saveSession, _ := driver.CreateApi()
		defer saveSession()

		data, err := sellerApi.CategoryAllListLite()
		if err != nil {
			pdc_common.ReportError(err)
		}

		fname := catapi.base.Path(catapi.fname)
		f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			pdc_common.ReportError(err)
		}
		defer f.Close()

		err = json.NewEncoder(f).Encode(&data)
		if err != nil {
			pdc_common.ReportError(err)
		}

	}()

	c.JSON(http.StatusOK, &Response{})

}

func RegisterCategoryApi(grp *v2_gots_sdk.SdkGroup, base baseInterface) {
	catapi := CategoryApi{
		fname: "tokopedia_categories.json",
		base:  base,
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
		Payload:      UpdateTopedCategoryPayload{},
		Response:     Response{},
	}, catapi.UpdateCateg)

}
