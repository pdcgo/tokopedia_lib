package api_public

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
)

var ClientApi *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 5,
	},
	Timeout: 30 * time.Second,
}

type SessionPublic struct {
	Cookies []*http.Cookie
	Ua      string
}

func (sess *SessionPublic) Update(cookies []*http.Cookie) error {
	for _, cookie := range cookies {
		err := sess.updateCookie(cookie)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sess *SessionPublic) updateCookie(cookie *http.Cookie) error {

	fixCookies := []*http.Cookie{}

	for _, oldCookie := range sess.Cookies {
		if oldCookie.Name == cookie.Name {
			fixCookies = append(fixCookies, cookie)
		} else {
			fixCookies = append(fixCookies, oldCookie)
		}
	}
	sess.Cookies = fixCookies
	return nil
}

type TokopediaApiPublic struct {
	Session SessionPublic
}

func (api *TokopediaApiPublic) NewRequest(method, ur string, query any, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, ur, body)
	if err != nil {
		pdc_common.ReportError(err)
	}
	// setting query'
	if query != nil {
		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()
	}

	// log.Info().Msg(req.URL.String())
	for _, cookie := range api.Session.Cookies {
		req.AddCookie(cookie)
	}

	return req
}

func (api *TokopediaApiPublic) SendRequest(req *http.Request, hasil any) error {
	res, err := ClientApi.Do(req)
	if err != nil {
		pdc_common.ReportError(err)
		return err
	}

	body, _ := io.ReadAll(res.Body)
	// log.Println(string(body))
	err = json.Unmarshal(body, hasil)
	if err != nil {
		return pdc_common.ReportError(err)
	}
	// return nil
	return api.Session.Update(res.Cookies())
}

func defaultHeader() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Origin":       "https://www.tokopedia.com",
		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	}
	return headers
}

func NewTokopediaApiPublic() (*TokopediaApiPublic, error) {
	headers := defaultHeader()
	req, err := http.NewRequest(http.MethodGet, "https://www.tokopedia.com/", nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := ClientApi.Do(req)
	if err != nil {
		return nil, err
	}

	pSession := SessionPublic{
		Cookies: resp.Cookies(),
		Ua:      resp.Request.UserAgent(),
	}

	return &TokopediaApiPublic{
		Session: pSession,
	}, nil
}
