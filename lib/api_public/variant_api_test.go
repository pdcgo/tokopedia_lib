package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestVarianCombinationQuery(t *testing.T) {

	api, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	variable := model_public.VariantCategoryCombinationVar{
		CategoryID:      2830,
		AllVariants:     "true",
		ProductVariants: "",
		Type:            "add",
	}

	hasil, err := api.VariantCombinationQuery(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
