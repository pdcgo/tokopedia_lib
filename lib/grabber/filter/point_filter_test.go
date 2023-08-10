package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestPointFilter(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabTokopedia(func(cfg *legacy.GrabTokopedia) error {

			cfg.Point = []int{15, 20}

			return nil
		}, func(cfg *legacy.GrabTokopedia) {

			api, err := api_public.NewTokopediaApiPublic()
			assert.Nil(t, err)

			filterPoint := filter.CreatePointFilter(api, cfg)

			t.Run("test filter point ok", func(t *testing.T) {

				// shop: https://www.tokopedia.com/indomieofficial
				layout.Data.PdpGetLayout.BasicInfo.ShopID = "4544078"

				cek, reason, err := filterPoint(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter point not ok", func(t *testing.T) {

				// shop: https://www.tokopedia.com/milenialbook-1
				layout.Data.PdpGetLayout.BasicInfo.ShopID = "10255126"

				cek, reason, err := filterPoint(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter point")
				assert.Nil(t, err)
			})

		})
	})

}
