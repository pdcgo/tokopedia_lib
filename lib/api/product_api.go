package api

import "strconv"

type ProductListMetaRes struct {
	Data struct {
		ProductListMeta struct {
			Header struct {
				ProcessTime float64 `json:"processTime"`
				Messages    []any   `json:"messages"`
				Reason      string  `json:"reason"`
				ErrorCode   string  `json:"errorCode"`
				Typename    string  `json:"__typename"`
			} `json:"header"`
			Data struct {
				Tab []struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Value    int    `json:"value"`
					Typename string `json:"__typename"`
				} `json:"tab"`
				Filter []struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Value    []any  `json:"value"`
					Typename string `json:"__typename"`
				} `json:"filter"`
				Sort []struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Value    string `json:"value"`
					Typename string `json:"__typename"`
				} `json:"sort"`
				ShopCategories []struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Typename string `json:"__typename"`
				} `json:"shopCategories"`
				Access []struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Value    string `json:"value"`
					Typename string `json:"__typename"`
				} `json:"access"`
				Typename string `json:"__typename"`
			} `json:"data"`
			Typename string `json:"__typename"`
		} `json:"ProductListMeta"`
	} `json:"data"`
}

type ProductListMetaVar struct {
	ShopID      string   `json:"shopID"`
	ExtraInfo   []string `json:"extraInfo"`
	WareHouseID string   `json:"warehouseID"`
}

func (api *TokopediaApi) ProductListMeta() (*ProductListMetaRes, error) {
	shopid := strconv.Itoa(int(api.AuthenticatedData.UserShopInfo.Info.ShopID))
	query := GraphqlPayload{
		OperationName: "ShopInfoByIDQuery",
		Variables: ProductListMetaVar{
			ShopID:    shopid,
			ExtraInfo: []string{"rbac", "access", "category"},
		},
		Query: `
		query ProductListMeta($shopID: String!, $warehouseID: String, $extraInfo: [String]) {
			ProductListMeta(shopID: $shopID, warehouseID: $warehouseID, extraInfo: $extraInfo) {
			  header {
				processTime
				messages
				reason
				errorCode
				__typename
			  }
			  data {
				tab {
				  id
				  name
				  value
				  __typename
				}
				filter {
				  id
				  name
				  value
				  __typename
				}
				sort {
				  id
				  name
				  value
				  __typename
				}
				shopCategories {
				  id
				  name
				  __typename
				}
				access {
				  id
				  name
				  value
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ProductListMetaRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
