package scenario

import (
	"sync"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

var apiclient *api.TokopediaApi
var saveSession func()
var apiclientOnce sync.Once

func GetTokopediaApiClient() (*api.TokopediaApi, func()) {

	apiclientOnce.Do(func() {

		driver, _ := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
		// driver, _ := tokopedia_lib.NewDriverAccount("mrwicaksono690@gmail.com", "Semogaberkah", "KQ2YZZJDS2DAC7Y6K4HCTTXSS5B7N4IN")

		api, save, err := driver.CreateApi()
		if err != nil {
			panic(err)
		}

		apiclient = api
		saveSession = save
	})

	return apiclient, saveSession
}
