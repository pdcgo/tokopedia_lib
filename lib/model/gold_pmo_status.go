package model

type PMOSStatusData struct {
	PowerMerchant struct {
		Status     string `json:"status"`
		AutoExtend struct {
			Status        string `json:"status"`
			TkpdProductID int    `json:"tkpd_product_id"`
			Typename      string `json:"__typename"`
		} `json:"auto_extend"`
		PmTier      int    `json:"pm_tier"`
		ExpiredTime string `json:"expired_time"`
		Typename    string `json:"__typename"`
	} `json:"power_merchant"`
	Typename string `json:"__typename"`
}

type GoldGetPMOSStatus struct {
	Data     *PMOSStatusData `json:"data"`
	Header   *Header         `json:"header"`
	Typename string          `json:"__typename"`
}

type GoldGetPMOSStatusData struct {
	GoldGetPMOSStatus *GoldGetPMOSStatus `json:"goldGetPMOSStatus"`
}
