package tokpedproxy

import (
	"strconv"
	"strings"

	"github.com/pdcgo/go-mitmproxy/proxy"
)

type DisableWebdriver struct {
	proxy.BaseAddon
}

func NewDisableWebdriver() proxy.Addon {
	return &DisableWebdriver{}
}

func (t *DisableWebdriver) Response(f *proxy.Flow) {

	url := f.Request.URL.String()
	contentType := f.Response.Header.Get("Content-Type")

	if strings.Contains(url, "/payment/deposit") && strings.Contains(contentType, "text/html") {

		f.Response.ReplaceToDecodedBody()

		body := string(f.Response.Body)
		bodies := strings.Split(body, "\n")

		for i, b := range bodies {
			if strings.Contains(b, "<script") && b[0:7] == "<script" {

				bodies = append(bodies[:i+1], bodies[i:]...)
				bodies[i] = `<script>
					const newProto = navigator.__proto__;
					delete newProto.webdriver;
					navigator.__proto__ = newProto;
				</script>`
				break
			}
		}

		body = strings.Join(bodies, "\n")
		f.Response.Body = []byte(body)
		f.Response.Header.Set("Content-Length", strconv.Itoa(len(body)))

		return
	}
}
