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
	data   *SocketData
	event  *common_concept.CoreEvent
	sio    *socketio.Server
}

func NewSocketGroup(
	config *config.AppConfig,
	event *common_concept.CoreEvent,
	sio *socketio.Server,
	data *SocketData,
) *SocketGroup {

	return &SocketGroup{
		config: config,
		data:   data,
		event:  event,
		sio:    sio,
	}
}

func (g *SocketGroup) socketEventHandler(adata *model.AccountData) chat.SocketEventhandler {
	return func(socket *chat.SocketClient, event *chat.RcvEventSocket) error {
		switch data := event.Data.(type) {

		case *chat.ReaduserChat:
			event := sio_event.NewReadChatEvent(adata.ShopID, data)
			g.sio.BroadcastToNamespace("", "rcv_read_event", event)
			g.event.Emit(event)

		case *chat.RcvChat:
			event := sio_event.NewSendChatEvent(adata.ShopID, data)
			g.sio.BroadcastToNamespace("", "rcv_message", event)
			g.event.Emit(event)

		case *chat.RcvStartTyping:
			event := sio_event.NewTypingStartChatEvent(adata.ShopID, data)
			g.sio.BroadcastToNamespace("", "rcv_start_typing_event", event)

		case *chat.RcvEndTyping:
			event := sio_event.NewTypingEndChatEvent(adata.ShopID, data)
			g.sio.BroadcastToNamespace("", "rcv_end_typing_event", event)
		}

		return nil
	}
}

var disconnectErrors = []error{
	io.EOF,
}

func (g *SocketGroup) socketErrHandler(adata *model.AccountData) chat.SocketErrorhandler {
	return func(socket *chat.SocketClient, err error) bool {

		event := sio_event.NewSocketDisconnectedEvent(adata.ShopID)
		g.sio.BroadcastToNamespace("", "disconnected_event", event)

		for _, expectErr := range disconnectErrors {
			if errors.Is(err, expectErr) {
				log.Printf("[ %s ] socket disconnected - %s", adata.Username, err)
				return true
			}
		}

		log.Printf("[ %s ] socket unhandle disconnected - %s", adata.Username, err)
		return true
	}
}

func (g *SocketGroup) getSyncActive(min, max float32) time.Duration {
	rand.Seed(time.Now().Unix())
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
		select {
		case <-ctx.Done():
			break Parent

		case <-syncTimer.C:
			g.event.Emit(sio_event.NewSocketSyncEvent(shopid))
			syncTimer.Reset(g.config.GetSync())

		case <-activeTimer.C:
			g.event.Emit(sio_event.NewAccountActiveEvent(shopid))
			activeTimer.Reset(g.getSyncActive(180, 300))
		}
	}
}

func (g *SocketGroup) disconnect(shopid int) {
	event := sio_event.NewSocketDisconnectedEvent(shopid)
	g.event.Emit(event)
	g.sio.BroadcastToNamespace("", "disconnected_event", event)
}

func (g *SocketGroup) AddSocket(ctx context.Context, adata *model.AccountData, api *api.TokopediaApi) error {
	g.Lock()
	defer g.Unlock()

	oldSocket, _ := g.data.Get(adata.Username)
	if oldSocket != nil {
		g.disconnect(adata.ShopID)
		oldSocket.Con.Close(websocket.StatusNormalClosure, "renew")
	}

	socket := chat.NewSocketClient(api)
	g.data.Add(adata, socket)

	eventHandler := g.socketEventHandler(adata)
	errorHandler := g.socketErrHandler(adata)

	event := sio_event.NewSocketConnectEvent(adata.ShopID)
	g.event.Emit(event)
	g.sio.BroadcastToNamespace("", "connected_event", &event)

	go socket.Connect(ctx, eventHandler, errorHandler)
	go g.syncSocket(ctx, adata.ShopID)

	return nil
}

func (g *SocketGroup) WithSocket(username string, handler func(socket *chat.SocketClient) error) error {
	g.RLock()
	defer g.RUnlock()

	socket, err := g.data.Get(username)
	if err != nil {
		return err
	}

	return handler(socket)
}

func (g *SocketGroup) WithSocketByShopid(shopid int, handler func(username string, socket *chat.SocketClient) error) error {
	g.RLock()
	defer g.RUnlock()

	username, socket, err := g.data.GetByShopid(shopid)
	if err != nil {
		return err
	}

	return handler(username, socket)
}

func (g *SocketGroup) DisconnectSocket(shopid int, cause string) error {
	g.RLock()
	defer g.RUnlock()

	return g.WithSocketByShopid(shopid, func(username string, sc *chat.SocketClient) error {
		g.disconnect(shopid)
		sc.Con.Close(websocket.StatusNormalClosure, cause)
		return nil
	})
}
