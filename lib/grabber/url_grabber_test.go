package grabber_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/zeebo/assert"
)

func TestUrlGrabber(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	grab := grabber.UrlGrabber{
		Api: api,
		Urls: []string{
			"https://www.tokopedia.com/acomeindone/acome-keyboard-mouse-wireless-portable-1600dpi-silikon-akm2000-keyboard-only?extParam=cmp%3D1%26ivf%3Dfalse&src=topads",
			"https://www.tokopedia.com/gojeteindonesia/jete-x-mouse-gaming-msx1-rgb-wired-6-programmable-buttons-original?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D3474367",
			"https://www.tokopedia.com/razer/razer-deathadder-essential-white-essential-gaming-mouse?extParam=ivf%3Dtrue%26src%3Dsearch%26whid%3D1778422",
		},
	}

	_, err = grab.Run()
	assert.Nil(t, err)
}
