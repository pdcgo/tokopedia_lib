package category_mapper

import (
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

type ItemMap interface {
	GetName() string
	SetTokopediaID(categid int)
}

type Mapper struct {
	limitGuard chan int
	papi       *api_public.TokopediaApiPublic
}

func NewMapper(papi *api_public.TokopediaApiPublic) *Mapper {
	return &Mapper{
		limitGuard: make(chan int),
		papi:       papi,
	}
}

func (mapi *Mapper) RunMapper(datas []ItemMap) {
	var wg sync.WaitGroup
	for _, item := range datas {
		mapi.limitGuard <- 1

		shopee := item

		wg.Add(1)
		go func() {
			defer func() {
				<-mapi.limitGuard
				wg.Done()
			}()
			name := shopee.GetName()
			res, err := mapi.papi.JarvisRecommendation(name)
			if err != nil {
				pdc_common.ReportError(err)
			}
			panjang := len(res.Data.GetJarvisRecommendation.Categories)

			if panjang == 0 {
				log.Println(name, "tidak dapat rekomendasi category tokopedia")
				return
			}

			tokped := res.Data.GetJarvisRecommendation.Categories[0]
			shopee.SetTokopediaID(tokped.ID)
		}()
	}

	wg.Wait()
}
