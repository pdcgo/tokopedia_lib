package repo_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAgregateproduct(t *testing.T) {
	db := scenario.GetMongoDatabase(t)

	agg := repo.ProductAggregateIpml{
		Collection: db.Collection("item"),
	}

	agg.IterCategory("default", func(tokopediaID, count int, name []string) error {
		t.Log(name, tokopediaID, count)

		// assert.NotEmpty(t, tokopediaID)
		// assert.NotEqual(t, 0, tokopediaID)

		assert.NotEmpty(t, name)
		return nil
	})

}
