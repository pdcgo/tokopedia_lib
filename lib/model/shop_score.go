package model

type ShopScoreLevelError struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}

type ShopScoreDetail struct {
	Title        string  `json:"title"`
	Identifier   string  `json:"identifier"`
	Value        float32 `json:"value"`
	RawValue     float32 `json:"rawValue"`
	NextMinValue float64 `json:"nextMinValue"`
	ColorText    string  `json:"colorText"`
	Typename     string  `json:"__typename"`
}

type ShopScoreResult struct {
	ShopID          string             `json:"shopID"`
	ShopScore       float32            `json:"shopScore"`
	ShopLevel       int                `json:"shopLevel"`
	ShopScoreDetail []*ShopScoreDetail `json:"shopScoreDetail"`
	Period          string             `json:"period"`
	NextUpdate      string             `json:"nextUpdate"`
	Typename        string             `json:"__typename"`
}

type ShopScoreLevel struct {
	Result   *ShopScoreResult     `json:"result"`
	Error    *ShopScoreLevelError `json:"error"`
	Typename string               `json:"__typename"`
}

type ShopLevelResult struct {
	ShopID     string `json:"shopID"`
	Period     string `json:"period"`
	NextUpdate string `json:"nextUpdate"`
	ShopLevel  int    `json:"shopLevel"`
	ItemSold   int    `json:"itemSold"`
	Niv        int    `json:"niv"`
	Typename   string `json:"__typename"`
}

type ShopLevel struct {
	Result   *ShopLevelResult     `json:"result"`
	Error    *ShopScoreLevelError `json:"error"`
	Typename string               `json:"__typename"`
}

type ShopScoreData struct {
	ShopScoreLevel *ShopScoreLevel `json:"shopScoreLevel"`
	ShopLevel      *ShopLevel      `json:"shopLevel"`
}
