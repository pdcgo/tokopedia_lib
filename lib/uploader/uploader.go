package uploader

import (
	"context"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type TokopediaUploader struct {
	Api *api.TokopediaApi
	Ctx context.Context
}

type PayloadUpload struct {
	sync.Mutex
	Input               *model.InputVariable
	HaveVariant         bool
	Variant             *model.Variant
	NovariantStockPrice *model.NoVariantStockPrice
}

type EmitFunc func(event interface{})
type UploadHandler func(eventcore EmitFunc, tokpedup *TokopediaUploader, payload *PayloadUpload, sub *common_concept.Subscriber) error

func (upload *TokopediaUploader) UploadProduct(payload *PayloadUpload) (*model.ProductAddResp, error) {
	paydata := model.ProductAddVar{}

	if payload.HaveVariant {
		paydata.Input = model.InputVariant{
			InputVariable: payload.Input,
			Variant:       payload.Variant,
		}

	} else {
		data := model.InputNoVariant{
			InputVariable:       payload.Input,
			NoVariantStockPrice: payload.NovariantStockPrice,
		}
		paydata.Input = data

	}

	return upload.Api.ProductAdd(&paydata)

}

func (upload *TokopediaUploader) RunUploader(handlers ...UploadHandler) (*model.ProductAddResp, error) {
	event := common_concept.NewCoreEvent()
	defer event.Close()

	payload := PayloadUpload{}
	handlerLen := len(handlers)
	waitchan := make(chan error, handlerLen)

	subsevents := make([]*common_concept.Subscriber, handlerLen)
	for cc := 0; cc < handlerLen; cc += 1 {
		sub := event.CreateSubscriber()
		subsevents[cc] = sub
	}

	for ind, handle := range handlers {
		sub := subsevents[ind]
		hand := handle
		go func() {
			err := hand(func(eventdata interface{}) {
				event.Emit(eventdata)
			}, upload, &payload, sub)

			waitchan <- err
		}()
	}

	for c := 0; c < handlerLen; c += 1 {
		select {
		case err := <-waitchan:
			if err != nil {
				return nil, err
			}
		case <-upload.Ctx.Done():
			log.Println("context Done")
			return nil, nil
		}
	}

	return upload.UploadProduct(&payload)
}
