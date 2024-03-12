package model

import (
	"fmt"
	"strconv"
	"time"
)

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
	IsAdult       bool   `json:"isAdult,omitempty"`
	Typename      string `json:"__typename"`
}

type Category struct {
	ID            int             `json:"id,string"`
	Name          string          `json:"name,omitempty"`
	Title         string          `json:"title,omitempty"`
	Detail        *CategoryDetail `json:"detail,omitempty"`
	BreadcrumbURL string          `json:"breadcrumbURL,omitempty"`
	IsAdult       bool            `json:"isAdult,omitempty"`
	IsKyc         bool            `json:"isKyc,omitempty"`
	MinAge        int             `json:"minAge,omitempty"`
	Typename      string          `json:"__typename,omitempty"`
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

type InputPicture struct {
	Data []Pictures `json:"data"`
}

type ConditionStatus string

const (
	NewCondition ConditionStatus = "NEW"
)

type MenuInput struct {
	MenuID string `json:"menuID"`
}

type InputVariable struct {
	Pictures      InputPicture    `json:"pictures"`
	ProductName   string          `json:"productName"`
	Category      Category        `json:"category"`
	Condition     ConditionStatus `json:"condition"`
	MinOrder      int64           `json:"minOrder"`
	PriceCurrency string          `json:"minCurrency"`
	Weight        int64           `json:"weight"`
	WeightUnit    WeightUnit      `json:"weightUnit"`
	MustInsurance bool            `json:"mustInsurance"`
	Menus         []*MenuInput    `json:"menus,omitempty"`
	Annotations   []string        `json:"annotations,omitempty"`
	Description   string          `json:"description"`
	Dimention     *Dimension      `json:"dimension,omitempty"`
	Catalog       *Catalog        `json:"catalog,omitempty"`
	PreOrder      *PreOrder       `json:"preorder,omitempty"`
}

type ProductStatus string

const (
	LimitedStatus   ProductStatus = "LIMITED"
	DeletedStatus   ProductStatus = "DELETED"
	ViolationStatus ProductStatus = "VIOLATION"
	ActiveStatus    ProductStatus = "ACTIVE"
	InActiveStatus  ProductStatus = "INACTIVE"
)

type NoVariantStockPrice struct {
	Sku    string        `json:"sku"`
	Stock  int64         `json:"stock"`
	Price  int64         `json:"price"`
	Status ProductStatus `json:"status"`
}

type InputNoVariant struct {
	*InputVariable
	*NoVariantStockPrice
}

type WeightUnit string

const (
	GramUnit WeightUnit = "GR"
)

type ProductVariant struct {
	Combination []int         `json:"combination"`
	IsPrimary   bool          `json:"isPrimary"`
	Price       int           `json:"price"`
	Sku         string        `json:"sku"`
	Status      ProductStatus `json:"status"`
	Stock       int           `json:"stock"`
	Pictures    []Pictures    `json:"pictures"`
	Weight      int           `json:"weight"`
	WeightUnit  WeightUnit    `json:"weightUnit"`
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

type ProductVariants []*ProductVariant

func (prods ProductVariants) MinPrice() int {
	var min int
	for ind, prod := range prods {
		if ind == 0 || prod.Price < min {
			min = prod.Price
		}
	}

	return min
}

func (prods ProductVariants) MaxPrice() int {
	var max int
	for ind, prod := range prods {
		if ind == 0 || prod.Price > max {
			max = prod.Price
		}
	}

	return max
}

func (prods ProductVariants) MinWeight() int {
	var min int
	for ind, prod := range prods {
		if ind == 0 || prod.Weight < min {
			min = prod.Weight
		}
	}

	return min
}

type Variant struct {
	Products   ProductVariants    `json:"products"`
	Selections []ProductSelection `json:"selections"`
	Sizecharts []interface{}      `json:"sizeChart"`
	Typename   string             `json:"__typename,omitempty"`
}

type InputVariant struct {
	*InputVariable
	Variant *Variant `json:"variant"`
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
	IsStockGuaranteed bool    `json:"isStockGuaranteed"`
	ScoreV3           float32 `json:"scoreV3"`
	Typename          string  `json:"__typename"`
}

type SellerProductItemShop struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type SellerProductItem struct {
	ID                 int                   `json:"id,string"`
	Name               string                `json:"name"`
	Price              Price                 `json:"price"`
	Stock              int                   `json:"stock"`
	Status             ProductStatus         `json:"status"`
	MinOrder           int                   `json:"minOrder"`
	MaxOrder           int                   `json:"maxOrder"`
	Weight             int                   `json:"weight"`
	WeightUnit         string                `json:"weightUnit"`
	Condition          string                `json:"condition"`
	IsMustInsurance    bool                  `json:"isMustInsurance"`
	IsKreasiLokal      bool                  `json:"isKreasiLokal"`
	IsCOD              bool                  `json:"isCOD"`
	IsCampaign         bool                  `json:"isCampaign"`
	IsVariant          bool                  `json:"isVariant"`
	URL                string                `json:"url"`
	Sku                string                `json:"sku"`
	Cashback           int                   `json:"cashback"`
	Featured           int                   `json:"featured"`
	HasStockReserved   bool                  `json:"hasStockReserved"`
	HasInbound         bool                  `json:"hasInbound"`
	WarehouseCount     int                   `json:"warehouseCount"`
	IsEmptyStock       bool                  `json:"isEmptyStock"`
	Score              Score                 `json:"score"`
	Pictures           []Pictures            `json:"pictures"`
	Shop               SellerProductItemShop `json:"shop"`
	Wholesale          []interface{}         `json:"wholesale"`
	Stats              Stats                 `json:"stats"`
	TxStats            TxStats               `json:"txStats"`
	Topads             interface{}           `json:"topads"`
	PriceSuggestion    interface{}           `json:"priceSuggestion"`
	CampaignType       []interface{}         `json:"campaignType"`
	SuspendLevel       int                   `json:"suspendLevel"`
	HasStockAlert      bool                  `json:"hasStockAlert"`
	StockAlertCount    int                   `json:"stockAlertCount"`
	StockAlertActive   bool                  `json:"stockAlertActive"`
	HaveNotifyMeOOS    bool                  `json:"haveNotifyMeOOS"`
	NotifyMeOOSCount   int                   `json:"notifyMeOOSCount"`
	NotifyMeOOSWording string                `json:"notifyMeOOSWording"`
	ManageProductData  ManageProductData     `json:"manageProductData"`
	CreateTime         time.Time             `json:"createTime"`
	Typename           string                `json:"__typename"`
}

type ProductList struct {
	Header   *Header              `json:"header"`
	Data     []*SellerProductItem `json:"data"`
	Typename string               `json:"__typename"`
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
	ID       string `json:"id,omitempty"`
	MenuID   string `json:"menuID,omitempty"`
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
	ProductID     int64         `json:"productID,string"`
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

type BulkProductEditShop struct {
	ID string `json:"id"`
}

type BulkProductEditV3Input struct {
	ProductID int                 `json:"productID,string"`
	Shop      BulkProductEditShop `json:"shop"`
	Status    ProductStatus       `json:"status"`
}

type ProductListVar struct {
	ShopID      string   `json:"shopID"`
	Filter      []Filter `json:"filter"`
	Sort        Sort     `json:"sort"`
	ExtraInfo   []string `json:"extraInfo"`
	WarehouseID string   `json:"warehouseID"`
}

func NewProductListVar(shopid int64, filter []Filter) *ProductListVar {
	return &ProductListVar{
		ShopID: strconv.Itoa(int(shopid)),
		Filter: filter,
		Sort: Sort{
			ID:    "DEFAULT",
			Value: "DESC",
		},
		ExtraInfo:   []string{"view", "topads", "rbac", "price-suggestion"},
		WarehouseID: "",
	}
}

type ProductListResp struct {
	Data struct {
		ProductList ProductList `json:"ProductList"`
	} `json:"data"`
}

type ProductAddVar struct {
	Input interface{} `json:"input"`
}

type ProductAddV3 struct {
	Header    *Header `json:"header"`
	IsSuccess bool    `json:"isSuccess"`
	ProductId int     `json:"productID,string"`
	TypeName  string  `json:"__typename"`
}

type ProductAddData struct {
	ProductAddV3 *ProductAddV3 `json:"ProductAddV3"`
}

type ProductAddResp struct {
	Data *ProductAddData `json:"data"`
}

type ExtraInfo struct {
	Aggregate bool `json:"aggregate"`
	Event     bool `json:"event"`
}

type GetProductV3Var struct {
	ProductID string    `json:"productID"`
	Options   Options   `json:"options"`
	ExtraInfo ExtraInfo `json:"extraInfo"`
}

func NewProductV3Var(productId int64) *GetProductV3Var {

	productIdStr := strconv.FormatInt(productId, 10)

	return &GetProductV3Var{
		ProductID: productIdStr,
		Options: Options{
			Basic:       true,
			Menu:        true,
			Shop:        true,
			Category:    true,
			Wholesale:   true,
			Preorder:    true,
			Picture:     true,
			Sku:         true,
			Lock:        true,
			Variant:     true,
			Video:       true,
			Edit:        true,
			TxStats:     true,
			Dimension:   true,
			CustomVideo: true,
		},
		ExtraInfo: ExtraInfo{
			Aggregate: true,
		},
	}
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
	Input []*BulkProductEditV3Input `json:"input"`
}

type BulkProductEditV3Item struct {
	ProductID string `json:"productID"`
	Result    struct {
		Header    Header `json:"header"`
		IsSuccess bool   `json:"isSuccess"`
		Typename  string `json:"__typename"`
	} `json:"result"`
	Typename string `json:"__typename"`
}

func (item *BulkProductEditV3Item) Error() string {
	return fmt.Sprintf("gagal delete product %s", item.ProductID)
}

type BulkProductEditV3Resp struct {
	Data struct {
		BulkProductEditV3 []*BulkProductEditV3Item `json:"BulkProductEditV3"`
	} `json:"data"`
}
