package model_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestOrderModel(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {

			t.Run("test order model", func(t *testing.T) {

				account := model.Account{
					ID:       9899,
					ShopName: "Markas beras",
				}
				orderItem := model.OrderItem{
					ProductName: "air bersih 90ml",
				}
				orderSheet := model.OrderSheet{
					Status: "pending",
				}
				order := model.Order{
					ID:         2728,
					ShopID:     9800,
					Account:    account,
					OrderItems: []model.OrderItem{orderItem},
					OrderSheet: &orderSheet,
				}
				db.Create(&order)

				res := model.Order{
					ID: order.ID,
				}
				tx := db.First(&res)

				assert.Nil(t, tx.Error)
				assert.Equal(t, res.ID, order.ID)
				assert.Equal(t, res.ShopID, order.ShopID)

				t.Run("get from account", func(t *testing.T) {

					res := model.Account{
						ID: account.ID,
					}
					tx := db.Preload("Orders").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.ID, account.ID)
					assert.NotEmpty(t, res.Orders[0])
					assert.Equal(t, res.Orders[0].ShopID, order.ShopID)
				})

				t.Run("get from order item", func(t *testing.T) {

					res := model.OrderItem{OrderID: order.ID}
					tx := db.Preload("Order").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.OrderID, order.ID)
					assert.Equal(t, res.Order.ShopID, order.ShopID)
				})

				t.Run("get from order sheet", func(t *testing.T) {

					res := model.OrderSheet{OrderID: order.ID}
					tx := db.Preload("Order").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.OrderID, order.ID)
					assert.Equal(t, res.Order.ShopID, order.ShopID)
				})
			})

		})
	})
}
