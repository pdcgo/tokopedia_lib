package tokopedia_lib

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
)

type SessionPublic struct {
	Cookies []*http.Cookie
	Ua      string
}

var pClient http.Client = http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   "localhost:8888",
		}),
	},
	Timeout: 30 * time.Second,
}

func defaultHeader() map[string]string {
	headers := map[string]string{
		"content-type": "application/json",
		"origin":       "https://www.tokopedia.com",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	}
	return headers
}

func (sess *SessionPublic) SetCookiesToReq(req *http.Request) {
	for _, cookie := range sess.Cookies {
		req.AddCookie(cookie)
	}
}

func (sess *SessionPublic) PublicUa() string {
	return sess.Ua
}

func NewSessionPublic() *SessionPublic {
	headers := defaultHeader()
	req, err := http.NewRequest(http.MethodGet, "https://www.tokopedia.com/", nil)
	if err != nil {
		pdc_common.ReportError(err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := pClient.Do(req)
	if err != nil {
		pdc_common.ReportError(err)
	}

	pSession := SessionPublic{
		Cookies: resp.Cookies(),
		Ua:      resp.Request.UserAgent(),
	}
	return &pSession
}
