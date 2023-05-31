package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestDiscussionDataByProductId(t *testing.T) {
	api := api_public.NewTokopediaApiPublic()

	variable := model_public.DiscussionDataProductByIDVar{
		ProductID: "4991611086",
		ShopID:    "13033117",
		Page:      1,
		Limit:     10,
		SortBy:    "time",
		Category:  "",
	}

	hasil, err := api.DiscussionDataByProductID(&variable)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
}
