package group

import (
	"errors"
	"sync"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type DriverGroup struct {
	sync.RWMutex
	data map[string]*tokopedia_lib.DriverAccount
}

func NewDriverGroup() *DriverGroup {
	return &DriverGroup{
		data: map[string]*tokopedia_lib.DriverAccount{},
	}
}

func (g *DriverGroup) AddDriver(username string, password string, secret string) error {
	g.Lock()
	defer g.Unlock()

	driver, err := tokopedia_lib.NewDriverAccount(username, password, secret)
	if err != nil {
		return err
	}

	g.data[username] = driver
	return nil
}

var ErrNoDriver = errors.New("driver not found")

type DriverApiHandler func(api *api.TokopediaApi) error

func (g *DriverGroup) WithDriverApi(username string, handler DriverApiHandler) error {
	g.RLock()
	defer g.RUnlock()

	driver := g.data[username]
	if driver == nil {
		return ErrNoDriver
	}

	api, saveSession, err := driver.CreateApi()
	if err != nil {
		return err
	}
	defer saveSession()

	err = handler(api)
	return err
}

func (g *DriverGroup) Reset() {
	g.Lock()
	defer g.Unlock()

	g.data = map[string]*tokopedia_lib.DriverAccount{}
}
