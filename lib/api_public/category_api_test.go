package api_public_test

import (
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
