package api

import "strconv"

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

func (api *TokopediaApi) ShopInfoByID() ([]*ShopInfoByIDRes, error) {
	query := GraphqlPayload{
		OperationName: "ShopInfoByIDQuery",
		Variables: ShopInfoByIDVar{
			ShopIDs: []int64{api.AuthenticatedData.UserShopInfo.Info.ShopID},
			Fields:  []string{"assets", "core", "favorite", "location", "other-goldos", "other-shiploc", "status", "allow_manage", "is_owner", "closed_info", "status", "assets"},
		},
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

	var hasil []*ShopInfoByIDRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
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

func (api *TokopediaApi) GoldGetPMOSStatus() ([]*GoldGetPMOSStatusRes, error) {
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

	var hasil []*GoldGetPMOSStatusRes
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

func (api *TokopediaApi) GetShopScoreLevel() ([]*GetShopScoreLevelRes, error) {
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

	var hasil []*GetShopScoreLevelRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
}
