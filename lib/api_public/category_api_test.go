package api_public_test

import (
	"encoding/json"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/stretchr/testify/assert"
)

func TestCategoryAllListLite(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	hasil, err := api.CategoryAllListLite()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

func TestJarvisRecommendation(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	hasil, err := api.JarvisRecommendation("New Arrival Gamis Polos Simpel | Size S M L XL XXL | Dress Polos Gamis Jumbo BIg Size")
	t.Log(hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

func TestHeaderMainData(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	hasil, err := api.HeaderMainData()
	t.Log(hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

	categ := hasil.Data.CategoryAllListLite.Categories[0]

	assert.NotEmpty(t, categ.IconImageUrl, "categ level pertama ada png")
	assert.NotEmpty(t, categ.IsCrawlable, "categ level pertama ada field is crawlable")

	t.Run("testing get bulk category", func(t *testing.T) {
		categsData := hasil.Data.CategoryAllListLite

		hasil, err := categsData.GetBulkCats([]int{
			984,
			3882,
			2772,
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)

		for _, cat := range hasil {
			data, err := json.Marshal(cat)
			assert.Nil(t, err)
			t.Log(string(data))
		}

	})
}
