package model

type ShipmentInfoProduct struct {
	IsAvailable int    `json:"isAvailable"`
	ProductName string `json:"productName"`
	UIHidden    bool   `json:"uiHidden"`
	Typename    string `json:"__typename"`
}

type Error struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}

type RecentOneMonth struct {
	Bad      int    `json:"bad"`
	Good     int    `json:"good"`
	Neutral  int    `json:"neutral"`
	Typename string `json:"__typename"`
}

type ShopSatisfaction struct {
	RecentOneMonth RecentOneMonth `json:"recentOneMonth"`
	Typename       string         `json:"__typename"`
}

type RatingDetail struct {
	FormattedTotalReviews string  `json:"formattedTotalReviews"`
	Rate                  int     `json:"rate"`
	Percentage            string  `json:"percentage,omitempty"`
	PercentageFloat       float64 `json:"percentageFloat"`
	TotalReviews          int     `json:"totalReviews"`
	Typename              string  `json:"__typename"`
}

type ShopRating struct {
	Detail      []RatingDetail `json:"detail"`
	TotalRating int            `json:"totalRating"`
	RatingScore string         `json:"ratingScore"`
	Typename    string         `json:"__typename"`
}

type ShopReputation struct {
	Badge    string `json:"badge"`
	Score    string `json:"score"`
	ScoreMap string `json:"score_map"`
	Typename string `json:"__typename"`
}

type ShopStatisticQueryData struct {
	ShopSatisfaction ShopSatisfaction `json:"shopSatisfaction"`
	ShopRating       ShopRating       `json:"shopRating"`
	ShopReputation   []ShopReputation `json:"shopReputation"`
}

type GetShopOperationalHourStatus struct {
	Timestamp        string `json:"timestamp"`
	StatusActive     bool   `json:"statusActive"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	TimestampUTCUnix string `json:"timestampUTCUnix"`
	TickerTitle      string `json:"tickerTitle"`
	TickerMessage    string `json:"tickerMessage"`
	Typename         string `json:"__typename"`
}

type Amount struct {
	Amount          int    `json:"amount"`
	AmountFormatted string `json:"amountFormatted"`
	AmountType      int    `json:"amountType"`
	Typename        string `json:"__typename"`
}

type Owner struct {
	Identifier string `json:"identifier"`
	Typename   string `json:"__typename"`
}

type Status struct {
	Identifier string `json:"identifier"`
	Status     int    `json:"status"`
	Typename   string `json:"__typename"`
}
type VoucherType struct {
	Identifier  string `json:"identifier"`
	VoucherType int    `json:"voucherType"`
	Typename    string `json:"__typename"`
}

type Banner struct {
	DesktopURL string `json:"desktopUrl"`
	Typename   string `json:"__typename"`
}

type Vouchers struct {
	Amount                Amount      `json:"amount"`
	InUseExpiry           string      `json:"inUseExpiry"`
	MinimumSpend          int         `json:"minimumSpend"`
	MinimumSpendFormatted string      `json:"minimumSpendFormatted"`
	Owner                 Owner       `json:"owner"`
	Status                Status      `json:"status"`
	ValidThru             string      `json:"validThru"`
	VoucherID             int         `json:"voucherID"`
	Name                  string      `json:"name"`
	VoucherCode           string      `json:"voucherCode"`
	VoucherType           VoucherType `json:"voucherType"`
	Banner                Banner      `json:"banner"`
	Tnc                   string      `json:"tnc"`
	Typename              string      `json:"__typename"`
}

type GetPublicMerchantVoucherList struct {
	Vouchers []Vouchers `json:"vouchers"`
	Typename string     `json:"__typename"`
}

type ShopNoteResults struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Link       string `json:"link"`
	UpdateTime string `json:"updateTime"`
	Typename   string `json:"__typename"`
}

type ShopNotesByShopID struct {
	Result   []ShopNoteResults `json:"result"`
	Error    Error             `json:"error"`
	Typename string            `json:"__typename"`
}

type WidgetHeader struct {
	Ratio             string `json:"ratio"`
	SizeOption        string `json:"sizeOption,omitempty"`
	Title             string `json:"title"`
	Subtitle          string `json:"subtitle,omitempty"`
	CtaText           string `json:"ctaText,omitempty"`
	CtaLink           string `json:"ctaLink,omitempty"`
	Cover             string `json:"cover,omitempty"`
	IsATC             int    `json:"isATC,omitempty"`
	IsActive          int    `json:"isActive,omitempty"`
	EtalaseID         string `json:"etalaseID,omitempty"`
	IsShowEtalaseName int    `json:"isShowEtalaseName,omitempty"`
	Typename          string `json:"__typename"`
}

type WidgetRequest struct {
	WidgetID       int          `json:"widgetID"`
	WidgetMasterID int          `json:"widgetMasterID"`
	WidgetType     string       `json:"widgetType"`
	WidgetName     string       `json:"widgetName"`
	Header         WidgetHeader `json:"header"`
	Typename       string       `json:"__typename"`
}

type StockWording struct {
	Title    string `json:"title"`
	Typename string `json:"__typename"`
}

type Products struct {
	ID                 int64         `json:"id"`
	Name               string        `json:"name"`
	URL                string        `json:"url"`
	URLApps            string        `json:"urlApps"`
	URLMobile          string        `json:"urlMobile"`
	ImageURL           string        `json:"imageURL"`
	Price              string        `json:"price"`
	CountSold          int           `json:"countSold"`
	Stock              int           `json:"stock"`
	Status             string        `json:"status"`
	DiscountedPrice    string        `json:"discountedPrice"`
	DiscountPercentage int           `json:"discountPercentage"`
	Position           int           `json:"position"`
	Rating             float64       `json:"rating"`
	StockWording       StockWording  `json:"stockWording"`
	HideGimmick        bool          `json:"hideGimmick"`
	LabelGroups        []LabelGroups `json:"labelGroups"`
	Typename           string        `json:"__typename"`
}

type BackgroundGradientColor struct {
	FirstColor  string `json:"firstColor"`
	SecondColor string `json:"secondColor"`
	Typename    string `json:"__typename"`
}

type DynamicRule struct {
	DynamicRoleData   []interface{} `json:"dynamicRoleData"`
	DescriptionHeader string        `json:"descriptionHeader"`
	Typename          string        `json:"__typename"`
}

type Banners struct {
	ImageID    string `json:"imageID"`
	ImageURL   string `json:"imageURL"`
	BannerType string `json:"bannerType"`
	Typename   string `json:"__typename"`
}

type WidgetData struct {
	CampaignID              string                  `json:"campaignID"`
	Name                    string                  `json:"name"`
	Description             string                  `json:"description"`
	StartDate               string                  `json:"startDate"`
	EndDate                 string                  `json:"endDate"`
	StatusCampaign          string                  `json:"statusCampaign"`
	TimeDescription         string                  `json:"timeDescription"`
	TimeCounter             string                  `json:"timeCounter"`
	TotalNotify             int                     `json:"totalNotify"`
	TotalNotifyWording      string                  `json:"totalNotifyWording"`
	TotalProduct            int                     `json:"totalProduct"`
	TotalProductWording     string                  `json:"totalProductWording"`
	BackgroundGradientColor BackgroundGradientColor `json:"backgroundGradientColor"`
	DynamicRule             DynamicRule             `json:"dynamicRule"`
	Banners                 []Banners               `json:"banners"`
	Products                []Products              `json:"products"`
	Typename                string                  `json:"__typename"`
}

type Widget struct {
	WidgetID       int          `json:"widgetID"`
	WidgetMasterID int          `json:"widgetMasterID"`
	LayoutOrder    int          `json:"layoutOrder"`
	Name           string       `json:"name"`
	Type           string       `json:"type"`
	Header         WidgetHeader `json:"header"`
	Data           []WidgetData `json:"data"`
	Typename       string       `json:"__typename"`
}

type ShopPageGetLayout struct {
	Widgets  []Widget `json:"widgets"`
	Typename string   `json:"__typename"`
}

////////////////////////////////////////////////////

type ShopStatisticQueryVar struct {
	ShopID    int    `json:"shopID"`
	ShopIDStr string `json:"shopIDStr"`
}

type ShopStatisticQueryResp struct {
	Data ShopStatisticQueryData `json:"data"`
}

type ShopIdVarInt struct {
	ShopID int64 `json:"shopID"`
}

type ShopIdVar struct {
	ShopID string `json:"shopID"`
}

type GetShopOperationalHourStatusResp struct {
	Data struct {
		GetShopOperationalHourStatus GetShopOperationalHourStatus `json:"getShopOperationalHourStatus"`
	} `json:"data"`
}

type ShopVoucherQueryResp struct {
	Data struct {
		GetPublicMerchantVoucherList GetPublicMerchantVoucherList `json:"getPublicMerchantVoucherList"`
	} `json:"data"`
}

type ShopNoteVar struct {
	ID  string `json:"id"`
	Sid string `json:"sid"`
}

type ShopNoteResp struct {
	Data struct {
		ShopNotesByShopID ShopNotesByShopID `json:"shopNotesByShopID"`
	} `json:"data"`
}

type ShopPageGetLayoutV2Var struct {
	ShopID        string          `json:"shopID"`
	DistrictID    string          `json:"districtID"`
	CityID        string          `json:"cityID"`
	Latitude      string          `json:"latitude"`
	Longitude     string          `json:"longitude"`
	WidgetRequest []WidgetRequest `json:"widgetRequest"`
}

type ShopPageGetLayoutV2Resp []struct {
	Data struct {
		ShopPageGetLayout ShopPageGetLayout `json:"shopPageGetLayout"`
	} `json:"data"`
}
