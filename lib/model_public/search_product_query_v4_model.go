package model_public

type Violation struct {
	HeaderText      string `json:"headerText"`
	DescriptionText string `json:"descriptionText"`
	ImageURL        string `json:"imageURL"`
	CtaURL          string `json:"ctaURL"`
	CtaApplink      string `json:"ctaApplink"`
	ButtonText      string `json:"buttonText"`
	ButtonType      string `json:"buttonType"`
	Typename        string `json:"__typename"`
}

type HeaderSearchProduct struct {
	TotalData        int     `json:"totalData"`
	TotalDataText    string  `json:"totalDataText"`
	ProcessTime      float64 `json:"processTime"`
	ResponseCode     int     `json:"responseCode"`
	ErrorMessage     string  `json:"errorMessage"`
	AdditionalParams string  `json:"additionalParams"`
	KeywordProcess   string  `json:"keywordProcess"`
	ComponentID      string  `json:"componentId"`
	Typename         string  `json:"__typename"`
}

type SearchProductBanner struct {
	Position       int    `json:"position"`
	Text           string `json:"text"`
	ImageURL       string `json:"imageUrl"`
	URL            string `json:"url"`
	ComponentID    string `json:"componentId"`
	TrackingOption int    `json:"trackingOption"`
	Typename       string `json:"__typename"`
}

type Ticker struct {
	Text           string `json:"text"`
	Query          string `json:"query"`
	TypeID         int    `json:"typeId"`
	ComponentID    string `json:"componentId"`
	TrackingOption int    `json:"trackingOption"`
	Typename       string `json:"__typename"`
}

type Redirection struct {
	RedirectURL  string `json:"redirectUrl"`
	DepartmentID int    `json:"departmentId"`
	Typename     string `json:"__typename"`
}

type Related struct {
	Position       int           `json:"position"`
	TrackingOption int           `json:"trackingOption"`
	RelatedKeyword string        `json:"relatedKeyword"`
	OtherRelated   []interface{} `json:"otherRelated"`
	Typename       string        `json:"__typename"`
}

type Suggestion struct {
	CurrentKeyword  string `json:"currentKeyword"`
	Suggestion      string `json:"suggestion"`
	SuggestionCount int    `json:"suggestionCount"`
	Instead         string `json:"instead"`
	InsteadCount    int    `json:"insteadCount"`
	Query           string `json:"query"`
	Text            string `json:"text"`
	ComponentID     string `json:"componentId"`
	TrackingOption  int    `json:"trackingOption"`
	Typename        string `json:"__typename"`
}

type Ads struct {
	AdsID              string `json:"adsId"`
	ProductClickURL    string `json:"productClickUrl"`
	ProductWishlistURL string `json:"productWishlistUrl"`
	ProductViewURL     string `json:"productViewUrl"`
	Typename           string `json:"__typename"`
}

type ProductShop struct {
	ShopID       int         `json:"shopId"`
	Name         string      `json:"name"`
	City         string      `json:"city"`
	IsOfficial   bool        `json:"isOfficial"`
	IsPowerBadge bool        `json:"isPowerBadge"`
	URL          string      `json:"url,omitempty"`
	Domain       string      `json:"domain,omitempty"`
	Tagline      string      `json:"tagline,omitempty"`
	URI          string      `json:"uri,omitempty"`
	Badges       []ShopBadge `json:"badges,omitempty"`
	Typename     string      `json:"__typename"`
}

type ProductSearch struct {
	ID                 int64         `json:"id"`
	Name               string        `json:"name"`
	Ads                Ads           `json:"ads"`
	Badges             []interface{} `json:"badges"`
	Category           int           `json:"category"`
	CategoryBreadcrumb string        `json:"categoryBreadcrumb"`
	CategoryID         int           `json:"categoryId"`
	CategoryName       string        `json:"categoryName"`
	CountReview        int           `json:"countReview"`
	CustomVideoURL     string        `json:"customVideoURL"`
	DiscountPercentage int           `json:"discountPercentage"`
	GaKey              string        `json:"gaKey"`
	ImageURL           string        `json:"imageUrl"`
	LabelGroups        []LabelGroups `json:"labelGroups"`
	OriginalPrice      string        `json:"originalPrice"`
	Price              string        `json:"price"`
	PriceRange         string        `json:"priceRange"`
	Rating             int           `json:"rating"`
	RatingAverage      string        `json:"ratingAverage"`
	Shop               ProductShop   `json:"shop"`
	URL                string        `json:"url"`
	Wishlist           bool          `json:"wishlist"`
	SourceEngine       string        `json:"sourceEngine"`
	Typename           string        `json:"__typename"`
}

type SearchProductData struct {
	Banner         SearchProductBanner `json:"banner"`
	BackendFilters string              `json:"backendFilters"`
	IsQuerySafe    bool                `json:"isQuerySafe"`
	Ticker         Ticker              `json:"ticker"`
	Redirection    Redirection         `json:"redirection"`
	Related        Related             `json:"related"`
	Suggestion     Suggestion          `json:"suggestion"`
	Products       []*ProductSearch    `json:"products"`
	Violation      Violation           `json:"violation"`
	Typename       string              `json:"__typename"`
}

type SearchProductQueryV4Resp struct {
	Data struct {
		AceSearchProductV4 struct {
			Header   HeaderSearchProduct `json:"header"`
			Data     SearchProductData   `json:"data"`
			Typename string              `json:"__typename"`
		} `json:"ace_search_product_v4"`
	} `json:"data"`
}
