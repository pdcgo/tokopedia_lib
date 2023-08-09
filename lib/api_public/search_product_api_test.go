package api_public_test

import (
	"encoding/json"
	"fmt"
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

	variable := model_public.ParamsVar{
		Params: "ob=8&page=1&rows=100&device=desktop&related=true&safe_search=false&scheme=https&user_districtId=176&user_cityId=1759&source=search&topads_bucket=true&pmin=10000&pmax=100000&rt=0%231%232%233%234%235&condition=1&start=1&identifier=fashion-pria&navsource=&unique_id=&shipping=%23%23%2310%2312%2313&page=1&sc=1759",
	}

	hasil, err := api.SearchProductQueryV4(&variable)
	assert.Nil(t, err)
	rawResp, _ := json.MarshalIndent(hasil, "", "	")
	fmt.Println(string(rawResp))
	assert.NotEmpty(t, hasil)
}
