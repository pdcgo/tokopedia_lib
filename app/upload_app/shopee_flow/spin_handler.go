package shopee_flow

import (
	"log"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/spin"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api/public_model"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createSpinHandler(akun *repo.AkunItem, spinner shopeeuploader.SpinFunc) uploader.UploadHandler {
	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		var source *public_model.PublicProduct
		var distance *mongorepo.ShopDistance

	Parent:
		for {
			ev := <-sub.Chan
			switch event := ev.(type) {
			case *public_model.PublicProduct:
				source = event
			case *mongorepo.ShopDistance:
				distance = event

			}
			if source != nil && distance != nil {
				break Parent
			}
		}

		havevariant := len(source.Models) > 1
		ratio := flow.ConfigFlow.RatioWeightPredict
		berat := float64(distance.Price) / distance.Km / ratio

		fixprice := source.GetPrice(true) // TODO: fixing
		priceprofit := spinner.Price(int(fixprice))

		productName := spinner.Title(source.Name, spin.MaxTokpedTitle)
		description := spinner.Description(productName, source.Description, spin.MaxTokpedDesc)

		payload.Lock()
		defer payload.Unlock()
		payload.HaveVariant = havevariant
		input := payload.Input
		input.ProductName = productName
		input.Description = description
		input.Weight = int64(berat)
		input.WeightUnit = model.GramUnit
		if !havevariant {
			payload.NovariantStockPrice.Price = int64(priceprofit)
			payload.NovariantStockPrice.Status = model.LimitedStatus

			payload.NovariantStockPrice.Stock = int64(source.Stock)

			if flow.ConfigFlow.VariantHandlerConfig.RandomStock {
				stock := flow.ConfigFlow.VariantHandlerConfig.StockSpin.GenerateSpin(1)
				payload.NovariantStockPrice.Stock = int64(stock)
			}

		}

		log.Println("setup spin")

		return nil
	}
}
