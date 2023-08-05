package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestFilterStock(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	stockFilter := filter.CreateStockFilter(&base)
	cek, reason, err := stockFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, "", reason)
	assert.False(t, cek)

}

func TestFilterProsentage(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	prosentageFilter := filter.CreateSoldPercentageFilter(&base)
	cek, reason, err := prosentageFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterSold(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	soldFilter := filter.CreateSoldFilter(&base)
	cek, reason, err := soldFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterPoint(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	pointFilter := filter.CreatePointFilter(api, &base)
	cek, reason, err := pointFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterReview(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	lastReviewFilter := filter.CreateLastReviewFilter(api, &base)
	cek, reason, err := lastReviewFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterLasLogin(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	lastLoginFilter := filter.CreateLastLoginFilter(&base)
	cek, reason, err := lastLoginFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterBlacklistUsername(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	blacklistFilter := filter.CreateBlacklistUsernameFilter(&base)
	cek, reason, err := blacklistFilter(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)
}

func TestFilterLimiter(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	base := legacy_source.BaseConfig{
		BaseData: "../../..",
	}

	prodUrl := "https://www.tokopedia.com/baseus/baseus-true-wireless-bluetooth-earphone-mini-earbuds-tws-wm01-hitam?extParam=src%3Dmultiloc%26whid%3D4895&source=homepage.left_carousel.0.280472"
	layoutVar, err := grabber.ParseProductDetailParamsFromUrl(prodUrl)
	assert.Nil(t, err)
	layout, err := api.PdpGetlayoutQuery(layoutVar)
	assert.Nil(t, err)

	pdpVar := &model_public.PdpGetDataP2Var{
		ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
		PdpSession: layout.Data.PdpGetLayout.PdpSession,
	}
	pdp, err := api.PdpGetDataP2(pdpVar)
	assert.Nil(t, err)

	grabBasic := legacy.NewGrabBasic(&base)

	filterLimit, addCount := filter.CreateLimiter(&base)
	cek, reason, err := filterLimit(layout, pdp)
	assert.Nil(t, err)
	assert.Equal(t, reason, "")
	assert.False(t, cek)

	for i := 1; i <= grabBasic.LimitGrab; i++ {
		addCount()
	}
}
