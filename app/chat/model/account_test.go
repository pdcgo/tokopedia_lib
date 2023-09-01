package model_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAccountModel(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {

			t.Run("test account model", func(t *testing.T) {

				account := model.Account{
					ID:       510,
					ShopName: "Anhar Geming",
				}
				db.Create(&account)

				res := model.Account{
					ID: account.ID,
				}
				tx := db.First(&res)

				assert.Nil(t, tx.Error)
				assert.Equal(t, res.ID, account.ID)
				assert.Equal(t, res.ShopName, account.ShopName)
			})

			t.Run("test account data model", func(t *testing.T) {

				group := model.Group{
					Name: "cokro",
				}
				account := model.Account{
					ID:       830,
					ShopName: "Budi Pangestu",
				}
				accountData := model.AccountData{
					ID:       250,
					Username: "Sopyan Randu",
					Account:  account,
					Groups:   []model.Group{group},
				}
				accountData.AppendGroup(model.Group{
					Name: "cokro",
				})
				db.Create(&accountData)

				res := model.AccountData{
					ID: accountData.ID,
				}
				tx := db.First(&res)

				assert.Nil(t, tx.Error)
				assert.Equal(t, res.ID, accountData.ID)
				assert.Equal(t, res.Username, accountData.Username)

				t.Run("get account data from group", func(t *testing.T) {

					res := model.Group{
						Name: group.Name,
					}
					tx := db.Preload("AccountDatas").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.Name, group.Name)
					assert.NotEmpty(t, res.AccountDatas)
					assert.Equal(t, res.AccountDatas[0].ID, accountData.ID)
				})

				t.Run("get account data from account", func(t *testing.T) {

					res := model.Account{
						ID: account.ID,
					}
					tx := db.Preload("AccountData").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.ID, account.ID)
					assert.NotEmpty(t, res.AccountData)
					assert.Equal(t, res.AccountData.ShopID, account.ID)

				})
			})

		})
	})
}
