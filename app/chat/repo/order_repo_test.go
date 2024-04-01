package repo_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestOrderRepo(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {
			orderRepo := repo.NewOrderRepo(db)

			t.Run("test create order", func(t *testing.T) {
				orderid := 6669
				shopid := 6999
				err := orderRepo.CreateOrUpdateOrder(orderid, func(order *model.Order) error {
					order.ShopID = 6999
					order.StatusID = 220
					order.Status = "status"
					return nil
				})
				assert.Nil(t, err)

				order, err := orderRepo.GetOrder(orderid)
				assert.Nil(t, err)
				assert.Equal(t, order.ID, orderid)
				assert.Equal(t, order.ShopID, 6999)
				assert.Equal(t, order.Status, "status")

				t.Run("test create order item", func(t *testing.T) {
					err := orderRepo.CreateOrUpdateOrderItem(6666, func(item *model.OrderItem) error {
						item.OrderID = orderid
						item.ShopID = shopid
						item.ProductName = "product test"
						item.Price = 50000
						return nil
					})
					assert.Nil(t, err)

					order, err := orderRepo.GetOrder(orderid)
					assert.Nil(t, err)
					assert.NotEmpty(t, order.OrderItems)
				})

				t.Run("test update order", func(t *testing.T) {
					err := orderRepo.CreateOrUpdateOrder(orderid, func(order *model.Order) error {
						order.Status = "updated"
						return nil
					})
					assert.Nil(t, err)

					order, err := orderRepo.GetOrder(orderid)
					assert.Nil(t, err)
					assert.Equal(t, order.ID, orderid)
					assert.Equal(t, order.ShopID, 6999)
					assert.Equal(t, order.Status, "updated")
				})
			})

			t.Run("test paginate", func(t *testing.T) {

				res, err := orderRepo.Paginate(&repo.ListOrderFilter{
					Page: 1,
					Size: 10,
				})
				assert.Nil(t, err)
				assert.NotEmpty(t, res.Items)

				t.Run("test filter product name", func(t *testing.T) {
					t.Run("test filter ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:        1,
							Size:        10,
							ProductName: "test",
						})
						assert.Nil(t, err)
						assert.NotEmpty(t, res.Items)
					})

					t.Run("test filter not ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:        1,
							Size:        10,
							ProductName: "notok",
						})
						assert.Nil(t, err)
						assert.Empty(t, res.Items)
					})
				})

				t.Run("test filter status", func(t *testing.T) {
					t.Run("test filter ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:   1,
							Size:   10,
							Status: model.NewOrder,
						})
						assert.Nil(t, err)
						assert.NotEmpty(t, res.Items)
					})

					t.Run("test filter not ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:   1,
							Size:   10,
							Status: model.Done,
						})
						assert.Nil(t, err)
						assert.Empty(t, res.Items)
					})
				})

				t.Run("test filter price", func(t *testing.T) {
					t.Run("test filter ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:     1,
							Size:     10,
							PriceMin: 40000,
						})
						assert.Nil(t, err)
						assert.NotEmpty(t, res.Items)

						res, err = orderRepo.Paginate(&repo.ListOrderFilter{
							Page:     1,
							Size:     10,
							PriceMax: 60000,
						})
						assert.Nil(t, err)
						assert.NotEmpty(t, res.Items)
					})

					t.Run("test filter not ok", func(t *testing.T) {
						res, err := orderRepo.Paginate(&repo.ListOrderFilter{
							Page:     1,
							Size:     10,
							PriceMin: 60000,
						})
						assert.Nil(t, err)
						assert.Empty(t, res.Items)

						res, err = orderRepo.Paginate(&repo.ListOrderFilter{
							Page:     1,
							Size:     10,
							PriceMax: 40000,
						})
						assert.Nil(t, err)
						assert.Empty(t, res.Items)
					})
				})
			})
		})
	})
}
