package api_public

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/sethvargo/go-retry"
)

// var CountClientPool int = 5

type TokopediaApiPublic struct {
	sync.RWMutex
	Client *http.Client
	guard  chan int
}

func (api *TokopediaApiPublic) NewRequest(method, ur string, query any, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, ur, body)
	if err != nil {
		pdc_common.ReportError(err)
	}
	// setting query
	if query != nil {
		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()
	}

	// log.Info().Msg(req.URL.String())
	// for _, cookie := range api.Session.Cookies {
	// 	req.AddCookie(cookie)
	// }

	return req
}

func (api *TokopediaApiPublic) SendRequest(req *http.Request, hasil any) error {
	res, err := api.Client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 403 {
		log.Println("retry 403")

		return errors.New("request api limit")
	}

	body, _ := io.ReadAll(res.Body)
	// log.Println(string(body))

	if res.StatusCode != 200 {
		log.Println(string(body))
		return retry.RetryableError(errors.New("request api limit"))
	}

	err = json.Unmarshal(body, hasil)
	if err != nil {
		log.Println(string(body))
		return err
	}

	return nil
}

// func defaultHeader() map[string]string {
// 	headers := map[string]string{
// 		"Content-Type": "application/json",
// 		"Origin":       "https://www.tokopedia.com",
// 		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
// 		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
// 	}
// 	return headers
// }

func NewTokopediaApiPublic() (*TokopediaApiPublic, error) {

	// headers := defaultHeader()
	// req, err := http.NewRequest(http.MethodGet, "https://www.tokopedia.com/", nil)
	// if err != nil {
	// 	return nil, err
	// }
	// for key, value := range headers {
	// 	req.Header.Set(key, value)
	// }
	// _, err = client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	return nil, err
	// }

	// transport := &http2.Transport{
	// 	DisableCompression: false,
	// }
	transport := &http.Transport{
		ForceAttemptHTTP2:  false,
		DisableCompression: false,
		// DisableKeepAlives:  true,
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}

	return &TokopediaApiPublic{
		guard: make(chan int, 100),

		Client: &http.Client{
			Transport: transport,
			// Transport: &http.Transport{
			// 	MaxIdleConnsPerHost: 10,
			// },
			// Jar:     jar,
			Timeout: 60 * time.Second,
		},
	}, nil
}
