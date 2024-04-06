package group

import (
	"context"
	"errors"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/rs/zerolog"
)

type ChatGroup struct {
	sio         *socketio.Server
	initConfig  *config.InitConfig
	accountRepo *repo.AccountRepo
	driverGroup *DriverGroup
	socketGroup *SocketGroup

	connectCtx    context.Context
	connectCancel context.CancelFunc
	reconnectLock sync.Mutex
}

func NewChatGroup(
	sio *socketio.Server,
	initConfig *config.InitConfig,
	accountRepo *repo.AccountRepo,
	driverGroup *DriverGroup,
	socketGroup *SocketGroup,
) *ChatGroup {

	return &ChatGroup{
		sio:           sio,
		initConfig:    initConfig,
		accountRepo:   accountRepo,
		driverGroup:   driverGroup,
		socketGroup:   socketGroup,
		connectCancel: func() {},
	}
}

func (g *ChatGroup) applyConnect(account *model.AccountData) error {
	disconnectEvent := sio_event.NewSocketDisconnectedEvent(account.ShopID)
	dapi, err := g.driverGroup.AddDriverApi(account.Username, account.Password, account.OtpPassword)
	if err != nil {
		g.sio.BroadcastToNamespace("", "disconnected_event", disconnectEvent)
		return err
	}

	err = g.socketGroup.AddSocket(g.connectCtx, account, dapi.Api)
	if err != nil {
		g.sio.BroadcastToNamespace("", "disconnected_event", disconnectEvent)
		return err
	}

	return nil
}

func (g *ChatGroup) Connect(groupName string) {

	data := map[string]string{
		"event":      "connect",
		"group_name": groupName,
	}

	if g.initConfig.CheckGroup(groupName) {
		return
	}
	err := g.initConfig.SetGroup(groupName)
	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
			return event.Interface("data", data)
		})
	}

	g.connectCancel()
	g.connectCtx, g.connectCancel = context.WithCancel(context.Background())

	err = g.accountRepo.IterateGroupAccount(groupName, func(account model.AccountData) error {
		data["username"] = account.Username

		select {

		case <-g.connectCtx.Done():
			return g.connectCtx.Err()

		default:
			return g.applyConnect(&account)
		}
	})

	if err != nil {
		pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
			return event.Interface("data", data)
		})
	}
}

func (g *ChatGroup) Reconnect(shopid int) error {
	g.reconnectLock.Lock()
	defer g.reconnectLock.Unlock()

	err := g.socketGroup.DisconnectSocket(shopid, "reconnect")
	if err != nil && !errors.Is(err, ErrNoSocket) {
		return err
	}

	return g.accountRepo.WithAccount(g.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
		// disconnect socket if exist
		err := g.socketGroup.DisconnectSocket(shopid, "reconnect")
		if err != nil && !errors.Is(err, ErrNoSocket) {
			return err
		}

		err = g.driverGroup.WithDriverApi(shopid, func(dapi *DriverApi) error {
			return g.socketGroup.AddSocket(g.connectCtx, account.AccountData, dapi.Api)
		})

		switch err {

		case ErrNoSocket, ErrNoDriver:
			return g.applyConnect(account.AccountData)

		default:
			return err
		}
	})
}
