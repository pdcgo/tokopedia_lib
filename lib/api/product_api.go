package api

import (
	"log"
	"strconv"
	"time"
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
		OperationName: "ProductListMeta",
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
	var hasil []*ProductListMetaRes
	err := api.SendRequest(req, &hasil)

	return hasil[0], err
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
	PreOrder struct {
		Duration int    `json:"duration"`
		IsActive bool   `json:"isActive"`
		TimeUnit string `json:"timeUnit"`
	} `json:"preorder,omitempty"`
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

type ProductAddVar struct {
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

func (api *TokopediaApi) ProductAdd(variables *ProductAddVar) (*ProductAddResp, error) {
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

	var hasil []*ProductAddResp
	err := api.SendRequest(req, &hasil)

	return hasil[0], err
}

type ProductListVar struct {
	ShopID string `json:"shopID"`
	Filter []struct {
		ID    string   `json:"id"`
		Value []string `json:"value"`
	} `json:"filter"`
	Sort struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"sort"`
	ExtraInfo   []string `json:"extraInfo"`
	WarehouseID string   `json:"warehouseID"`
}

type ProductListResp []struct {
	Data struct {
		ProductList struct {
			Header struct {
				ProcessTime float64       `json:"processTime"`
				Messages    []interface{} `json:"messages"`
				Reason      string        `json:"reason"`
				ErrorCode   string        `json:"errorCode"`
				Typename    string        `json:"__typename"`
			} `json:"header"`
			Data []struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Price struct {
					Min      int    `json:"min"`
					Max      int    `json:"max"`
					Typename string `json:"__typename"`
				} `json:"price"`
				Stock            int    `json:"stock"`
				Status           string `json:"status"`
				MinOrder         int    `json:"minOrder"`
				MaxOrder         int    `json:"maxOrder"`
				Weight           int    `json:"weight"`
				WeightUnit       string `json:"weightUnit"`
				Condition        string `json:"condition"`
				IsMustInsurance  bool   `json:"isMustInsurance"`
				IsKreasiLokal    bool   `json:"isKreasiLokal"`
				IsCOD            bool   `json:"isCOD"`
				IsCampaign       bool   `json:"isCampaign"`
				IsVariant        bool   `json:"isVariant"`
				URL              string `json:"url"`
				Sku              string `json:"sku"`
				Cashback         int    `json:"cashback"`
				Featured         int    `json:"featured"`
				HasStockReserved bool   `json:"hasStockReserved"`
				HasInbound       bool   `json:"hasInbound"`
				WarehouseCount   int    `json:"warehouseCount"`
				IsEmptyStock     bool   `json:"isEmptyStock"`
				Score            struct {
					Total    int    `json:"total"`
					Typename string `json:"__typename"`
				} `json:"score"`
				Pictures []struct {
					URLThumbnail string `json:"urlThumbnail"`
					Typename     string `json:"__typename"`
				} `json:"pictures"`
				Shop struct {
					ID       string `json:"id"`
					Typename string `json:"__typename"`
				} `json:"shop"`
				Wholesale []interface{} `json:"wholesale"`
				Stats     struct {
					CountView   int    `json:"countView"`
					CountReview int    `json:"countReview"`
					CountTalk   int    `json:"countTalk"`
					Typename    string `json:"__typename"`
				} `json:"stats"`
				TxStats struct {
					Sold     int    `json:"sold"`
					Typename string `json:"__typename"`
				} `json:"txStats"`
				Topads             interface{}   `json:"topads"`
				PriceSuggestion    interface{}   `json:"priceSuggestion"`
				CampaignType       []interface{} `json:"campaignType"`
				SuspendLevel       int           `json:"suspendLevel"`
				HasStockAlert      bool          `json:"hasStockAlert"`
				StockAlertCount    int           `json:"stockAlertCount"`
				StockAlertActive   bool          `json:"stockAlertActive"`
				HaveNotifyMeOOS    bool          `json:"haveNotifyMeOOS"`
				NotifyMeOOSCount   int           `json:"notifyMeOOSCount"`
				NotifyMeOOSWording string        `json:"notifyMeOOSWording"`
				ManageProductData  struct {
					IsStockGuaranteed bool   `json:"isStockGuaranteed"`
					ScoreV3           int    `json:"scoreV3"`
					Typename          string `json:"__typename"`
				} `json:"manageProductData"`
				CreateTime time.Time `json:"createTime"`
				Typename   string    `json:"__typename"`
			} `json:"data"`
			Typename string `json:"__typename"`
		} `json:"ProductList"`
	} `json:"data"`
}

func (api *TokopediaApi) ProductList(payload *ProductListVar) (*ProductListResp, error) {

	query := GraphqlPayload{
		OperationName: "ProductList",
		Variables:     payload,
		Query: `query ProductList($shopID: String!, $filter: [GoodsFilterInput], $sort: GoodsSortInput, $extraInfo: [String], $warehouseID: String) {
			  ProductList(shopID: $shopID, filter: $filter, sort: $sort, extraInfo: $extraInfo, warehouseID: $warehouseID) {
			    header {
			      processTime
			      messages
			      reason
			      errorCode
			      __typename
			    }
		    data {
			      id
			      name
			      price {
			        min
			        max
			        __typename
			      }
		      stock
		      status
		      minOrder
		      maxOrder
		      weight
		      weightUnit
		      condition
		      isMustInsurance
		      isKreasiLokal
		      isCOD
		      isCampaign
		      isVariant
		      url
		      sku
		      cashback
		      featured
		      hasStockReserved
		      hasInbound
		      warehouseCount
		      isEmptyStock
		      score {
			        total
			        __typename
			      }
		      pictures {
			        urlThumbnail
			        __typename
			      }
		      shop {
			        id
			        __typename
			      }
		      wholesale {
			        minQty
			        __typename
			      }
		      stats {
			        countView
			        countReview
			        countTalk
			        __typename
			      }
		      txStats {
			        sold
			        __typename
			      }
		      topads {
			        status
			        management
			        __typename
			      }
		      priceSuggestion {
			        suggestedPrice
			        suggestedPriceTreshold
			        suggestedPriceMin
			        suggestedPriceMax
			        label
			        productRecommendation {
			          title
			          productID
			          price
			          imageURL
			          sold
			          rating
			          __typename
			        }
		        __typename
		      }
		      campaignType {
			        id
			        name
			        iconURL
			        __typename
			      }
		      suspendLevel
		      hasStockAlert
		      stockAlertCount
		      stockAlertActive
		      haveNotifyMeOOS
		      notifyMeOOSCount
		      notifyMeOOSWording
		      manageProductData {
			        isStockGuaranteed
			        scoreV3
			        __typename
			      }
		      createTime
		      __typename
		    }
		    __typename
		  }
		}`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil *ProductListResp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}
