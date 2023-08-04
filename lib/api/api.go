package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/pdcgo/common_conf/pdc_common"
)

var ClientApi *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 5,
	},
	Timeout: 30 * time.Second,
}

type Session interface {
	Sync() error
	Update(cookies []*http.Cookie) error
	AddToHttpRequest(req *http.Request)
	UserAgent() string
	GetCookies() []*http.Cookie
}

type TokopediaApi struct {
	Session           Session
	encoder           *schema.Encoder
	AuthenticatedData *IsAtuheticatedData
}

func (api *TokopediaApi) NewRequest(method, ur string, query any, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, ur, body)
	if err != nil {
		pdc_common.ReportError(err)
	}
	// setting query'
	if query != nil {
		q := req.URL.Query()
		api.encoder.Encode(query, q)
		req.URL.RawQuery = q.Encode()
	}

	// log.Info().Msg(req.URL.String())
	api.Session.AddToHttpRequest(req)

	return req
}

func (api *TokopediaApi) SendRequest(req *http.Request, hasil any) error {
	res, err := ClientApi.Do(req)
	if err != nil {
		return err
	}

	body, _ := io.ReadAll(res.Body)
	// log.Println(string(body))
	err = json.Unmarshal(body, hasil)
	if err != nil {
		return err
	}

	return api.Session.Update(res.Cookies())
}

func NewTokopediaApi(session Session) *TokopediaApi {
	encoder := schema.NewEncoder()

	return &TokopediaApi{
		Session: session,
		encoder: encoder,
	}
}
