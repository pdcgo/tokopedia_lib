package model_public

type State struct {
	IsMasked    bool   `json:"isMasked"`
	IsLiked     bool   `json:"isLiked,omitempty"`
	AllowLike   bool   `json:"allowLike,omitempty"`
	IsYours     bool   `json:"isYours"`
	AllowReport bool   `json:"allowReport"`
	AllowDelete bool   `json:"allowDelete"`
	AllowReply  bool   `json:"allowReply,omitempty"`
	IsFollowed  bool   `json:"isFollowed,omitempty"`
	AllowFollow bool   `json:"allowFollow,omitempty"`
	Typename    string `json:"__typename"`
}

type Answer struct {
	AnswerID             string `json:"answerID"`
	Content              string `json:"content"`
	MaskedContent        string `json:"maskedContent"`
	UserName             string `json:"userName"`
	UserThumbnail        string `json:"userThumbnail"`
	UserID               string `json:"userID"`
	IsSeller             bool   `json:"isSeller"`
	CreateTime           string `json:"createTime"`
	CreateTimeFormatted  string `json:"createTimeFormatted"`
	LikeCount            int    `json:"likeCount"`
	State                State  `json:"state"`
	AttachedProductCount int    `json:"attachedProductCount"`
	Typename             string `json:"__typename"`
}

type Question struct {
	QuestionID          string `json:"questionID"`
	Content             string `json:"content"`
	MaskedContent       string `json:"maskedContent"`
	UserName            string `json:"userName"`
	UserThumbnail       string `json:"userThumbnail"`
	UserID              string `json:"userID"`
	CreateTime          string `json:"createTime"`
	CreateTimeFormatted string `json:"createTimeFormatted"`
	State               State  `json:"state"`
	TotalAnswer         int    `json:"totalAnswer"`
	Answer              Answer `json:"answer"`
	Typename            string `json:"__typename"`
}

type DiscussionDataByProductID struct {
	ShopID        string     `json:"shopID"`
	ShopURL       string     `json:"shopURL"`
	ProductID     string     `json:"productID"`
	HasNext       bool       `json:"hasNext"`
	TotalQuestion int        `json:"totalQuestion"`
	Question      []Question `json:"question"`
	Typename      string     `json:"__typename"`
}

type DiscussionDataProductByIDVar struct {
	ProductID string `json:"productID"`
	ShopID    string `json:"shopID"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sortBy"`
	Category  string `json:"category"`
}

type DiscussionDataProductByIDResp struct {
	Data struct {
		DiscussionDataByProductID DiscussionDataByProductID `json:"discussionDataByProductID"`
	} `json:"data"`
}
