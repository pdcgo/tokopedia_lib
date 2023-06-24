package api_test

import (
	"testing"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestSavingMapItem(t *testing.T) {

	db := scenario.GetDb()

	item := api.NewShopeeMapSuggestItem(db, &mongorepo.ProductCategoryAgg{
		ID: 123,
	})

	err := item.SetTokopediaID(1233)
	assert.Nil(t, err)
}
