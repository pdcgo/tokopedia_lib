package main

import (
	"log"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func RunReqBrutal() error {

	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	headerData, err := pubapi.HeaderMainData()
	if err != nil {
		return err
	}

	return headerData.Data.CategoryAllListLite.Categories.Iterate(func(parents []*model_public.Categories, category *model_public.Categories) (stop bool, err error) {

		searchVar := model_public.NewSearchProductVar()
		searchVar.CategoryId = category.ID

		log.Println("request category ", category.Name)

		itemCount := searchVar.Rows
		currentCount := 0
		for currentCount < itemCount {

			variable := &model_public.ParamsVar{
				Params: searchVar.GetQuery(),
			}
			resp, err := pubapi.SearchProductQueryV4(variable)
			if err != nil {
				return false, err
			}

			products := resp.Data.AceSearchProductV4.Data.Products

			for _, p := range products {
				go func(url string) {
					query, err := model_public.NewPdpGetlayoutQueryVar(url)
					if err != nil {
						log.Println(err)
						return
					}

					playout, _ := pubapi.PdpGetlayoutQuery(query)
					if err != nil {
						log.Println(err)
						return
					}

					log.Println(playout.Data.PdpGetLayout.GetProductName())
				}(p.URL)
			}

			log.Println("page", searchVar.Page, ", products :", len(products))

			itemCount = resp.Data.AceSearchProductV4.Header.TotalData
			currentCount = searchVar.Rows * searchVar.Page

			start := searchVar.Page * searchVar.Rows
			searchVar.Page += 1
			searchVar.Start = start

			if itemCount == 0 {
				break
			}
		}

		return false, err
	})
}
