package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/v2_gots_sdk"
)

type CategoryDumpApi struct {
	base   baseInterface
	pubapi *api_public.TokopediaApiPublic
}

func NewCategoryDumpApi(base baseInterface, pubapi *api_public.TokopediaApiPublic) *CategoryDumpApi {
	return &CategoryDumpApi{
		base:   base,
		pubapi: pubapi,
	}
}

type CategoryDumpHandler func(parent, category *model_public.Categories) error

func IterateNestedCategory(parent *model_public.Categories, categories []*model_public.Categories, handler CategoryDumpHandler) error {

	for _, category := range categories {
		if parent != nil {

			err := handler(parent, category)
			if err != nil {
				return err
			}

			err = IterateNestedCategory(parent, category.Children, handler)
			if err != nil {
				return err
			}

		} else {

			err := handler(&model_public.Categories{}, category)
			if err != nil {
				return err
			}

			err = IterateNestedCategory(category, category.Children, handler)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (api *CategoryDumpApi) IterateCategory(handler CategoryDumpHandler) error {

	res, err := api.pubapi.HeaderMainData()
	if err != nil {
		return err
	}

	categories := res.Data.CategoryAllListLite.Categories
	err = IterateNestedCategory(nil, categories, handler)

	return err
}

type DumpCategoryResponse struct {
	ErrCode int    `json:"errcode"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

func (api *CategoryDumpApi) DumpCategory(c *gin.Context) {

	items := []*csv.CategoryCsv{}
	err := api.IterateCategory(func(parent, category *model_public.Categories) error {

		item := csv.NewCategoryCsv(parent, category)
		items = append(items, item)

		err := csv.SaveCategoryCsv(api.base, items)
		return err
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, DumpCategoryResponse{
			ErrCode: 500,
			Message: err.Error(),
			Status:  "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, DumpCategoryResponse{
		ErrCode: 200,
		Message: "success",
		Status:  "OK",
	})
}

func (api *CategoryDumpApi) RegisterApi(grp *v2_gots_sdk.SdkGroup) {

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "category_dump",
		Response:     DumpCategoryResponse{},
	}, api.DumpCategory)
}
