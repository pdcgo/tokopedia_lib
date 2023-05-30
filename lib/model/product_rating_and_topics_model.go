package model

type ProductIdVar struct {
	ProductId string `json:"productID"`
}

type Rating struct {
	PositivePercentageFmt   string         `json:"positivePercentageFmt"`
	RatingScore             string         `json:"ratingScore"`
	TotalRating             int            `json:"totalRating"`
	TotalRatingWithImage    int            `json:"totalRatingWithImage"`
	TotalRatingTextAndImage int            `json:"totalRatingTextAndImage"`
	Detail                  []RatingDetail `json:"detail"`
	Typename                string         `json:"__typename"`
}

type AvailableFilter struct {
	WithAttachment bool   `json:"withAttachment"`
	Rating         bool   `json:"rating"`
	Topics         bool   `json:"topics"`
	Helpfulness    bool   `json:"helpfulness"`
	Typename       string `json:"__typename"`
}

type Topic struct {
	Rating         float64 `json:"rating"`
	RatingFmt      string  `json:"ratingFmt"`
	Formatted      string  `json:"formatted"`
	Key            string  `json:"key"`
	ReviewCount    int     `json:"reviewCount"`
	ReviewCountFmt string  `json:"reviewCountFmt"`
	Show           bool    `json:"show"`
	Typename       string  `json:"__typename"`
}

type ProductrevGetProductRatingAndTopics struct {
	ProductID        string          `json:"productID"`
	Rating           Rating          `json:"rating"`
	Topics           []Topic         `json:"topics"`
	AvailableFilters AvailableFilter `json:"availableFilters"`
	Typename         string          `json:"__typename"`
}

type ProductRatingandTopicsResp struct {
	Data struct {
		ProductrevGetProductRatingAndTopics ProductrevGetProductRatingAndTopics `json:"productrevGetProductRatingAndTopics"`
	} `json:"data"`
}
