package group

import (
	"context"
	"errors"
	"sync"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"nhooyr.io/websocket"
)

type SocketGroup struct {
	sync.RWMutex
	data  map[string]*chat.SocketClient
	event *common_concept.CoreEvent
}

func NewSocketGroup(event *common_concept.CoreEvent) *SocketGroup {
	return &SocketGroup{
		data:  map[string]*chat.SocketClient{},
		event: event,
	}
}

type SocketChatEvent struct {
	Event    *chat.RcvEventSocket
	Username string
}

func (g *SocketGroup) socketEventHandler(username string) chat.SocketEventhandler {
	return func(socket *chat.SocketClient, event *chat.RcvEventSocket) error {

		g.event.Emit(&SocketChatEvent{
			Username: username,
			Event:    event,
		})

		return nil
	}
}

func (g *SocketGroup) AddSocket(ctx context.Context, username string, api *api.TokopediaApi) error {
	g.Lock()
	defer g.Unlock()

	oldSocket := g.data[username]
	if oldSocket != nil {
		oldSocket.Con.Close(websocket.StatusNormalClosure, "renew")
	}

	socket := chat.NewSocketClient(api)
	g.data[username] = socket

	eventHandler := g.socketEventHandler(username)
	go socket.Connect(ctx, eventHandler)

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

	err := handler(socket)
	return err
}
