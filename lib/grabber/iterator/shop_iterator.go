package iterator

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ShopHandler func(shopCore *model_public.ShopCoreInfoResp) error

func generateShopCoreInfoParamsFormUrl(uri string) (*model_public.ShopCoreInfoVar, error) {
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

func IterateShops(api *api_public.TokopediaApiPublic, fname string, handler ShopHandler) error {
	shops, err := helper.FileLoadLineString(fname)
	if err != nil {

		if os.IsNotExist(err) {
			log.Printf("[ warning ] file %s not found", fname)
			return nil
		}

		pdc_common.ReportError(err)
		return err
	}

	setShop := func(shops []string) error {
		return helper.FileSaveLineString(fname, shops)
	}

	for index, shop := range shops {

		variable, err := generateShopCoreInfoParamsFormUrl(shop)
		if err != nil {
			pdc_common.ReportError(err)
		}
		shopCoreInfo, err := api.ShopCoreInfo(variable)
		if err != nil {
			pdc_common.ReportError(err)
		}

		err = handler(shopCoreInfo)
		if err != nil {
			pdc_common.ReportError(err)
		}

		err = setShop(shops[index+1:])
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	return nil
}
