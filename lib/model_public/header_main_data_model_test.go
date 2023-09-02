package model_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestHeaderMainDataModel(t *testing.T) {

	t.Run("test categories data model", func(t *testing.T) {

		categories := model_public.CategoriesData{
			{
				BaseCategories: &model_public.BaseCategories{
					ID:   1,
					Name: "cat1",
					URL:  "https://www.tokopedia.com/1",
				},
				Children: []*model_public.Categories{
					{
						BaseCategories: &model_public.BaseCategories{
							ID:   11,
							Name: "cat11",
							URL:  "https://www.tokopedia.com/11",
						},
					},
				},
			},
			{
				BaseCategories: &model_public.BaseCategories{
					ID:   2,
					Name: "cat2",
					URL:  "https://www.tokopedia.com/2",
				},
			},
			{
				BaseCategories: &model_public.BaseCategories{
					ID:   3,
					Name: "cat2",
					URL:  "https://www.tokopedia.com/3",
				},
			},
		}

		t.Run("test categories data model iterate", func(t *testing.T) {

			count := 0
			err := categories.Iterate(func(parents []*model_public.Categories, category *model_public.Categories) (stop bool, err error) {
				count++
				return false, nil
			})

			assert.Nil(t, err)
			assert.Equal(t, 4, count)
		})

		t.Run("test categories data model get category from url", func(t *testing.T) {

			category, err := categories.GetCategoryByUrl("https://www.tokopedia.com/2")
			assert.Nil(t, err)
			assert.Equal(t, 2, category.ID)
		})

		t.Run("test categories data model get category children from url", func(t *testing.T) {

			category, err := categories.GetCategoryByUrl("https://www.tokopedia.com/11")
			assert.Nil(t, err)
			assert.Equal(t, 11, category.ID)
		})

	})

}
