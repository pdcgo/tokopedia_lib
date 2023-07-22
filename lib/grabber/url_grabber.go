package grabber

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type UrlGrabber struct {
	Api    *api_public.TokopediaApiPublic
	Urls   []string
	Filter *filter.BaseFilter
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

func (grab *UrlGrabber) Run() ([]*grab_handler.UrlGrabberResp, error) {
	results := []*grab_handler.UrlGrabberResp{}
	for _, uri := range grab.Urls {
		params, err := parseProductDetailParamsFromUrl(uri)
		if err != nil {
			return nil, err
		}

		product, err := grab.Api.PdpGetlayoutQuery(params)
		if err != nil {
			return nil, err
		}
		if product.Data.PdpGetLayout.BasicInfo.Alias == "" {
			fmt.Printf("error [ produk ] : produk [ %s ] tidak mempunyai data yang lengkap\n", params.ProductKey)
			continue
		}
		shopId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ShopID)
		shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
			Id:     shopId,
			Domain: params.ShopDomain,
		})

		if shopFilter.ApplyFilter() {
			continue
		}

		payload := &model_public.PdpGetDataP2Var{
			PdpSession: product.Data.PdpGetLayout.PdpSession,
			ProductID:  product.Data.PdpGetLayout.BasicInfo.ID,
		}
		product_p2, err := grab.Api.PdpGetDataP2(payload)
		if err != nil {
			return nil, err
		}
		productFilter := filter.CreateProductDetailFilter(*grab.Filter, product.Data.PdpGetLayout, product_p2.Data.PdpGetData)
		if productFilter.ApplyFilter() {
			continue
		}

		fmt.Printf("grab [ url ] : mendapat produk [ %s ]\n", product.Data.PdpGetLayout.BasicInfo.Alias)
		result := &grab_handler.UrlGrabberResp{
			Product:   product,
			ProductP2: product_p2,
		}
		results = append(results, result)
	}

	return results, nil
}

func CreateUrlGrabber(urls []string) (*UrlGrabber, error) {
	base := legacy_source.BaseConfig{
		BaseData: "../..",
	}
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}

	urlGrabber := &UrlGrabber{
		Api:    api,
		Urls:   urls,
		Filter: filter.CreateBaseFilter(api, &base),
	}

	return urlGrabber, nil
}
