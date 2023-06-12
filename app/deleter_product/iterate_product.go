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

func IterateProduct(ctx context.Context, sellerapi *api.TokopediaApi) error {
	// errorHelp := NewErrorHandler(ctx)

	sID := sellerapi.AuthenticatedData.UserShopInfo.Info.ShopID
	shopId := strconv.Itoa(int(sID))

	page := 0

	for {
		page += 1
		log.Println(sellerapi.AuthenticatedData.User.Email, "request product page", page)
		pagestr := strconv.Itoa(page)

		query := model.ProductListVar{
			ShopID: shopId,
			Filter: []model.Filter{
				{
					ID:    "pageSize",
					Value: []string{"20"},
				}, {
					ID:    "keyword",
					Value: []string{""},
				}, {
					ID:    "status",
					Value: []string{},
				},
				{
					ID:    "page",
					Value: []string{pagestr},
				},
			},
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
	}

	return nil
}
