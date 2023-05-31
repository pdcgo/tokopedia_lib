package upload_handler

import (
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

// email : pdcthoni@gmail.com
// password : SilentIsMyMantra
// otp : IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ

func CreateCloneHandler() []uploader.UploadHandler {
	handlers := []uploader.UploadHandler{}
	handlers = append(handlers, func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		return nil
	})

	return handlers
}
