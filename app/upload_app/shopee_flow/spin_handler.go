package shopee_flow

import (
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createSpinHandler(akun *repo.AkunItem, spin shopeeuploader.SpinFunc) uploader.UploadHandler {
	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		// defer sub.Cancel()

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

		havevariant := len(source.Models) > 1

		fixprice := source.GetPrice(true) // TODO: fixing
		priceprofit := spin.Price(int(fixprice))

		payload.Lock()
		defer payload.Unlock()
		payload.HaveVariant = havevariant
		input := payload.Input
		input.ProductName = spin.Title(source.Name)
		input.Description = spin.Description(input.ProductName, source.Description)
		if !havevariant {
			payload.NovariantStockPrice.Price = int64(priceprofit)
			// TODO: payload.NovariantStockPrice.Stock =
		}

		return nil
	}
}
