package grabber_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestParseProductDetailParamsFromUrl(t *testing.T) {
	url := "https://www.tokopedia.com/acomeindone/acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only?extParam=cmp%3D1%26ivf%3Dfalse&src=topads"
	params, err := grabber.ParseProductDetailParamsFromUrl(url)
	assert.Nil(t, err)
	assert.Equal(t, params, &model_public.PdpGetlayoutQueryVar{
		ShopDomain: "acomeindone",
		ProductKey: "acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only",
		APIVersion: 1,
		ExtParam:   "cmp%3D1%26ivf%3Dfalse",
	})
}

func TestGenerateShopCoreInfoParamsFormUrl(t *testing.T) {
	t.Run("test parse with url", func(t *testing.T) {
		shopUrl := "https://www.tokopedia.com/acomeindone"
		params, err := grabber.GenerateShopCoreInfoParamsFormUrl(shopUrl)
		assert.Nil(t, err)
		assert.Equal(t, params, &model_public.ShopCoreInfoVar{
			ID:     0,
			Domain: "acomeindone",
		})
	})
}
