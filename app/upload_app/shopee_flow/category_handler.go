package shopee_flow

import (
	"errors"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api/public_model"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

type CategoryEvent struct {
	CatID int
}

func (flow *ShopeeToTopedFlow) createCategoryHandler() uploader.UploadHandler {

	configMap := config.ShopeeMapperConfig{}
	err := flow.configrepo.GetConfig(&configMap)
	if err != nil {
		pdc_common.ReportError(err)
		panic(err)
	}

	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		var source *public_model.PublicProduct
	Parent:
		for {
			ev := <-sub.Chan
			switch event := ev.(type) {
			case *public_model.PublicProduct:
				source = event
				sub.Cancel()
				break Parent
			}
		}

		var fixid int

		if configMap.UseMapper {
			categories := source.Categories
			catid := categories[len(categories)-1].Catid
			mapitem, err := flow.mapper.GetTokopediaID(catid)
			if err != nil {
				return err
			}

			eventcore(&CategoryEvent{
				CatID: mapitem.TokopediaID,
			})

			fixid = mapitem.TokopediaID
		} else {
			title := source.Name
			catrmd, err := flow.TopedPublicApi.JarvisRecommendation(title)

			if err != nil {
				return err
			}

			if len(catrmd.Data.GetJarvisRecommendation.Categories) == 0 {
				return errors.New("recomendation category kosong untuk produk " + source.Name)
			}

			catID := catrmd.Data.GetJarvisRecommendation.Categories[0].ID

			eventcore(&CategoryEvent{
				CatID: catID,
			})

			fixid = catID
		}

		payload.Lock()
		defer payload.Unlock()
		payload.Input.Category = model.Category{
			ID: fixid,
		}

		return nil
	}

}
