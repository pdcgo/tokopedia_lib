package iterator_test

import (
	"strings"
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestShopIterator(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)
	scen := scenario.NewScenario(t)

	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithBaseConfig(func(base *legacy_source.BaseConfig) {

			t.Run("test shop iterator no file", func(t *testing.T) {

				shops := []*model_public.ShopCoreInfoResp{}
				err := iterator.IterateShops(api, "nofile.txt", func(shopCore *model_public.ShopCoreInfoResp) error {
					shops = append(shops, shopCore)
					return nil
				})

				assert.Nil(t, err)
				assert.Equal(t, 0, len(shops))
			})

			t.Run("test shop iterator with file", func(t *testing.T) {

				// create file keyword
				fname := base.Path("shops.txt")
				fdata := []string{
					"milenialbook-1",
					"indomieofficial",
					"baseus",
				}
				fbdata := []byte(strings.Join(fdata, "\n"))

				scen.CreateFile(fbdata, fname)

				shops := []*model_public.ShopCoreInfoResp{}
				err := iterator.IterateShops(api, fname, func(shopCore *model_public.ShopCoreInfoResp) error {
					shops = append(shops, shopCore)
					return nil
				})

				shopNames := []string{}
				for _, shop := range shops {
					shopNames = append(shopNames, shop.Data.Result[0].ShopCore.Domain)
				}

				assert.Nil(t, err)
				assert.Equal(t, len(fdata), len(shops))
				assert.Equal(t, fdata, shopNames)
			})
		})
	})
}
