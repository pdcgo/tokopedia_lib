package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestAceSearchShop(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.ParamsVar{
		Params: "fcity=174%2C175%2C176%2C177%2C178%2C179&q=monitor&rows=30&shopname=true&srp_component_id=04.06.00.00&start=0&user_addressId=&user_cityId=176&user_districtId=2274&user_id=215038983&user_lat=&user_long=&user_postCode=&user_warehouseId=12210375&warehouses=12210375%232h%2C0%2315m",
	}

	hasil, err := api.AceSearchShop(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
