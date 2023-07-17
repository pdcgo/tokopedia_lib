package grabber

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductGrabber struct {
	Api    *api_public.TokopediaApiPublic
	params *model_public.SearchProductVar
	// Config
}

func (grab *ProductGrabber) generateProductSearchParams() *model_public.SearchProductVar {
	params := model_public.SearchProductVar{
		Device:         "desktop",
		Sort:           23,
		Page:           1,
		Rows:           100,
		UserDistrictID: "176",
		UserCityID:     "2274",
		Related:        true,
		Scheme:         "https",
		SafeSearch:     false,
		TopadsBucket:   true,
		Source:         "search",
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
	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "")
	validParams := replacer.Replace(stringParams)

	fmt.Println(validParams[1 : len(validParams)-1])
	params := &model_public.ParamsVar{
		Params: validParams[1 : len(validParams)-1],
	}
	resp, err := grab.Api.SearchProductQueryV4(params)
	if err != nil {
		return nil, err
	}
	products := resp.Data.AceSearchProductV4.Data.Products
	if len(products) == 0 {
		return nil, errors.New("produk kosong")
	}

	return products, nil
}

// func (grab *ProductGrabber) filterProduct(product model_public.ProductSearch) (string, error) {
// 	shopFilter, err := filter.CreateShopProductSearchFilter(product)
// 	if err != nil {
// 		return "", err
// 	}

// 	filtered, err := shopFilter.ApplyFilter()
// 	if err != nil {
// 		return "", err
// 	}
// 	if filtered {
// 		return shopFilter.CreateMessage(), nil
// 	}

// 	return "", nil
// }

func (grab *ProductGrabber) IteratePages(prodResp chan<- model_public.ProductSearch) error {
	defer close(prodResp)

	for {

		products, err := grab.RunProductGrabber()
		if err != nil {
			fmt.Println(err, "Error gais")
			return err
		}
		for _, product := range products {
			// Implement Filter
			//
			//
			//
			//

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
}

func (grab *ProductListGrabber) RunProductListGrabber() {

	for _, keyword := range grab.Keywords {
		grab.params = grab.generateProductSearchParams()
		grab.params.Query = url.QueryEscape(keyword)

		products := make(chan model_public.ProductSearch)
		go grab.IteratePages(products)
		for product := range products {
			fmt.Println(product)
		}
	}

}

func NewProductListGrabber(keywords []string) (*ProductListGrabber, error) {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}
	productListGrabber := &ProductListGrabber{
		ProductGrabber: ProductGrabber{
			Api: api,
		},
		Keywords: keywords,
	}
	return productListGrabber, nil
}

// Product Category
type ProductCategoryGrabber struct {
	ProductGrabber
	CategoryId   int
	CategoryName string
	CategoryUrl  string
	params       *model_public.SearchProductVar
	adParams     *model_public.SearchProductAdParams
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
	u, err := url.Parse(grab.CategoryUrl)
	if err != nil {
		return "", err
	}
	path := u.EscapedPath()
	paths := strings.Split(path, "/")

	identifier := strings.Join(paths[2:], "_")
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
		AdParams: adParams[1:len(adParams)-1] + fmt.Sprintf("&page=%s", strconv.Itoa(grab.adParams.Page)),
		Params:   params[1:len(params)-1] + fmt.Sprintf("&page=%s", strconv.Itoa(grab.adParams.Page)),
	}

	resp, err := grab.Api.SearchProductQuery(variable)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Data.CategoryProducts.Count, "counting")
	products := resp.Data.CategoryProducts.Data
	if len(products) == 0 {
		return nil, errors.New("produk kosong")
	}

	return products, nil
}

func (grab *ProductCategoryGrabber) IteratePages(prodResp chan<- model_public.CategoryProduct) error {
	defer close(prodResp)

	for {
		products, err := grab.RunProductGrabber()
		if err != nil {
			fmt.Println(err, "Error gais")
			return err
		}
		for _, product := range products {
			// Implement Filter
			//
			//
			//
			//

			prodResp <- product
		}
		start := grab.params.Page * grab.params.Rows
		grab.params.Page += 1
		grab.params.Start = start + 1
		grab.adParams.Page += 1
		grab.adParams.Start = start + 1

	}
}

func (grab *ProductCategoryGrabber) RunProductCategoryGrabber() {
	params := grab.generateProductSearchParams()
	params.CategoryId = grab.CategoryId
	params.Start = 1
	identifier, _ := grab.parseIdentifierFromCategoryUrl()
	params.Identifier = identifier

	adParams := grab.generateAdParams()
	adParams.DepId = grab.CategoryId

	grab.params = params
	grab.adParams = adParams

	products := make(chan model_public.CategoryProduct)
	go grab.IteratePages(products)
	count := 0
	for product := range products {
		fmt.Println(product)
		count += 1
	}
	fmt.Println(count, "dapat ini")
}

func CreateProductCategoryGrabber(categoryId int, categoryName string, categoryUrl string) (*ProductCategoryGrabber, error) {
	api, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return nil, err
	}
	productCategoryGrabber := &ProductCategoryGrabber{
		ProductGrabber: ProductGrabber{
			Api: api,
		},
		CategoryId:   categoryId,
		CategoryName: categoryName,
		CategoryUrl:  categoryUrl,
	}
	return productCategoryGrabber, nil
}
