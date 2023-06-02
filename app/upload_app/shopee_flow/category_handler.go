package shopee_flow

import (
	"strconv"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createCategoryHandler() uploader.UploadHandler {

	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		var source *public_api.PublicProduct
	Parent:
		for {
			ev := <-sub.Chan
			switch event := ev.(type) {
			case *public_api.PublicProduct:
				source = event
				sub.Cancel()
				break Parent
			}
		}

		title := source.Name
		catrmd, err := flow.TopedPublicApi.JarvisRecommendation(title)

		if err != nil {
			return err
		}

		fixid := strconv.Itoa(catrmd.Data.GetJarvisRecommendation.Categories[0].ID)

		payload.Lock()
		defer payload.Unlock()
		payload.Input.Category = model.Category{
			ID: fixid,
		}

		return nil
	}

}
