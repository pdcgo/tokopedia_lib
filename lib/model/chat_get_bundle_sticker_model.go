package model

type BundleSticker struct {
	ImageURL    string `json:"imageUrl"`
	Intention   string `json:"intention"`
	StickerUUID string `json:"stickerUUID"`
	GroupUUID   string `json:"groupUUID"`
	Typename    string `json:"__typename"`
}

type ChatBundleSticker struct {
	List     []BundleSticker `json:"list"`
	HasNext  bool            `json:"hasNext"`
	MaxUUID  string          `json:"maxUUID"`
	Typename string          `json:"__typename"`
}

type ChatGetBundleStickerVar struct {
	ID    string      `json:"id"`
	Limit int         `json:"limit"`
	MaxID interface{} `json:"maxId"`
}

type ChatGetBundleStickerResp struct {
	Data struct {
		ChatBundleSticker ChatBundleSticker `json:"chatBundleSticker"`
	} `json:"data"`
}
