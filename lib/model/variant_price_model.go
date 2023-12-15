package model

// VPV: Variant Price Validation
type VPVCategory struct {
	ID int `json:"id,string"`
}

type VPVInput struct {
	Variant  *Variant     `json:"variant"`
	Category *VPVCategory `json:"category"`
}

type VPVVariable struct {
	Input *VPVInput `json:"input"`
}

type VPVProductValidateV3DataVariants struct {
	Messages []string `json:"messages"`
	Typename string   `json:"__typename"`
}

type VPVProductValidateV3Data struct {
	Variants *VPVProductValidateV3DataVariants `json:"variants"`
	Typename string                            `json:"__typename"`
}

type VPVProductValidateV3 struct {
	Header    *Header                   `json:"header"`
	IsSuccess bool                      `json:"isSuccess"`
	Data      *VPVProductValidateV3Data `json:"data"`
	Typename  string                    `json:"__typename"`
}

type VPVData struct {
	ProductValidateV3 *VPVProductValidateV3 `json:"ProductValidateV3"`
}

type VPVRes struct {
	Data VPVData `json:"data"`
}
