package deleter_product

import (
	"errors"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type IterateFilter struct {
	CategoryID string
	Page       int
	PageSize   int
	Status     model.ProductStatus
}

func (f *IterateFilter) GetFilters() []model.Filter {
	pagestr := strconv.Itoa(f.Page)
	pagesizestr := strconv.Itoa(f.PageSize)
	queryFilter := []model.Filter{
		{
			ID:    "pageSize",
			Value: []string{pagesizestr},
		}, {
			ID:    "keyword",
			Value: []string{""},
		},
		{
			ID:    "page",
			Value: []string{pagestr},
		},
	}

	if f.Status != "" {
		queryFilter = append(queryFilter, model.Filter{
			ID:    "status",
			Value: []string{string(f.Status)},
		})
	}

	if f.CategoryID != "" {
		queryFilter = append(queryFilter, model.Filter{
			ID:    "category",
			Value: []string{f.CategoryID},
		})
	}

	return queryFilter
}

func IterateProduct(
	sellerapi *api.TokopediaApi,
	filter *IterateFilter,
	handleItem func(page int, product *model.SellerProductItem, delete func() int) error,
) error {

	shopId := sellerapi.AuthenticatedData.UserShopInfo.Info.ShopID
	meta, err := sellerapi.ProductListMeta()
	if err != nil {
		return err
	}

	username := sellerapi.AuthenticatedData.User.Email
	tab := meta.Data.ProductListMeta.Data.Tab.GetTab(filter.Status)

	if tab == nil {
		log.Println(username, "produk status tab tidak ditemukan")
		return nil
	}

	if tab.Value == 0 {
		log.Println(username, "delete selesai ...")
		return nil
	}

	countDelete := 0
	filpage := float64(tab.Value) / float64(filter.PageSize)
	filpage = math.Ceil(filpage)
	fixpage := int(filpage)

	for page := fixpage; page > 0; page-- {
		filter.Page = page
		deletedProducts := []*model.SellerProductItem{}

		log.Println(username, "request product page", page)
		filters := filter.GetFilters()

		query := model.NewProductListVar(shopId, filters)
		hasilList, err := sellerapi.ProductList(query)
		if err != nil {
			return err
		}

		if len(hasilList.Data.ProductList.Data) == 0 {
			break
		}

		for _, item := range hasilList.Data.ProductList.Data {
			err := handleItem(page, item, func() int {
				countDelete += 1
				deletedProducts = append(deletedProducts, item)
				return countDelete
			})

			if errors.Is(err, ErrDeleteLimitExcedeed) {
				err := DeleteProducts(sellerapi, deletedProducts)
				return err
			}

			if err != nil {
				return err
			}
		}

		err = DeleteProducts(sellerapi, deletedProducts)
		if err != nil {
			return err
		}

		log.Println(username, "sleep for 3 seconds")
		time.Sleep(time.Second * 3)
	}

	return nil
}
