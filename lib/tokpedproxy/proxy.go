package tokpedproxy

import (
	"context"
	"log"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"github.com/pdcgo/common_conf/pdc_common"
	logrus "github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel) // TODO: EXPERIMENTAL REMOVE LOG MITM PROXY GOLANG
}

// TODO: kalau ada waktu hilangkan log
type InspectProxy struct {
	Ctx    context.Context
	Addr   string
	Addons []proxy.Addon
	proxy  *proxy.Proxy
}

func (i *InspectProxy) RunProxy() {
	opts := &proxy.Options{
		Addr:              i.Addr,
		StreamLargeBodies: 1024 * 1024 * 5,
	}

	p, err := proxy.NewProxy(opts)

	if err != nil {
		log.Fatal(err)
	}

	for _, addon := range i.Addons {
		p.AddAddon(addon)
	}

	i.proxy = p

	select {
	case <-i.Ctx.Done():
		err := i.proxy.Close()
		if err != nil {
			pdc_common.ReportError(err)
		}
		return

	default:
		err := i.proxy.Start()
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

}

func NewInspectProxy(listen string, ctx context.Context) *InspectProxy {

	addons := []proxy.Addon{
		NewTamperingTracking(),
	}

	return &InspectProxy{
		Ctx:    ctx,
		Addr:   listen,
		Addons: addons,
	}
}
