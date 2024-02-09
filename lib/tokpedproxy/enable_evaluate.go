package tokpedproxy

import (
	"log"
	"regexp"
	"strings"

	"github.com/pdcgo/go-mitmproxy/proxy"
)

type EnbaleEvaluate struct {
	proxy.BaseAddon
}

func NewEnableEvaluate() proxy.Addon {
	return &EnbaleEvaluate{}
}

func (t *EnbaleEvaluate) Responseheaders(f *proxy.Flow) {
	url := f.Request.URL.String()
	contentType := f.Response.Header.Get("Content-Type")

	if strings.Contains(url, "tokopedia.com/") && strings.Contains(contentType, "text/html") {
		log.Println("tampering evaluate")

		pattern := "'unsafe-inline'"
		headerName := "Content-Security-Policy"
		strReplacer := "'unsafe-eval' 'unsafe-inline'"

		regex := regexp.MustCompile(pattern)

		csp := f.Response.Header.Get(headerName)
		newCsp := regex.ReplaceAllString(csp, strReplacer)

		f.Response.Header.Set(headerName, newCsp)
	}
}
