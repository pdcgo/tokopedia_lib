package model_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGroupModel(t *testing.T) {

	scen := scenario.NewScenario(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {

			t.Run("test group model", func(t *testing.T) {

				accountData := model.AccountData{
					ID:       110,
					Username: "Udin Pedeih",
				}
				group := model.Group{
					Name:         "trunopo",
					AccountDatas: []model.AccountData{accountData},
				}
				db.Create(&group)

				res := model.Group{
					Name: group.Name,
				}
				tx := db.First(&res)

				assert.Nil(t, tx.Error)
				assert.Equal(t, res.Name, group.Name)

				t.Run("get from account data", func(t *testing.T) {

					res := model.AccountData{
						ID: accountData.ID,
					}
					tx := db.Preload("Groups").First(&res)

					assert.Nil(t, tx.Error)
					assert.Equal(t, res.ID, accountData.ID)
					assert.NotEmpty(t, res.Groups)
					assert.Equal(t, res.Groups[0].Name, group.Name)
				})
			})

		})
	})
}
