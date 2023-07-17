package deleter_product

import (
	"context"
	"log"
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type ErrorHandler struct {
	err chan error
	ctx context.Context
}

func NewErrorHandler(ctx context.Context) *ErrorHandler {
	return &ErrorHandler{
		err: make(chan error, 1),
		ctx: ctx,
	}
}

func (erh *ErrorHandler) Error() error {
	return <-erh.err
}

func (erh *ErrorHandler) SetError(err error) {
	if err == nil {
		return
	}

	select {
	case erh.err <- err:
		return
	case <-erh.ctx.Done():
		return
	}
}

func deleteProducts(sapi *api.TokopediaApi, datas []*model.SellerProductItem) error {
	deletes := make([]*model.BulkProductEditV3Input, len(datas))

	for ind, data := range datas {
		deletes[ind] = &model.BulkProductEditV3Input{
			ProductID: data.ID,
			Status:    model.DeletedStatus,
			Shop: model.BulkProductEditShop{
				ID: data.Shop.ID,
			},
		}
	}

	payload := model.BulkProductEditV3Var{
		Input: deletes,
	}

	hasils, err := sapi.BulkProductEditV3(&payload)
	if err != nil {
		return err
	}

	for _, hasil := range hasils.Data.BulkProductEditV3 {
		if !hasil.Result.IsSuccess {
			return hasil
		}
	}

	return nil
}

func IterateProduct(sellerapi *api.TokopediaApi, handleItem func(page int, product *model.SellerProductItem, delete func() int) error, filters ...model.Filter) error {
	// errorHelp := NewErrorHandler(ctx)

	sID := sellerapi.AuthenticatedData.UserShopInfo.Info.ShopID
	shopId := strconv.Itoa(int(sID))

	page := 0
	countDelete := 0
	for {
		page += 1
		deletedProducts := []*model.SellerProductItem{}

		log.Println(sellerapi.AuthenticatedData.User.Email, "request product page", page)
		pagestr := strconv.Itoa(page)

		queryFilter := []model.Filter{
			{
				ID:    "pageSize",
				Value: []string{"20"},
			}, {
				ID:    "keyword",
				Value: []string{""},
			},
			{
				ID:    "page",
				Value: []string{pagestr},
			},
		}

		queryFilter = append(queryFilter, filters...)

		query := model.ProductListVar{
			ShopID: shopId,
			Filter: queryFilter,
			Sort: model.Sort{
				ID:    "DEFAULT",
				Value: "DESC",
			},
			ExtraInfo:   []string{"view", "topads", "rbac", "price-suggestion"},
			WarehouseID: "",
		}

		hasilList, err := sellerapi.ProductList(&query)
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
			if err != nil {
				return err
			}
		}

		err = deleteProducts(sellerapi, deletedProducts)
		if err != nil {
			return err
		}
	}

	return nil
}
