package grabber

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ShopGrabber struct {
	Api    *api_public.TokopediaApiPublic
	params *model_public.ShopProductVar
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

			// Implement Product Filter
			//
			//
			//

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

func (grab *ShopListGrabber) Run() error {
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

		// Implement Shop Filter
		//
		//
		//

		products := make(chan model_public.ShopProductData)
		go grab.RunShopGrabber(products)
		for product := range products {
			fmt.Println(product)
		}

		grab.params.Page = 1
	}
	return nil
}

func CreateShopListGrabber(shops []string, pathfile string) (*ShopListGrabber, error) {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}
	shopListGrabber := &ShopListGrabber{
		ShopGrabber: ShopGrabber{
			Api:    api,
			params: generateShopProductVar(),
		},
		Shops:    shops,
		Pathfile: pathfile,
	}
	return shopListGrabber, nil
}
