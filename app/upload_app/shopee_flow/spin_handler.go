package shopee_flow

import (
	"log"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createSpinHandler(akun *repo.AkunItem, spin shopeeuploader.SpinFunc) uploader.UploadHandler {
	return func(eventcore uploader.EmitFunc, tokpedup *uploader.TokopediaUploader, payload *uploader.PayloadUpload, sub *common_concept.Subscriber) error {

		var source *public_api.PublicProduct
		var distance *mongorepo.ShopDistance

	Parent:
		for {
			ev := <-sub.Chan
			switch event := ev.(type) {
			case *public_api.PublicProduct:
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
		priceprofit := spin.Price(int(fixprice))

		productName := spin.Title(source.Name)
		if len(productName) > 70 {
			productName = productName[:70]
		}

		description := spin.Description(productName, source.Description)
		if len(description) > 2000 {
			description = description[:2000]
		}

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
