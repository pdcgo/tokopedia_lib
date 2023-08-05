package iterator

import (
	"context"
	"net/url"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type BatchLayoutHandler func(layout *model_public.PdpGetlayoutQueryResp) error

func ParseProductDetailParamsFromUrl(uri string) (*model_public.PdpGetlayoutQueryVar, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	path := u.EscapedPath()
	query := u.Query()

	splitPath := strings.Split(path, "/")
	shopDomain := splitPath[len(splitPath)-2]
	productKey := splitPath[len(splitPath)-1]

	payload := &model_public.PdpGetlayoutQueryVar{
		ShopDomain: shopDomain,
		ProductKey: productKey,
		APIVersion: 1,
		ExtParam:   url.QueryEscape(query.Get("extParam")),
	}
	return payload, nil
}

func IterateBatchLayout(
	api *api_public.TokopediaApiPublic,
	ctx context.Context,
	urls []string, handler BatchLayoutHandler,
) error {

	var layoutVars []*model_public.PdpGetlayoutQueryVar
Urls:
	for _, url := range urls {
		select {
		case <-ctx.Done():
			break Urls
		default:
			layoutVar, err := ParseProductDetailParamsFromUrl(url)
			if err != nil {
				return err
			}
			layoutVars = append(layoutVars, layoutVar)
		}
	}

	resp, err := api.PdpGetlayoutQueryBatch(layoutVars)
	if err != nil {
		pdc_common.ReportError(err)
		return err
	}

Resp:
	for _, resp := range resp {
		select {
		case <-ctx.Done():
			break Resp
		default:
			err := handler(resp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
