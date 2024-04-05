package group

import (
	"context"
	"errors"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

var ErrNoSocket = errors.New("socket not found")

type Socket struct {
	*chat.SocketClient
	Account *model.AccountData
	Cancel  context.CancelFunc
}

type SocketData struct {
	sync.Mutex
	data map[int]*Socket
	sio  *socketio.Server
}

func NewSocketData(sio *socketio.Server) *SocketData {
	return &SocketData{
		data: map[int]*Socket{},
		sio:  sio,
	}
}

func (s *SocketData) Add(
	account *model.AccountData,
	socketClient *chat.SocketClient,
	cancel context.CancelFunc,
) *Socket {

	s.Lock()
	defer s.Unlock()

	if oldSocket := s.data[account.ShopID]; oldSocket != nil {
		oldSocket.Cancel()
	}

	socket := Socket{
		SocketClient: socketClient,
		Account:      account,
		Cancel:       cancel,
	}

	s.data[account.ShopID] = &socket
	return &socket
}

func (s *SocketData) Get(shopid int) (*Socket, error) {
	if socket := s.data[shopid]; socket != nil {
		return socket, nil
	}

	event := sio_event.NewSocketDisconnectedEvent(shopid)
	s.sio.BroadcastToNamespace("", "disconnected_event", event)
	return nil, ErrNoSocket
}
