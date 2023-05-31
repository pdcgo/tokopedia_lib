package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestCategoryAllListLite(t *testing.T) {
	api := api_public.NewTokopediaApiPublic()

	hasil, err := api.CategoryAllListLite()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

func TestJarvisRecommendation(t *testing.T) {
	api := api_public.NewTokopediaApiPublic()

	variable := model_public.JarvisRecommendationVar{
		ProductName: "sepatu",
	}
	hasil, err := api.JarvisRecommendation(&variable)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}
