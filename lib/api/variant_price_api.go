package api

import "github.com/pdcgo/tokopedia_lib/lib/model"

func (api *TokopediaApi) VariantPriceValidation(catId int, variant *model.Variant) (*model.VPVRes, error) {

	cpvariant := &model.Variant{
		Products:   model.ProductVariants{},
		Selections: variant.Selections,
		Sizecharts: []any{},
	}

	for _, prod := range variant.Products {
		prod.Pictures = []model.Pictures{}
		cpvariant.Products = append(cpvariant.Products, prod)
	}

	variable := model.VPVVariable{
		Input: &model.VPVInput{
			Variant: cpvariant,
			Category: &model.VPVCategory{
				ID: catId,
			},
		},
	}

	query := GraphqlPayload{
		OperationName: "GetVariantPriceValidation",
		Variables:     variable,
		Query: `mutation GetVariantPriceValidation($input: ProductInputV3!) {
			ProductValidateV3(input: $input) {
			  header {
				messages
				reason
				errorCode
				__typename
			  }
			  isSuccess
			  data {
				variants {
				  messages
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil model.VPVRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
