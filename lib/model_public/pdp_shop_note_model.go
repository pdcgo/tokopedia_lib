package model_public

type Result struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	URL        string `json:"url"`
	UpdateTime string `json:"updateTime"`
	Typename   string `json:"__typename"`
}

type PdpShopNoteResp struct {
	Data struct {
		ShopNotesByShopID struct {
			Result   []Result `json:"result"`
			Error    Error    `json:"error"`
			Typename string   `json:"__typename"`
		} `json:"shopNotesByShopID"`
	} `json:"data"`
}
