package api

import "github.com/pdcgo/tokopedia_lib/lib/model"

func (api *TokopediaApi) VariantPricevValidation(catId int, variant *model.Variant) (*model.VPVRes, error) {
	variable := model.VPVVariable{
		Input: &model.VPVInput{
			Variant: variant,
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
