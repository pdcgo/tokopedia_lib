package grabber

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductGrabber struct {
	Api    *api_public.TokopediaApiPublic
	params *model_public.SearchProductVar
	Filter *filter.BaseFilter
}

func arrayConverter(datas []interface{}) []string {
	results := make([]string, len(datas))
	for _, data := range datas {
		switch value := data.(type) {
		case int:
			results = append(results, fmt.Sprint(value))
		case string:
			results = append(results, value)
		}
	}
	return results
}

func (grab *ProductGrabber) generateProductSearchParams() *model_public.SearchProductVar {
	locs := arrayConverter(grab.Filter.GrabTokopedia.Query.Fcity)
	shippings := arrayConverter(grab.Filter.GrabTokopedia.Query.Shipping)

	conditions := strings.Split(grab.Filter.GrabTokopedia.Query.Condition, ",")
	shopRating := strings.Split(grab.Filter.GrabTokopedia.Query.Rt, ",")

	shopTier := []string{}
	if grab.Filter.GrabTokopedia.Query.Official {
		shopTier = append(shopTier, "2")
	}
	if grab.Filter.GrabTokopedia.Query.Goldmerchant {
		shopTier = append(shopTier, "3")
	}

	params := model_public.SearchProductVar{
		Device:         "desktop",
		Sort:           grab.Filter.GrabTokopedia.Query.Ob,
		Page:           1,
		Rows:           100,
		UserDistrictID: "176",
		UserCityID:     "2274",
		Related:        true,
		Scheme:         "https",
		SafeSearch:     false,
		TopadsBucket:   true,
		Source:         "search",
		PriceMin:       grab.Filter.GrabTokopedia.Query.Pmin,
		PriceMax:       grab.Filter.GrabTokopedia.Query.Pmax,
		PreOrder:       grab.Filter.GrabTokopedia.Query.Preorder,
		Locations:      url.QueryEscape(strings.Join(locs, ",")),
		Rate:           url.QueryEscape(strings.Join(shopRating, "#")),
		Condition:      url.QueryEscape(strings.Join(conditions, "#")),
		Shipping:       url.QueryEscape(strings.Join(shippings, "#")),
		ShopTier:       url.QueryEscape(strings.Join(shopTier, "#")),
	}
	return &params
}

func (grab *ProductGrabber) RunProductGrabber() ([]model_public.ProductSearch, error) {
	if grab.params == nil {
		grab.params = grab.generateProductSearchParams()
	}

	rawParams, err := json.Marshal(grab.params)
	if err != nil {
		return nil, err
	}
	stringParams := string(rawParams)
	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "", "[", "", "]", "")
	validParams := replacer.Replace(stringParams)

	params := &model_public.ParamsVar{
		Params: validParams[1 : len(validParams)-1],
	}
	resp, err := grab.Api.SearchProductQueryV4(params)
	if err != nil {
		return nil, err
	}
	products := resp.Data.AceSearchProductV4.Data.Products
	if len(products) == 0 {
		fmt.Println("error [ error ] : produk telah habis")
		return nil, errors.New("produk kosong")
	}

	return products, nil
}

func (grab *ProductGrabber) IteratePages(prodResp chan<- model_public.ProductSearch) error {
	defer close(prodResp)

	for {

		products, err := grab.RunProductGrabber()
		if err != nil {
			fmt.Println(err)
			return err
		}
		for _, product := range products {
			prodResp <- product
		}
		grab.params.Page += 1
		grab.params.Start = grab.params.Page * grab.params.Rows
	}
}

// Product List
type ProductListGrabber struct {
	ProductGrabber
	Keywords []string
	wg       *sync.WaitGroup
}

func (grab *ProductListGrabber) Run(prodResp chan<- grab_handler.ProductListGrabberResp) {
	defer close(prodResp)

	for _, keyword := range grab.Keywords {
		grab.wg.Add(1)

		go func(keyword string, prodResp chan<- grab_handler.ProductListGrabberResp) {
			defer grab.wg.Done()

			fmt.Printf("keyword [ %s ] : memulai grab keyword [ %s ]\n", keyword, keyword)
			grab.params = grab.generateProductSearchParams()
			grab.params.Query = url.QueryEscape(keyword)

			products := make(chan model_public.ProductSearch)
			go grab.IteratePages(products)
			for product := range products {
				prodVar, _ := parseProductDetailParamsFromUrl(product.URL)
				shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
					Id:     product.Shop.ShopID,
					Domain: prodVar.ShopDomain,
				})

				if shopFilter.ApplyFilter() {
					continue
				}

				variable := &model_public.PdpGetlayoutQueryVar{
					ShopDomain: prodVar.ShopDomain,
					ProductKey: prodVar.ProductKey,
					APIVersion: 1,
				}
				product_detail, err := grab.Api.PdpGetlayoutQuery(variable)
				if err != nil {
					fmt.Printf("error [ produk ] : error  mendapatkan produk [ %s ]\n", product.Name)
					return
				}

				if product_detail.Data.PdpGetLayout.BasicInfo.Alias == "" {
					fmt.Printf("error [ produk ] : produk [ %s ] tidak mempunyai data yang lengkap\n", product.Name)
					continue
				}

				productFilter := filter.CreateProductLayoutFilter(*grab.Filter, product_detail.Data.PdpGetLayout)
				if productFilter.ApplyFilter() {
					continue
				}

				fmt.Printf("grab [ %s ] : mendapatkan produk [ %s ]\n", grab.params.Query, product.Name)
				res := grab_handler.ProductListGrabberResp{
					Product:       product,
					ProductDetail: *product_detail,
				}
				prodResp <- res
			}

		}(keyword, prodResp)

		grab.wg.Wait()
	}

}

func NewProductListGrabber(keywords []string) (*ProductListGrabber, error) {
	base := legacy_source.BaseConfig{
		BaseData: "../..",
	}
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}

	productListGrabber := &ProductListGrabber{
		ProductGrabber: ProductGrabber{
			Api:    api,
			Filter: filter.CreateBaseFilter(api, &base),
		},
		Keywords: keywords,
		wg:       &sync.WaitGroup{},
	}
	return productListGrabber, nil
}
