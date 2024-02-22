package api_public_test

import (
	"testing"

	"github.com/pdcgo/common_conf/pdc_common"
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
			ProductID:  2873462702,
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
			ProductID: 2873462702,
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

		t.Run("test api layout batch urls", func(t *testing.T) {
			urls := []string{
				"https://www.tokopedia.com/alvaboard/alvaboard-kardus-kotak-pengiriman-box-pindahan-60x40x30-cm-hitam?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dtrue%26src%3Dsearch%26whid%3D14634330",
				"https://www.tokopedia.com/aceindonesiaid/ace-proclean-alat-pel-spray-putih-abu-abu?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dtrue%26src%3Dsearch%26whid%3D15257554",
				"https://www.tokopedia.com/aceindonesiaid/ace-krisbow-dispenser-sabun-cair-silver?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dtrue%26src%3Dsearch%26whid%3D15257554",
				// "https://www.tokopedia.com/aceindonesiaid/ace-ace-set-10-pcs-lap-microfiber?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D15257554",
				// "https://www.tokopedia.com/aceindonesiaid/ace-krisbow-set-2-pcs-baterai-cr2032?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D15257554",
				// "https://www.tokopedia.com/aceindonesiaid/ace-astonish-750-ml-cairan-pembersih-stainless-steel-n-clear?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D15257554",
				// "https://www.tokopedia.com/aceindonesiaid/ace-susino-53-5-cm-payung-lipat?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D15257554",
				// "https://www.tokopedia.com/informa/informa-set-8-pcs-hanger-kayu-cokelat-tua?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D14326045",
				// "https://www.tokopedia.com/alvaboard/kardus-alvaboard-xl-heavy-duty-box-besar-ukuran-70x50x40-cm-ab-xl-polos-15ad1?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dtrue%26src%3Dsearch%26whid%3D14634330",
				// "https://www.tokopedia.com/informa/informa-rak-sepatu-compact-5-tier-shoe-rack-black-42x20x79cm?extParam=fcity%3D258%2C259%2C260%2C261%2C262%2C263%2C264%2C265%2C476%2C266%26ivf%3Dfalse%26src%3Dsearch%26whid%3D14326045",
			}

			var layoutVars []*model_public.PdpGetlayoutQueryVar
			for _, url := range urls {
				layoutVar, err := model_public.NewPdpGetlayoutQueryVar(url)
				if err != nil {
					pdc_common.ReportError(err)
					return
				}

				layoutVars = append(layoutVars, layoutVar)
			}

			hasil, err := api.PdpGetlayoutQueryBatch(layoutVars)
			assert.Nil(t, err)
			assert.NotEmpty(t, hasil)

			for _, v := range hasil {
				assert.NotEmpty(t, v)
			}
		})
	})

	t.Run("test dengan url", func(t *testing.T) {
		uri := "https://www.tokopedia.com/botolminumviral/botol-minum-anak-sedotan-tali-anti-tumpah-henoor-2in1-aesthetic-viral-yellow-80292?source=homepage.left_carousel.0.280151"
		hasil, err := api.PdpGetlayoutQueryFromUrl(uri)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)

		t.Run("test parsing get component", func(t *testing.T) {
			layout := hasil.Data.PdpGetLayout

			com, err := model_public.GetComponent[model_public.ProductDetailComponent](&layout)

			assert.Nil(t, err)
			assert.NotEmpty(t, com)
			// t.Log(com)
		})

	})

}

func TestPdpGetDataP2(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.PdpGetDataP2Var{
		Affiliate:  nil,
		ProductID:  2873462702,
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
		ProductID: 2873462702,
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
