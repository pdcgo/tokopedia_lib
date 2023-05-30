package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestCategoryAllListLite(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	hasil, err := api.CategoryAllListLite()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

func TestJarvisRecommendation(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	variable := model.JarvisRecommendationVar{
		ProductName: "sepatu",
	}
	hasil, err := api.JarvisRecommendation(&variable)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}
