package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestFilterSortProductQuery(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.ParamsVar{
		Params: "navsource=&q=a&source=search_product&srp_component_id=01.07.00.00&srp_page_id=&srp_page_title=&st=product&user_addressId=&user_cityId=176&user_districtId=2274&user_lat=&user_long=&user_postCode=&user_warehouseId=12210375&warehouses=12210375%232h%2C0%2315m",
	}

	hasil, err := api.FilterSortProductQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
