package grab_handler

import (
	"math/rand"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductParser struct {
	ShopId         int
	ProductId      int
	CatId          int
	ProductSold    int
	Stock          int
	ProductContent model_public.ProductContentComponent
	ProductDetail  model_public.ProductDetailComponent
	ProductMedia   model_public.MediaComponent
	ProductDesc    model_public.ProductDetailContent
	Images         []string
	Category       []int64
	Categories     []mongorepo.ProductCategory
}

func parseComponentsPDPProductLayout(product model_public.PdpGetlayoutQueryResp) ProductParser {
	shopId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ShopID)
	productId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.ID)
	catId, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.Category.ID)
	productSold, _ := strconv.Atoi(product.Data.PdpGetLayout.BasicInfo.TxStats.TransactionSuccess)

	productComponentParse := helper.ParseProductLayoutComponents(product.Data.PdpGetLayout.Components)
	productContent := productComponentParse.ProductContent
	productDetail := productComponentParse.ProductDetail
	productMedia := productComponentParse.ProductMedia

	stock, _ := strconv.Atoi(productContent.Data[0].Stock.Value)
	var productDesc model_public.ProductDetailContent
	for _, content := range productDetail.Data[0].Content {
		if content.Title == "Deskripsi" {
			productDesc = content
		}
	}

	var images = []string{}
	for _, media := range productMedia.Data[0].Media {
		if media.Type == "image" {
			images = append(images, media.URLOriginal)
		}
	}

	var category []int64
	var categories []mongorepo.ProductCategory
	for _, cat := range *product.Data.PdpGetLayout.BasicInfo.Category.Detail {
		Id, _ := strconv.Atoi(cat.ID)
		category = append(category, int64(Id))
		categories = append(categories, mongorepo.ProductCategory{
			Catid:       int64(Id),
			DisplayName: cat.Name,
		})
	}

	return ProductParser{
		ShopId:         shopId,
		ProductId:      productId,
		CatId:          catId,
		Stock:          stock,
		ProductSold:    productSold,
		ProductContent: *productContent,
		ProductDetail:  *productDetail,
		ProductMedia:   *productMedia,
		ProductDesc:    productDesc,
		Images:         images,
		Category:       category,
		Categories:     categories,
	}
}

func createCacheProduct(product model_public.PdpGetlayoutQueryResp) mongorepo.CacheProduct {
	productParser := parseComponentsPDPProductLayout(product)

	oriPrice := productParser.ProductContent.Data[0].Campaign.OriginalPrice
	if productParser.ProductContent.Data[0].Campaign.OriginalPrice == 0 {
		oriPrice = productParser.ProductContent.Data[0].Price.Value
	}

	res := mongorepo.CacheProduct{
		Shop: mongorepo.ProductShop{
			Shopid: int64(productParser.ShopId),
			// Location: product.ProductP2.Data.PdpGetData.ShopInfo.Location,
		},

		Marketplace: mongorepo.MP_TOKOPEDIA,
		Id:          int64(productParser.ProductId),
		Productid:   int64(productParser.ProductId),
		// Namespace:           namespace,
		Rnd:                 rand.Float64(),
		Name:                string(productParser.ProductContent.Data[0].Name),
		Price:               int64(oriPrice),
		PriceBeforeDiscount: int64(oriPrice),
		PriceAfterDiscount:  int64(productParser.ProductContent.Data[0].Price.Value),
		Image:               productParser.Images[0],
		Images:              productParser.Images,
		Sold:                int32(productParser.ProductSold),
		// ShopLocation:        product.ProductP2.Data.PdpGetData.ShopInfo.Location,
		Catid:      int32(productParser.CatId),
		CatName:    product.Data.PdpGetLayout.BasicInfo.Category.Name,
		Stock:      int32(productParser.Stock),
		Desc:       productParser.ProductDesc.Subtitle,
		Url:        product.Data.PdpGetLayout.BasicInfo.URL,
		Category:   productParser.Category,
		Categories: productParser.Categories,
	}
	return res
}

func CreateCacheProductUrl(namespace string, product *UrlGrabberResp) mongorepo.CacheProduct {
	cacheProduct := createCacheProduct(*product.Product)
	cacheProduct.Namespace = namespace
	cacheProduct.Shop.Location = product.ProductP2.Data.PdpGetData.ShopInfo.Location
	cacheProduct.ShopLocation = product.ProductP2.Data.PdpGetData.ShopInfo.Location

	return cacheProduct
}

func CreateCacheProductCategory(namespace string, product *ProductCategoryGrabResp) mongorepo.CacheProduct {
	cacheProduct := createCacheProduct(product.ProductDetail)
	cacheProduct.Namespace = namespace
	cacheProduct.Shop.Location = product.ProductCategory.Shop.Location
	cacheProduct.ShopLocation = product.ProductCategory.Shop.Location

	return cacheProduct
}

func CreateCacheProductShop(namespace string, product *ShopGrabberResp) mongorepo.CacheProduct {
	cacheProduct := createCacheProduct(product.Product)
	cacheProduct.Namespace = namespace
	cacheProduct.Shop.Location = product.Shop.Data.ShopInfoByID.Result[0].Location
	cacheProduct.ShopLocation = product.Shop.Data.ShopInfoByID.Result[0].Location

	return cacheProduct
}

func CreateCacheProductSearch(namespace string, product *ProductListGrabberResp) mongorepo.CacheProduct {
	cacheProduct := createCacheProduct(product.ProductDetail)
	cacheProduct.Namespace = namespace
	cacheProduct.Shop.Location = product.Product.Shop.City
	cacheProduct.ShopLocation = product.Product.Shop.City

	return cacheProduct
}
