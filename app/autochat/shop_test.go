package autochat_test

import (
	"errors"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/autochat"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestShop(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

		limit := 5
		config, err := autochat.NewAutochatConfig(scen)
		assert.Nil(t, err)

		t.Run("get akuns file not exist", func(t *testing.T) {

			shopdata, err := autochat.NewShopData(scen, config)
			assert.Nil(t, err)
			assert.Zero(t, len(shopdata.Data))
		})

		t.Run("save akuns", func(t *testing.T) {

			shopdata, err := autochat.NewShopData(scen, config)
			assert.Nil(t, err)

			for i := 0; i < limit; i++ {
				shopdata.Data = append(shopdata.Data, &autochat.Shop{
					ShopName: "test",
				})
			}
			err = shopdata.Save()
			assert.Nil(t, err)
		})

		t.Run("iterate shop", func(t *testing.T) {

			shopdata, err := autochat.NewShopData(scen, config)
			assert.Nil(t, err)

			count := 0
			err = shopdata.Iterate(func(shop *autochat.Shop) error {
				count++
				assert.Equal(t, shop.ShopName, "test")
				return nil
			})
			assert.Nil(t, err)
			assert.Equal(t, count, limit)
		})

		t.Run("get shop", func(t *testing.T) {

			shopdata, err := autochat.NewShopData(scen, config)
			assert.Nil(t, err)

			count := 0
			for {
				shop, err := shopdata.Get()
				if errors.Is(err, autochat.ErrNoShopMore) || errors.Is(err, autochat.ErrNoShop) {
					break
				}

				count++
				assert.Nil(t, err)
				assert.Equal(t, shop.ShopName, "test")
			}
			assert.Nil(t, err)
			assert.Equal(t, count, limit)
		})
	})
}
