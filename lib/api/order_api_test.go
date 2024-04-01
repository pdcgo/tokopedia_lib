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

	var orderid int

	t.Run("test order list", func(t *testing.T) {
		payload := query.NewOrderListQuery()
		payload.SetYear(2023)

		hasil, err := tapi.OrderList(payload)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.OrderList.List)
		orderid = hasil.Data.OrderList.List[0].ID
	})

	t.Run("test order pending list", func(t *testing.T) {
		payload := query.NewOrderPendingListQuery()
		hasil, err := tapi.OrderPendingList(payload)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test order income detail", func(t *testing.T) {
		hasil, err := tapi.OrderIncomeDetail(orderid)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.GetSomIncomeDetail.Sections)
	})
}
