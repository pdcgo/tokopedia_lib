package model_public

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/gorilla/schema"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
)

var encoder = schema.NewEncoder()

type SearchProductVar struct {
	Sort            string   `json:"ob,omitempty" schema:"ob,omitempty" form:"ob,omitempty"`
	Page            int      `json:"page,omitempty" schema:"page,omitempty" form:"page,omitempty"`
	Rows            int      `json:"rows,omitempty" schema:"rows,omitempty" form:"rows,omitempty"`
	Device          string   `json:"device,omitempty" schema:"device,omitempty" form:"device,omitempty"`
	Related         bool     `json:"related,omitempty" schema:"related,omitempty" form:"related,omitempty"`
	SafeSearch      bool     `json:"safe_search,omitempty" schema:"safe_search,omitempty" form:"safe_search,omitempty"`
	Scheme          string   `json:"scheme,omitempty" schema:"scheme,omitempty" form:"scheme,omitempty"`
	UserDistrictID  string   `json:"user_districtId,omitempty" schema:"user_districtId,omitempty" form:"user_districtId,omitempty"`
	UserCityID      string   `json:"user_cityId,omitempty" schema:"user_cityId,omitempty" form:"user_cityId,omitempty"`
	Source          string   `json:"source,omitempty" schema:"source,omitempty" form:"source,omitempty"`
	TopadsBucket    bool     `json:"topads_bucket,omitempty" schema:"topads_bucket,omitempty" form:"topads_bucket,omitempty"`
	Start           int      `json:"start,omitempty" schema:"start,omitempty" form:"start,omitempty"`
	PriceMin        int      `json:"pmin,omitempty" schema:"pmin,omitempty" form:"pmin,omitempty"`
	PriceMax        int      `json:"pmax,omitempty" schema:"pmax,omitempty" form:"pmax,omitempty"`
	Rate            []string `json:"rt,omitempty" schema:"rt,omitempty" form:"rt,omitempty"`
	Query           string   `json:"q,omitempty" schema:"q,omitempty" form:"q,omitempty"`
	Fcity           []string `json:"fcity,omitempty" schema:"fcity,omitempty" form:"fcity,omitempty"`
	Condition       []string `json:"condition,omitempty" schema:"condition,omitempty" form:"condition,omitempty"`
	ShopTier        []string `json:"shop_tier,omitempty" schema:"shop_tier,omitempty" form:"shop_tier,omitempty"`
	CategoryId      int      `json:"sc,omitempty" schema:"sc,omitempty" form:"sc,omitempty"`
	Identifier      string   `json:"identifier,omitempty" schema:"identifier,omitempty" form:"identifier,omitempty"`
	Navsource       string   `json:"navsource,omitempty" schema:"navsource,omitempty" form:"navsource,omitempty"`
	UniqueId        string   `json:"unique_id,omitempty" schema:"unique_id,omitempty" form:"unique_id,omitempty"`
	Shipping        []string `json:"shipping,omitempty" schema:"shipping,omitempty" form:"shipping,omitempty"`
	PreOrder        bool     `json:"preorder,omitempty" schema:"preorder,omitempty" form:"preorder,omitempty"`
	Etalase         string   `json:"etalase,omitempty" schema:"etalase,omitempty" form:"etalase,omitempty"`
	ShopId          int      `json:"shop_id,omitempty" schema:"shop_id,omitempty" form:"shop_id,omitempty"`
	UserWarehouseId int      `json:"user_warehouseId,omitempty" schema:"user_warehouseId,omitempty" form:"user_warehouseId,omitempty"`
}

func (v *SearchProductVar) GetQuery() string {

	encoder.RegisterEncoder([]string{}, func(v reflect.Value) string {
		return strings.Join(v.Interface().([]string), "#")
	})

	encoder.RegisterEncoder("", func(v reflect.Value) string {
		return url.QueryEscape(v.Interface().(string))
	})

	query := url.Values{}
	encoder.Encode(v, query)

	return query.Encode()
}

func NewSearchProductVar() *SearchProductVar {
	productVar := SearchProductVar{
		Device:         "desktop",
		Sort:           "8",
		Page:           1,
		Rows:           100,
		UserDistrictID: "176",
		UserCityID:     "2274",
		Related:        true,
		Scheme:         "https",
		SafeSearch:     false,
		TopadsBucket:   true,
		Source:         "search",
	}

	return &productVar
}

func NewGrabSearchProductVar(grabTokopedia *legacy.GrabTokopedia) *SearchProductVar {

	searchVar := NewSearchProductVar()

	searchVar.PriceMin = grabTokopedia.Query.Pmin
	searchVar.PriceMax = grabTokopedia.Query.Pmax
	searchVar.PreOrder = grabTokopedia.Query.Preorder
	searchVar.Fcity = grabTokopedia.Query.Fcity
	searchVar.Shipping = grabTokopedia.Query.Shipping
	searchVar.Sort = grabTokopedia.Query.Ob
	searchVar.Rate = strings.Split(grabTokopedia.Query.Rt, ",")
	searchVar.Condition = strings.Split(grabTokopedia.Query.Condition, ",")

	if grabTokopedia.Query.Official {
		searchVar.ShopTier = append(searchVar.ShopTier, "2")
	}

	if grabTokopedia.Query.Goldmerchant {
		searchVar.ShopTier = append(searchVar.ShopTier, "3")
	}

	return searchVar
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
