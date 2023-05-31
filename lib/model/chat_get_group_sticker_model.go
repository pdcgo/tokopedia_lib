package model

type GroupSticker struct {
	GroupUUID  string `json:"groupUUID"`
	LastUpdate string `json:"lastUpdate"`
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	Typename   string `json:"__typename"`
}

type ChatGetGroupSticker struct {
	List     []GroupSticker `json:"list"`
	Typename string         `json:"__typename"`
}

type TypeVar struct {
	Type int `json:"type"`
}

type ChatGetGroupStickerResp struct {
	Data struct {
		ChatListGroupSticker ChatGetGroupSticker `json:"chatListGroupSticker"`
	} `json:"data"`
}
