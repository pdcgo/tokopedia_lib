package shopee_flow

import (
	"sync"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
)

type CacheApi struct {
	Api         *api.TokopediaApi
	SaveCookies func()
}

type CacheApiDriver struct {
	sync.Mutex
	Data map[string]*CacheApi
}

func NewCacheApiDriver() *CacheApiDriver {
	return &CacheApiDriver{
		Data: map[string]*CacheApi{},
	}
}

func (cache *CacheApiDriver) Get(akun *repo.AkunItem) (*api.TokopediaApi, func(), error) {
	cache.Lock()
	defer cache.Unlock()

	cacheItem := cache.Data[akun.Username]
	if cacheItem != nil {
		return cacheItem.Api, cacheItem.SaveCookies, nil
	}

	driver, err := tokopedia_lib.NewDriverAccount(akun.Username, akun.Password, akun.Secret)
	if err != nil {
		return nil, func() {}, err
	}

	api, saveCookies, err := driver.CreateApi()

	cache.Data[akun.Username] = &CacheApi{
		Api:         api,
		SaveCookies: saveCookies,
	}

	return api, saveCookies, err

}
