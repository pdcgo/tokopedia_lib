package grabber_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/grabber"
	"github.com/stretchr/testify/assert"
)

func TestProductGrabber(t *testing.T) {
	prodGrabber, err := grabber.NewProductListGrabber([]string{"mouse gaming", "mousepad"})
	assert.Nil(t, err)
	prodGrabber.RunProductListGrabber()
}

func TestProductCategoryGrabber(t *testing.T) {
	prodGrabber, err := grabber.CreateProductCategoryGrabber(2099, "Gaming", "https://www.tokopedia.com/p/gaming")
	assert.Nil(t, err)
	prodGrabber.RunProductCategoryGrabber()
}
