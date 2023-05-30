package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestPdpGetlayoutQuery(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.PdpGetlayoutQueryVar{
		ShopDomain: "toko-cctv-1",
		ProductKey: "aevision-monitor-led-pc-19-inch-garansi-resmi",
		LayoutID:   "",
		APIVersion: 1,
		Tokonow: model.Tokonow{
			ShopID:      "11530573",
			WhID:        "12210375",
			ServiceType: "2h",
		},
		UserLocation: model.UserLocation{
			CityID:     "176",
			AddressID:  "0",
			DistrictID: "2274",
			PostalCode: "",
			Latlon:     "",
		},
		ExtParam: "",
	}
	hasil, err := api.PdpGetlayoutQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestPdpGetDataP2(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.PdpGetDataP2Var{
		Affiliate:  nil,
		ProductID:  "2873462702",
		PdpSession: "{\"sid\":6218809,\"sd\":\"toko-cctv-1\",\"cat\":{\"id\":3958},\"cp\":{\"lr\":{}},\"opr\":755000,\"pr\":755000,\"whid\":6476614,\"mo\":1,\"pn\":\"AEVISION MONITOR LED PC 19 Inch GARANSI RESMI\",\"purl\":\"https://www.tokopedia.com/toko-cctv-1/aevision-monitor-led-pc-19-inch-garansi-resmi\",\"st\":5,\"cn\":\"new\",\"li\":1,\"ln\":\"Default Layout Desktop\",\"w\":7,\"sf\":{},\"nid\":3,\"stat\":{\"cv\":10917,\"cr\":33,\"ct\":5,\"r\":4.8,\"cs\":78,\"mcs\":\"70+\"},\"fst\":[{\"FSID\":0,\"PartnerName\":\"\",\"FSType\":0,\"ShopID\":0}],\"upsn\":\"NON_SUBSCRIBER\",\"v\":1,\"pi\":2873462702,\"pse\":1,\"ps\":\"ACTIVE\",\"fc\":[\"new_variant_options\"],\"cui\":{}}",
		DeviceID:   "",
		UserLocation: model.UserLocation{
			CityID:     "176",
			AddressID:  "0",
			DistrictID: "2274",
			PostalCode: "",
		},
		Tokonow: model.Tokonow{
			ShopID:      "11530573",
			WhID:        "12210375",
			ServiceType: "2h",
		},
	}
	hasil, err := api.PdpGetDataP2(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestPdpShopNote(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.ShopIdVar{
		ShopID: "6218809",
	}
	hasil, err := api.PdpShopNote(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestProductRatingandTopics(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.ProductIdVar{
		ProductId: "2873462702",
	}
	hasil, err := api.ProductRatingandTopics(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestPdpGetReiewImageQuery(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.PdpGetReiewImageQueryVar{
		ProductID: "2873462702",
		Page:      1,
		Limit:     15,
	}
	hasil, err := api.PdpGetReiewImageQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestProductReviewList(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.ProductReviewListVar{
		ProductID: "2873462702",
		Page:      1,
		Limit:     15,
		SortBy:    "create_time desc",
		FilterBy:  "",
	}
	hasil, err := api.ProductReviewList(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestRecomWidget(t *testing.T) {
	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.RecomWidgetVar{
		UserID:         229210063,
		XDevice:        "desktop",
		PageName:       "pdp_3",
		Ref:            "",
		ProductIDs:     "2873462702",
		TokoNow:        false,
		CategoryIDs:    "",
		Keyword:        []interface{}{},
		LayoutPageType: "",
		QueryParam:     "user_addressId=0&user_cityId=176&user_districtId=2274&user_lat=&user_long=&user_postCode=&warehouse_ids=12210375",
	}
	hasil, err := api.RecomWidget(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
