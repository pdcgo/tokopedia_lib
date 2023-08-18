package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestFilterTitle(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithFilterText(func(cfg *legacy_source.FilterText) error {

			cfg.Title = []string{
				"tidak",
				"regex-->(rokok)",
			}

			return nil
		}, func(cfg *legacy_source.FilterText) {

			filterTitle := filter.CreateTitleFilter(cfg)

			t.Run("test filter title ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Name = "Window Update"
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = append(layout.Data.PdpGetLayout.Components, &productContentComponent)

				cek, reason, err := filterTitle(&layout, &pdp)
				assert.Nil(t, err)
				assert.False(t, cek, reason)
				assert.Empty(t, reason)

			})

			t.Run("test filter title not ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Name = "Window Tidak Update"
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterTitle(&layout, &pdp)
				assert.Nil(t, err)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter title")

			})

			t.Run("test filter title regex not ok", func(t *testing.T) {

				productContent := model_public.ProductContentData{}
				productContent.Name = "Rokok Marboro murah 10k/pak"
				productContentComponent := model_public.ProductContentComponent{}
				productContentComponent.Data = append(productContentComponent.Data, productContent)
				layout.Data.PdpGetLayout.Components = model_public.PDPListComponents{&productContentComponent}

				cek, reason, err := filterTitle(&layout, &pdp)
				assert.Nil(t, err)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter title")

			})
		})
	})
}
