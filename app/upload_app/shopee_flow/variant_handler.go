package shopee_flow

import (
	"log"
	"strconv"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/go_v2_shopeelib/lib/public_api/public_model"
	shopeeuploader "github.com/pdcgo/go_v2_shopeelib/lib/uploader"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/uploader"
)

func (flow *ShopeeToTopedFlow) createVariantHandler(spin shopeeuploader.SpinFunc) uploader.UploadHandler {

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

		if !havevariant {
			log.Println("setup no variant")
			return nil
		}
		// variant list
		selections := []model.ProductSelection{}
		for tierind, tier := range source.TierVariations {
			options := make([]model.SelectionsOptions, len(tier.Options))

			for optind, name := range tier.Options {
				option := model.SelectionsOptions{
					UnitValueID: "0",
					Value:       name,
					HexCode:     "",
				}
				options[optind] = option
			}

			varID := strconv.Itoa(tierind)
			selection := model.ProductSelection{
				UnitID:    "",
				VariantID: varID,
				Name:      tier.Name,
				Options:   options,
			}

			selections = append(selections, selection)
		}

		// product variant
		ratio := flow.ConfigFlow.RatioWeightPredict
		berat := float64(distance.Price) / distance.Km / ratio

		products := []model.ProductVariant{}
		for ind, smodel := range source.Models {

			price := int(smodel.GetPrice(flow.ConfigFlow.MarkupConfig.UseDiscount) / 100000)
			price = spin.Price(price)

			product := model.ProductVariant{
				Combination: smodel.Extinfo.TierIndex,
				IsPrimary:   ind == 0,
				Price:       price,
				Stock:       smodel.Stock,
				Pictures:    []model.Pictures{},
				WeightUnit:  model.GramUnit,
				Weight:      int(berat),
				Status:      model.LimitedStatus,
			}

			if flow.ConfigFlow.VariantHandlerConfig.RandomStock {
				product.Stock = flow.ConfigFlow.VariantHandlerConfig.StockSpin.GenerateSpin(len(source.Models))
			}

			products = append(products, product)
		}

		payload.Lock()
		defer payload.Unlock()

		payload.HaveVariant = havevariant
		payload.Variant = &model.Variant{
			Products:   products,
			Selections: selections,
			Sizecharts: []interface{}{},
		}

		log.Println("setup variant")
		return nil
	}
}
