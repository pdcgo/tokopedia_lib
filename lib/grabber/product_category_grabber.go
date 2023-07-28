package grabber

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/url"
// 	"strings"

// 	"github.com/pdcgo/go_v2_shopeelib/helper"
// 	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
// 	"github.com/pdcgo/tokopedia_lib/lib/api_public"
// 	"github.com/pdcgo/tokopedia_lib/lib/filter"
// 	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
// 	"github.com/pdcgo/tokopedia_lib/lib/model_public"
// )

// type Category struct {
// 	Id   int
// 	Name string
// 	Url  string
// }

// // Product Category
// type ProductCategoryGrabber struct {
// 	Grabber
// 	Category Category
// }

// func (grab *ProductCategoryGrabber) generateAdParams() *model_public.SearchProductAdParams {
// 	adParams := &model_public.SearchProductAdParams{
// 		Page:        1,
// 		NoAutoFill:  "5-14",
// 		Start:       1,
// 		Ep:          "product",
// 		Src:         "directory",
// 		Device:      "desktop",
// 		MinimumItem: 15,
// 		Item:        15,
// 		UserId:      0,
// 	}
// 	return adParams
// }

// func (grab *ProductCategoryGrabber) parseIdentifierFromCategoryUrl() (string, error) {
// 	u, err := url.Parse(grab.Category.Url)
// 	if err != nil {
// 		return "", err
// 	}
// 	path := u.EscapedPath()
// 	paths := strings.Split(path, "/")

// 	identifier := strings.Join(paths[2:], "_")
// 	fmt.Println(identifier, "identifier", paths, path)
// 	return identifier, nil
// }

// func (grab *ProductCategoryGrabber) getSearchProductQueryVar(
// 	params *model_public.SearchProductVar,
// 	adParams *model_public.SearchProductAdParams) (*model_public.SearchProductQueryVar, error) {
// 	rawParams, err := json.Marshal(params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stringParams := string(rawParams)
// 	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "")
// 	strParams := replacer.Replace(stringParams)

// 	rawAdParams, err := json.Marshal(adParams)
// 	if err != nil {
// 		return nil, err
// 	}
// 	strAdParams := replacer.Replace(string(rawAdParams))

// 	variable := &model_public.SearchProductQueryVar{
// 		AdParams: strAdParams[1:len(strAdParams)-1] + fmt.Sprintf("&page=%d", adParams.Page),
// 		Params:   strParams[1:len(strParams)-1] + fmt.Sprintf("&page=%d", params.Page),
// 	}
// 	return variable, nil
// }

// func (grab *ProductCategoryGrabber) RunProductGrabber(
// 	params *model_public.SearchProductVar,
// 	adParams *model_public.SearchProductAdParams,
// ) ([]model_public.CategoryProduct, error) {
// 	variable, err := grab.getSearchProductQueryVar(params, adParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := grab.Api.SearchProductQuery(variable)
// 	if err != nil {
// 		return nil, err
// 	}
// 	products := resp.Data.CategoryProducts.Data
// 	return products, nil
// }

// func (grab *ProductCategoryGrabber) IteratePages(
// 	ctx context.Context,
// 	params *model_public.SearchProductVar,
// 	adParams *model_public.SearchProductAdParams,
// ) (<-chan grab_handler.ProductCategoryGrabResp, *helper.ChannelError) {

// 	resProducts := make(chan grab_handler.ProductCategoryGrabResp)
// 	errChan := helper.NewChannelError()
// 	fmt.Printf("grab [ kategori ] : memulai grab kategori [ %s ]\n", grab.Category.Name)

// 	go func() {
// 		defer close(resProducts)
// 		defer errChan.Raise()

// 	Parent:
// 		for {
// 			fmt.Printf("grab [ kategori ] : mencari kategori [ %s ] dihalaman [ %d ]\n", grab.Category.Name, params.Page)
// 			products, err := grab.RunProductGrabber(params, adParams)
// 			if err != nil {
// 				errChan.SetError(err)
// 				break Parent
// 			}

// 			for _, product := range products {
// 				prodVar, _ := ParseProductDetailParamsFromUrl(product.URL)
// 				shopFilter := filter.CreateShopFilter(*grab.Filter, filter.Shop{
// 					Id:     product.Shop.ID,
// 					Domain: prodVar.ShopDomain,
// 				})
// 				if shopFilter.ApplyFilter() {
// 					continue
// 				}

// 				variable := &model_public.PdpGetlayoutQueryVar{
// 					ShopDomain: prodVar.ShopDomain,
// 					ProductKey: prodVar.ProductKey,
// 					APIVersion: 1,
// 				}
// 				product_detail, err := grab.Api.PdpGetlayoutQuery(variable)
// 				if err != nil {
// 					fmt.Printf("error mendapatkan produk [ %s ]", product.Name)
// 					fmt.Println(err)
// 					continue
// 				}

// 				if product_detail.Data.PdpGetLayout.BasicInfo.Alias == "" {
// 					fmt.Printf("error [ produk ] : produk [ %s ] tidak mempunyai data yang lengkap\n", product.Name)
// 					continue
// 				}
// 				productFilter := filter.CreateProductLayoutFilter(*grab.Filter, product_detail.Data.PdpGetLayout)
// 				if productFilter.ApplyFilter() {
// 					continue
// 				}

// 				res := grab_handler.ProductCategoryGrabResp{
// 					ProductCategory: &product,
// 					ProductDetail:   product_detail,
// 				}
// 				resProducts <- res

// 				start := params.Page * params.Rows
// 				params.Page += 1
// 				params.Start = start + 1
// 				adParams.Page += 1
// 				adParams.Start = start + 1
// 			}
// 		}
// 	}()
// 	return resProducts, errChan
// }

// func (grab *ProductCategoryGrabber) Run() {

// 	params := grab.GenerateProductSearchParams()
// 	params.CategoryId = grab.Category.Id
// 	params.Start = 1
// 	identifier, _ := grab.parseIdentifierFromCategoryUrl()
// 	params.Identifier = identifier

// 	adParams := grab.generateAdParams()
// 	adParams.DepId = grab.Category.Id

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	limit := int32(grab.Filter.GrabBasic.LimitGrab)
// 	limiter := helper.NewLimiter(limit)

// 	products, errChan := grab.IteratePages(ctx, params, adParams)
// 	for product := range products {
// 		if limiter.LimitReached() {
// 			fmt.Printf("filter [ limit ] : grab ketegori [ %s ] telah mencapai limit\n", grab.Category.Name)
// 			cancel()
// 			return
// 		}
// 		// resp <- product
// 		grab.Save("", &product)
// 		fmt.Printf("grab [ kategori ] : mendapatkan produk [ %s ] dari kategori [ %s ]\n", product.ProductCategory.Name, grab.Category.Name)

// 		limiter.Add()
// 	}

// 	err := errChan.GetError()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// func (grab *ProductCategoryGrabber) Save(namespace string, product *grab_handler.ProductCategoryGrabResp) error {
// 	return grab.CacheHandler.AddItemProductCategory(namespace, product)
// }

// func CreateProductCategoryGrabber(
// 	api *api_public.TokopediaApiPublic,
// 	filter *filter.BaseFilter,
// 	repo *mongorepo.ProductRepo,
// 	category Category,
// ) *ProductCategoryGrabber {

// 	cacheHandler := grab_handler.NewCacheProductHandler(repo)
// 	productCategoryGrabber := &ProductCategoryGrabber{
// 		Grabber: Grabber{
// 			Api:          api,
// 			Filter:       filter,
// 			CacheHandler: cacheHandler,
// 		},
// 		Category: category,
// 	}
// 	return productCategoryGrabber
// }
