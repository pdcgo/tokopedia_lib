package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestGetDrogonAnnotation(t *testing.T) {
	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	varibale := model_public.GetDrogonAnnotationVar{
		CategoryID:       2830,
		ExcludeSensitive: "true",
		ProductID:        7010430068,
		VendorName:       "merchant",
	}

	hasil, err := api.GetDrogonAnnotation(&varibale)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
