package model

import (
	"regexp"
	"strconv"
	"strings"
)

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

var priceGrabRX = regexp.MustCompile(`[0-9]`)

func (v *VPVProductValidateV3DataVariants) GetPriceGab() int {

	for _, message := range v.Messages {
		if strings.Contains(message, "perbedaan harga") {
			num := priceGrabRX.FindString(message)
			fixnum, _ := strconv.Atoi(num)
			return fixnum
		}
	}

	return 0
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
