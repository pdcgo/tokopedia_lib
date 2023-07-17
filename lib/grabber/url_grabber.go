package grabber

import (
	"net/url"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type UrlGrabber struct {
	Api  *api_public.TokopediaApiPublic
	Urls []string
}

func parseProductDetailParamsFromUrl(uri string) (*model_public.PdpGetlayoutQueryVar, error) {
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

type UrlGrabberResp struct {
	Product   *model_public.PdpGetlayoutQueryResp
	ProductP2 *model_public.PdpGetDataP2Resp
}

func (grab *UrlGrabber) Run() ([]*UrlGrabberResp, error) {
	results := []*UrlGrabberResp{}
	for _, uri := range grab.Urls {
		params, err := parseProductDetailParamsFromUrl(uri)
		if err != nil {
			return nil, err
		}

		product, err := grab.Api.PdpGetlayoutQuery(params)
		if err != nil {
			return nil, err
		}

		payload := &model_public.PdpGetDataP2Var{
			PdpSession: product.Data.PdpGetLayout.PdpSession,
			ProductID:  product.Data.PdpGetLayout.BasicInfo.ID,
		}
		product_p2, err := grab.Api.PdpGetDataP2(payload)
		if err != nil {
			return nil, err
		}

		// Implement Filter
		//
		//
		//
		//

		result := &UrlGrabberResp{
			Product:   product,
			ProductP2: product_p2,
		}
		results = append(results, result)
	}

	return results, nil
}

func CreateUrlGrabber(urls []string) (*UrlGrabber, error) {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}
	urlGrabber := &UrlGrabber{
		Api:  api,
		Urls: urls,
	}

	return urlGrabber, nil
}
