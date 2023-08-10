package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestProductDetailApi(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	t.Run("test api layout query", func(t *testing.T) {

		variable := model_public.PdpGetlayoutQueryVar{
			ShopDomain: "toko-cctv-1",
			ProductKey: "aevision-monitor-led-pc-19-inch-garansi-resmi",
			LayoutID:   "",
			APIVersion: 1,
			// Tokonow: model_public.Tokonow{
			// 	ShopID:      "0",
			// 	WhID:        "0",
			// 	ServiceType: "",
			// },
			// UserLocation: &model_public.UserLocation{
			// 	CityID:     "",
			// 	AddressID:  "",
			// 	DistrictID: "",
			// 	PostalCode: "",
			// 	Latlon:     "",
			// },
			ExtParam: "",
		}
		hasil, err := api.PdpGetlayoutQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api pdp data p2", func(t *testing.T) {

		variable := model_public.PdpGetDataP2Var{
			Affiliate:  nil,
			ProductID:  "2873462702",
			PdpSession: "{\"sid\":6218809,\"sd\":\"toko-cctv-1\",\"cat\":{\"id\":3958},\"cp\":{\"lr\":{}},\"opr\":755000,\"pr\":755000,\"whid\":6476614,\"mo\":1,\"pn\":\"AEVISION MONITOR LED PC 19 Inch GARANSI RESMI\",\"purl\":\"https://www.tokopedia.com/toko-cctv-1/aevision-monitor-led-pc-19-inch-garansi-resmi\",\"st\":5,\"cn\":\"new\",\"li\":1,\"ln\":\"Default Layout Desktop\",\"w\":7,\"sf\":{},\"nid\":3,\"stat\":{\"cv\":10917,\"cr\":33,\"ct\":5,\"r\":4.8,\"cs\":78,\"mcs\":\"70+\"},\"fst\":[{\"FSID\":0,\"PartnerName\":\"\",\"FSType\":0,\"ShopID\":0}],\"upsn\":\"NON_SUBSCRIBER\",\"v\":1,\"pi\":2873462702,\"pse\":1,\"ps\":\"ACTIVE\",\"fc\":[\"new_variant_options\"],\"cui\":{}}",
			DeviceID:   "",
			// UserLocation: model_public.UserLocation{
			// 	CityID:     "176",
			// 	AddressID:  "0",
			// 	DistrictID: "2274",
			// 	PostalCode: "",
			// },
			// Tokonow: model_public.Tokonow{
			// 	ShopID:      "0",
			// 	WhID:        "0",
			// 	ServiceType: "",
			// },
		}
		hasil, err := api.PdpGetDataP2(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api pdp shop note", func(t *testing.T) {

		variable := model_public.ShopIdVar{
			ShopID: "6218809",
		}
		hasil, err := api.PdpShopNote(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api product rating and topics", func(t *testing.T) {

		variable := model_public.ProductIdVar{
			ProductId: "2873462702",
		}
		hasil, err := api.ProductRatingandTopics(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api pdp get review image query", func(t *testing.T) {

		variable := model_public.PdpGetReiewImageQueryVar{
			ProductID: "2873462702",
			Page:      1,
			Limit:     15,
		}
		hasil, err := api.PdpGetReiewImageQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api product review list", func(t *testing.T) {

		variable := model_public.ProductReviewListVar{
			ProductID: "2873462702",
			Page:      1,
			Limit:     15,
			SortBy:    "create_time desc",
			FilterBy:  "",
		}
		hasil, err := api.ProductReviewList(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test api recom widget", func(t *testing.T) {

		variable := model_public.RecomWidgetVar{
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
	})

	t.Run("test api layout query batch", func(t *testing.T) {

		variables := []*model_public.PdpGetlayoutQueryVar{
			{
				ShopDomain: "toko-cctv-1",
				ProductKey: "aevision-monitor-led-pc-19-inch-garansi-resmi",
				APIVersion: 1,
			},
			{
				ShopDomain: "wonderlandvalen",
				ProductKey: "baseus-bowie-d05-headphone-wireless-dual-connection-low-latency-bt5-3",
				APIVersion: 1,
			},
			{
				ShopDomain: "hebohstore",
				ProductKey: "keyboard-piano-angelet-xts-690-original",
				APIVersion: 1,
			},
		}

		expectIds := []string{"2873462702", "9180330473", "740565721"}

		hasil, err := api.PdpGetlayoutQueryBatch(variables)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)

		for ind, v := range hasil {
			assert.Equal(t, expectIds[ind], v.Data.PdpGetLayout.BasicInfo.ID)
		}
	})

	t.Run("test dengan url", func(t *testing.T) {
		uri := "https://www.tokopedia.com/botolminumviral/botol-minum-anak-sedotan-tali-anti-tumpah-henoor-2in1-aesthetic-viral-yellow-80292?source=homepage.left_carousel.0.280151"
		hasil, err := api.PdpGetlayoutQueryFromUrl(uri)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)

	})

}

func TestPdpGetDataP2(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.PdpGetDataP2Var{
		Affiliate:  nil,
		ProductID:  "2873462702",
		PdpSession: "{\"sid\":6218809,\"sd\":\"toko-cctv-1\",\"cat\":{\"id\":3958},\"cp\":{\"lr\":{}},\"opr\":755000,\"pr\":755000,\"whid\":6476614,\"mo\":1,\"pn\":\"AEVISION MONITOR LED PC 19 Inch GARANSI RESMI\",\"purl\":\"https://www.tokopedia.com/toko-cctv-1/aevision-monitor-led-pc-19-inch-garansi-resmi\",\"st\":5,\"cn\":\"new\",\"li\":1,\"ln\":\"Default Layout Desktop\",\"w\":7,\"sf\":{},\"nid\":3,\"stat\":{\"cv\":10917,\"cr\":33,\"ct\":5,\"r\":4.8,\"cs\":78,\"mcs\":\"70+\"},\"fst\":[{\"FSID\":0,\"PartnerName\":\"\",\"FSType\":0,\"ShopID\":0}],\"upsn\":\"NON_SUBSCRIBER\",\"v\":1,\"pi\":2873462702,\"pse\":1,\"ps\":\"ACTIVE\",\"fc\":[\"new_variant_options\"],\"cui\":{}}",
		DeviceID:   "",
		// UserLocation: model_public.UserLocation{
		// 	CityID:     "176",
		// 	AddressID:  "0",
		// 	DistrictID: "2274",
		// 	PostalCode: "",
		// },
		// Tokonow: model_public.Tokonow{
		// 	ShopID:      "0",
		// 	WhID:        "0",
		// 	ServiceType: "",
		// },
	}
	hasil, err := api.PdpGetDataP2(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestPdpShopNote(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.ShopIdVar{
		ShopID: "6218809",
	}
	hasil, err := api.PdpShopNote(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestProductRatingandTopics(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.ProductIdVar{
		ProductId: "2873462702",
	}
	hasil, err := api.ProductRatingandTopics(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestPdpGetReiewImageQuery(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.PdpGetReiewImageQueryVar{
		ProductID: "2873462702",
		Page:      1,
		Limit:     15,
	}
	hasil, err := api.PdpGetReiewImageQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}

func TestProductReviewList(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.ProductReviewListVar{
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

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.RecomWidgetVar{
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
