package grabber

import (
	"errors"
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

type ShopGrabber struct {
	Api    *api_public.TokopediaApiPublic
	params *model_public.ShopProductVar
	Filter *filter.BaseFilter
}

func generateShopProductVar() *model_public.ShopProductVar {
	params := &model_public.ShopProductVar{
		Page:           1,
		PerPage:        100,
		EtalaseID:      "etalase",
		Sort:           1,
		Sid:            "",
		UserDistrictID: "176",
		UserCityID:     "2274",
		UserLat:        "",
		UserLong:       "",
	}
	return params
}

func generateShopCoreInfoParams(uri string) (*model_public.ShopCoreInfoVar, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	params := &model_public.ShopCoreInfoVar{
		ID:     0,
		Domain: strings.Replace(u.Path, "/", "", -1),
	}
	return params, nil
}

func (grab *ShopGrabber) RunShopGrabber(prodResp chan<- model_public.ShopProductData) error {
	defer close(prodResp)

	hasMore := true
	for hasMore {
		resp, err := grab.Api.ShopProducts(grab.params)
		if err != nil {
			return err
		}

		products := resp.Data.GetShopProduct.Data
		if len(products) == 0 {
			return errors.New("halaman kosong")
		}

		for _, product := range products {
			prodResp <- product
		}
		if resp.Data.GetShopProduct.Links.Next == "" {
			hasMore = false
		}
		grab.params.Page += 1
	}
	return nil
}

type ShopListGrabber struct {
	ShopGrabber
	Shops    []string
	Pathfile string
}

func (grab *ShopListGrabber) Run(prodResp chan<- grab_handler.ShopGrabberResp) error {
	if grab.params == nil {
		grab.params = generateShopProductVar()
	}

	for _, domain := range grab.Shops {
		variable, err := generateShopCoreInfoParams(domain)
		if err != nil {
			return err
		}
		shopCoreInfo, err := grab.Api.ShopCoreInfo(variable)
		if err != nil {
			return err
		}
		grab.params.Sid = shopCoreInfo.Data.ShopInfoByID.Result[0].ShopCore.ShopID

		shopId, _ := strconv.Atoi(grab.params.Sid)
		shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
			Id:     shopId,
			Domain: variable.Domain,
		})

		if shopFilter.ApplyFilter() {
			continue
		}

		products := make(chan model_public.ShopProductData)
		go grab.RunShopGrabber(products)
		for product := range products {
			prodVar, _ := parseProductDetailParamsFromUrl(product.ProductURL)
			variable := &model_public.PdpGetlayoutQueryVar{
				ShopDomain: variable.Domain,
				ProductKey: prodVar.ProductKey,
				APIVersion: 1,
			}
			product_detail, err := grab.Api.PdpGetlayoutQuery(variable)
			if err != nil {
				fmt.Printf("error [ produk ] : error  mendapatkan produk [ %s ]\n", product.Name)
				continue
			}

			if product_detail.Data.PdpGetLayout.BasicInfo.Alias == "" {
				fmt.Printf("error [ produk ] : produk [ %s ] tidak mempunyai data yang lengkap\n", product.Name)
				continue
			}

			productFilter := filter.CreateProductLayoutFilter(*grab.Filter, product_detail.Data.PdpGetLayout)
			if productFilter.ApplyFilter() {
				continue
			}
			res := grab_handler.ShopGrabberResp{
				Shop:    *shopCoreInfo,
				Product: *product_detail,
			}
			prodResp <- res

		}

		grab.params.Page = 1
	}
	return nil
}

func CreateShopListGrabber(shops []string, pathfile string) (*ShopListGrabber, error) {
	base := legacy_source.BaseConfig{
		BaseData: "../..",
	}
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}
	shopListGrabber := &ShopListGrabber{
		ShopGrabber: ShopGrabber{
			Api:    api,
			params: generateShopProductVar(),
			Filter: filter.CreateBaseFilter(api, &base),
		},
		Shops:    shops,
		Pathfile: pathfile,
	}
	return shopListGrabber, nil
}
