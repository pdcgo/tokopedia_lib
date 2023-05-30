package model_public

type RecomWidgetVar struct {
	UserID         int           `json:"userID"`
	XDevice        string        `json:"xDevice"`
	PageName       string        `json:"pageName"`
	Ref            string        `json:"ref"`
	ProductIDs     string        `json:"productIDs"`
	TokoNow        bool          `json:"tokoNow"`
	CategoryIDs    string        `json:"categoryIDs"`
	Keyword        []interface{} `json:"keyword"`
	LayoutPageType string        `json:"LayoutPageType"`
	QueryParam     string        `json:"queryParam"`
}

type Pagination struct {
	HasNext  bool   `json:"hasNext"`
	Typename string `json:"__typename"`
}

type RecomShop struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	City     string `json:"city"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
}

type ProductLabel struct {
	Title    string `json:"title"`
	Color    string `json:"color"`
	Typename string `json:"__typename"`
}

type Recomendation struct {
	ProductSlashedPrice       string         `json:"productSlashedPrice"`
	SlashedPriceInt           int            `json:"slashedPriceInt"`
	ProductDiscountPercentage int            `json:"productDiscountPercentage"`
	ProductReviewCount        int            `json:"productReviewCount"`
	IsWishlist                bool           `json:"isWishlist"`
	ProductImageURL           string         `json:"productImageUrl"`
	IsTopads                  bool           `json:"isTopads"`
	ClickURL                  string         `json:"clickUrl"`
	TrackerImageURL           string         `json:"trackerImageUrl"`
	ProductURL                string         `json:"productUrl"`
	ProductRating             int            `json:"productRating"`
	ProductPrice              string         `json:"productPrice"`
	PriceInt                  int            `json:"priceInt"`
	ID                        int            `json:"id"`
	ProductName               string         `json:"productName"`
	CategoryBreadcrumbs       string         `json:"categoryBreadcrumbs"`
	RecommendationType        string         `json:"recommendationType"`
	Stock                     int            `json:"stock"`
	DepartmentID              int            `json:"departmentID"`
	Shop                      RecomShop      `json:"shop"`
	ProductLabels             []ProductLabel `json:"productLabels"`
	LabelGroup                []LabelGroups  `json:"labelGroup"`
	WholesalePrice            []interface{}  `json:"wholesalePrice"`
	Badges                    []Badge        `json:"badges"`
	Typename                  string         `json:"__typename"`
}

type ProductRecommendationWidget struct {
	Data []struct {
		TID            string          `json:"tID"`
		Source         string          `json:"source"`
		Title          string          `json:"title"`
		ForeignTitle   string          `json:"foreignTitle"`
		SeeMoreURLLink string          `json:"seeMoreUrlLink"`
		LayoutType     string          `json:"layoutType"`
		PageName       string          `json:"pageName"`
		WidgetURL      string          `json:"widgetUrl"`
		Pagination     Pagination      `json:"pagination"`
		Recommendation []Recomendation `json:"recommendation"`
		Typename       string          `json:"__typename"`
	} `json:"data"`
	Meta struct {
		Recommendation    string `json:"recommendation"`
		Size              int    `json:"size"`
		FailSize          int    `json:"failSize"`
		ProcessTime       string `json:"processTime"`
		ExperimentVersion string `json:"experimentVersion"`
		Typename          string `json:"__typename"`
	} `json:"meta"`
	Typename string `json:"__typename"`
}

type RecomWidgetResp struct {
	Data struct {
		ProductRecommendationWidget ProductRecommendationWidget `json:"productRecommendationWidget"`
	} `json:"data"`
}
