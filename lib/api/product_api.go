package api

import (
	"log"
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/model"
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
	var hasil *ProductListMetaRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
}

func (api *TokopediaApi) ProductAdd(variables *model.ProductAddVar) (*model.ProductAddResp, error) {
	gqlQuery := GraphqlPayload{
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

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ProductAddResp
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return &hasil, err
	}

	if !hasil.Data.ProductAddV3.IsSuccess {
		return &hasil, hasil.Data.ProductAddV3.Header
	}

	return &hasil, nil
}

func (api *TokopediaApi) ProductList(payload *model.ProductListVar) (*model.ProductListResp, error) {

	gqlQuery := GraphqlPayload{
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

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.ProductListResp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}

func (api *TokopediaApi) GetProductV3(payload *model.GetProductV3Var) (*model.GetProductV3Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "getProductV3",
		Variables:     payload,
		Query: `query getProductV3($productID: String!, $options: OptionV3!, $extraInfo: ExtraInfoV3) {
			  getProductV3(productID: $productID, options: $options, extraInfo: $extraInfo) {
			    lock {
			      full
			      partial {
			        price
			        status
			        stock
			        wholesale
			        name
			        __typename
			      }
		      __typename
		    }
		    txStats {
			      itemSold
			      __typename
			    }
		    shop {
			      id
			      __typename
			    }
		    productID
		    productName
		    status
		    stock
		    price
		    minOrder
		    description
		    weightUnit
		    weight
		    condition
		    mustInsurance
		    sku
		    category {
			      id
			      name
			      title
			      detail {
			        id
			        name
			        breadcrumbURL
			        __typename
			      }
		      __typename
		    }
		    menu {
			      menuID
			      name
			      url
			      __typename
			    }
		    menus
		    video {
			      url
			      __typename
			    }
		    customVideo {
			      id
			      fileName
			      url
			      __typename
			    }
		    pictures {
			      picID
			      filePath
			      fileName
			      width
			      height
			      urlOriginal
			      __typename
			    }
		    wholesale {
			      minQty
			      price
			      __typename
			    }
		    dimension {
			      length
			      width
			      height
			      __typename
			    }
		    preorder {
			      duration
			      timeUnit
			      isActive
			      __typename
			    }
		    variant {
			      products {
			        status
			        combination
			        isPrimary
			        price
			        sku
			        stock
			        weight
			        weightUnit
			        pictures {
			          picID
			          filePath
			          fileName
			          width
			          height
			          urlOriginal
			          __typename
			        }
		        __typename
		      }
		      selections {
			        variantID
			        variantName
			        unitID
			        unitName
			        identifier
			        options {
			          unitValueID
			          value
			          hexCode
			          __typename
			        }
		        __typename
		      }
		      sizecharts {
			        picID
			        filePath
			        fileName
			        width
			        height
			        urlOriginal
			        __typename
			      }
		      __typename
		    }
		    cpl {
			      shipperServices
			      __typename
			    }
		    __typename
		  }
		}`,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.GetProductV3Resp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}

func (api *TokopediaApi) ProductUpdate(payload *model.ProductUpdateVar) (*model.ProductUpdateResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "productUpdate",
		Variables:     payload,
		Query: `mutation productUpdate($input: ProductInputV3!) {
			  ProductUpdateV3(input: $input) {
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
		}`,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.ProductUpdateResp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}

func (api *TokopediaApi) BulkProductEditV3(payload *model.BulkProductEditV3Var) (*model.BulkProductEditV3Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "BulkProductEditV3",
		Variables:     payload,
		Query: `mutation BulkProductEditV3($input: [ProductInputV3]!) {
			  BulkProductEditV3(input: $input) {
			    productID
			    result {
			      header {
			        messages
			        reason
			        errorCode
			        __typename
			      }
		      isSuccess
		      __typename
		    }
		    __typename
		  }
		}`,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.BulkProductEditV3Resp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}
