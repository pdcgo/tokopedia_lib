package group

import (
	"errors"
	"sync"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

var ErrNoDriver = errors.New("driver not found")

type DriverApi struct {
	Api    *api.TokopediaApi
	Driver *tokopedia_lib.DriverAccount
}

type DriverApiData struct {
	sync.RWMutex
	data        map[string]*DriverApi
	usernamemap map[int]string
}

func NewDriverApiData() *DriverApiData {
	return &DriverApiData{
		data:        map[string]*DriverApi{},
		usernamemap: map[int]string{},
	}
}

func (s *DriverApiData) Add(
	shopid int,
	username string,
	driver *tokopedia_lib.DriverAccount,
	dapi *api.TokopediaApi,
) {

	s.Lock()
	defer s.Unlock()

	s.usernamemap[shopid] = username
	s.data[username] = &DriverApi{
		Api:    dapi,
		Driver: driver,
	}
}

func (s *DriverApiData) Get(username string) (*DriverApi, error) {
	s.RLock()
	defer s.RUnlock()

	if dapi := s.data[username]; dapi != nil {
		return dapi, nil
	}
	return nil, ErrNoDriver
}

func (s *DriverApiData) GetByShopid(shopid int) (string, *DriverApi, error) {
	s.RLock()
	defer s.RUnlock()

	username := s.usernamemap[shopid]
	if dapi := s.data[username]; dapi != nil {
		return username, dapi, nil
	}
	return username, nil, ErrNoDriver
}
