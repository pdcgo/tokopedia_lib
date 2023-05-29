package model

import "time"

type Options struct {
	Basic       bool `json:"basic"`
	Menu        bool `json:"menu"`
	Shop        bool `json:"shop"`
	Category    bool `json:"category"`
	Wholesale   bool `json:"wholesale"`
	Preorder    bool `json:"preorder"`
	Picture     bool `json:"picture"`
	Sku         bool `json:"sku"`
	Lock        bool `json:"lock"`
	Variant     bool `json:"variant"`
	Video       bool `json:"video"`
	Edit        bool `json:"edit"`
	TxStats     bool `json:"txStats"`
	Dimension   bool `json:"dimension"`
	CustomVideo bool `json:"custom_video"`
}

type Pictures struct {
	PicID        string `json:"picID,omitempty"`
	FilePath     string `json:"filePath,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	URLOriginal  string `json:"urlOriginal,omitempty"`
	URLThumbnail string `json:"urlThumbnail,omitempty"` // untuk productList
	UploadIds    string `json:"uploadIds,omitempty"`    // untuk create or update
	Typename     string `json:"__typename,omitempty"`
}

type Filter struct {
	ID    string   `json:"id"`
	Value []string `json:"value"`
}

type Sort struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type CategoryDetail []struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	BreadcrumbURL string `json:"breadcrumbURL"`
	Typename      string `json:"__typename"`
}

type Category struct {
	ID       string          `json:"id"`
	Name     string          `json:"name,omitempty"`
	Title    string          `json:"title,omitempty"`
	Detail   *CategoryDetail `json:"detail,omitempty"`
	Typename string          `json:"__typename"`
}

type Dimension struct {
	Length   int    `json:"length"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Typename string `json:"__typename"`
}

type Catalog struct {
	CatalogID string `json:"catalogID"`
	IsActive  bool   `json:"isActive"`
}

type PreOrder struct {
	Duration int    `json:"duration,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
	IsActive bool   `json:"isActive"`
	Typename string `json:"__typename,omitempty"`
}

type InputVariable struct {
	Pictures struct {
		Data []Pictures `json:"data"`
	} `json:"pictures"`
	ProductName   string   `json:"productName"`
	Category      Category `json:"category"`
	Condition     string   `json:"condition"`
	MinOrder      int64    `json:"minOrder"`
	PriceCurrency string   `json:"minCurrency"`
	Weight        int64    `json:"weight"`
	WeightUnit    string   `json:"weightUnit"`
	MustInsurance bool     `json:"mustInsurance"`
	Menus         []struct {
		MenuID string `json:"menuID"`
	} `json:"menus"`
	Annotations []string  `json:"annotations"`
	Description string    `json:"description"`
	Dimention   Dimension `json:"dimension"`
	Catalog     *Catalog  `json:"catalog,omitempty"`
	PreOrder    PreOrder  `json:"preorder,omitempty"`
}

type InputNoVariant struct {
	InputVariable
	Sku    string `json:"sku"`
	Stock  int64  `json:"stock"`
	Price  int64  `json:"price"`
	Status string `json:"status"`
}

type ProdctVariant struct {
	Combination []int      `json:"combination"`
	IsPrimary   bool       `json:"isPrimary"`
	Price       int        `json:"price"`
	Sku         string     `json:"sku"`
	Status      string     `json:"status"`
	Stock       int        `json:"stock"`
	Pictures    []Pictures `json:"pictures"`
	Weight      int        `json:"weight"`
	WeightUnit  string     `json:"weightUnit"`
}

type SelectionsOptions struct {
	UnitValueID string `json:"unitValueID"`
	Value       string `json:"value"`
	HexCode     string `json:"hexCode"`
}

type ProductSelection struct {
	UnitID    string              `json:"unitID"`
	VariantID string              `json:"variantID"`
	Name      string              `json:"name"`
	Options   []SelectionsOptions `json:"options"`
}

type Variant struct {
	Products   []ProdctVariant    `json:"products"`
	Selections []ProductSelection `json:"selections"`
	Sizecharts []interface{}      `json:"sizecharts"`
	Typename   string             `json:"__typename,omitempty"`
}

type InputVariant struct {
	InputVariable
	Variant Variant `json:"variant"`
}

type Header struct {
	ProcessTime float64       `json:"processTime,omitempty"`
	Messages    []interface{} `json:"messages"`
	Reason      string        `json:"reason"`
	ErrorCode   string        `json:"errorCode"`
	Typename    string        `json:"__typename"`
}

type Price struct {
	Min      int    `json:"min,omitempty"`
	Max      int    `json:"max,omitempty"`
	TextIdr  string `json:"text_idr,omitempty"`
	Typename string `json:"__typename"`
}

type Score struct {
	Total    int    `json:"total"`
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

type TxStats struct {
	ItemSold int    `json:"itemSold,omitempty"`
	Sold     int    `json:"sold,omitempty"`
	Typename string `json:"__typename"`
}

type ManageProductData struct {
	IsStockGuaranteed bool   `json:"isStockGuaranteed"`
	ScoreV3           int    `json:"scoreV3"`
	Typename          string `json:"__typename"`
}

type ProductList struct {
	Header Header `json:"header"`
	Data   []struct {
		ID               string     `json:"id"`
		Name             string     `json:"name"`
		Price            Price      `json:"price"`
		Stock            int        `json:"stock"`
		Status           string     `json:"status"`
		MinOrder         int        `json:"minOrder"`
		MaxOrder         int        `json:"maxOrder"`
		Weight           int        `json:"weight"`
		WeightUnit       string     `json:"weightUnit"`
		Condition        string     `json:"condition"`
		IsMustInsurance  bool       `json:"isMustInsurance"`
		IsKreasiLokal    bool       `json:"isKreasiLokal"`
		IsCOD            bool       `json:"isCOD"`
		IsCampaign       bool       `json:"isCampaign"`
		IsVariant        bool       `json:"isVariant"`
		URL              string     `json:"url"`
		Sku              string     `json:"sku"`
		Cashback         int        `json:"cashback"`
		Featured         int        `json:"featured"`
		HasStockReserved bool       `json:"hasStockReserved"`
		HasInbound       bool       `json:"hasInbound"`
		WarehouseCount   int        `json:"warehouseCount"`
		IsEmptyStock     bool       `json:"isEmptyStock"`
		Score            Score      `json:"score"`
		Pictures         []Pictures `json:"pictures"`
		Shop             struct {
			ID       string `json:"id"`
			Typename string `json:"__typename"`
		} `json:"shop"`
		Wholesale          []interface{}     `json:"wholesale"`
		Stats              Stats             `json:"stats"`
		TxStats            TxStats           `json:"txStats"`
		Topads             interface{}       `json:"topads"`
		PriceSuggestion    interface{}       `json:"priceSuggestion"`
		CampaignType       []interface{}     `json:"campaignType"`
		SuspendLevel       int               `json:"suspendLevel"`
		HasStockAlert      bool              `json:"hasStockAlert"`
		StockAlertCount    int               `json:"stockAlertCount"`
		StockAlertActive   bool              `json:"stockAlertActive"`
		HaveNotifyMeOOS    bool              `json:"haveNotifyMeOOS"`
		NotifyMeOOSCount   int               `json:"notifyMeOOSCount"`
		NotifyMeOOSWording string            `json:"notifyMeOOSWording"`
		ManageProductData  ManageProductData `json:"manageProductData"`
		CreateTime         time.Time         `json:"createTime"`
		Typename           string            `json:"__typename"`
	} `json:"data"`
	Typename string `json:"__typename"`
}

type Partial struct {
	Price     bool   `json:"price"`
	Status    bool   `json:"status"`
	Stock     bool   `json:"stock"`
	Wholesale bool   `json:"wholesale"`
	Name      bool   `json:"name"`
	Typename  string `json:"__typename"`
}

type Lock struct {
	Full     bool    `json:"full"`
	Partial  Partial `json:"partial"`
	Typename string  `json:"__typename"`
}

type Cpl struct {
	ShipperServices []interface{} `json:"shipperServices"`
	Typename        string        `json:"__typename"`
}

type Menu struct {
	MenuID   string `json:"menuID"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
}

type GetProductV3 struct {
	Lock    Lock    `json:"lock"`
	TxStats TxStats `json:"txStats"`
	Shop    struct {
		ID       string `json:"id"`
		Typename string `json:"__typename"`
	} `json:"shop"`
	ProductID     string        `json:"productID"`
	ProductName   string        `json:"productName"`
	Status        string        `json:"status"`
	Stock         int           `json:"stock"`
	Price         int           `json:"price"`
	MinOrder      int           `json:"minOrder"`
	Description   string        `json:"description"`
	WeightUnit    string        `json:"weightUnit"`
	Weight        int           `json:"weight"`
	Condition     string        `json:"condition"`
	MustInsurance bool          `json:"mustInsurance"`
	Sku           string        `json:"sku"`
	Category      Category      `json:"category"`
	Menu          Menu          `json:"menu"`
	Menus         []interface{} `json:"menus"`
	Video         []interface{} `json:"video"`
	CustomVideo   []interface{} `json:"customVideo"`
	Pictures      []Pictures    `json:"pictures"`
	Wholesale     []interface{} `json:"wholesale"`
	Dimension     Dimension     `json:"dimension"`
	Preorder      PreOrder      `json:"preorder"`
	Variant       Variant       `json:"variant"`
	Cpl           Cpl           `json:"cpl"`
	Typename      string        `json:"__typename"`
}

type BulkProductEditV3Input struct {
	ProductID string `json:"productID"`
	Shop      struct {
		ID string `json:"id"`
	} `json:"shop"`
	Status string `json:"status"`
}

///////////////////////////////////////////////////////

type ProductListVar struct {
	ShopID      string   `json:"shopID"`
	Filter      []Filter `json:"filter"`
	Sort        Sort     `json:"sort"`
	ExtraInfo   []string `json:"extraInfo"`
	WarehouseID string   `json:"warehouseID"`
}

type ProductListResp struct {
	Data struct {
		ProductList ProductList `json:"ProductList"`
	} `json:"data"`
}

type ProductAddVar struct {
	Input interface{} `json:"input"`
}

type ProductAddResp struct {
	Data struct {
		ProductAddV3 struct {
			Header    Header `json:"header"`
			IsSuccess bool   `json:"isSuccess"`
			ProductId string `json:"productID"`
			TypeName  string `json:"__typename"`
		} `json:"ProductAddV3"`
	} `json:"data"`
}

type GetProductV3Var struct {
	ProductID string  `json:"productID"`
	Options   Options `json:"options"`
	ExtraInfo struct {
		Event bool `json:"event"`
	} `json:"extraInfo"`
}

type GetProductV3Resp struct {
	Data struct {
		GetProductV3 GetProductV3 `json:"getProductV3"`
	} `json:"data"`
}

type ProductUpdateVar struct {
	Input struct {
		Pictures struct {
			Data []Pictures `json:"data"`
		} `json:"pictures"`
		Status       string      `json:"status"`
		Catalog      interface{} `json:"catalog"`
		CustomVideos struct {
			Data []interface{} `json:"data"`
		} `json:"customVideos"`
		Wholesale struct {
			Data []interface{} `json:"data"`
		} `json:"wholesale"`
		Dimension Dimension `json:"dimension"`
		ProductID string    `json:"productID"`
		Shop      struct {
			ID string `json:"id"`
		} `json:"shop"`
		Stock    int      `json:"stock"`
		Preorder PreOrder `json:"preorder"`
	} `json:"input"`
}

type ProductUpdateResp struct {
	Data struct {
		ProductUpdateV3 struct {
			Header    Header `json:"header"`
			IsSuccess bool   `json:"isSuccess"`
			ProductId string `json:"productID"`
			TypeName  string `json:"__typename"`
		} `json:"ProductUpdateV3"`
	} `json:"data"`
}

type BulkProductEditV3Var struct {
	Input []BulkProductEditV3Input `json:"input"`
}

type BulkProductEditV3Resp struct {
	Data struct {
		BulkProductEditV3 []struct {
			ProductID string `json:"productID"`
			Result    struct {
				Header    Header `json:"header"`
				IsSuccess bool   `json:"isSuccess"`
				Typename  string `json:"__typename"`
			} `json:"result"`
			Typename string `json:"__typename"`
		} `json:"BulkProductEditV3"`
	} `json:"data"`
}
