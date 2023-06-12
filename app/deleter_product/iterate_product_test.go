package deleter_product_test

import (
	"context"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestIterateProduct(t *testing.T) {
	sellerapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	err := deleter_product.IterateProduct(context.Background(), sellerapi)

	assert.Nil(t, err)

}
