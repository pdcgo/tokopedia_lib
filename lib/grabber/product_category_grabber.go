package grabber

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Category struct {
	Id   int
	Name string
	Url  string
}

// Product Category
type ProductCategoryGrabber struct {
	Grabber
	Category Category
	params   *model_public.SearchProductVar
	adParams *model_public.SearchProductAdParams
}

func (grab *ProductCategoryGrabber) generateAdParams() *model_public.SearchProductAdParams {
	adParams := &model_public.SearchProductAdParams{
		Page:        1,
		NoAutoFill:  "5-14",
		Start:       1,
		Ep:          "product",
		Src:         "directory",
		Device:      "desktop",
		MinimumItem: 15,
		Item:        15,
		UserId:      0,
	}
	return adParams
}

func (grab *ProductCategoryGrabber) parseIdentifierFromCategoryUrl() (string, error) {
	u, err := url.Parse(grab.Category.Url)
	if err != nil {
		return "", err
	}
	path := u.EscapedPath()
	paths := strings.Split(path, "/")

	identifier := strings.Join(paths[2:], "_")
	fmt.Println(identifier, "identifier", paths, path)
	return identifier, nil
}

func (grab *ProductCategoryGrabber) RunProductGrabber() ([]model_public.CategoryProduct, error) {
	rawParams, err := json.Marshal(grab.params)
	if err != nil {
		return nil, err
	}
	stringParams := string(rawParams)
	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "")
	params := replacer.Replace(stringParams)

	rawAdParams, err := json.Marshal(grab.adParams)
	if err != nil {
		return nil, err
	}
	adParams := replacer.Replace(string(rawAdParams))

	variable := &model_public.SearchProductQueryVar{
		AdParams: adParams[1:len(adParams)-1] + fmt.Sprintf("&page=%d", grab.adParams.Page),
		Params:   params[1:len(params)-1] + fmt.Sprintf("&page=%d", grab.adParams.Page),
	}

	resp, err := grab.Api.SearchProductQuery(variable)
	if err != nil {
		return nil, err
	}
	products := resp.Data.CategoryProducts.Data
	return products, nil
}

func (grab *ProductCategoryGrabber) IteratePages(prodResp chan<- grab_handler.ProductCategoryGrabResp) {
	defer close(prodResp)

	for {
		fmt.Printf("########  mencari product dari kategori [ %s ] di halaman [ %d ]  ########\n", grab.Category.Name, grab.adParams.Page)
		products, err := grab.RunProductGrabber()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(products) == 0 {
			fmt.Printf("produk dari [ %s ] telah habis\n", grab.Category.Name)
			return
		}

		for _, product := range products {
			prodVar, _ := parseProductDetailParamsFromUrl(product.URL)
			shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
				Id:     product.Shop.ID,
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
				fmt.Printf("error mendapatkan produk [ %s ]", product.Name)
				fmt.Println(err)
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

			fmt.Printf("kategori [ %s ] : mendapatkan produk [ %s ]\n", grab.Category.Name, product.Name)
			res := grab_handler.ProductCategoryGrabResp{
				ProductCategory: product,
				ProductDetail:   *product_detail,
			}
			prodResp <- res
		}
		start := grab.params.Page * grab.params.Rows
		grab.params.Page += 1
		grab.params.Start = start + 1
		grab.adParams.Page += 1
		grab.adParams.Start = start + 1
	}
}

func (grab *ProductCategoryGrabber) Run(resp chan<- grab_handler.ProductCategoryGrabResp) {
	params := grab.generateProductSearchParams()
	params.CategoryId = grab.Category.Id
	params.Start = 1
	identifier, _ := grab.parseIdentifierFromCategoryUrl()
	params.Identifier = identifier

	adParams := grab.generateAdParams()
	adParams.DepId = grab.Category.Id

	grab.params = params
	grab.adParams = adParams

	go grab.IteratePages(resp)
}

func (grab *ProductCategoryGrabber) Save(namespace string, product *grab_handler.ProductCategoryGrabResp) {
	grab.CacheHandler.AddItemProductCategory(namespace, product)
}

func CreateProductCategoryGrabber(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	repo *mongorepo.ProductRepo,
	category Category) (*ProductCategoryGrabber, error) {

	productCategoryGrabber := &ProductCategoryGrabber{
		Grabber:  *CreateBaseGrabber(api, base, repo),
		Category: category,
	}
	return productCategoryGrabber, nil
}
