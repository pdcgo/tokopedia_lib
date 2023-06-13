package deleter_product_test

import (
	"log"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/deleter_product"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestIterateProduct(t *testing.T) {
	sellerapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	err := deleter_product.IterateProduct(sellerapi, func(page int, product *model.SellerProductItem, delete func() int) error {
		log.Println(product.Name)

		return nil
	})

	assert.Nil(t, err)

}
