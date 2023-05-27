package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strconv"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
)

func getBoundary() string {
	var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < 16; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return "----WebKitFormBoundary" + s
}

func headerUploadImage(fileSize int64, boundary string) map[string]string {
	contentLength := strconv.Itoa(int(fileSize))

	headers := map[string]string{
		"Content-Length":     contentLength,
		"Sec-Ch-Ua":          `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`,
		"Sec-Ch-Ua-Platform": `"Windows"`,
		"Sec-Ch-Ua-Mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		"Content-Type":       "multipart/form-data; boundary=" + boundary,
		"Accept":             "*/*",
		"Origin":             "https://seller.tokopedia.com",
		"Sec-Fetch-Site":     "cross-site",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Dest":     "empty",
		"Referer":            "https://seller.tokopedia.com/add-product",
		"Accept-Encoding":    "gzip, deflate, br",
		"Accept-Language":    "en-US,en;q=0.9",
	}
	return headers
}

func CreateFormFile(w *multipart.Writer, name string, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, name, filename))
	h.Set("Content-Type", "image/jpeg")
	return w.CreatePart(h)
}

type UploadMediaResp struct {
	Header struct {
		ProcessTime float64       `json:"process_time"`
		Reason      string        `json:"reason"`
		ErrorCode   string        `json:"error_code"`
		IsSuccess   bool          `json:"is_success"`
		Messages    []interface{} `json:"messages"`
	} `json:"header"`
	Data struct {
		UploadID string `json:"upload_id"`
		ImageURL string `json:"image_url"`
	} `json:"data"`
}

func (api *TokopediaApi) UploadProductImage(locfile string) (*UploadMediaResp, error) {
	uri := "https://upedia.tokopedia.net/v1/upload/image/VqbcmM"

	file, err := os.Open(locfile)
	if err != nil {
		pdc_common.ReportError(err)
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	boundary := getBoundary()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary)

	part, _ := CreateFormFile(writer, "file_upload", "blob")
	_, errCopy := io.Copy(part, file)
	if err != nil {
		log.Println(errCopy)
	}
	writer.Close()

	r, _ := http.NewRequest("POST", uri, body)
	headers := headerUploadImage(fileStat.Size(), boundary)
	for key, value := range headers {
		r.Header.Set(key, value)
	}

	client := &http.Client{
		// Transport: &http.Transport{
		// 	Proxy: http.ProxyURL(&url.URL{
		// 		Scheme: "http",
		// 		Host:   "localhost:8888",
		// 	})},
	}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	resBody, _ := io.ReadAll(res.Body)
	var hasil UploadMediaResp

	json.Unmarshal(resBody, &hasil)
	return &hasil, nil
}

type ImageChatRes struct {
	Data struct {
		URLImage string `json:"url_image"`
	} `json:"data"`
	ServerProcessTime float64 `json:"server_process_time"`
	Server            string  `json:"server"`
	Status            string  `json:"status"`
	Success           int     `json:"success"`
}

func (api *TokopediaApi) UploadImageChat(msgId string, locfile string) (*ImageChatRes, error) {
	uri := "https://chat.tokopedia.com/tc/v1/upload_secure"

	file, err := os.Open(locfile)
	if err != nil {
		pdc_common.ReportError(err)
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	boundary := getBoundary()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary)
	// msg_id
	writer.WriteField("msg_id", msgId)

	// file
	part, _ := CreateFormFile(writer, "file", fileStat.Name())
	log.Println(boundary)
	_, errCopy := io.Copy(part, file)
	if err != nil {
		pdc_common.ReportError(errCopy)
	}

	// source
	writer.WriteField("source", "topchat")
	writer.Close()

	r := api.NewRequest("POST", uri, nil, body)
	headers := headerUploadImage(fileStat.Size(), boundary)
	for key, value := range headers {
		r.Header.Set(key, value)
	}

	var hasil ImageChatRes
	api.SendRequest(r, &hasil)
	return &hasil, nil
}
