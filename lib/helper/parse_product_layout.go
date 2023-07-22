package helper

import "github.com/pdcgo/tokopedia_lib/lib/model_public"

type ProductlayoutParse struct {
	ProductContent *model_public.ProductContentComponent
	ProductDetail  *model_public.ProductDetailComponent
	ProductMedia   *model_public.MediaComponent
}

func ParseProductLayoutComponents(components model_public.PDPListComponents) *ProductlayoutParse {
	var productContent *model_public.ProductContentComponent
	var productDetail *model_public.ProductDetailComponent
	var productMedia *model_public.MediaComponent
	for _, component := range components {
		switch value := component.(type) {
		case *model_public.MediaComponent:
			productMedia = value
		case *model_public.ProductContentComponent:
			productContent = value
		case *model_public.ProductDetailComponent:
			productDetail = value
		default:
			continue
		}
	}

	return &ProductlayoutParse{
		ProductContent: productContent,
		ProductDetail:  productDetail,
		ProductMedia:   productMedia,
	}
}
