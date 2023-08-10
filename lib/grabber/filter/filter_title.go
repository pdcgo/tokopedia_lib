package filter

import (
	"regexp"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func CreateTitleFilter(filterText *legacy_source.FilterText) FilterHandler {
	regexs := []*regexp.Regexp{}
	texts := []string{}

	for _, s := range filterText.Title {
		if strings.Contains(s, "regex-->") {
			rx := regexp.MustCompile(s[8:])
			regexs = append(regexs, rx)
		} else {
			texts = append(texts, strings.ToLower(s))
		}
	}

	return func(layout *model_public.PdpGetlayoutQueryResp, pdp *model_public.PdpGetDataP2Resp) (cek bool, reason string, err error) {
		productName := layout.Data.PdpGetLayout.GetProductName()
		productName = strings.ToLower(productName)

		for _, t := range texts {
			if strings.Contains(productName, t) {
				return true, "filter title", nil
			}
		}

		for _, rx := range regexs {
			if rx.MatchString(productName) {
				return true, "filter title", nil
			}
		}

		return false, "", nil

	}

}
