package shopee_flow

import (
	"errors"
	"log"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
	"github.com/rs/zerolog"
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
		}

		if err != nil {

			return pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
				return event.Int("cat id", catevent.CatID).
					Interface("showcase", showcase).
					Interface("catevert", catevent).
					Str("akun", tokpedup.Api.AuthenticatedData.User.Name)
			})
		}
		if showcase == nil {
			log.Println(err, "showcase")
			return pdc_common.ReportErrorCustom(errors.New("showcase error"), func(event *zerolog.Event) *zerolog.Event {
				return event.Int("cat id", catevent.CatID).
					Interface("showcase", showcase).
					Interface("catevert", catevent).
					Str("akun", tokpedup.Api.AuthenticatedData.User.Name)
			})
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
