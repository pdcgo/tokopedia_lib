package model

type ShopCore struct {
	Description string `json:"description"`
	Domain      string `json:"domain"`
	ShopID      string `json:"shopID"`
	Name        string `json:"name"`
	TagLine     string `json:"tagLine"`
	DefaultSort int    `json:"defaultSort"`
	Typename    string `json:"__typename"`
}

type CreateInfo struct {
	OpenSince        string `json:"openSince,omitempty"`
	EpochShopCreated string `json:"epochShopCreated,omitempty"`
	Typename         string `json:"__typename"`
}

type FavoriteData struct {
	TotalFavorite    int    `json:"totalFavorite"`
	AlreadyFavorited int    `json:"alreadyFavorited"`
	Typename         string `json:"__typename"`
}

type ShopAssets struct {
	Avatar   string `json:"avatar"`
	Cover    string `json:"cover,omitempty"`
	Typename string `json:"__typename"`
}

type ShipmentInfo struct {
	IsAvailable int                   `json:"isAvailable"`
	Image       string                `json:"image"`
	Name        string                `json:"name"`
	Product     []ShipmentInfoProduct `json:"product"`
	Typename    string                `json:"__typename"`
}

type ShippingLoc struct {
	DistrictName string `json:"districtName"`
	CityName     string `json:"cityName"`
	Typename     string `json:"__typename"`
}

type ShopStats struct {
	ProductSold    string `json:"productSold"`
	TotalTxSuccess string `json:"totalTxSuccess"`
	TotalShowcase  string `json:"totalShowcase"`
	Typename       string `json:"__typename"`
}

type StatusInfo struct {
	ShopStatus    int    `json:"shopStatus"`
	StatusMessage string `json:"statusMessage"`
	StatusTitle   string `json:"statusTitle,omitempty"`
	TickerType    string `json:"tickerType,omitempty"`
	IsIdle        bool   `json:"isIdle,omitempty"`
	Typename      string `json:"__typename"`
}

type ClosedInfoDetail struct {
	OpenDate string `json:"openDate,omitempty"`
	Status   int    `json:"status,omitempty"`
	Typename string `json:"__typename"`
}
type ClosedInfo struct {
	ClosedNote string           `json:"closedNote"`
	Until      string           `json:"until,omitempty"`
	Reason     string           `json:"reason"`
	Detail     ClosedInfoDetail `json:"detail"`
	Typename   string           `json:"__typename"`
}

type GoldOS struct {
	IsGold      int    `json:"isGold"`
	IsGoldBadge int    `json:"isGoldBadge"`
	IsOfficial  int    `json:"isOfficial"`
	Badge       string `json:"badge"`
	ShopTier    int    `json:"shopTier"`
	Typename    string `json:"__typename"`
}

type CustomSEO struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	BottomContent string `json:"bottomContent"`
	Typename      string `json:"__typename"`
}

type PartnerInfo struct {
	FsType   int    `json:"fsType"`
	Typename string `json:"__typename"`
}

type EpharmacyInfo struct {
	SiaNumber  string `json:"siaNumber"`
	SipaNumber string `json:"sipaNumber"`
	Apj        string `json:"apj"`
	Typename   string `json:"__typename"`
}

type ShopInfoByIDResult struct {
	ShopCore         ShopCore       `json:"shopCore"`
	CreateInfo       CreateInfo     `json:"createInfo"`
	FavoriteData     FavoriteData   `json:"favoriteData"`
	ActiveProduct    int            `json:"activeProduct"`
	ShopAssets       ShopAssets     `json:"shopAssets"`
	Location         string         `json:"location"`
	IsAllowManage    int            `json:"isAllowManage"`
	BranchLinkDomain string         `json:"branchLinkDomain"`
	IsOpen           int            `json:"isOpen"`
	ShipmentInfo     []ShipmentInfo `json:"shipmentInfo"`
	ShippingLoc      ShippingLoc    `json:"shippingLoc"`
	ShopStats        ShopStats      `json:"shopStats"`
	StatusInfo       StatusInfo     `json:"statusInfo"`
	ClosedInfo       ClosedInfo     `json:"closedInfo"`
	BbInfo           []interface{}  `json:"bbInfo"`
	GoldOS           GoldOS         `json:"goldOS"`
	ShopSnippetURL   string         `json:"shopSnippetURL"`
	CustomSEO        CustomSEO      `json:"customSEO"`
	IsQA             bool           `json:"isQA"`
	IsGoApotik       bool           `json:"isGoApotik"`
	PartnerInfo      []PartnerInfo  `json:"partnerInfo"`
	EpharmacyInfo    EpharmacyInfo  `json:"epharmacyInfo"`
	Typename         string         `json:"__typename"`
}

type ShopCoreInfoVar struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
}

type ShopInfoByID struct {
	Result   []ShopInfoByIDResult `json:"result"`
	Error    Error                `json:"error"`
	Typename string               `json:"__typename"`
}

type ShopCoreInfoResp struct {
	Data struct {
		ShopInfoByID `json:"shopInfoByID"`
	} `json:"data"`
}
