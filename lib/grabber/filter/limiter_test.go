package filter_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {

	scen := scenario.NewScenario(t)
	layout := model_public.PdpGetlayoutQueryResp{}
	pdp := model_public.PdpGetDataP2Resp{}

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

			cfg.LimitGrab = 100

			return nil
		}, func(cfg *legacy.GrabBasic) {

			limiter, addCount := filter.CreateLimiter(cfg)

			for i := 1; i <= 200; i++ {

				cek, reason, err := limiter(&layout, &pdp)

				if i <= 100 {
					assert.False(t, cek)
					assert.Empty(t, reason)
					assert.Nil(t, err)

				} else {
					assert.True(t, cek)
					assert.Equal(t, "limit reached", reason)
					assert.Equal(t, err, filter.ErrLimiterReached)
				}

				addCount()
			}

		})
	})
}
