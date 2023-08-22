package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestSearchProductQuery(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.SearchProductQueryVar{
		Params:   "ob=8&page=1&rows=100&device=desktop&related=true&safe_search=false&scheme=https&user_districtId=176&user_cityId=1759&source=search&topads_bucket=true&pmin=10000&pmax=100000&rt=0%231%232%233%234%235&condition=1&sc=1759&start=1&identifier=fashion-pria&navsource=&unique_id=&shipping=%23%23%2310%2312%2313&page=1",
		AdParams: "page=1&ep=product&item=15&src=directory&device=desktop&user_id=0&minimum_item=15&start=1&no_autofill_range=5-14&dep_id=1759&ob=0&page=1",
	}

	hasil, err := api.SearchProductQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestSearchProductQueryV4(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	searchVar := model_public.NewSearchProductVar()
	searchVar.Query = "ayam"
	variable := model_public.ParamsVar{
		Params: searchVar.GetQuery(),
	}

	hasil, err := api.SearchProductQueryV4(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil.Data.AceSearchProductV4.Data.Products)

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
}
