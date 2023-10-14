package benchmark_test

import (
	"context"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/core_concept"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCategorypage(b *testing.B) {
	// b.Skip("skip tidak perlu")

	pubapi, err := api_public.NewTokopediaApiPublic()
	if err != nil {
		b.Error(err)
	}

	headerData, err := pubapi.HeaderMainData()
	assert.Nil(b, err)

	categories := headerData.Data.CategoryAllListLite.Categories

	uri := "https://www.tokopedia.com/p/fashion-muslim/atasan-muslim-wanita"
	category, _ := categories.GetCategoryByUrl(uri)
	searchVar := model_public.NewSearchProductVar()
	searchVar.CategoryId = category.ID

	ctx := core_concept.NewTaskContext(context.Background())
	defer ctx.Cancel()
	chunks, err := iterator.V2IterateSearchPage(ctx, 10, pubapi, searchVar)
	assert.Nil(b, err)

	b.Run("bench per chunk page", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			<-chunks
		}

	})
	// mulai iterasi per page
	total := 0
	for chunk := range chunks {
		total += len(chunk)
	}
	b.Log(total)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	layouts, err := iterator.V2GetBatchLayout(chunks, 1000, 5, ctx, pubapi)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	count := 0
	// Parent:
	// 	for layout := range layouts {
	// 		select {
	// 		case <-ctx.Ctx.Done():
	// 			break Parent
	// 		default:
	// 			count++
	// 			name, _ := layout.Data.PdpGetLayout.GetProductName()
	// 			log.Println(count, name)
	// 		}

	// 	}

}
