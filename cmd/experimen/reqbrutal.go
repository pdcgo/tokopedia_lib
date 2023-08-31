package main

import (
	"log"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

func RunReqBrutal() error {

	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		return err
	}

	// batchlayout, err := api_public.NewTokopediaApiPublic()
	// if err != nil {
	// 	return err
	// }

	headerData, err := pubapi.HeaderMainData()
	if err != nil {
		return err
	}

	return headerData.Data.CategoryAllListLite.Categories.Iterate(func(parents []*model_public.Categories, category *model_public.Categories) (stop bool, err error) {

		searchVar := model_public.NewSearchProductVar()
		searchVar.CategoryId = category.ID

		log.Println("request category ", category.Name)

		ctx := iterator.NewContextError()

		// mulai iterasi per page
		chunks, err := iterator.V2IterateSearchPage(ctx, 10, pubapi, searchVar)

		if err != nil {
			panic(err)
		}

		layouts, err := iterator.V2GetBatchLayout(chunks, 1000, 10, ctx, pubapi)

		if err != nil {
			panic(err)
		}

		count := 0
	Parent:
		for layout := range layouts {
			select {
			case <-ctx.Ctx.Done():
				break Parent
			default:
				count++
				name, _ := layout.Data.PdpGetLayout.GetProductName()
				log.Println(count, name)
			}

		}

		log.Println(count)
		panic("asdasdasd")

		return false, err
	})
}
