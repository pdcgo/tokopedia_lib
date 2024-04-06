package group_test

import (
	"context"
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

func TestSocketGroup(t *testing.T) {

	scen := scenario.NewScenario(t)
	scenChat := scenario.NewScenarioChat(t)
	scen.WithBase(func(dirbase string, scen *scenario.Scenario) {
		scen.WithChatSqliteDatabase(func(db *gorm.DB) {
			scenChat.WithChatEvent(func(event *common_concept.CoreEvent, summary *scenario.EventSummary) {
				scenChat.WithAccountRepo(db, func(scenChat *scenario.ScenarioChat, accountRepo *repo.AccountRepo) {

					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					sio := socketio.NewServer(nil)
					appconfig := config.AppConfig{}
					driverData := group.NewDriverApiData(sio)
					socketData := group.NewSocketData(sio)
					driverGroup := group.NewDriverGroup(driverData)
					socketGroup := group.NewSocketGroup(&appconfig, event, sio, socketData)

					account, err := accountRepo.GetAccount("test", 7125740)
					assert.Nil(t, err)

					adata := account.AccountData
					dapi, err := driverGroup.AddDriverApi(
						adata.Username,
						adata.Password,
						adata.OtpPassword,
					)
					assert.Nil(t, err)

					t.Run("test add socket", func(t *testing.T) {

						err := socketGroup.AddSocket(ctx, account.AccountData, dapi.Api)
						time.Sleep(time.Second)
						assert.Nil(t, err)
						assert.Equal(t, 1, summary.ConnectEvent)
					})

					t.Run("test with socket", func(t *testing.T) {

						err := socketGroup.WithSocket(7125740, func(socket *group.Socket) error {
							assert.NotEmpty(t, socket)
							assert.Equal(t, "pdcthoni@gmail.com", socket.Account.Username)
							return nil
						})
						assert.Nil(t, err)
					})

					t.Run("test with socket", func(t *testing.T) {

						err := socketGroup.DisconnectSocket(7125740, "test")
						time.Sleep(time.Second)
						assert.Nil(t, err)
						assert.Equal(t, 1, summary.DisconnectEvent)
					})
				})
			})
		})
	})
}
