package shopee_flow

import (
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createAnnotationHandler() uploader.UploadHandler {
	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		payload.Lock()
		defer payload.Unlock()
		input := payload.Input
		input.Annotations = []string{}

		return nil
	}
}
