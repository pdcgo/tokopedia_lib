package group

import (
	"context"
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"nhooyr.io/websocket"
)

type SocketGroup struct {
	sync.RWMutex
	config *config.AppConfig
	data   map[string]*chat.SocketClient
	event  *common_concept.CoreEvent
	sio    *socketio.Server
}

func NewSocketGroup(
	config *config.AppConfig,
	event *common_concept.CoreEvent,
	sio *socketio.Server,
) *SocketGroup {

	return &SocketGroup{
		config: config,
		data:   map[string]*chat.SocketClient{},
		event:  event,
		sio:    sio,
	}
}

func (g *SocketGroup) socketEventHandler(accountData *model.AccountData) chat.SocketEventhandler {
	return func(socket *chat.SocketClient, event *chat.RcvEventSocket) error {

		switch data := event.Data.(type) {

		case *chat.ReaduserChat:
			event := sio_event.ReadChatEvent{
				Shopid: accountData.ShopID,
				Event:  data,
			}
			g.sio.BroadcastToNamespace("", "rcv_read_event", &event)
			g.event.Emit(&event)

		case *chat.RcvChat:
			event := sio_event.SendChatEvent{
				Shopid: accountData.ShopID,
				Event:  data,
			}
			g.sio.BroadcastToNamespace("", "rcv_message", &event)
			g.event.Emit(&event)

		case *chat.RcvStartTyping:
			g.sio.BroadcastToNamespace("", "rcv_start_typing_event", &sio_event.TypingStartChatEvent{
				Shopid: accountData.ShopID,
				Event:  data,
			})

		case *chat.RcvEndTyping:
			g.sio.BroadcastToNamespace("", "rcv_end_typing_event", &sio_event.TypingEndChatEvent{
				Shopid: accountData.ShopID,
				Event:  data,
			})
		}

		return nil
	}
}

var disconnectErrors = []error{
	io.EOF,
}

func (g *SocketGroup) socketErrHandler(accountData *model.AccountData) chat.SocketErrorhandler {
	return func(socket *chat.SocketClient, err error) bool {

		g.sio.BroadcastToNamespace("", "disconnected_event", sio_event.SocketDisconnectedEvent{
			Shopid: accountData.ShopID,
		})

		for _, expectErr := range disconnectErrors {
			if errors.Is(err, expectErr) {
				log.Printf("[ %s ] socket disconnected - %s", accountData.Username, err)
				return true
			}
		}

		log.Printf("[ %s ] socket unhandle disconnected - %s", accountData.Username, err)
		return true
	}
}

func (g *SocketGroup) getSyncActive(min, max float32) time.Duration {
	r := min + rand.Float32()*(max-min)
	return time.Second * time.Duration(r)
}

func (g *SocketGroup) syncSocket(ctx context.Context, shopid int) {

	syncTimer := time.NewTimer(g.config.GetSync())
	defer syncTimer.Stop()

	activeTimer := time.NewTimer(g.getSyncActive(180, 300))
	defer syncTimer.Stop()

Parent:
	for {
		rand.Seed(time.Now().Unix())

		select {
		case <-ctx.Done():
			break Parent

		case <-syncTimer.C:
			g.event.Emit(&sio_event.SocketSyncEvent{
				Shopid: shopid,
			})
			syncTimer.Reset(g.config.GetSync())

		case <-activeTimer.C:
			g.event.Emit(&sio_event.AccountActiveEvent{
				Shopid: shopid,
			})
			activeTimer.Reset(g.getSyncActive(180, 300))
		}
	}
}

func (g *SocketGroup) disconnect(shopid int) {
	event := sio_event.SocketDisconnectedEvent{
		Shopid: shopid,
	}
	g.event.Emit(&event)
	g.sio.BroadcastToNamespace("", "disconnected_event", &event)
}

func (g *SocketGroup) AddSocket(ctx context.Context, accountData *model.AccountData, api *api.TokopediaApi) error {
	g.Lock()
	defer g.Unlock()

	oldSocket := g.data[accountData.Username]
	if oldSocket != nil {
		oldSocket.Con.Close(websocket.StatusNormalClosure, "renew")
		g.disconnect(int(api.AuthenticatedData.UserShopInfo.Info.ShopID))
	}

	socket := chat.NewSocketClient(api)
	g.data[accountData.Username] = socket

	eventHandler := g.socketEventHandler(accountData)
	errorHandler := g.socketErrHandler(accountData)

	event := sio_event.SocketConnectEvent{
		Shopid: accountData.ShopID,
	}
	g.event.Emit(&event)
	g.sio.BroadcastToNamespace("", "connected_event", &event)

	go socket.Connect(ctx, eventHandler, errorHandler)
	go g.syncSocket(ctx, accountData.ShopID)

	return nil
}

var ErrNoSocket = errors.New("socket not found")

type SocketHandler func(*chat.SocketClient) error

func (g *SocketGroup) WithSocket(username string, handler SocketHandler) error {
	g.RLock()
	defer g.RUnlock()

	socket := g.data[username]
	if socket == nil {
		return ErrNoSocket
	}

	return handler(socket)
}

func (g *SocketGroup) DisconnectSocket(username, cause string) error {
	return g.WithSocket(username, func(sc *chat.SocketClient) error {
		g.disconnect(int(sc.Api.AuthenticatedData.UserShopInfo.Info.ShopID))
		sc.Con.Close(websocket.StatusNormalClosure, cause)
		return nil
	})
}
