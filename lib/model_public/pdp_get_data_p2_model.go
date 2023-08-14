package model_public

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

type PDPComponentName string

const (
	MediaComponentName               PDPComponentName = "product_media"
	TickerInfoComponentName          PDPComponentName = "ticker_info"
	VariantOptionsComponentName      PDPComponentName = "variant_options"
	ProductContentComponentName      PDPComponentName = "product_content"
	NewVariantOptionsComponentName   PDPComponentName = "new_variant_options"
	ProductDetailComponentName       PDPComponentName = "product_detail"
	ObatKerasComponentName           PDPComponentName = "obat_keras"
	ShopCredibilityComponentName     PDPComponentName = "shop_credibility"
	ShipmentComponentName            PDPComponentName = "shipment"
	ShippingComponentName            PDPComponentName = "shipping"
	ShopVoucerComponentName          PDPComponentName = "shop_voucher"
	OfferingComponentName            PDPComponentName = "offerings"
	InstallmentPaylaterComponentName PDPComponentName = "installment_paylater"
	QRCodeComponentName              PDPComponentName = "QRCode"
	WholesaleComponentName           PDPComponentName = "wholesale"
	ProtectionComponentName          PDPComponentName = "protection"
	ReportComponentName              PDPComponentName = "report"
	ReviewComponentName              PDPComponentName = "review"
	TdnTopadsComponentName           PDPComponentName = "tdn_topads"
	DiscussionFaqComponentName       PDPComponentName = "discussion_faq"
	Pdp1ComponentName                PDPComponentName = "pdp_1"
	Pdp2ComponentName                PDPComponentName = "pdp_2"
	Pdp3ComponentName                PDPComponentName = "pdp_3"
	Pdp4ComponentName                PDPComponentName = "pdp_4"
)

type Component struct {
	Name     PDPComponentName `json:"name"`
	Type     string           `json:"type"`
	Position string           `json:"position"`
	Typename string           `json:"__typename"`
}
type MediaType string

const (
	MediaVideo MediaType = "video"
	MediaImage MediaType = "image"
)

type Media struct {
	Type            MediaType `json:"type"`
	URLOriginal     string    `json:"urlOriginal"`
	URLThumbnail    string    `json:"urlThumbnail"`
	URLMaxRes       string    `json:"urlMaxRes"`
	VideoURL        string    `json:"videoUrl"`
	Prefix          string    `json:"prefix"`
	Suffix          string    `json:"suffix"`
	Description     string    `json:"description"`
	VariantOptionID string    `json:"variantOptionID"`
	Typename        string    `json:"__typename"`
}

type MediaComponentData struct {
	Media    []Media       `json:"media"`
	Videos   []interface{} `json:"videos"`
	Typename string        `json:"__typename"`
}

type MediaComponent struct {
	Component
	Data     []MediaComponentData `json:"data"`
	Typename string               `json:"__typename"`
}

type InterfaceDataComp struct {
	Component
	Data     []interface{} `json:"data"`
	Typename string        `json:"__typename"`
}

type GlobalDataComponent struct {
	Icon      string        `json:"icon"`
	Title     string        `json:"title"`
	IsApplink bool          `json:"isApplink"`
	Applink   string        `json:"applink"`
	Content   []interface{} `json:"content"`
	Typename  string        `json:"__typename"`
}

type ProdDetailComTitle string

const (
	KondisiTitle      ProdDetailComTitle = "Kondisi"
	BeratSatuanTitle  ProdDetailComTitle = "Berat Satuan"
	MinPemesananTitle ProdDetailComTitle = "Min. Pemesanan"
	CategoryTitle     ProdDetailComTitle = "Kategori"
	EtalaseTitle      ProdDetailComTitle = "Etalase"
	DeskripsiTitle    ProdDetailComTitle = "Deskripsi"
)

type ProductDetailContent struct {
	Title        ProdDetailComTitle `json:"title"`
	Subtitle     string             `json:"subtitle"`
	Applink      string             `json:"applink"`
	ShowAtFront  bool               `json:"showAtFront"`
	IsAnnotation bool               `json:"isAnnotation"`
	Typename     string             `json:"__typename"`
}

type ProductDetailData struct {
	Content  []ProductDetailContent `json:"content"`
	Typename string                 `json:"__typename"`
}

func (detail *ProductDetailData) GetContent(title ProdDetailComTitle) (*ProductDetailContent, error) {
	for _, cont := range detail.Content {
		if cont.Title == title {
			return &cont, nil
		}
	}

	return nil, errors.New(string(title) + " tidak ada")
}

type Preorder struct {
	Duration       int    `json:"duration"`
	TimeUnit       string `json:"timeUnit"`
	IsActive       bool   `json:"isActive"`
	PreorderInDays int    `json:"preorderInDays"`
	Typename       string `json:"__typename"`
}

type ProductCampaign struct {
	CampaignID          string `json:"campaignID"`
	CampaignType        string `json:"campaignType"`
	CampaignTypeName    string `json:"campaignTypeName"`
	CampaignIdentifier  int    `json:"campaignIdentifier"`
	Background          string `json:"background"`
	OriginalPrice       int    `json:"originalPrice"`
	Stock               int    `json:"stock"`
	StockSoldPercentage int    `json:"stockSoldPercentage"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	EndDateUnix         string `json:"endDateUnix"`
	AppLinks            string `json:"appLinks"`
	IsAppsOnly          bool   `json:"isAppsOnly"`
	IsActive            bool   `json:"isActive"`
	HideGimmick         bool   `json:"hideGimmick"`
	PercentageAmount    int    `json:"percentageAmount"`
	DiscountedPrice     int    `json:"discountedPrice"`
	OriginalStock       int    `json:"originalStock"`
	Threshold           int    `json:"threshold"`
	IsCheckImei         bool   `json:"isCheckImei,omitempty"`
	DiscountPercentage  int    `json:"discountPercentage,omitempty"`
	DiscountPrice       int    `json:"discountPrice,omitempty"`
	MinOrder            int    `json:"minOrder,omitempty"`
	Typename            string `json:"__typename"`
}

type ThematicCampaign struct {
	AdditionalInfo string `json:"additionalInfo"`
	Background     string `json:"background"`
	CampaignName   string `json:"campaignName"`
	Icon           string `json:"icon"`
	Typename       string `json:"__typename"`
}

type ProductStock struct {
	UseStock     bool   `json:"useStock"`
	Value        string `json:"value"`
	StockWording string `json:"stockWording"`
	Typename     string `json:"__typename"`
}
type IsCashback struct {
	Percentage int    `json:"percentage"`
	Typename   string `json:"__typename"`
}

type ProductPrice struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
	Typename string `json:"__typename"`
}

type ProductVariant struct {
	IsVariant bool   `json:"isVariant"`
	ParentID  string `json:"parentID"`
	Typename  string `json:"__typename"`
}

type ProductContentData struct {
	Name             string           `json:"name"`
	Price            ProductPrice     `json:"price"`
	Campaign         ProductCampaign  `json:"campaign"`
	ThematicCampaign ThematicCampaign `json:"thematicCampaign"`
	Stock            ProductStock     `json:"stock"`
	Variant          ProductVariant   `json:"variant"`
	Wholesale        []interface{}    `json:"wholesale"`
	IsCashback       IsCashback       `json:"isCashback"`
	IsTradeIn        bool             `json:"isTradeIn"`
	IsOS             bool             `json:"isOS"`
	IsPowerMerchant  bool             `json:"isPowerMerchant"`
	IsWishlist       bool             `json:"isWishlist"`
	IsCOD            bool             `json:"isCOD"`
	Preorder         Preorder         `json:"preorder"`
	Typename         string           `json:"__typename"`
}

type VariantStock struct {
	Stock            string `json:"stock"`
	IsBuyable        bool   `json:"isBuyable"`
	StockWordingHTML string `json:"stockWordingHTML"`
	MinimumOrder     string `json:"minimumOrder"`
	MaximumOrder     string `json:"maximumOrder"`
	Typename         string `json:"__typename"`
}

type Picture struct {
	URLOriginal  string `json:"urlOriginal"`
	URLThumbnail string `json:"urlThumbnail"`
	Typename     string `json:"__typename"`
}

type VariantOption struct {
	Picture                Picture `json:"picture"`
	ProductVariantOptionID string  `json:"productVariantOptionID"`
	VariantUnitValueID     string  `json:"variantUnitValueID"`
	Value                  string  `json:"value"`
	Hex                    string  `json:"hex"`
	Stock                  string  `json:"stock"`
	Typename               string  `json:"__typename"`
}

type Variant struct {
	ProductVariantID string          `json:"productVariantID"`
	VariantID        string          `json:"variantID"`
	Name             string          `json:"name"`
	Identifier       string          `json:"identifier"`
	Option           []VariantOption `json:"option"`
	Typename         string          `json:"__typename"`
}

type NewVariantChild struct {
	ProductID        string           `json:"productID"`
	Price            int              `json:"price"`
	PriceFmt         string           `json:"priceFmt"`
	OptionID         []int            `json:"optionID"`
	OptionName       []string         `json:"optionName"`
	ProductName      string           `json:"productName"`
	ProductURL       string           `json:"productURL"`
	Picture          Picture          `json:"picture"`
	Stock            VariantStock     `json:"stock"`
	IsCOD            bool             `json:"isCOD"`
	IsWishlist       bool             `json:"isWishlist"`
	CampaignInfo     ProductCampaign  `json:"campaignInfo"`
	ThematicCampaign ThematicCampaign `json:"thematicCampaign"`
	Typename         string           `json:"__typename"`
}
type NewVariantOptionData struct {
	ErrorCode     int               `json:"errorCode"`
	ParentID      string            `json:"parentID"`
	DefaultChild  string            `json:"defaultChild"`
	SizeChart     string            `json:"sizeChart"`
	TotalStockFmt string            `json:"totalStockFmt"`
	Variants      []Variant         `json:"variants"`
	Children      []NewVariantChild `json:"children"`
	Typename      string            `json:"__typename"`
}

type NewVariantOptionsComponent struct {
	Component
	Data     []NewVariantOptionData `json:"data"`
	Typename string                 `json:"__typename"`
}

type ProductContentComponent struct {
	Component
	Data     []ProductContentData `json:"data"`
	Typename string               `json:"__typename"`
}

type GlobalComponent struct {
	Component
	Data     []GlobalDataComponent `json:"data"`
	Typename string                `json:"__typename"`
}

type ProductDetailComponent struct {
	Component
	Data     []ProductDetailData `json:"data"`
	Typename string              `json:"__typename"`
}

type ComponentParser struct {
	Name      PDPComponentName `json:"name"`
	Type      string           `json:"type"`
	Component interface{}
}

func (p *ComponentParser) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Name PDPComponentName `json:"name"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch aux.Name {
	case MediaComponentName:
		component := MediaComponent{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	case ProductContentComponentName:
		component := ProductContentComponent{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	case NewVariantOptionsComponentName:
		component := NewVariantOptionsComponent{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	case ProductDetailComponentName:
		component := ProductDetailComponent{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	case ShippingComponentName, InstallmentPaylaterComponentName, WholesaleComponentName, ProtectionComponentName:
		component := GlobalComponent{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	default:
		component := InterfaceDataComp{}
		if err := json.Unmarshal(data, &component); err != nil {
			log.Println("struct", aux.Name, err.Error(), string(data))
			return err
		}
		p.Component = &component
	}

	return nil
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

type CategoryDetail []struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	BreadcrumbURL string `json:"breadcrumbURL"`
	IsAdult       bool   `json:"isAdult"`
	Typename      string `json:"__typename"`
}

type Category struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Title         string          `json:"title"`
	Detail        *CategoryDetail `json:"detail"`
	BreadcrumbURL string          `json:"breadcrumbURL"`
	IsAdult       bool            `json:"isAdult"`
	IsKyc         bool            `json:"isKyc"`
	MinAge        int             `json:"minAge"`
	Typename      string          `json:"__typename"`
}

type Menu struct {
	ID       string `json:"id,omitempty"`
	MenuID   string `json:"menuID,omitempty"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
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

type PDPListComponents []interface{}

func (p *PDPListComponents) UnmarshalJSON(data []byte) error {
	aux := []*ComponentParser{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	fixcomponents := make([]interface{}, len(aux))
	for ind, component := range aux {

		fixcomponents[ind] = component.Component

	}

	*p = fixcomponents

	return nil
}

type PdpGetLayout struct {
	RequestID  string            `json:"requestID"`
	Name       string            `json:"name"`
	PdpSession string            `json:"pdpSession"`
	BasicInfo  BasicInfo         `json:"basicInfo"`
	Components PDPListComponents `json:"components"`
	Typename   string            `json:"__typename"`
}

func GetComponent[V any](layout *PdpGetLayout) (*V, error) {
	for _, com := range layout.Components {
		switch fcom := com.(type) {
		case *V:
			return fcom, nil
		}
	}

	return nil, errors.New("component not found")
}

func (layout *PdpGetLayout) GetProductName() (string, error) {

	productContent, err := GetComponent[ProductContentComponent](layout)
	if err != nil {
		return "", err
	}

	return productContent.Data[0].Name, nil
}

func (layout *PdpGetLayout) GetPercentageAmount() (int, error) {

	productContent, err := GetComponent[ProductContentComponent](layout)
	if err != nil {
		return 0, err
	}

	return productContent.Data[0].Campaign.PercentageAmount, nil
}

func (layout *PdpGetLayout) GetStock() (int, error) {

	productContent, err := GetComponent[ProductContentComponent](layout)
	if err != nil {
		return 0, err
	}

	stockVal := productContent.Data[0].Stock
	stock, err := strconv.Atoi(stockVal.Value)
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (layout *PdpGetLayout) GetPrice() (int, error) {

	productContent, err := GetComponent[ProductContentComponent](layout)
	if err != nil {
		return 0, err
	}

	return productContent.Data[0].Price.Value, nil
}

func (layout *PdpGetLayout) GetPriceBeforeDiscount() (int, error) {
	productContent, err := GetComponent[ProductContentComponent](layout)
	if err != nil {
		return 0, err
	}

	oriPrice := productContent.Data[0].Campaign.OriginalPrice
	if oriPrice > 0 {
		return oriPrice, nil
	}

	return layout.GetPrice()
}

func (layout *PdpGetLayout) GetImages() ([]string, error) {
	mediaComponent, err := GetComponent[MediaComponent](layout)
	if err != nil {
		return []string{}, err
	}

	images := []string{}
	for _, media := range mediaComponent.Data[0].Media {
		if media.Type == MediaImage {
			images = append(images, media.URLOriginal)
		}
	}

	return images, nil
}

func (layout *PdpGetLayout) GetDescription() (string, error) {
	productDetail, err := GetComponent[ProductDetailComponent](layout)
	if err != nil {
		return "", err
	}

	desc, err := productDetail.Data[0].GetContent(DeskripsiTitle)
	if err != nil {
		return "", err
	}

	return desc.Subtitle, nil
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
	Affiliate    interface{}   `json:"affiliate"`
	ProductID    string        `json:"productID"`
	PdpSession   string        `json:"pdpSession"`
	DeviceID     string        `json:"deviceID"`
	UserLocation *UserLocation `json:"userLocation,omitempty"`
	Tokonow      *Tokonow      `json:"tokonow,omitempty"`
}

func NewPdpGetDataP2Var(layout PdpGetLayout) *PdpGetDataP2Var {
	return &PdpGetDataP2Var{
		PdpSession: layout.PdpSession,
		ProductID:  layout.BasicInfo.ID,
	}
}

type PdpGetDataP2Data struct {
	PdpGetData PdpGetData `json:"pdpGetData"`
}

type PdpGetDataP2Resp struct {
	Data PdpGetDataP2Data `json:"data"`
}
