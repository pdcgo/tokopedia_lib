package tokpedproxy

import (
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type TamperingTracking struct {
	proxy.BaseAddon
}

func NewTamperingTracking() *TamperingTracking {

	return &TamperingTracking{}
}

func (t *TamperingTracking) Response(f *proxy.Flow) {

}
