package upload_handler_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/upload_handler"
	"github.com/stretchr/testify/assert"
)

func TestParsingUrl(t *testing.T) {
	uri := "https://www.tokopedia.com/evoutdoor/hoodie-sweater-cowok-hi-comby-fleece-hoodie-korean-style-cream"
	data := upload_handler.NewPublicUrl(uri)
	t.Log(data.ShopCoreInfoVar)
	assert.NotEmpty(t, data)
}

func TestCloneHandler(t *testing.T) {
	uri := "https://www.tokopedia.com/evoutdoor/hoodie-sweater-cowok-hi-comby-fleece-hoodie-korean-style-cream"
	handlers := upload_handler.CreateCloneHandler(uri)

	t.Log(handlers)

}
