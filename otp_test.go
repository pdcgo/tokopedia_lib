package tokopedia_lib_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestTOTP(t *testing.T) {
	data, err := tokopedia_lib.GetTotp("OGH7 THDG 3KAA ROWM MXEC 3USB U62Z NLWP")
	t.Log(data)
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
}
