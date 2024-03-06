package deleter_product_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestScenarioView(t *testing.T) {
	cfg := deleter_product.TokopediaDeleteConfig{
		Title: []string{
			"ayam",
			"regex-->pisang",
		},
		SoldFilter: &deleter_product.SoldConfig{
			Min: 10,
			Max: 20,
		},
		ViewFilter: &deleter_product.ViewConfig{
			Min: 10,
			Max: 20,
		},
		PriceFilter: &deleter_product.PriceConfig{
			Min: 10000,
			Max: 20000,
		},
		TStartTime: time.Time{},
		TEndTime:   time.Now(),
	}

	t.Run("test filter title", func(t *testing.T) {

		filter := cfg.GenerateFilterTitle()
		for name, expect := range map[string]bool{
			"ayam kecap manis":   true,
			"pisang keju biadab": true,
			"nasi padang":        false,
		} {

			t.Run(fmt.Sprintf("test filter name should be %t", expect), func(t *testing.T) {
				cek := filter(&model.SellerProductItem{
					Name: name,
				})
				assert.Equal(t, cek, expect)
			})
		}
	})

	t.Run("test filter view", func(t *testing.T) {

		filter := cfg.ViewFilter.GenerateFilter()
		for expect, view := range map[bool]int{
			true:  15,
			false: 50,
		} {

			t.Run(fmt.Sprintf("test filter view should be %t", expect), func(t *testing.T) {
				cek := filter(&model.SellerProductItem{
					Stats: model.Stats{
						CountView: view,
					},
				})
				assert.Equal(t, cek, expect)
			})
		}
	})

	t.Run("test filter sold", func(t *testing.T) {

		filter := cfg.SoldFilter.GenerateFilter()
		for expect, sold := range map[bool]int{
			true:  15,
			false: 50,
		} {

			t.Run(fmt.Sprintf("test filter sold should be %t", expect), func(t *testing.T) {
				cek := filter(&model.SellerProductItem{
					TxStats: model.TxStats{
						Sold: sold,
					},
				})
				assert.Equal(t, cek, expect)
			})
		}
	})

	t.Run("test filter time", func(t *testing.T) {

		filter := cfg.GenerateFilterTime()
		for expect, ctime := range map[bool]time.Time{
			true:  time.Now().Add(-time.Hour),
			false: time.Now().Add(time.Hour),
		} {

			t.Run(fmt.Sprintf("test filter time should be %t", expect), func(t *testing.T) {
				cek := filter(&model.SellerProductItem{
					CreateTime: ctime,
				})
				assert.Equal(t, cek, expect)
			})
		}
	})

	t.Run("test filter price", func(t *testing.T) {

		filter := cfg.PriceFilter.GenerateFilter()
		for expect, price := range map[bool]int{
			true:  15000,
			false: 50000,
		} {

			t.Run(fmt.Sprintf("test filter price should be %t", expect), func(t *testing.T) {
				cek := filter(&model.SellerProductItem{
					Price: model.Price{
						Min: price,
						Max: price,
					},
				})
				assert.Equal(t, cek, expect)
			})
		}
	})

	t.Run("test filter all true", func(t *testing.T) {

		filter := cfg.GenerateFilter()
		cek, _ := filter(&model.SellerProductItem{
			Name: "mie ayam",
			Stats: model.Stats{
				CountView: 15,
			},
			TxStats: model.TxStats{
				Sold: 15,
			},
			CreateTime: time.Now().Add(-time.Hour),
			Price: model.Price{
				Min: 15000,
				Max: 15000,
			},
		})
		assert.True(t, cek)
	})
}
