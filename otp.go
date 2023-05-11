package tokopedia_lib

import (
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func GetTotp(secret string) (string, error) {
	data := strings.ReplaceAll(secret, " ", "")
	return totp.GenerateCode(data, time.Now())
}
