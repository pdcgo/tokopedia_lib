package api

type UserShopInfo struct {
	Info struct {
		ShopID         int64  `json:"shop_id,string"`
		ShopDomain     string `json:"shop_domain"`
		ShopName       string `json:"shop_name"`
		ShopAvatar     string `json:"shop_avatar"`
		ShopIsOfficial string `json:"shop_is_official"`
		ShopScore      int    `json:"shop_score"`
		ShopLocation   string `json:"shop_location"`
		Typename       string `json:"__typename"`
	} `json:"info"`
	Owner struct {
		OwnerID        int    `json:"owner_id"`
		IsGoldMerchant bool   `json:"is_gold_merchant"`
		PmStatus       string `json:"pm_status"`
		Typename       string `json:"__typename"`
	} `json:"owner"`
	Typename string `json:"__typename"`
}

type IsAtuheticatedData struct {
	IsAuthenticated int           `json:"isAuthenticated"`
	UserShopInfo    *UserShopInfo `json:"userShopInfo"`
	User            struct {
		Email          string `json:"email"`
		ID             string `json:"id"`
		Name           string `json:"name"`
		FullName       string `json:"full_name"`
		ProfilePicture string `json:"profilePicture"`
		Typename       string `json:"__typename"`
	} `json:"user"`
}

type IsAuthenticatedRes struct {
	Data IsAtuheticatedData `json:"data"`
}

func (api *TokopediaApi) IsAutheticated() ([]*IsAuthenticatedRes, error) {
	var variable struct{}

	query := GraphqlPayload{
		OperationName: "isAuthenticatedQuery",
		Variables:     variable,
		Query: `
		query isAuthenticatedQuery {
			isAuthenticated
			userShopInfo {
			  info {
				shop_id
				shop_domain
				shop_name
				shop_avatar
				shop_is_official
				shop_score
				shop_location
				__typename
			  }
			  owner {
				owner_id
				is_gold_merchant
				pm_status
				__typename
			  }
			  __typename
			}
			user {
			  email
			  id
			  name
			  full_name
			  profilePicture
			  __typename
			}
		}`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil []*IsAuthenticatedRes
	err := api.SendRequest(req, &hasil)

	if err != nil {
		return hasil, err
	}

	api.AuthenticatedData = &hasil[0].Data

	return hasil, err
}
