package grabber_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/zeebo/assert"
)

func TestShopGrab(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	grab := grabber.ShopListGrabber{
		ShopGrabber: grabber.ShopGrabber{
			Api: api,
		},
		Shops: []string{
			"logitech-g",
			"https://www.tokopedia.com/jbl-official",
			"https://www.tokopedia.com/ellipses",
		},
	}

	grab.Run()

}
