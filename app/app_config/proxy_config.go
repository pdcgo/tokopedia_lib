package app_config

import (
	"context"
	"errors"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/pdcgo/go-mitmproxy/proxy"
	"github.com/pdcgo/go_v2_shopeelib/lib/shopee_dp"
	"github.com/pdcgo/go_v2_shopeelib/lib/shopee_dp/addon"
	"github.com/pdcgo/tokopedia_lib/lib/tokpedproxy"
)

type ProxyKey string

const (
	WithdrawProxyKey ProxyKey = "main_proxy"
)

type addonCreator func() proxy.Addon

type ProxyConf struct {
	Addr   string
	Addons []addonCreator
}

func NewChatFrontendCacher(path string) proxy.Addon {
	return addon.NewFrontendCacher([]string{"https://deo.shopeemobile.com"}, addon.NewCacheStore())
}

var proxyConfig map[ProxyKey]*ProxyConf = map[ProxyKey]*ProxyConf{

	WithdrawProxyKey: {
		Addr: "127.0.0.1:8081",
		Addons: []addonCreator{
			tokpedproxy.NewDisableWebdriver,
		},
	},
}

type ProxyItem struct {
	Instance  *shopee_dp.InspectProxy
	CancelCtx context.CancelFunc
}

var proxyList map[ProxyKey]*ProxyItem = map[ProxyKey]*ProxyItem{}
var maplock sync.Mutex

func GetProxy(key ProxyKey) (*shopee_dp.InspectProxy, context.CancelFunc) {
	maplock.Lock()
	defer maplock.Unlock()

	prox := proxyList[key]

	if prox == nil {
		config := proxyConfig[key]
		if config == nil {
			panic("no config for " + key)
		}

		err := raw_connect(config.Addr)
		if err != nil {
			panic(err)
		}

		ctx, cancel := context.WithCancel(context.TODO())
		proxInstance := shopee_dp.NewInspectProxy(config.Addr, ctx, []proxy.Addon{})

		for _, handler := range config.Addons {
			proxInstance.Addons = append(proxInstance.Addons, handler())
		}

		prox = &ProxyItem{
			Instance:  proxInstance,
			CancelCtx: cancel,
		}

		go proxInstance.RunProxy()

		proxyList[key] = prox

	}

	return prox.Instance, prox.CancelCtx
}

func raw_connect(rawuri string) error {

	uri, _ := url.Parse("http://" + rawuri)
	port := uri.Port()
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(uri.Host, port), timeout)
	if err != nil {
		return nil
	}
	if conn != nil {
		defer conn.Close()
		return errors.New("port terpakai")
	}

	return nil
}
