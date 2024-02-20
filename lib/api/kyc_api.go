package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type InfoKycType struct {
	Typename   string `json:"__typename"`
	TypeID     int    `json:"TypeID"`
	Status     int    `json:"Status"`
	StatusName string `json:"StatusName"`
}

type InfoKycRes struct {
	Typename   string         `json:"__typename"`
	Status     int            `json:"Status"`
	StatusName string         `json:"StatusName"`
	Message    string         `json:"Message"`
	Reason     []string       `json:"Reason"`
	TypeList   []*InfoKycType `json:"TypeList"`
}

var KycRx = regexp.MustCompile(`{"__typename":"KycProjectInfoResult",(.*)}]}`)
var ErrKycInfoNotFound = errors.New("kyc info not found")

func (api *TokopediaApi) GetInfoKyc() (*InfoKycRes, error) {

	req := api.NewRequest(http.MethodGet, "https://mitra.tokopedia.com/user/akun-saya", nil, nil)

	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language":           "en-US,en;q=0.9",
		"Cache-Control":             "no-cache",
		"Connection":                "keep-alive",
		"Host":                      "mitra.tokopedia.com",
		"Pragma":                    "no-cache",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"ect":                       "4g",
		"sec-ch-ua":                 `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Windows"`,
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	var hasil InfoKycRes
	err := api.SendRequestCustomParse(req, func(body []byte) error {
		data := KycRx.Find(body)

		if len(data) == 0 {
			return ErrKycInfoNotFound
		}

		return json.Unmarshal(data, &hasil)
	})

	return &hasil, err
}

type SubmitKycHeader struct {
	model.Header
}

type SubmitKycDataApp struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Button   string `json:"button"`
}

type SubmitKycData struct {
	IsSuccessRegister bool             `json:"is_success_register"`
	ListRetake        any              `json:"list_retake"`
	ListMessage       any              `json:"list_message"`
	Apps              SubmitKycDataApp `json:"apps"`
}

type SubmitKycRes struct {
	Header SubmitKycHeader `json:"header"`
	Data   *SubmitKycData  `json:"data"`
}

// TODO: response success
func (api *TokopediaApi) SubmitKyc(imgKtpFile, imgSelfieFile io.Reader) (res *SubmitKycRes, err error) {

	res = &SubmitKycRes{}
	body := &bytes.Buffer{}
	boundary := getBoundary()
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary)

	writer.WriteField("project_id", "5")
	writer.WriteField("params", `[{"kyc_type":1,"param":"imgKtp"},{"kyc_type":2,"param":"imgSelfie"}]`)

	part, _ := CreateFormFile(writer, "imgKtp", "imgKtp.jpeg")
	_, err = io.Copy(part, imgKtpFile)
	if err != nil {
		return
	}

	part, _ = CreateFormFile(writer, "imgSelfie", "imgSelfie.jpeg")
	_, err = io.Copy(part, imgKtpFile)
	if err != nil {
		return
	}

	f, _ := os.Create("test.daaa")
	f.Write(body.Bytes())
	defer f.Close()

	req := api.NewRequest(http.MethodPost, "https://accounts.tokopedia.com/kycapp/api/v1/validate-register", nil, body)
	headers := map[string]string{
		"Accept":          "application/json",
		"Accept-Language": "en-US,en;q=0.9",
		"Connection":      "keep-alive",
		// "Content-Length":     fmt.Sprint(len(body.Bytes())),
		"Content-Type":       fmt.Sprintf("multipart/form-data; boundary=%s", boundary),
		"Host":               "accounts.tokopedia.com",
		"Origin":             "https://mitra.tokopedia.com",
		"Referer":            "https://mitra.tokopedia.com/kyc/3",
		"Sec-Ch-Ua":          `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": `"Windows"`,
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	err = api.SendRequest(req, res)
	return
}
