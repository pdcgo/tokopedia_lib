package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestFilterDiscount(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithMarkupConfig(func(cfg *legacy.LegacyMarkupConfig) error {

			cfg.UsePriceDiscount = true

			return nil
		}, func(cfg *legacy.LegacyMarkupConfig) {

			filterDiscount := filter.CreateFilterDiscount(cfg)

			t.Run("test filter discount ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Campaign.PercentageAmount = 0
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterDiscount(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter discount not ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Campaign.PercentageAmount = 10
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterDiscount(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter discount")
				assert.Nil(t, err)
			})

		})
	})

}
