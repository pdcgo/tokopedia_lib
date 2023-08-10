package filter_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestLastLoginFilterTest(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}
	now := time.Now()

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabTokopedia(func(cfg *legacy.GrabTokopedia) error {

			cfg.LastLoginActive = true
			cfg.LastLoginDays = 7

			return nil
		}, func(cfg *legacy.GrabTokopedia) {

			filterLastLogin := filter.CreateLastLoginFilter(cfg)

			t.Run("test filter last login ok", func(t *testing.T) {

				lastActive := strconv.FormatInt(now.Unix(), 10)
				pdp.Data.PdpGetData.ShopInfo.ShopLastActive = lastActive

				cek, reason, err := filterLastLogin(&layout, &pdp)
				assert.False(t, cek)
				assert.Empty(t, reason)
				assert.Nil(t, err)
			})

			t.Run("test filter last login not ok", func(t *testing.T) {

				date := now.AddDate(0, 0, -10)
				lastActive := strconv.FormatInt(date.Unix(), 10)
				pdp.Data.PdpGetData.ShopInfo.ShopLastActive = lastActive

				cek, reason, err := filterLastLogin(&layout, &pdp)
				assert.True(t, cek)
				assert.Equal(t, reason, "filter last login")
				assert.Nil(t, err)
			})

		})
	})
}
