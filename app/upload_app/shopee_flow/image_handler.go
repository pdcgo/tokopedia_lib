package shopee_flow

import (
	"bytes"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api/public_model"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createImageHandler() uploader.UploadHandler {
	config := flow.ConfigFlow.ImageHandlerConfig
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

		errChan := make(chan error)
		var waitup sync.WaitGroup
		api := tokpedup.Api

		imagedatas := source.Images

		if len(source.Images) > 5 {
			imagedatas = source.Images[:5]
		}

		pictures := model.InputPicture{
			Data: make([]model.Pictures, len(imagedatas)),
		}

		for ind, imguri := range imagedatas {
			uri := imguri
			idex := ind
			waitup.Add(1)
			go func() {
				defer waitup.Done()
				log.Println(api.AuthenticatedData.User.Email, "uploading image", uri)
				imgdata, err := public_api.GetShopeeImageFromID(uri)
				if err != nil {
					errChan <- err
					return
				}

				fiximage := bytes.NewBuffer(nil)

				err = shopeeuploader.CropImage(config.CropValue(), imgdata, fiximage)
				if err != nil {
					errChan <- err
					return
				}

				imgres, err := api.UploadProductImage(fiximage)

				if err != nil {
					errChan <- err
					return
				}

				pictures.Data[idex] = model.Pictures{
					UploadIds: imgres.Data.UploadID,
				}
			}()
		}
		waitup.Wait()

		payload.Lock()
		defer payload.Unlock()
		payload.Input.Pictures = pictures

		log.Println("setup image")
		return nil
	}
}
