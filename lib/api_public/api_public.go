package api_public

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
)

var ClientApi *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   "localhost:8888",
		}),
	},
	Timeout: 30 * time.Second,
}

type Session interface {
	PublicUa() string
	SetCookiesToReq(*http.Request)
}

type TokopediaApiPublic struct {
	Session Session
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
	api.Session.SetCookiesToReq(req)

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
	return nil
	// return api.Session.Update(res.Cookies())
}

func NewTokopediaApiPublic() *TokopediaApiPublic {

	return &TokopediaApiPublic{
		Session: session,
	}
}
