package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
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

func headerUploadImage(boundary string) map[string]string {

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		"Content-Type": "multipart/form-data; boundary=" + boundary,
		"Accept":       "*/*",
		"Origin":       "https://seller.tokopedia.com",
	}
	return headers
}

func CreateFormFile(w *multipart.Writer, name string, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, name, filename))
	h.Set("Content-Type", "image/jpeg")
	return w.CreatePart(h)
}

type UpImageHeader struct {
	ProcessTime float64  `json:"process_time"`
	Reason      string   `json:"reason"`
	ErrorCode   string   `json:"error_code"`
	IsSuccess   bool     `json:"is_success"`
	Messages    []string `json:"messages"`
}

func (head *UpImageHeader) Error() string {
	return strings.Join(head.Messages, "|")

}

type UploadMediaResp struct {
	Header *UpImageHeader `json:"header"`
	Data   struct {
		UploadID string `json:"upload_id"`
		ImageURL string `json:"image_url"`
	} `json:"data"`
}

func (api *TokopediaApi) UploadImageFromUrl(uri string) (*UploadMediaResp, error) {
	res, err := ClientApi.Get(uri)

	if err != nil {
		return nil, err
	}

	data := res.Body
	return api.UploadProductImage(data)

}

func (api *TokopediaApi) UploadProductImage(content io.Reader) (*UploadMediaResp, error) {
	uri := "https://upedia.tokopedia.net/v1/upload/image/VqbcmM"

	boundary := getBoundary()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary)

	part, _ := CreateFormFile(writer, "file_upload", "blob")
	_, err := io.Copy(part, content)
	if err != nil {
		return nil, err
	}
	writer.Close()

	r, _ := http.NewRequest("POST", uri, body)
	headers := headerUploadImage(boundary)
	for key, value := range headers {
		r.Header.Set(key, value)
	}

	res, err := ClientApi.Do(r)
	if err != nil {
		return nil, err
	}
	resBody, _ := io.ReadAll(res.Body)
	var hasil UploadMediaResp

	err = json.Unmarshal(resBody, &hasil)
	if err != nil {
		return nil, err
	}

	if !hasil.Header.IsSuccess {
		return &hasil, hasil.Header
	}
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
	_, errCopy := io.Copy(part, file)
	if err != nil {
		pdc_common.ReportError(errCopy)
	}

	// source
	writer.WriteField("source", "topchat")
	writer.Close()

	r := api.NewRequest("POST", uri, nil, body)
	headers := headerUploadImage(boundary)
	for key, value := range headers {
		r.Header.Set(key, value)
	}

	var hasil ImageChatRes
	api.SendRequest(r, &hasil)
	return &hasil, nil
}
