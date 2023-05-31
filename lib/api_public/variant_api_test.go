package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestVarianCombinatioNQuery(t *testing.T) {

	api := api_public.NewTokopediaApiPublic()

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
