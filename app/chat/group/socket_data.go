package group

import (
	"errors"
	"sync"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

var ErrNoSocket = errors.New("socket not found")

type SocketData struct {
	sync.Mutex
	data        map[string]*chat.SocketClient
	usernamemap map[int]string
}

func NewSocketData() *SocketData {
	return &SocketData{
		data:        map[string]*chat.SocketClient{},
		usernamemap: map[int]string{},
	}
}

func (s *SocketData) Add(adata *model.AccountData, socket *chat.SocketClient) {

	s.Lock()
	defer s.Unlock()

	s.usernamemap[adata.ShopID] = adata.Username
	s.data[adata.Username] = socket
}

func (s *SocketData) Get(username string) (*chat.SocketClient, error) {
	if socket := s.data[username]; socket != nil {
		return socket, nil
	}
	return nil, ErrNoSocket
}

func (s *SocketData) GetByShopid(shopid int) (string, *chat.SocketClient, error) {
	username := s.usernamemap[shopid]
	if socket := s.data[username]; socket != nil {
		return username, socket, nil
	}
	return username, nil, ErrNoSocket
}
