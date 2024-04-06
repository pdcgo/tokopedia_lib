package group_test

import (
	"testing"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestChatGroup(t *testing.T) {

	scen := scenario.NewScenario(t)
	scenChat := scenario.NewScenarioChat(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {
			scenChat.WithChatEvent(func(event *common_concept.CoreEvent, summary *scenario.EventSummary) {
				scenChat.WithAccountRepo(db, func(scenChat *scenario.ScenarioChat, accountRepo *repo.AccountRepo) {

					sio := socketio.NewServer(nil)
					appconfig := config.AppConfig{}
					initConfig := config.InitConfig{}
					driverData := group.NewDriverApiData(sio)
					socketData := group.NewSocketData(sio)
					driverGroup := group.NewDriverGroup(driverData)
					socketGroup := group.NewSocketGroup(&appconfig, event, sio, socketData)
					chatGroup := group.NewChatGroup(sio, &initConfig, accountRepo, driverGroup, socketGroup)

					account, err := accountRepo.GetAccount("test", 7125740)
					assert.Nil(t, err)

					shopid := account.AccountData.ShopID
					groupName := account.AccountData.Groups[0].Name

					t.Run("test chat group connect", func(t *testing.T) {

						chatGroup.Connect(groupName)
						time.Sleep(time.Second)

						assert.Equal(t, 1, summary.ConnectEvent)
						assert.Zero(t, summary.DisconnectEvent)
						assert.Equal(t, initConfig.ActiveGroup, groupName)

						dapi, err := driverData.Get(shopid)
						assert.Nil(t, err)
						assert.Equal(t, "schmart", dapi.GetUsername())

						socket, err := socketData.Get(shopid)
						assert.Nil(t, err)
						assert.Equal(t, "pdcthoni@gmail.com", socket.Account.Username)
					})

					t.Run("test chat group reconnect", func(t *testing.T) {

						chatGroup.Reconnect(shopid)
						time.Sleep(time.Second)

						assert.Equal(t, 2, summary.ConnectEvent)
						assert.Greater(t, summary.DisconnectEvent, 1)
						assert.Equal(t, initConfig.ActiveGroup, groupName)

						dapi, err := driverData.Get(shopid)
						assert.Nil(t, err)
						assert.Equal(t, "schmart", dapi.GetUsername())

						socket, err := socketData.Get(shopid)
						assert.Nil(t, err)
						assert.Equal(t, "pdcthoni@gmail.com", socket.Account.Username)
					})
				})
			})
		})
	})
}
