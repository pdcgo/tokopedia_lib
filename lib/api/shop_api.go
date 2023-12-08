package api

import (
	"errors"
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

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
	Data struct {
		GoldGetPMOSStatus struct {
			Data struct {
				PowerMerchant struct {
					Status     string `json:"status"`
					AutoExtend struct {
						Status        string `json:"status"`
						TkpdProductID int    `json:"tkpd_product_id"`
						Typename      string `json:"__typename"`
					} `json:"auto_extend"`
					PmTier      int    `json:"pm_tier"`
					ExpiredTime string `json:"expired_time"`
					Typename    string `json:"__typename"`
				} `json:"power_merchant"`
				Typename string `json:"__typename"`
			} `json:"data"`
			Header struct {
				ErrorCode string `json:"error_code"`
				Typename  string `json:"__typename"`
			} `json:"header"`
			Typename string `json:"__typename"`
		} `json:"goldGetPMOSStatus"`
	} `json:"data"`
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
	Data struct {
		ShopScoreLevel struct {
			Result struct {
				ShopID          string  `json:"shopID"`
				ShopScore       float32 `json:"shopScore"`
				ShopLevel       int     `json:"shopLevel"`
				ShopScoreDetail []struct {
					Title        string  `json:"title"`
					Identifier   string  `json:"identifier"`
					Value        float32 `json:"value"`
					RawValue     float32 `json:"rawValue"`
					NextMinValue float64 `json:"nextMinValue"`
					ColorText    string  `json:"colorText"`
					Typename     string  `json:"__typename"`
				} `json:"shopScoreDetail"`
				Period     string `json:"period"`
				NextUpdate string `json:"nextUpdate"`
				Typename   string `json:"__typename"`
			} `json:"result"`
			Error struct {
				Message  string `json:"message"`
				Typename string `json:"__typename"`
			} `json:"error"`
			Typename string `json:"__typename"`
		} `json:"shopScoreLevel"`
		ShopLevel struct {
			Result struct {
				ShopID     string `json:"shopID"`
				Period     string `json:"period"`
				NextUpdate string `json:"nextUpdate"`
				ShopLevel  int    `json:"shopLevel"`
				ItemSold   int    `json:"itemSold"`
				Niv        int    `json:"niv"`
				Typename   string `json:"__typename"`
			} `json:"result"`
			Error struct {
				Message  string `json:"message"`
				Typename string `json:"__typename"`
			} `json:"error"`
			Typename string `json:"__typename"`
		} `json:"shopLevel"`
	} `json:"data"`
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
