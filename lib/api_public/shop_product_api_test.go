package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestShopProducts(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	t.Run("test ShopProduct", func(t *testing.T) {
		variable := model_public.ShopProductVar{
			Sid:            "1409816",
			Page:           1,
			PerPage:        80,
			EtalaseID:      "etalase",
			Sort:           1,
			UserDistrictID: "2274",
			UserCityID:     "176",
			UserLat:        "",
			UserLong:       "",
		}
		hasil, err := api.ShopProducts(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}
