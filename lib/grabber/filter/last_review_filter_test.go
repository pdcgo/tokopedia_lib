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

func TestLastReviewFilterTest(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

			cfg.LastReviewActive = true
			cfg.LastReviewDays = 40

			return nil
		}, func(cfg *legacy.GrabBasic) {

			api, err := api_public.NewTokopediaApiPublic()
			assert.Nil(t, err)

			filterLastReview := filter.CreateLastReviewFilter(api, cfg)

			t.Run("test filter last review ok", func(t *testing.T) {

				// product: https://www.tokopedia.com/indomieofficial/10-pcs-indomie-goreng-spesial
				layout.Data.PdpGetLayout.BasicInfo.ID = 448842257

				cek, reason, err := filterLastReview(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter last review not ok", func(t *testing.T) {

				// product: https://www.tokopedia.com/milenialbook-1/novel-d-bijis
				layout.Data.PdpGetLayout.BasicInfo.ID = 1771154061

				cek, reason, err := filterLastReview(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter last review")
				assert.Nil(t, err)
			})

		})
	})
}
