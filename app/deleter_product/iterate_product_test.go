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

	err := deleter_product.IterateProduct(
		sellerapi,
		&deleter_product.IterateFilter{
			PageSize: 20,
		},
		func(page int, product *model.SellerProductItem, delete func() int) error {
			log.Println(product.Name)

			return nil
		},
	)

	assert.Nil(t, err)

}

func TestIterateViolation(t *testing.T) {
	sellerapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Skip()
	names := []string{}

	err := deleter_product.IterateProduct(
		sellerapi,
		&deleter_product.IterateFilter{
			PageSize: 20,
			Status:   model.ViolationStatus,
		},
		func(page int, product *model.SellerProductItem, delete func() int) error {
			names = append(names, product.Name)

			t.Log(product)

			return nil
		},
	)

	assert.NotEmpty(t, names)
	assert.Nil(t, err)
}

// func TestIterateDeleteDapetSedikit(t *testing.T) {
// 	driver, err := tokopedia_lib.NewDriverAccount("acetylenetqshop@umkmevent.my.id", "Manis123", "U4WRB7E4DSCRVXDCRJ47MZDDGE77WMW4")
// 	assert.Nil(t, err)
// 	sellerapi, save, err := driver.CreateApi()
// 	assert.Nil(t, err)
// 	defer save()

// 	err = deleter_product.IterateProduct(sellerapi, func(page int, product *model.SellerProductItem, delete func() int) error {

// 		t.Log(product.Name)

// 		return nil
// 	}, model.Filter{
// 		ID:    "status",
// 		Value: []string{string(model.ActiveStatus)},
// 	})

// 	assert.Nil(t, err)
// }
