package deleter_product_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func scenarioConfig1(t *testing.T) *deleter_product.DeleteConfig {
	var hasil deleter_product.DeleteConfig

	data := "{\"limit_concurent\":3,\"limit_product\":4,\"title\":[\"elemen\"],\"product_status\":\"\",\"category_id\":\"\",\"start_time\":0,\"end_time\":1686982546,\"akuns\":[{\"username\":\"pdcthoni@gmail.com\",\"password\":\"Muhammad123!`\",\"secret\":\"IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ\"}]}\r\n"

	err := json.Unmarshal([]byte(data), &hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

	return &hasil
}

func scenarioConfigWithSold(t *testing.T) *deleter_product.DeleteConfig {
	var hasil deleter_product.DeleteConfig

	data := "{\"sold_filter\": {\"min\": 10, \"max\":20 }, \"limit_concurent\":3,\"limit_product\":4,\"title\":[],\"product_status\":\"\",\"category_id\":\"\",\"start_time\":0,\"end_time\":1686982546,\"akuns\":[{\"username\":\"pdcthoni@gmail.com\",\"password\":\"Muhammad123!`\",\"secret\":\"IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ\"}]}\r\n"

	err := json.Unmarshal([]byte(data), &hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

	return &hasil
}

func scenarioConfigWithView(t *testing.T) *deleter_product.DeleteConfig {
	var hasil deleter_product.DeleteConfig

	data := "{\"view_filter\": {\"min\": 10, \"max\":20 }, \"limit_concurent\":3,\"limit_product\":4,\"title\":[],\"product_status\":\"\",\"category_id\":\"\",\"start_time\":0,\"end_time\":1686982546,\"akuns\":[{\"username\":\"pdcthoni@gmail.com\",\"password\":\"Muhammad123!`\",\"secret\":\"IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ\"}]}\r\n"

	err := json.Unmarshal([]byte(data), &hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

	return &hasil
}

func TestScenarioView(t *testing.T) {
	cfg := scenarioConfigWithView(t)

	filterFunc := cfg.GenerateFilter()

	t.Run("test masuk view", func(t *testing.T) {
		product := model.SellerProductItem{
			Name: "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			Stats: model.Stats{
				CountView: 15,
			},
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.True(t, cek, key)
	})

	t.Run("test masuk diluar view", func(t *testing.T) {
		product := model.SellerProductItem{
			Name: "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			Stats: model.Stats{
				CountView: 50,
			},
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.False(t, cek, key)
	})
}

func TestScenarioSold(t *testing.T) {
	cfg := scenarioConfigWithSold(t)

	filterFunc := cfg.GenerateFilter()

	t.Run("test masuk sold", func(t *testing.T) {
		product := model.SellerProductItem{
			Name: "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			TxStats: model.TxStats{
				Sold: 20,
			},
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.True(t, cek, key)
	})

	t.Run("test masuk diluar sold", func(t *testing.T) {
		product := model.SellerProductItem{
			Name: "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			TxStats: model.TxStats{
				Sold: 50,
			},
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.False(t, cek, key)
	})
}

func TestFilterProduct(t *testing.T) {
	cfg := scenarioConfig1(t)

	filterFunc := cfg.GenerateFilter()

	t.Run("test sold", func(t *testing.T) {
		product := model.SellerProductItem{
			Name:       "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.True(t, cek, key)
	})

	t.Run("test diluar tanggal", func(t *testing.T) {
		product := model.SellerProductItem{
			Name:       "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			CreateTime: time.Now().Add(time.Hour * 24 * 3000),
		}

		cek, key := filterFunc(&product)
		assert.False(t, cek, key)
	})

	t.Run("test didalam tanggal", func(t *testing.T) {
		product := model.SellerProductItem{
			Name:       "Elemen Tutup Rice Cooker / Element Heater Pemanas Magicom / Alumunium",
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.True(t, cek, key)
	})

	t.Run("test didalam tanggal tapi name tidak sesuai", func(t *testing.T) {
		product := model.SellerProductItem{
			Name:       "Alumunium",
			CreateTime: cfg.TStartTime.Add(time.Hour),
		}

		cek, key := filterFunc(&product)
		assert.False(t, cek, key)
	})

}
