package model

type ImageAttachment struct {
	AttachmentID      string `json:"attachmentID"`
	ImageThumbnailURL string `json:"imageThumbnailUrl"`
	ImageURL          string `json:"imageUrl"`
	Typename          string `json:"__typename"`
}

type ReviewResponse struct {
	Message    string `json:"message"`
	CreateTime string `json:"createTime"`
	Typename   string `json:"__typename"`
}

type ReviewStats struct {
	Key       string `json:"key"`
	Formatted string `json:"formatted"`
	Count     int    `json:"count"`
	Typename  string `json:"__typename"`
}

type Shop struct {
	ShopID   string `json:"shopID"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Image    string `json:"image"`
	Typename string `json:"__typename"`
}

type User struct {
	UserID   string `json:"userID"`
	FullName string `json:"fullName"`
	Image    string `json:"image"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
}

type LikeDislike struct {
	TotalLike  int    `json:"totalLike"`
	LikeStatus int    `json:"likeStatus"`
	Typename   string `json:"__typename"`
}

type Review struct {
	ID                    string            `json:"id"`
	VariantName           string            `json:"variantName"`
	Message               string            `json:"message"`
	ProductRating         int               `json:"productRating"`
	ReviewCreateTime      string            `json:"reviewCreateTime"`
	ReviewCreateTimestamp string            `json:"reviewCreateTimestamp"`
	IsReportable          bool              `json:"isReportable"`
	IsAnonymous           bool              `json:"isAnonymous"`
	ImageAttachments      []ImageAttachment `json:"imageAttachments"`
	VideoAttachments      []interface{}     `json:"videoAttachments"`
	ReviewResponse        ReviewResponse    `json:"reviewResponse"`
	User                  User              `json:"user"`
	LikeDislike           LikeDislike       `json:"likeDislike"`
	Stats                 []ReviewStats     `json:"stats"`
	BadRatingReasonFmt    string            `json:"badRatingReasonFmt"`
	Typename              string            `json:"__typename"`
}

type ProductrevGetProductReviewList struct {
	ProductID    string   `json:"productID"`
	List         []Review `json:"list"`
	Shop         Shop     `json:"shop"`
	HasNext      bool     `json:"hasNext"`
	TotalReviews int      `json:"totalReviews"`
	Typename     string   `json:"__typename"`
}

type ProductReviewListVar struct {
	ProductID string `json:"productID"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sortBy"`
	FilterBy  string `json:"filterBy"`
}

type ProductReviewListResp struct {
	Data struct {
		ProductrevGetProductReviewList ProductrevGetProductReviewList `json:"productrevGetProductReviewList"`
	} `json:"data"`
}
