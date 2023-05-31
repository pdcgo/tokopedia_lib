package model_public

type ShopSuggestion struct {
	CurrentKeyword      string `json:"currentKeyword"`
	SuggestionText      string `json:"suggestionText"`
	SuggestionTextQuery string `json:"suggestionTextQuery"`
	Typename            string `json:"__typename"`
}

type VoucherCashback struct {
	CashbackValue int    `json:"cashbackValue"`
	IsPercentage  bool   `json:"isPercentage"`
	Typename      string `json:"__typename"`
}

type ShopVoucher struct {
	FreeShipping bool            `json:"freeShipping"`
	Cashback     VoucherCashback `json:"cashback"`
	Typename     string          `json:"__typename"`
}

type Header struct {
	KeywordProcess string `json:"keywordProcess"`
	ResponseCode   int    `json:"responseCode"`
	TotalData      int    `json:"totalData"`
	Typename       string `json:"__typename"`
}

type ShopProduct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Price      int    `json:"price"`
	ProductImg string `json:"productImg"`
	PriceText  string `json:"priceText"`
	Typename   string `json:"__typename"`
}

type SearchShop struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Domain             string        `json:"domain"`
	OwnerID            int           `json:"ownerId"`
	City               string        `json:"city"`
	ShopStatus         int           `json:"shopStatus"`
	TagLine            string        `json:"tagLine"`
	Desc               string        `json:"desc"`
	ReputationScore    int           `json:"reputationScore"`
	TotalFave          string        `json:"totalFave"`
	IsPowerBadge       int           `json:"isPowerBadge"`
	IsPMPro            bool          `json:"isPMPro"`
	IsOfficial         bool          `json:"isOfficial"`
	URL                string        `json:"url"`
	ImageURL           string        `json:"imageURL"`
	ReputationImageURL string        `json:"reputationImageURL"`
	ShopLucky          string        `json:"shopLucky"`
	Products           []ShopProduct `json:"products"`
	GAKey              string        `json:"GAKey"`
	Favorited          bool          `json:"favorited"`
	Voucher            ShopVoucher   `json:"voucher"`
	Typename           string        `json:"__typename"`
}

type AceSearchShopResp struct {
	Data struct {
		AceSearchShop struct {
			TotalData  int            `json:"totalData"`
			Shops      []SearchShop   `json:"shops"`
			Suggestion ShopSuggestion `json:"suggestion"`
			Header     Header         `json:"header"`
			Typename   string         `json:"__typename"`
		} `json:"aceSearchShop"`
	} `json:"data"`
}
