package model_public

type Links struct {
	Prev     string `json:"prev"`
	Next     string `json:"next"`
	Typename string `json:"__typename"`
}

type PrimaryImage struct {
	Original  string `json:"original"`
	Thumbnail string `json:"thumbnail"`
	Resize300 string `json:"resize300"`
	Typename  string `json:"__typename"`
}

type Flags struct {
	IsSold      bool   `json:"isSold"`
	IsPreorder  bool   `json:"isPreorder"`
	IsWholesale bool   `json:"isWholesale"`
	IsWishlist  bool   `json:"isWishlist"`
	Typename    string `json:"__typename"`
}

type Campaign struct {
	DiscountedPercentage string `json:"discounted_percentage"`
	OriginalPriceFmt     string `json:"original_price_fmt"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	Typename             string `json:"__typename"`
}

type LabelGroups struct {
	Position string `json:"position"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
}

type Badge struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Typename string `json:"__typename"`
}

type Label struct {
	ColorHex string `json:"color_hex"`
	Content  string `json:"content"`
	Typename string `json:"__typename"`
}

type Stats struct {
	CountView     int    `json:"countView,omitempty"`
	CountReview   int    `json:"countReview,omitempty"`
	CountTalk     int    `json:"countTalk,omitempty"`
	ReviewCount   int    `json:"reviewCount,omitempty"`
	Rating        int    `json:"rating,omitempty"`
	AverageRating string `json:"averageRating,omitempty"`
	Typename      string `json:"__typename"`
}

type Price struct {
	Min      int    `json:"min,omitempty"`
	Max      int    `json:"max,omitempty"`
	TextIdr  string `json:"text_idr,omitempty"`
	Typename string `json:"__typename"`
}

type ShopProductData struct {
	Name         string        `json:"name"`
	ProductURL   string        `json:"product_url"`
	ProductID    string        `json:"product_id"`
	Price        Price         `json:"price"`
	PrimaryImage PrimaryImage  `json:"primary_image"`
	Flags        Flags         `json:"flags"`
	Campaign     Campaign      `json:"campaign"`
	Label        []Label       `json:"label"`
	LabelGroups  []LabelGroups `json:"label_groups"`
	Badge        []Badge       `json:"badge"`
	Stats        Stats         `json:"stats"`
	Category     Category      `json:"category"`
	Typename     string        `json:"__typename"`
}

type GetShopProduct struct {
	Status   string             `json:"status"`
	Errors   string             `json:"errors"`
	Links    Links              `json:"links"`
	Data     []*ShopProductData `json:"data"`
	Typename string             `json:"__typename"`
}

///////////////////////////////////////////////////////

type ShopProductVar struct {
	Sid            string `json:"sid"`
	Page           int    `json:"page"`
	PerPage        int    `json:"perPage"`
	EtalaseID      string `json:"etalaseId"`
	Sort           int    `json:"sort"`
	UserDistrictID string `json:"user_districtId"`
	UserCityID     string `json:"user_cityId"`
	UserLat        string `json:"user_lat"`
	UserLong       string `json:"user_long"`
}

func NewShopProductVar(shopId string) *ShopProductVar {
	params := &ShopProductVar{
		Page:           1,
		PerPage:        100,
		EtalaseID:      "etalase",
		Sort:           1,
		Sid:            shopId,
		UserDistrictID: "176",
		UserCityID:     "2274",
		UserLat:        "",
		UserLong:       "",
	}

	return params
}

type ShopProductResp struct {
	Data struct {
		GetShopProduct GetShopProduct `json:"GetShopProduct"`
	} `json:"data"`
}
