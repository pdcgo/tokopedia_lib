package model_public

type SearchProductVar struct {
	Sort           string `json:"ob,omitempty"`    // 23
	Page           int    `json:"page"`            // 1
	Rows           int    `json:"rows,omitempty"`  // 80
	Device         string `json:"device"`          // desktop
	Related        bool   `json:"related"`         // true
	SafeSearch     bool   `json:"safe_search"`     // false
	Scheme         string `json:"scheme"`          // https
	UserDistrictID string `json:"user_districtId"` // 176
	UserCityID     string `json:"user_cityId"`     // 2274
	Source         string `json:"source"`          // search
	TopadsBucket   bool   `json:"topads_bucket"`   // true
	Start          int    `json:"start"`
	PriceMin       int    `json:"pmin,omitempty"`
	PriceMax       int    `json:"pmax,omitempty"`
	Rate           string `json:"rt,omitempty"`
	Query          string `json:"q,omitempty"`
	Locations      string `json:"fcity,omitempty"`
	Condition      string `json:"condition,omitempty"`
	ShopTier       string `json:"shop_tier,omitempty"`
	CategoryId     int    `json:"sc,omitempty"`
	Identifier     string `json:"identifier,omitempty"`
	Navsource      string `json:"navsource"`
	UniqueId       string `json:"unique_id"`
	Shipping       string `json:"shipping,omitempty"`
	PreOrder       bool   `json:"preorder,omitempty"`
}

type SearchProductAdParams struct {
	Page          int     `json:"page"`              // 1
	Ep            string  `json:"ep"`                // product
	Item          int     `json:"item"`              // 15
	Src           string  `json:"src"`               // directory
	Device        string  `json:"device"`            // desktop
	UserId        int     `json:"user_id"`           // 0
	MinimumItem   int     `json:"minimum_item"`      // 15
	Start         int     `json:"start"`             // 1
	NoAutoFill    string  `json:"no_autofill_range"` // 5-14
	Query         string  `json:"q,omitempty"`
	Fcity         []int   `json:"fcity,omitempty"`
	DepId         int     `json:"dep_id"`
	Sort          int     `json:"ob"`
	Shipping      []int   `json:"shipping,omitempty"`
	Cashback      bool    `json:"cashbackm,omitempty"`
	FreeShipping  bool    `json:"free_shipping,omitempty"`
	COD           bool    `json:"cod,omitempty"`
	IsDiscount    bool    `json:"is_discount,omitempty"`
	Bundling      bool    `json:"bundling,omitempty"`
	Wholesale     bool    `json:"wholesale,omitempty"`
	IsMvcDiscount bool    `json:"is_mvc_discount,omitempty"`
	ShopTier      []int   `json:"shop_tier,omitempty"`
	PriceMin      int     `json:"pmin,omitempty"`
	PriceMax      int     `json:"pmax,omitempty"`
	PreOrder      bool    `json:"preorder,omitempty"`
	Condition     []int   `json:"condition,omitempty"`
	Rating        float64 `json:"rt,omitempty"`
}

////////////////////////////////////////////

type SearchProductQueryVar struct {
	AdParams string `json:"adParams"`
	Params   string `json:"params"`
}

type CategoryProduct struct {
	ID                 int64         `json:"id"`
	URL                string        `json:"url"`
	ImageURL           string        `json:"imageUrl"`
	ImageURLLarge      string        `json:"imageUrlLarge"`
	CatID              int           `json:"catId"`
	GaKey              string        `json:"gaKey"`
	CountReview        int           `json:"countReview"`
	DiscountPercentage int           `json:"discountPercentage"`
	Preorder           bool          `json:"preorder"`
	Name               string        `json:"name"`
	Price              string        `json:"price"`
	PriceInt           int           `json:"priceInt"`
	OriginalPrice      string        `json:"original_price"`
	Rating             int           `json:"rating"`
	Wishlist           bool          `json:"wishlist"`
	Labels             []interface{} `json:"labels"`
	Badges             []struct {
		ImageURL string `json:"imageUrl"`
		Show     bool   `json:"show"`
		Typename string `json:"__typename"`
	} `json:"badges"`
	Shop struct {
		ID           int    `json:"id"`
		URL          string `json:"url"`
		Name         string `json:"name"`
		Goldmerchant bool   `json:"goldmerchant"`
		Official     bool   `json:"official"`
		Reputation   string `json:"reputation"`
		Clover       string `json:"clover"`
		Location     string `json:"location"`
		Typename     string `json:"__typename"`
	} `json:"shop"`
	LabelGroups []LabelGroups `json:"labelGroups"`
	Typename    string        `json:"__typename"`
}

type CategoryProducts struct {
	Count    int               `json:"count"`
	Data     []CategoryProduct `json:"data"`
	Typename string            `json:"__typename"`
}

type SearchProductQueryResp struct {
	Data struct {
		CategoryProducts CategoryProducts `json:"CategoryProducts"`
		DisplayAdsV3     struct {
			Data     []interface{} `json:"data"`
			Template []interface{} `json:"template"`
			Typename string        `json:"__typename"`
		} `json:"displayAdsV3"`
	} `json:"data"`
}
