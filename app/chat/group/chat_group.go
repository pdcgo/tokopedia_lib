package group

import (
	"context"
	"errors"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/rs/zerolog"
	"nhooyr.io/websocket"
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

func (g *ChatGroup) reportErr(err error, ev string, data any) error {
	return pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
		return event.Str("event", ev).Interface("data", data)
	})
}

func (g *ChatGroup) Connect(groupName string) {

	if groupName == g.initConfig.ActiveGroup {
		return
	}
	g.initConfig.ActiveGroup = groupName

	g.connectCancel()
	g.driverGroup.Reset()

	g.connectCtx, g.connectCancel = context.WithCancel(context.Background())

	err := g.accountRepo.IterateGroupAccount(groupName, func(account model.AccountData) error {
		select {

		case <-g.connectCtx.Done():
			return g.connectCtx.Err()

		default:
			err := g.driverGroup.AddDriver(account.Username, account.Password, account.OtpPassword)
			if err != nil {
				g.reportErr(err, "connect", map[string]string{
					"group_name": groupName,
					"username":   account.Username,
					"on":         "get_driver",
				})
				return nil
			}

			err = g.driverGroup.WithDriverApi(account.Username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
				return g.socketGroup.AddSocket(g.connectCtx, &account, api)
			})
			if err != nil {
				g.reportErr(err, "connect", map[string]string{
					"group_name": groupName,
					"username":   account.Username,
					"on":         "get_socket",
				})
				return nil
			}

			return nil
		}
	})

	if err != nil {
		g.reportErr(err, "connect", map[string]string{"group_name": groupName})
	}
}

func (g *ChatGroup) Reconnect(shopid int) error {

	account, err := g.accountRepo.GetChatAccount(g.initConfig.ActiveGroup, shopid)
	if err != nil {
		return err
	}

	g.reconnectLock.Lock()
	defer g.reconnectLock.Unlock()

	// disconnect socket if exist
	username := account.GetUsername()
	err = g.socketGroup.WithSocket(username, func(sc *chat.SocketClient) error {

		g.sio.BroadcastToNamespace("", "disconnected_event", &sio_event.SocketDisconnectedEvent{
			Shopid: int(sc.Api.AuthenticatedData.UserShopInfo.Info.ShopID),
		})

		err := sc.Con.Close(websocket.StatusNormalClosure, "reconnect")
		return err
	})
	if err != nil && !errors.Is(err, ErrNoSocket) {
		return err
	}

	return g.driverGroup.WithDriverApi(username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
		return g.socketGroup.AddSocket(g.connectCtx, account.AccountData, api)
	})
}
