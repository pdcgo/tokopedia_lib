package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/common_conf/pdc_common"
)

type PurnaSession struct {
	Data []*network.Cookie
}

func (sess *PurnaSession) SetCookieToDriver(ctx context.Context) error {
	return chromedp.Run(
		ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {

			for _, cookie := range sess.Data {
				// expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))

				// log.Println(cookie.Name, cookie.Value)

				err := network.SetCookie(cookie.Name, cookie.Value).
					WithDomain(cookie.Domain).
					// WithPath(cookie.Path).
					// WithHTTPOnly(cookie.HTTPOnly).
					// WithSecure(cookie.Secure).
					// WithExpires(&expr).
					Do(ctx)

				if err != nil {
					if !errors.Is(context.Canceled, err) {
						pdc_common.ReportError(err)
					}

				}
			}
			return nil
		}),
	)
}
func (sess *PurnaSession) Load() error                         { return nil }
func (sess *PurnaSession) DeleteSession() error                { return nil }
func (sess *PurnaSession) SaveSession() error                  { return nil }
func (sess *PurnaSession) Sync() error                         { return nil }
func (sess *PurnaSession) Update(cookies []*http.Cookie) error { return nil }
func (sess *PurnaSession) AddToHttpRequest(req *http.Request)  {}
func (sess *PurnaSession) UserAgent() string {
	return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
}
func (sess *PurnaSession) GetCookies() []*http.Cookie {
	log.Println("not Implemented")
	return []*http.Cookie{}
}
func (sess *PurnaSession) SaveFromDriver(cookies []*network.Cookie, ua string) error { return nil }

func NewPurnaSession(fname string) (*PurnaSession, error) {
	purnaSess := PurnaSession{
		Data: []*network.Cookie{},
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		return &purnaSess, err
	}

	err = json.Unmarshal(data, &purnaSess.Data)
	if err != nil {
		return &purnaSess, err
	}
	return &purnaSess, nil
}

type SessionHandler func(fname string, sess *PurnaSession) error

func WalkPurnaSession(base string, handler SessionHandler) {
	openBase := filepath.Join(base, "opened")
	os.MkdirAll(openBase, 0775)

	basecookies := filepath.Join(base, "*.json")

	files, err := filepath.Glob(basecookies)

	if err != nil {
		log.Panicln(err)
	}

	for _, file := range files {
		sess, err := NewPurnaSession(file)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		err = handler(file, sess)
		if err != nil {
			pdc_common.ReportError(err)
		}

		_, fname := filepath.Split(file)
		openedpath := filepath.Join(openBase, fname)
		os.Rename(file, openedpath)

	}

}
