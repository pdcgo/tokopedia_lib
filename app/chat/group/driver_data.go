package group

import (
	"errors"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

var ErrNoDriver = errors.New("driver not found")

type DriverApi struct {
	Api    *api.TokopediaApi
	Driver *tokopedia_lib.DriverAccount
}

func (d *DriverApi) GetUsername() string {
	return d.Api.AuthenticatedData.UserShopInfo.Info.ShopDomain
}

type DriverApiData struct {
	sync.RWMutex
	data map[int]*DriverApi
	sio  *socketio.Server
}

func NewDriverApiData(sio *socketio.Server) *DriverApiData {
	return &DriverApiData{
		data: map[int]*DriverApi{},
		sio:  sio,
	}
}

func (s *DriverApiData) Add(driver *tokopedia_lib.DriverAccount, tapi *api.TokopediaApi) *DriverApi {

	s.Lock()
	defer s.Unlock()

	dapi := DriverApi{
		Api:    tapi,
		Driver: driver,
	}

	shopid := tapi.AuthenticatedData.UserShopInfo.Info.ShopID
	s.data[int(shopid)] = &dapi
	return &dapi
}

func (s *DriverApiData) Get(shopid int) (*DriverApi, error) {
	s.RLock()
	defer s.RUnlock()

	if dapi := s.data[shopid]; dapi != nil {
		return dapi, nil
	}

	event := sio_event.NewSocketDisconnectedEvent(shopid)
	s.sio.BroadcastToNamespace("", "disconnected_event", event)
	return nil, ErrNoDriver
}
