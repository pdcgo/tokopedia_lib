package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/query"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestOrderApi(t *testing.T) {
	tapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	payload := query.NewOrderListQuery()
	payload.SetYear(2023)

	hasil, err := tapi.OrderList(payload)
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
	assert.NotEmpty(t, hasil.Data.OrderList.List)

	t.Run("test order income detail", func(t *testing.T) {
		orderid := hasil.Data.OrderList.List[0].ID
		hasil, err := tapi.OrderIncomeDetail(orderid)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.GetSomIncomeDetail.Sections)
	})
}
