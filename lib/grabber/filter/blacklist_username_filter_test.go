package filter_test

import (
	"strings"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestBlacklistUsernameFilter(t *testing.T) {

	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			// create txt file
			fname := base.Path("blacklist.txt")
			fdata := []string{
				"cederamata",
				"bluehat",
			}
			fbdata := []byte(strings.Join(fdata, "\n"))

			removeFile := scen.CreateFile(fbdata, fname)
			defer removeFile()

			scen.WithGrabBasic(func(cfg *legacy.GrabBasic) error {

				cfg.BlacklistUsername.Active = true
				cfg.BlacklistUsername.Tokopedia.Filename = "blacklist.txt"

				return nil
			}, func(cfg *legacy.GrabBasic) {

				blacklistFilter := filter.CreateBlacklistUsernameFilter(base, cfg)

				t.Run("test filter blacklist username filter not ok", func(t *testing.T) {

					layout := model_public.PdpGetlayoutQueryResp{}
					pdp := model_public.PdpGetDataP2Resp{}
					pdp.Data.PdpGetData.ShopInfo.ShopCore.Domain = "tokosate"

					cek, reason, err := blacklistFilter(&layout, &pdp)
					assert.False(t, cek)
					assert.Empty(t, reason)
					assert.Nil(t, err)
				})

				t.Run("test filter blacklist username filter ok", func(t *testing.T) {

					layout := model_public.PdpGetlayoutQueryResp{}
					pdp := model_public.PdpGetDataP2Resp{}
					pdp.Data.PdpGetData.ShopInfo.ShopCore.Domain = "bluehat"

					cek, reason, err := blacklistFilter(&layout, &pdp)
					assert.True(t, cek)
					assert.Equal(t, reason, "filter shopname")
					assert.Nil(t, err)
				})

			})
		})
	})
}
