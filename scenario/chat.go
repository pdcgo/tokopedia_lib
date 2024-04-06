package scenario

import (
	"testing"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type ScenarioChat struct {
	t *testing.T
}

func NewScenarioChat(t *testing.T) *ScenarioChat {
	scen := ScenarioChat{
		t: t,
	}

	return &scen
}

type EventSummary struct {
	ConnectEvent    int
	DisconnectEvent int
}

func (c *ScenarioChat) WithChatEvent(handler func(event *common_concept.CoreEvent, summary *EventSummary)) {

	event := common_concept.NewCoreEvent()
	summary := EventSummary{}
	go func() {
		for event := range event.GetEvent() {
			switch event.(type) {

			case *sio_event.SocketConnectEvent:
				summary.ConnectEvent++

			case *sio_event.SocketDisconnectedEvent:
				summary.DisconnectEvent++
			}
		}
	}()

	handler(event, &summary)
}

func (c *ScenarioChat) WithAccountRepo(db *gorm.DB, handler func(scenChat *ScenarioChat, accountRepo *repo.AccountRepo)) {
	db.AutoMigrate(
		model.Account{},
		model.AccountData{},
		model.Group{},
	)

	account := model.AccountData{
		Username:    "pdcthoni@gmail.com",
		Password:    "SilentIsMyMantra",
		OtpPassword: "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ",
		ShopID:      7125740,
		Account: model.Account{
			ID: 7125740,
		},
	}

	accountRepo := repo.NewAccountRepo(db)
	err := accountRepo.AddAccountData("test", &account)
	assert.Nil(c.t, err)

	handler(c, accountRepo)
}
