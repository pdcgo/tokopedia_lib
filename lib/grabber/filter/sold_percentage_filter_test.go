package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestSoldPercentageFilter(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

			cfg.Prosentase = 50

			return nil
		}, func(cfg *legacy.GrabBasic) {

			filterSoldPercentage := filter.CreateSoldPercentageFilter(cfg)

			t.Run("test filter sold percentage ok", func(t *testing.T) {

				layout.Data.PdpGetLayout.BasicInfo.TxStats.TransactionSuccess = "30"
				layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold = "33"

				cek, reason, err := filterSoldPercentage(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter sold percentage not ok", func(t *testing.T) {

				layout.Data.PdpGetLayout.BasicInfo.TxStats.TransactionSuccess = "20"
				layout.Data.PdpGetLayout.BasicInfo.TxStats.CountSold = "50"

				cek, reason, err := filterSoldPercentage(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter prosentase")
				assert.Nil(t, err)
			})

		})
	})
}
