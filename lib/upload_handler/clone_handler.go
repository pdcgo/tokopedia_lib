package upload_handler

import (
	"net/url"
	"strings"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

// email : pdcthoni@gmail.com
// password : SilentIsMyMantra
// otp : IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ

func CreateCloneHandler(uri string) []uploader.UploadHandler {
	publicApi := api_public.NewTokopediaApiPublic()

	handlers := []uploader.UploadHandler{}

	handlers = append(handlers, func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		// semua image ada di component basic

		return nil
	})

	handlers = append(handlers, func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {
		dataurl := NewPublicUrl(uri)

		_, err := publicApi.PdpGetlayoutQuery(dataurl.LayoutVar)
		if err != nil {
			return err
		}

		payload.Lock()
		defer payload.Unlock()

		payload.Input.Pictures = model.InputPicture{
			Data: []model.Pictures{
				model.Pictures{
					UploadIds: "asd",
				},
				model.Pictures{},
			},
		}

		return nil
	})

	return handlers
}

type PublicUrl struct {
	U               *url.URL
	ShopCoreInfoVar *model_public.ShopCoreInfoVar
	LayoutVar       *model_public.PdpGetlayoutQueryVar
	Domain          string
	ProductPerma    string
}

func NewPublicUrl(uri string) *PublicUrl {
	ur, _ := url.Parse(uri)

	paths := strings.Split(ur.Path, "/")
	shopdomain := paths[len(paths)-2]
	productperma := paths[len(paths)-1]

	return &PublicUrl{
		U:            ur,
		ProductPerma: productperma,
		Domain:       shopdomain,
		ShopCoreInfoVar: &model_public.ShopCoreInfoVar{
			ID:     0,
			Domain: shopdomain,
		},
		LayoutVar: &model_public.PdpGetlayoutQueryVar{
			APIVersion: 1,
			ExtParam:   "",
			LayoutID:   "",
			ProductKey: productperma,
			ShopDomain: shopdomain,
		},
	}
}
