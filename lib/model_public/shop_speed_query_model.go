package model_public

type ShopIdStrVar struct {
	ShopID string `json:"shopID"`
}

type ShopSpeedQueryResp struct {
	Data struct {
		ShopSpeed struct {
			MessageResponseTime int    `json:"messageResponseTime"`
			Typename            string `json:"__typename"`
		} `json:"shopSpeed"`
	} `json:"data"`
}
