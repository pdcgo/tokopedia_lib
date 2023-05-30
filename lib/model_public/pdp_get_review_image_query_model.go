package model_public

type DetailImage struct {
	AttachmentID string `json:"attachmentID"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailURL"`
	FullsizeURL  string `json:"fullsizeURL"`
	FeedbackID   string `json:"feedbackID"`
	Typename     string `json:"__typename"`
}

type Reviewer struct {
	UserID         string `json:"userID"`
	FullName       string `json:"fullName"`
	ProfilePicture string `json:"profilePicture"`
	Typename       string `json:"__typename"`
}

type DetailReview struct {
	Reviewer           Reviewer `json:"reviewer"`
	ShopID             string   `json:"shopID"`
	FeedbackID         string   `json:"feedbackID"`
	VariantName        string   `json:"variantName"`
	ReviewText         string   `json:"reviewText"`
	Rating             int      `json:"rating"`
	ReviewTime         string   `json:"reviewTime"`
	LikeCount          int      `json:"likeCount"`
	BadRatingReasonFmt string   `json:"badRatingReasonFmt"`
	IsLiked            bool     `json:"isLiked"`
	IsAnonymous        bool     `json:"isAnonymous"`
	IsReportable       bool     `json:"isReportable"`
	Typename           string   `json:"__typename"`
}

type ReviewImageDetail struct {
	Reviews    []DetailReview `json:"reviews"`
	Images     []DetailImage  `json:"images"`
	Video      []interface{}  `json:"video"`
	MediaCount int            `json:"mediaCount"`
	Typename   string         `json:"__typename"`
}

type ReviewImageList struct {
	ImageID    string `json:"imageID"`
	FeedbackID string `json:"feedbackID"`
	VideoID    string `json:"videoID"`
	Typename   string `json:"__typename"`
}

type ProductrevGetReviewImage struct {
	List     []ReviewImageList `json:"list"`
	Detail   ReviewImageDetail `json:"detail"`
	HasNext  bool              `json:"hasNext"`
	Typename string            `json:"__typename"`
}

type PdpGetReiewImageQueryVar struct {
	ProductID string `json:"productID"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
}

type PdpGetReiewImageQueryResp struct {
	Data struct {
		ProductrevGetReviewImage ProductrevGetReviewImage `json:"productrevGetReviewImage"`
	} `json:"data"`
}
