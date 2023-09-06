package repo_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAccountModel(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {

			accountRepo := repo.NewAccountRepo(db)
			testAccountData := model.AccountData{
				ID:          100,
				Username:    "test",
				Password:    "test",
				OtpPassword: "test",
				ShopID:      100,
				Account: model.Account{
					ID:       100,
					ShopName: "test",
				},
			}

			t.Run("test add account data", func(t *testing.T) {
				err := accountRepo.AddAccountData("test", testAccountData)
				assert.Nil(t, err)

				t.Run("test account data ada", func(t *testing.T) {
					account := model.AccountData{
						ID: testAccountData.ID,
					}
					tx := db.Preload("Account").First(&account)

					assert.Nil(t, tx.Error)
					assert.Equal(t, testAccountData.ID, account.ID)
					assert.Equal(t, testAccountData.Account.ID, account.Account.ID)
				})
			})

			t.Run("test update account", func(t *testing.T) {
				err := accountRepo.UpdateAccount(testAccountData.ShopID, func(account *model.Account) {
					account.ShopName = "updated"
				})
				assert.Nil(t, err)

				t.Run("test account data terupdate", func(t *testing.T) {
					account := model.AccountData{
						ID: testAccountData.ID,
					}
					tx := db.Preload("Account").First(&account)

					assert.Nil(t, tx.Error)
					assert.Equal(t, testAccountData.ID, account.ID)
					assert.Equal(t, testAccountData.Account.ID, account.Account.ID)
					assert.Equal(t, "updated", account.Account.ShopName)
				})
			})

		})
	})
}
