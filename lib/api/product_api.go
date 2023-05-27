package api

import (
	"log"
	"strconv"
)

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
	log.Println("create req success")
	var hasil ProductListMetaRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type InputVariable struct {
	Pictures struct {
		Data []struct {
			UploadIds string `json:"uploadIds"`
		} `json:"data"`
	} `json:"pictures"`
	ProductName string `json:"productName"`
	Category    struct {
		ID string `json:"id"`
	} `json:"category"`
	Condition     string `json:"condition"`
	MinOrder      int64  `json:"minOrder"`
	PriceCurrency string `json:"minCurrency"`
	Weight        int64  `json:"weight"`
	WeightUnit    string `json:"weightUnit"`
	MustInsurance bool   `json:"mustInsurance"`
	Menus         []struct {
		MenuID string `json:"menuID"`
	} `json:"menus"`
	Annotations []string `json:"annotations"`
	Description string   `json:"description"`
	Dimention   struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Length int `json:"length"`
	} `json:"dimension"`
	Catalog *struct {
		CatalogID string `json:"catalogID"`
		IsActive  bool   `json:"isActive"`
	} `json:"catalog,omitempty"`
}

type InputNoVariant struct {
	InputVariable
	Sku    string `json:"sku"`
	Stock  int64  `json:"stock"`
	Price  int64  `json:"price"`
	Status string `json:"status"`
}

type InputVariant struct {
	InputVariable
	Variant struct {
		Selections []struct {
			UnitID    string `json:"unitID"`
			VariantID string `json:"variantID"`
			Name      string `json:"name"`
			Options   []struct {
				UnitValueID string `json:"unitValueID"`
				Value       string `json:"value"`
				HexCode     string `json:"hexCode"`
			} `json:"options"`
		} `json:"selections"`
		Products []struct {
			Combination []int  `json:"combination"`
			IsPrimary   bool   `json:"isPrimary"`
			Price       int    `json:"price"`
			Sku         string `json:"sku"`
			Status      string `json:"status"`
			Stock       int    `json:"stock"`
			Pictures    []struct {
				UploadIds string `json:"uploadIds"`
			} `json:"pictures"`
			Weight     int    `json:"weight"`
			WeightUnit string `json:"weightUnit"`
		} `json:"products"`
		SizeChart []interface{} `json:"sizeChart"`
	} `json:"variant"`
}

type VariablesProductAdd struct {
	Input interface{} `json:"input"`
}

type ProductAddResp struct {
	Data struct {
		ProductAddV3 struct {
			Header struct {
				Message   []any  `json:"message"`
				Reason    string `json:"reason"`
				ErrorCode string `json:"errorCode"`
				TypeName  string `json:"__typename"`
			} `json:"header"`
			IsSuccess bool   `json:"isSuccess"`
			ProductId string `json:"productID"`
			TypeName  string `json:"__typename"`
		} `json:"ProductAddV3"`
	} `json:"data"`
}

func (api *TokopediaApi) ProductAdd(variables *VariablesProductAdd) (*ProductAddResp, error) {
	query := GraphqlPayload{
		OperationName: "productAdd",
		Variables:     variables,
		Query: `
		mutation productAdd($input: ProductInputV3!) {
			  ProductAddV3(input: $input) {
			    header {
			      messages
			      reason
			      errorCode
			      __typename
			    }
		    isSuccess
		    productID
		    __typename
		  }
		}
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ProductAddResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
