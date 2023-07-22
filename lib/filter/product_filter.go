package filter

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

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

type ProductFilterModel struct {
	ProductId   int
	ProductName string
	ProductUrl  string
}

type ProductFilter struct {
	BaseFilter
	Product ProductFilterModel
}

func (filter *ProductFilter) getProductReviews() ([]model_public.Review, error) {
	variable := model_public.ProductReviewListVar{
		ProductID: fmt.Sprintf("%d", filter.Product.ProductId),
		Page:      1,
		Limit:     15,
		SortBy:    "create_time desc",
	}
	resp, err := filter.api.ProductReviewList(&variable)
	return resp.Data.ProductrevGetProductReviewList.List, err
}

func (filter *ProductFilter) getPDPGetLayout() (model_public.PdpGetLayout, error) {

	variable, err := parseProductDetailParamsFromUrl(filter.Product.ProductUrl)
	if err != nil {
		fmt.Printf("error [ produk ] : terjadi kesalahan produk [ %s ]\n", filter.Product.ProductName)
		return model_public.PdpGetLayout{}, err
	}
	resp, err := filter.api.PdpGetlayoutQuery(variable)
	return resp.Data.PdpGetLayout, err
}

func (filter *ProductFilter) getPDPGetDataP2() (model_public.PdpGetData, error) {
	prodLayout, err := filter.getPDPGetLayout()
	if err != nil {
		fmt.Printf("error [ produk ] : terjadi kesalahan produk [ %s ]\n", filter.Product.ProductName)
		return model_public.PdpGetData{}, err
	}
	variable := model_public.PdpGetDataP2Var{
		PdpSession: prodLayout.PdpSession,
		ProductID:  prodLayout.BasicInfo.ID,
		Affiliate:  nil,
	}
	resp, err := filter.api.PdpGetDataP2(&variable)
	return resp.Data.PdpGetData, err
}

func (filter *ProductFilter) FilterLastLogin() bool {
	if !filter.GrabTokopedia.LastLoginActive {
		return false
	}
	productP2, err := filter.getPDPGetDataP2()
	if err != nil {
		fmt.Printf("error [ produk ] : terjadi kesalahan produk [ %s ]\n", filter.Product.ProductName)
		return true
	}
	lastLogin, _ := strconv.Atoi(productP2.ShopInfo.ShopLastActive)
	result := filter.LastLogin(int64(lastLogin))
	if result {
		fmt.Printf("filter [ shop ] : toko produk [ %s ] terakhir aktif lebih lama dari filter\n", filter.Product.ProductName)
	}
	return result
}

func (filter *ProductFilter) FilterLayout() bool {
	if filter.GrabBasic.Stock == 0 || filter.GrabBasic.Penjualan == 0 || filter.GrabBasic.Prosentase == 0 {
		return false
	}
	productLayout, err := filter.getPDPGetLayout()
	if err != nil {
		fmt.Printf("error [ produk ] : terjadi kesalahan produk [ %s ]\n", filter.Product.ProductName)
		return true
	}
	productSold, _ := strconv.Atoi(productLayout.BasicInfo.TxStats.CountSold)
	productSuccessSold, _ := strconv.Atoi(productLayout.BasicInfo.TxStats.TransactionSuccess)
	soldPercentage := (float64(productSuccessSold) / float64(productSold)) * 100
	productLayoutParse := helper.ParseProductLayoutComponents(productLayout.Components)

	if filter.Sold(productSold) {
		fmt.Printf("filter [ produk ] : penjualan poduk [ %s ] kurang dari [ %d ]\n", filter.Product.ProductName, filter.GrabBasic.Penjualan)
		return true
	}
	if filter.SoldPercentage(int(soldPercentage)) {
		fmt.Printf("filter [ produk ] : persentase penjualan poduk [ %s ] kurang dari [ %d persen ]\n", filter.Product.ProductName, filter.GrabBasic.Prosentase)
		return true
	}

	oriStock := productLayoutParse.ProductContent.Data[0].Campaign.OriginalStock
	if oriStock == 0 {
		oriStock, _ = strconv.Atoi(productLayoutParse.ProductContent.Data[0].Stock.Value)
	}
	if filter.Stock(oriStock) {
		fmt.Printf("filter [ produk ] : stok poduk [ %s ] kurang dari [ %d ]\n", filter.Product.ProductName, filter.GrabBasic.Stock)
		return true
	}

	return false
}

func (filter *ProductFilter) FilterLastReview() bool {
	if !filter.GrabBasic.LastReviewActive {
		return false
	}
	productReviews, err := filter.getProductReviews()
	if err != nil {
		fmt.Printf("filter [ produk ] : terjadi kesalahan [ %s ]\n", filter.Product.ProductName)
		return true
	}

	if len(productReviews) == 0 {
		fmt.Printf("filter [ produk ] : tidak ada review untuk produk [ %s ]\n", filter.Product.ProductName)
		return false
	}
	lastProductReview, _ := strconv.Atoi(productReviews[0].ReviewCreateTime)
	result := filter.LastReview(int64(lastProductReview))
	if result {
		fmt.Printf("filter [ produk ] : review terakhir produk [ %s ] kurang dari filter\n", filter.Product.ProductName)
	}
	return result
}

func (filter *ProductFilter) ApplyFilter() bool {
	filters := []func() bool{
		filter.FilterLayout,
		filter.FilterLastReview,
		filter.FilterLastLogin,
	}
	for _, filter := range filters {
		if filter() {
			return true
		}
	}
	return false
}

func CreateProductFilter(base BaseFilter, product ProductFilterModel) *ProductFilter {
	return &ProductFilter{
		base,
		product,
	}
}

type ProductLayoutFilter struct {
	ProductFilter
	ProductLayout model_public.PdpGetLayout
}

func (filter *ProductLayoutFilter) getPDPGetLayout() (model_public.PdpGetLayout, error) {
	return filter.ProductLayout, nil
}

func CreateProductLayoutFilter(base BaseFilter, product model_public.PdpGetLayout) *ProductLayoutFilter {
	prodId, _ := strconv.Atoi(product.BasicInfo.ID)
	return &ProductLayoutFilter{
		ProductFilter: ProductFilter{
			BaseFilter: base,
			Product: ProductFilterModel{
				ProductId:   prodId,
				ProductName: product.BasicInfo.Alias,
				ProductUrl:  product.BasicInfo.URL,
			},
		},
		ProductLayout: product,
	}
}

type ProductDetailFilter struct {
	ProductLayoutFilter
	ProductP2 model_public.PdpGetData
}

func (filter *ProductDetailFilter) getPDPGetDataP2() (model_public.PdpGetData, error) {
	return filter.ProductP2, nil
}

func CreateProductDetailFilter(base BaseFilter, product model_public.PdpGetLayout, productP2 model_public.PdpGetData) *ProductDetailFilter {
	prodId, _ := strconv.Atoi(product.BasicInfo.ID)
	return &ProductDetailFilter{
		ProductLayoutFilter: ProductLayoutFilter{
			ProductFilter: ProductFilter{
				BaseFilter: base,
				Product: ProductFilterModel{
					ProductId:   prodId,
					ProductName: product.BasicInfo.Alias,
					ProductUrl:  product.BasicInfo.URL,
				},
			},
			ProductLayout: product,
		},
		ProductP2: productP2,
	}
}
