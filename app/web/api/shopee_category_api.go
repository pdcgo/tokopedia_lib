package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/go_v2_shopeelib/controller"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type ShopeeCategoryApi struct {
	base repo.BaseInterface
}

func (c *ShopeeCategoryApi) extractCategory(
	tokpedCategories legacy.CategoryTokopedia,
	category *model_public.Categories,
	parent *legacy.CategoryTokopediaItem,
) {

	item := legacy.CategoryTokopediaItem{
		ID:        category.ID,
		DisplayID: category.ID,
		Category:  []string{category.Name},
		Chain:     []int{category.ID},
	}

	if parent != nil {
		item.Parentid = parent.ID
		item.Category = append(parent.Category, category.Name)
		item.Chain = append(parent.Chain, category.ID)
	}

	id := strconv.Itoa(category.ID)
	tokpedCategories.SetCategory(id, item)
	for _, child := range category.Children {
		c.extractCategory(tokpedCategories, child, &item)
	}
}

func (a *ShopeeCategoryApi) UpdateCategoryTokopedia(ctx *gin.Context) {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		controller.WebResponse.SetInternalServerError(ctx, err.Error())
		return
	}

	res, err := api.CategoryAllListLite()
	if err != nil {
		controller.WebResponse.SetInternalServerError(ctx, err.Error())
		return
	}

	categories := res.Data.CategoryAllListLite.Categories
	tokpedCategories := legacy.CategoryTokopedia{}
	for _, category := range categories {
		a.extractCategory(tokpedCategories, category, nil)
	}

	err = legacy.SaveCategoryTokopedia(a.base, &tokpedCategories)
	if err != nil {
		controller.WebResponse.SetInternalServerError(ctx, err.Error())
		return
	}

	controller.WebResponse.SetSuccess(ctx, "success")
}

func RegisterShopeeCategoryApi(sdk *v2_gots_sdk.ApiSdk, base repo.BaseInterface) {
	api := ShopeeCategoryApi{
		base: base,
	}

	sdk.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "/api/updateTokpedCategories",
		Response:     controller.BaseResponse{},
	}, api.UpdateCategoryTokopedia)
}
