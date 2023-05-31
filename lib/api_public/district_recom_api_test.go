package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestLocDisctricRecommendation(t *testing.T) {
	api := api_public.NewTokopediaApiPublic()

	variable := model_public.LocDisctricRecommendationVar{
		Page:  "1",
		Query: "blitar",
	}
	hasil, err := api.LocDisctricRecommendation(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
