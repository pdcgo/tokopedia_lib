package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestProductApi(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	t.Run("test search product api", func(t *testing.T) {
		searchVar := model_public.NewSearchProductVar()
		searchVar.CategoryId = 1759
		searchVar.Identifier = "fashion-pria"

		variable := model_public.SearchProductQueryVar{
			Params: searchVar.GetQuery(),
		}

		hasil, err := api.SearchProductQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.CategoryProducts)
	})

	t.Run("test search product v4", func(t *testing.T) {
		searchVar := model_public.NewSearchProductVar()
		searchVar.Query = "ayam"
		variable := model_public.ParamsVar{
			Params: searchVar.GetQuery(),
		}

		hasil, err := api.SearchProductQueryV4(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil.Data.AceSearchProductV4.Data.Products)

		t.Run("test iterate chunk items", func(t *testing.T) {

			count := 0
			err := hasil.Data.AceSearchProductV4.Data.Products.IterateChunks(10, func(ps []*model_public.ProductSearch) error {
				count++
				assert.Equal(t, len(ps), 10)
				return nil
			})

			assert.Nil(t, err)
			assert.Equal(t, count, 10)
		})

		t.Run("test ketika ada kota count items tidak sama dengan 0", func(t *testing.T) {

			searchVar := model_public.NewSearchProductVar()
			searchVar.Query = "ayam"
			searchVar.Fcity = []string{"174,175,176,177,178,179"}

			variable := model_public.ParamsVar{
				Params: searchVar.GetQuery(),
			}

			hasil, err := api.SearchProductQueryV4(&variable)
			assert.Nil(t, err)
			assert.NotEmpty(t, hasil.Data.AceSearchProductV4.Data.Products)
		})

		t.Run("test ketika ada kota dobel count items tidak sama dengan 0", func(t *testing.T) {

			searchVar := model_public.NewSearchProductVar()
			searchVar.Query = "ayam"
			searchVar.Fcity = []string{"174,175,176,177,178,179", "174", "258,259,260,261,262,263,264,265,476,266"}

			variable := model_public.ParamsVar{
				Params: searchVar.GetQuery(),
			}

			hasil, err := api.SearchProductQueryV4(&variable)
			assert.Nil(t, err)
			assert.NotEmpty(t, hasil.Data.AceSearchProductV4.Data.Products)
		})

		t.Run("test search attach product toko spesifik", func(t *testing.T) {

			searchVar := model_public.NewSearchProductVar()
			searchVar.ShopId = 7125740
			searchVar.Source = "attach_product"
			searchVar.Rows = 5

			variable := model_public.ParamsVar{
				Params: searchVar.GetQuery(),
			}

			hasil, err := api.SearchProductQueryV4(&variable)
			assert.Nil(t, err)
			assert.NotEmpty(t, hasil.Data.AceSearchProductV4.Data.Products)

			for _, p := range hasil.Data.AceSearchProductV4.Data.Products {
				assert.Equal(t, p.Shop.ShopID, searchVar.ShopId)
			}
		})
	})
}
