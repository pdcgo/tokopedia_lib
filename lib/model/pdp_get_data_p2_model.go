package model

type Media struct {
	Type            string `json:"type"`
	URLOriginal     string `json:"urlOriginal"`
	URLThumbnail    string `json:"urlThumbnail"`
	URLMaxRes       string `json:"urlMaxRes"`
	VideoURL        string `json:"videoUrl"`
	Prefix          string `json:"prefix"`
	Suffix          string `json:"suffix"`
	Description     string `json:"description"`
	VariantOptionID string `json:"variantOptionID"`
	Typename        string `json:"__typename"`
}

type ComponentData struct {
	Media    []Media       `json:"media"`
	Videos   []interface{} `json:"videos"`
	Typename string        `json:"__typename"`
}

type Component struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Position string          `json:"position"`
	Data     []ComponentData `json:"data"`
	Typename string          `json:"__typename"`
}

type ProductTxStats struct {
	TransactionSuccess string `json:"transactionSuccess"`
	TransactionReject  string `json:"transactionReject"`
	CountSold          string `json:"countSold"`
	PaymentVerified    string `json:"paymentVerified"`
	ItemSoldFmt        string `json:"itemSoldFmt"`
	Typename           string `json:"__typename"`
}

type BasicInfoStats struct {
	CountView   string  `json:"countView"`
	CountReview string  `json:"countReview"`
	CountTalk   string  `json:"countTalk"`
	Rating      float64 `json:"rating"`
	Typename    string  `json:"__typename"`
}

type BasicInfo struct {
	Alias            string         `json:"alias"`
	CreatedAt        string         `json:"createdAt"`
	IsQA             bool           `json:"isQA"`
	ID               string         `json:"id"`
	ShopID           string         `json:"shopID"`
	ShopName         string         `json:"shopName"`
	MinOrder         int            `json:"minOrder"`
	MaxOrder         int            `json:"maxOrder"`
	Weight           int            `json:"weight"`
	WeightUnit       string         `json:"weightUnit"`
	Condition        string         `json:"condition"`
	Status           string         `json:"status"`
	URL              string         `json:"url"`
	NeedPrescription bool           `json:"needPrescription"`
	CatalogID        string         `json:"catalogID"`
	IsLeasing        bool           `json:"isLeasing"`
	IsBlacklisted    bool           `json:"isBlacklisted"`
	IsTokoNow        bool           `json:"isTokoNow"`
	Menu             Menu           `json:"menu"`
	Category         Category       `json:"category"`
	TxStats          ProductTxStats `json:"txStats"`
	Stats            BasicInfoStats `json:"stats"`
	Typename         string         `json:"__typename"`
}

type PdpGetLayout struct {
	RequestID  string      `json:"requestID"`
	Name       string      `json:"name"`
	PdpSession string      `json:"pdpSession"`
	BasicInfo  BasicInfo   `json:"basicInfo"`
	Components []Component `json:"components"`
	Typename   string      `json:"__typename"`
}

type OwnerInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type ErrNearestWarehouse struct {
	Code     int    `json:"Code"`
	Message  string `json:"Message"`
	Typename string `json:"__typename"`
}

type NearestWarehouse struct {
	ProductID     string        `json:"product_id"`
	Stock         string        `json:"stock"`
	StockWording  string        `json:"stock_wording"`
	Price         string        `json:"price"`
	WarehouseInfo WarehouseInfo `json:"warehouse_info"`
	Typename      string        `json:"__typename"`
}

type ErrCartRedirection struct {
	Code     int    `json:"Code"`
	Message  string `json:"Message"`
	Typename string `json:"__typename"`
}

type ErrShopInfo struct {
	Code     int    `json:"Code"`
	Message  string `json:"Message"`
	Typename string `json:"__typename"`
}

type DevError struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	DevMessage string `json:"DevMessage"`
	Typename   string `json:"__typename"`
}

type CallsError struct {
	ShopInfo         ErrShopInfo         `json:"shopInfo"`
	CartRedirection  ErrCartRedirection  `json:"cartRedirection"`
	NearestWarehouse ErrNearestWarehouse `json:"nearestWarehouse"`
	Typename         string              `json:"__typename"`
}

type ShopFinishRate struct {
	FinishRate string `json:"finishRate"`
	Typename   string `json:"__typename"`
}

type MerchantVoucherType struct {
	VoucherType int    `json:"voucher_type"`
	Identifier  string `json:"identifier"`
	Typename    string `json:"__typename"`
}

type MerchantVoucherAmount struct {
	Amount          int    `json:"amount"`
	AmountType      int    `json:"amount_type"`
	AmountFormatted string `json:"amount_formatted"`
	Typename        string `json:"__typename"`
}

type MerchantVoucherBanner struct {
	DesktopURL string `json:"desktop_url"`
	MobileURL  string `json:"mobile_url"`
	Typename   string `json:"__typename"`
}

type WarehouseInfo struct {
	WarehouseID   string `json:"warehouse_id"`
	IsFulfillment bool   `json:"is_fulfillment"`
	DistrictID    string `json:"district_id"`
	PostalCode    string `json:"postal_code"`
	Geolocation   string `json:"geolocation"`
	Typename      string `json:"__typename"`
}

type ProductsBebasOngkir struct {
	ProductID string `json:"productID"`
	BoType    int    `json:"boType"`
	Typename  string `json:"__typename"`
}

type MerchantVoucher struct {
	VoucherID    int                   `json:"voucher_id"`
	VoucherName  string                `json:"voucher_name"`
	VoucherType  MerchantVoucherType   `json:"voucher_type"`
	VoucherCode  string                `json:"voucher_code"`
	Amount       MerchantVoucherAmount `json:"amount"`
	MinimumSpend int                   `json:"minimum_spend"`
	ValidThru    string                `json:"valid_thru"`
	Tnc          string                `json:"tnc"`
	Banner       MerchantVoucherBanner `json:"banner"`
	Status       Status                `json:"status"`
	InUseExpiry  string                `json:"in_use_expiry"`
	Typename     string                `json:"__typename"`
}

type InstallmentRecommendationData struct {
	Term           int    `json:"term"`
	MdrValue       int    `json:"mdr_value"`
	MdrType        string `json:"mdr_type"`
	InterestRate   int    `json:"interest_rate"`
	MinimumAmount  int    `json:"minimum_amount"`
	MaximumAmount  int    `json:"maximum_amount"`
	MonthlyPrice   int    `json:"monthly_price"`
	OsMonthlyPrice int    `json:"os_monthly_price"`
	PartnerCode    string `json:"partner_code"`
	PartnerName    string `json:"partner_name"`
	PartnerIcon    string `json:"partner_icon"`
	Subtitle       string `json:"subtitle"`
	Typename       string `json:"__typename"`
}

type FulfillmentData struct {
	Icon        string `json:"icon"`
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type Bottomsheet struct {
	Title      string `json:"title"`
	IconURL    string `json:"iconURL"`
	Subtitle   string `json:"subtitle"`
	ButtonCopy string `json:"buttonCopy"`
	Typename   string `json:"__typename"`
}

type ShopInfoFavoriteData struct {
	TotalFavorite    string `json:"totalFavorite"`
	AlreadyFavorited int    `json:"alreadyFavorited"`
	Typename         string `json:"__typename"`
}

type ShopInfo struct {
	ShopTier      int                  `json:"shopTier"`
	BadgeURL      string               `json:"badgeURL"`
	ClosedInfo    ClosedInfo           `json:"closedInfo"`
	IsOpen        int                  `json:"isOpen"`
	FavoriteData  ShopInfoFavoriteData `json:"favoriteData"`
	ActiveProduct string               `json:"activeProduct"`
	CreateInfo    CreateInfo           `json:"createInfo"`
	ShopAssets    ShopAssets           `json:"shopAssets"`
	ShopCore      struct {
		Domain    string `json:"domain"`
		ShopID    string `json:"shopID"`
		Name      string `json:"name"`
		ShopScore int    `json:"shopScore"`
		URL       string `json:"url"`
		OwnerID   string `json:"ownerID"`
		Typename  string `json:"__typename"`
	} `json:"shopCore"`
	ShopLastActive string        `json:"shopLastActive"`
	Location       string        `json:"location"`
	StatusInfo     StatusInfo    `json:"statusInfo"`
	IsAllowManage  int           `json:"isAllowManage"`
	IsOwner        int           `json:"isOwner"`
	OwnerInfo      OwnerInfo     `json:"ownerInfo"`
	IsCOD          bool          `json:"isCOD"`
	ShopType       int           `json:"shopType"`
	TickerData     []interface{} `json:"tickerData"`
	PartnerLabel   string        `json:"partnerLabel"`
	Typename       string        `json:"__typename"`
}

type CartRedirection struct {
	Status       string        `json:"status"`
	ErrorMessage []interface{} `json:"error_message"`
	Data         []struct {
		ProductID          string `json:"product_id"`
		ConfigName         string `json:"config_name"`
		HideFloatingButton bool   `json:"hide_floating_button"`
		AvailableButtons   []struct {
			Text               string `json:"text"`
			Color              string `json:"color"`
			CartType           string `json:"cart_type"`
			OnboardingMessage  string `json:"onboarding_message"`
			ShowRecommendation bool   `json:"show_recommendation"`
			Typename           string `json:"__typename"`
		} `json:"available_buttons"`
		UnavailableButtons []string `json:"unavailable_buttons"`
		Typename           string   `json:"__typename"`
	} `json:"data"`
	Typename string `json:"__typename"`
}

type RatesEstimateData struct {
	Destination           string          `json:"destination"`
	Title                 string          `json:"title"`
	Subtitle              string          `json:"subtitle"`
	ChipsLabel            []string        `json:"chipsLabel"`
	CourierLabel          string          `json:"courierLabel"`
	ETAText               string          `json:"eTAText"`
	CheapestShippingPrice int             `json:"cheapestShippingPrice"`
	FulfillmentData       FulfillmentData `json:"fulfillmentData"`
	Errors                []interface{}   `json:"errors"`
	Typename              string          `json:"__typename"`
}

type RatesEstimate struct {
	WarehouseID string            `json:"warehouseID"`
	Products    []string          `json:"products"`
	Data        RatesEstimateData `json:"data"`
	Bottomsheet Bottomsheet       `json:"bottomsheet"`
	Typename    string            `json:"__typename"`
}

type PdpGetData struct {
	Error           DevError       `json:"error"`
	CallsError      CallsError     `json:"callsError"`
	ProductView     string         `json:"productView"`
	WishlistCount   string         `json:"wishlistCount"`
	ShopFinishRate  ShopFinishRate `json:"shopFinishRate"`
	ShopInfo        ShopInfo       `json:"shopInfo"`
	MerchantVoucher struct {
		Vouchers []MerchantVoucher `json:"vouchers"`
		Typename string            `json:"__typename"`
	} `json:"merchantVoucher"`
	NearestWarehouse          []NearestWarehouse `json:"nearestWarehouse"`
	InstallmentRecommendation struct {
		Data     InstallmentRecommendationData `json:"data"`
		Typename string                        `json:"__typename"`
	} `json:"installmentRecommendation"`
	ProductWishlistQuery struct {
		Value    bool   `json:"value"`
		Typename string `json:"__typename"`
	} `json:"productWishlistQuery"`
	CartRedirection  CartRedirection `json:"cartRedirection"`
	ShopTopChatSpeed struct {
		MessageResponseTime string `json:"messageResponseTime"`
		Typename            string `json:"__typename"`
	} `json:"shopTopChatSpeed"`
	ShopRatingsQuery struct {
		RatingScore float64 `json:"ratingScore"`
		Typename    string  `json:"__typename"`
	} `json:"shopRatingsQuery"`
	ShopPackSpeed struct {
		SpeedFmt string `json:"speedFmt"`
		Hour     int    `json:"hour"`
		Typename string `json:"__typename"`
	} `json:"shopPackSpeed"`
	RatesEstimate   []RatesEstimate `json:"ratesEstimate"`
	RestrictionInfo struct {
		Message         string        `json:"message"`
		RestrictionData []interface{} `json:"restrictionData"`
		Typename        string        `json:"__typename"`
	} `json:"restrictionInfo"`
	Ticker struct {
		TickerInfo []interface{} `json:"tickerInfo"`
		Typename   string        `json:"__typename"`
	} `json:"ticker"`
	NavBar struct {
		Name     string        `json:"name"`
		Items    []interface{} `json:"items"`
		Typename string        `json:"__typename"`
	} `json:"navBar"`
	BebasOngkir struct {
		Products []ProductsBebasOngkir `json:"products"`
		Typename string                `json:"__typename"`
	} `json:"bebasOngkir"`
	Typename string `json:"__typename"`
}

type PdpGetDataP2Var struct {
	Affiliate    interface{}  `json:"affiliate"`
	ProductID    string       `json:"productID"`
	PdpSession   string       `json:"pdpSession"`
	DeviceID     string       `json:"deviceID"`
	UserLocation UserLocation `json:"userLocation"`
	Tokonow      Tokonow      `json:"tokonow"`
}

type PdpGetDataP2Resp struct {
	Data struct {
		PdpGetData PdpGetData `json:"pdpGetData"`
	} `json:"data"`
}
