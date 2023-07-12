package api

type DeleteShopShowcaseVar struct {
	Input *DeleteShopShowcaseInput `json:"input"`
}
type DeleteShopShowcaseInput struct {
	ID string `json:"id"`
}

type DeleteShopShowcaseRes struct {
	Data *Data `json:"data"`
}
type DeleteShopShowcase struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}
type Data struct {
	DeleteShopShowcase *DeleteShopShowcase `json:"deleteShopShowcase"`
}

func (api *TokopediaApi) DeleteShopShowcase(id string) (*DeleteShopShowcaseRes, error) {
	query := GraphqlPayload{
		OperationName: "DeleteShopShowcase",
		Variables: DeleteShopShowcaseVar{
			Input: &DeleteShopShowcaseInput{
				ID: id,
			},
		},
		Query: `mutation DeleteShopShowcase($input: ParamDeleteShopShowcase!) {
			deleteShopShowcase(input: $input) {
				success
				message
				__typename
			}
		}`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil DeleteShopShowcaseRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ShopShowcaseVar struct {
	WithDefault bool `json:"withDefault"`
}

type ShopShowcaseRes struct {
	Data *ShopShowcaseData `json:"data"`
}
type ShopShowcaseResult struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Count    int    `json:"count"`
	URI      string `json:"uri"`
	Typename string `json:"__typename"`
}
type ShopShowcaseError struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}
type ShopShowcases struct {
	Result   []*ShopShowcaseResult `json:"result"`
	Error    *ShopShowcaseError    `json:"error"`
	Typename string                `json:"__typename"`
}
type ShopShowcaseData struct {
	ShopShowcases ShopShowcases `json:"shopShowcases"`
}

func (api *TokopediaApi) ShopShowcase() (*ShopShowcaseRes, error) {
	query := GraphqlPayload{
		OperationName: "ShopShowcase",
		Variables: ShopShowcaseVar{
			WithDefault: false,
		},
		Query: `query ShopShowcase($withDefault: Boolean) {
			shopShowcases(withDefault: $withDefault) {
			  result {
				id
				name
				count
				uri
				__typename
			  }
			  error {
				message
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ShopShowcaseRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type AddShopShowcaseVar struct {
	Input *AddShopShowcaseInput `json:"input"`
}
type AddShopShowcaseInput struct {
	Name string `json:"name"`
}

type AddShopShowcaseRes struct {
	Data *AddShopShowcaseData `json:"data"`
}
type AddShopShowcase struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	CreatedID string `json:"createdId"`
	Typename  string `json:"__typename"`
}
type AddShopShowcaseData struct {
	AddShopShowcase AddShopShowcase `json:"addShopShowcase"`
}

func (api *TokopediaApi) AddShopShowcase(name string) (*AddShopShowcaseRes, error) {
	query := GraphqlPayload{
		OperationName: "addShopShowcase",
		Variables: AddShopShowcaseVar{
			Input: &AddShopShowcaseInput{
				Name: name,
			},
		},
		Query: `mutation addShopShowcase($input: ParamAddShopShowcase!) {
			addShopShowcase(input: $input) {
			  success
			  message
			  createdId
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil AddShopShowcaseRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
