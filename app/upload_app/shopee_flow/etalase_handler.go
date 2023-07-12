package shopee_flow

import (
	"errors"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
	"gorm.io/gorm"
)

func (flow *ShopeeToTopedFlow) createEtalaseHandler() uploader.UploadHandler {

	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		var catevent *CategoryEvent
	Parent:
		for {
			ev := <-sub.Chan
			switch event := ev.(type) {
			case *CategoryEvent:
				catevent = event
				sub.Cancel()
				break Parent
			}
		}

		// // testing
		// flow.etalasemap.AddMap(&services.EtalasePayload{
		// 	Etalase: "test Upload etalase",
		// 	CatIDs:  []int{catevent.CatID},
		// })

		akunetalase := services.NewAkunEtalaseService(tokpedup.Api, flow.etalasemap)
		err := akunetalase.RefreshShowCase()
		if err != nil {
			return err
		}

		showcase, err := akunetalase.GetEtalase(catevent.CatID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			if err != nil {
				return err
			}
		}

		payload.Lock()
		defer payload.Unlock()
		payload.Input.Menus = []*model.MenuInput{
			{
				MenuID: showcase.ID,
			},
		}

		return nil
	}

}
