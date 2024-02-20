package api

import (
	"errors"
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type ShopInfoByIDRes struct {
	Data struct {
		ShopInfoByID struct {
			Result []struct {
				FavoriteData struct {
					TotalFavorite int    `json:"totalFavorite"`
					Typename      string `json:"__typename"`
				} `json:"favoriteData"`
				GoldOS struct {
					IsGold           int    `json:"isGold"`
					IsOfficial       int    `json:"isOfficial"`
					Badge            string `json:"badge"`
					ShopTier         int    `json:"shopTier"`
					ShopTierWording  string `json:"shopTierWording"`
					ShopGrade        int    `json:"shopGrade"`
					ShopGradeWording string `json:"shopGradeWording"`
					Typename         string `json:"__typename"`
				} `json:"goldOS"`
				Location   string `json:"location"`
				ShopAssets struct {
					Avatar       string `json:"avatar"`
					Cover        string `json:"cover"`
					DefaultCover []struct {
						ID       string `json:"id"`
						Path     string `json:"path"`
						Typename string `json:"__typename"`
					} `json:"defaultCover"`
					Typename string `json:"__typename"`
				} `json:"shopAssets"`
				IsAllowManage int `json:"isAllowManage"`
				IsOwner       int `json:"isOwner"`
				ShopCore      struct {
					Name        string `json:"name"`
					ShopID      string `json:"shopID"`
					Domain      string `json:"domain"`
					Description string `json:"description"`
					TagLine     string `json:"tagLine"`
					Typename    string `json:"__typename"`
				} `json:"shopCore"`
				ShopHomeType string `json:"shopHomeType"`
				ClosedInfo   struct {
					ClosedNote string `json:"closedNote"`
					Until      string `json:"until"`
					Detail     struct {
						StartDate string `json:"startDate"`
						EndDate   string `json:"endDate"`
						OpenDate  string `json:"openDate"`
						Status    int    `json:"status"`
						Typename  string `json:"__typename"`
					} `json:"detail"`
					Typename string `json:"__typename"`
				} `json:"closedInfo"`
				StatusInfo struct {
					ShopStatus int    `json:"shopStatus"`
					StatusName string `json:"statusName"`
					Typename   string `json:"__typename"`
				} `json:"statusInfo"`
				Os struct {
					IsOfficial int    `json:"isOfficial"`
					Expired    string `json:"expired"`
					Typename   string `json:"__typename"`
				} `json:"os"`
				Typename string `json:"__typename"`
			} `json:"result"`
			Typename string `json:"__typename"`
		} `json:"shopInfoByID"`
	} `json:"data"`
}

type ShopInfoByIDVar struct {
	ShopIDs []int64  `json:"shopIDs"`
	Fields  []string `json:"fields"`
}

func NewShopInfoByIDVar(shopid int64) *ShopInfoByIDVar {
	return &ShopInfoByIDVar{
		ShopIDs: []int64{shopid},
		Fields: []string{
			"assets",
			"core",
			"favorite",
			"location",
			"other-goldos",
			"other-shiploc",
			"status",
			"allow_manage",
			"is_owner",
			"closed_info",
			"status",
			"assets",
		},
	}
}

var ErrIsNotAuthorized = errors.New("shop is not authorized")

func (api *TokopediaApi) ShopInfoByID() (*model.ShopInfoByIDRes, error) {
	variable := NewShopInfoByIDVar(api.AuthenticatedData.UserShopInfo.Info.ShopID)
	query := GraphqlPayload{
		OperationName: "ShopInfoByIDQuery",
		Variables:     variable,
		Query: `
		query ShopInfoByIDQuery($shopIDs: [Int!]!, $fields: [String!]!) {
			shopInfoByID(input: {shopIDs: $shopIDs, fields: $fields}) {
			  result {
				favoriteData {
				  totalFavorite
				  __typename
				}
				goldOS {
				  isGold
				  isOfficial
				  badge
				  shopTier
				  shopTierWording
				  shopGrade
				  shopGradeWording
				  __typename
				}
				location
				shopAssets {
				  avatar
				  cover
				  defaultCover {
					id
					path
					__typename
				  }
				  __typename
				}
				isAllowManage
				isOwner
				shopCore {
				  name
				  shopID
				  domain
				  description
				  tagLine
				  __typename
				}
				shopHomeType
				closedInfo {
				  closedNote
				  until
				  detail {
					startDate
					endDate
					openDate
					status
					__typename
				  }
				  __typename
				}
				statusInfo {
				  shopStatus
				  statusName
				  __typename
				}
				os {
				  isOfficial
				  expired
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil *model.ShopInfoByIDRes
	err := api.SendRequest(req, &hasil)

	if err != nil {
		return hasil, err
	}

	if hasil.Errors.IsNotAuthorized() {
		return hasil, ErrIsNotAuthorized
	}

	if len(hasil.Errors) > 0 {
		return hasil, hasil.Errors
	}

	return hasil, nil
}

type GoldGetPMOSStatusRes struct {
	Data *model.GoldGetPMOSStatusData `json:"data"`
}

type GoldGetPMOSStatusVar struct {
	ShopID int64 `json:"shopId"`
}

func (api *TokopediaApi) GoldGetPMOSStatus() (*GoldGetPMOSStatusRes, error) {
	query := GraphqlPayload{
		OperationName: "goldGetPMOSStatusQuery",
		Variables: GoldGetPMOSStatusVar{
			ShopID: api.AuthenticatedData.UserShopInfo.Info.ShopID,
		},
		Query: `
		query goldGetPMOSStatusQuery($shopId: Int!) {
			goldGetPMOSStatus(shopID: $shopId, includeOS: false) {
			  data {
				power_merchant {
				  status
				  auto_extend {
					status
					tkpd_product_id
					__typename
				  }
				  pm_tier
				  expired_time
				  __typename
				}
				__typename
			  }
			  header {
				error_code
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil *GoldGetPMOSStatusRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
}

type GetShopScoreLevelRes struct {
	Data *model.ShopScoreData `json:"data"`
}

type GetShopScoreLevelVar struct {
	ShopIDStr       string `json:"shopIDStr"`
	Source          string `json:"source"`
	CalculateScore  bool   `json:"calculateScore"`
	GetNextMinValue bool   `json:"getNextMinValue"`
	IncludeRawData  bool   `json:"includeRawData"`
}

func (api *TokopediaApi) GetShopScoreLevel() (*GetShopScoreLevelRes, error) {
	shopid := strconv.Itoa(int(api.AuthenticatedData.UserShopInfo.Info.ShopID))
	query := GraphqlPayload{
		OperationName: "GetShopScoreLevel",
		Variables: GetShopScoreLevelVar{
			ShopIDStr:       shopid,
			IncludeRawData:  true,
			GetNextMinValue: true,
			Source:          "icarus",
		},
		Query: `
		query GetShopScoreLevel($shopIDStr: String!, $source: String!, $calculateScore: Boolean!, $getNextMinValue: Boolean!, $includeRawData: Boolean) {
			shopScoreLevel(input: {shopID: $shopIDStr, source: $source, calculateScore: $calculateScore, getNextMinValue: $getNextMinValue, includeRawData: $includeRawData}) {
			  result {
				shopID
				shopScore
				shopLevel
				shopScoreDetail {
				  title
				  identifier
				  value
				  rawValue
				  nextMinValue
				  colorText
				  __typename
				}
				period
				nextUpdate
				__typename
			  }
			  error {
				message
				__typename
			  }
			  __typename
			}
			shopLevel(input: {shopID: $shopIDStr, source: $source}) {
			  result {
				shopID
				period
				nextUpdate
				shopLevel
				itemSold
				niv
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

	var hasil *GetShopScoreLevelRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
}

type GetShopLocationAllVar struct {
	ShopId int `json:"shop_id"`
}

type GetShopLocationAllRes struct {
	Data struct {
		ShopLocGetAllLocations model.ShopLocationAll
	} `json:"data"`
}

func (api *TokopediaApi) GetShopLocationAll(shopid int) (*GetShopLocationAllRes, error) {
	query := GraphqlPayload{
		OperationName: "getAllShopLocations",
		Variables: GetShopLocationAllVar{
			ShopId: shopid,
		},
		Query: `
		query getAllShopLocations($shop_id: Int!) {
			ShopLocGetAllLocations(input: {shop_id: $shop_id}) {
			  status
			  message
			  error {
				id
				description
				__typename
			  }
			  data {
				general_ticker {
				  header
				  body
				  body_link_text
				  body_link_url
				  __typename
				}
				warehouses {
				  warehouse_id
				  warehouse_name
				  warehouse_type
				  shop_id {
					int64
					valid
					__typename
				  }
				  partner_id {
					int64
					valid
					__typename
				  }
				  address_detail
				  postal_code
				  latlon
				  district_id
				  district_name
				  city_id
				  city_name
				  province_id
				  province_name
				  country
				  status
				  is_covered_by_couriers
				  ticker {
					text_inactive
					text_courier_setting
					link_courier_setting
					__typename
				  }
				  __typename
				}
				__typename
			  }
			  __typename
			}
		}`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil *GetShopLocationAllRes
	var test map[string]any
	api.SendRequest(req, &test)
	err := api.SendRequest(req, &hasil)

	return hasil, err
}
