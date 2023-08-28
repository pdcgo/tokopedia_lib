package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestSoldFilter(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

			cfg.Penjualan = 100

			return nil
		}, func(cfg *legacy.GrabBasic) {

			filterSold := filter.CreateSoldFilter(cfg)

			t.Run("test filter sold ok", func(t *testing.T) {

				layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold = 200

				cek, reason, err := filterSold(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter sold not ok", func(t *testing.T) {

				layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold = 91

				cek, reason, err := filterSold(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter sold")
				assert.Nil(t, err)
			})

		})
	})
}
