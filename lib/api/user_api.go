package api

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Info struct {
	ShopID         int64  `json:"shop_id,string"`
	ShopDomain     string `json:"shop_domain"`
	ShopName       string `json:"shop_name"`
	ShopAvatar     string `json:"shop_avatar"`
	ShopIsOfficial string `json:"shop_is_official"`
	ShopScore      int    `json:"shop_score"`
	ShopLocation   string `json:"shop_location"`
	Typename       string `json:"__typename"`
}

func (info *Info) UnmarshalJSON(data []byte) (err error) {
	type Alias Info
	aux := &struct {
		*Alias
		ShopID string `json:"shop_id"`
	}{
		Alias: (*Alias)(info),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return
	}

	if aux.ShopID == "" {
		info.ShopID = 0
	} else {
		info.ShopID, err = strconv.ParseInt(aux.ShopID, 10, 64)
		if err != nil {
			return
		}
	}

	return
}

type UserShopInfo struct {
	Info  Info `json:"info"`
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

var ErrNoShopid = errors.New("no shopid")

func (api *TokopediaApi) IsAutheticated() (*IsAuthenticatedRes, error) {
	var variable struct{}

	query := GraphqlPayload{
		OperationName: "isAuthenticatedQuery",
		Variables:     variable,
		Query: `query isAuthenticatedQuery {
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
	headers := map[string]string{
		"Content-Type":                     "application/json",
		"Access-Control-Allow-Headers":     "Content-type, Fingerprint-Data, Fingerprint-Hash, x-user-id, Webview-App-Version, Redirect, Access-Control-Allow-Origin, Content-MD5, Tkpd-UserId, X-Tkpd-UserId, Tkpd-SessionId, X-Device, X-Source, X-Method, X-Date, Authorization, Accounts-Authorization, flight-thirdparty, x-origin, Cshld-SessionID, X-Mitra-Device, x-tkpd-akamai, x-tkpd-lite-service, x-ga-id, x-device, Akamai-Bot, x-tkpd-app-name, x-tkpd-clc, x-return-hmac-md5, queryhash, Event, X-Element-ID, sid_intools, sonic_access_token, Referer, x-tkpd-ht, x-tkpd-htt, X-Auth-Signature, X-Auth-Timestamp, X-Auth-Hash, X-NewRelic-Id, newrelic, tracestate, traceparent",
		"Access-Control-Allow-Methods":     "GET, HEAD, PUT, PATCH, POST, DELETE",
		"x-tkpd-gql-dc":                    "gcp",
		"x-tkpd-srv-go":                    "go-graphql",
		"Access-Control-Allow-Origin":      "https://seller.tokopedia.com",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Expose-Headers":    "queryhash",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	var hasil IsAuthenticatedRes
	err := api.SendRequest(req, &hasil)

	if err != nil {
		return &hasil, err
	}

	api.AuthenticatedData = &hasil.Data
	if hasil.Data.UserShopInfo.Info.ShopID == 0 {
		return &hasil, ErrNoShopid
	}

	return &hasil, err
}
