package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestStockFilter(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

			cfg.Stock = 100

			return nil
		}, func(cfg *legacy.GrabBasic) {

			filterStock := filter.CreateStockFilter(cfg)

			t.Run("test filter stock ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Stock.Value = "2001"
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterStock(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter stock not ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Stock.Value = "36"
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterStock(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter stock")
				assert.Nil(t, err)
			})

		})
	})
}
