package report

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

type CekReport struct {
	*tokopedia_lib.DriverAccount
	ShopName         string
	ProductActive    int
	ProductInActive  int
	ProductViolation int
	ShopScore        float32
	UreadChat        int
	NewOrder         int
	PmStatus         string
	ExtendStatus     string
	Status           string
}

func SaveCekReport(fname string, akuns []*CekReport) error {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("username,password,secret,shopname,product_active,product_inactive,product_violation,shop_score,unread_chat,new_order,pm_status,extend_status,status\n")
	for _, driver := range akuns {

		f.WriteString(fmt.Sprintf("%s,%s,%s,%s,%d,%d,%d,%.2f,%d,%d,%s,%s,%s\n",
			driver.Username,
			driver.Password,
			driver.Secret,
			driver.ShopName,
			driver.ProductActive,
			driver.ProductInActive,
			driver.ProductViolation,
			driver.ShopScore,
			driver.UreadChat,
			driver.NewOrder,
			driver.PmStatus,
			driver.ExtendStatus,
			driver.Status,
		))
	}
	return nil
}

func NewCekReport(fname string) (akuns []*CekReport, save func() error, err error) {
	hasil := []*CekReport{}
	data, _ := os.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

Parent:
	for _, line := range lines {
		if line == "" {
			continue Parent
		}

		if strings.HasPrefix(line, "username") {
			continue
		}

		dataline := make([]string, 13)

		fixline := strings.ReplaceAll(line, "\r", "")

		for ind, value := range strings.Split(fixline, ",") {
			dataline[ind] = value
		}

		acdriver, err := tokopedia_lib.NewDriverAccount(dataline[0], dataline[1], dataline[2])
		driver := CekReport{
			DriverAccount: acdriver,
		}

		if err != nil {
			if !errors.Is(err, tokopedia_lib.ErrSessionNotFound) {
				pdc_common.ReportError(err)
			}

		}

		hasil = append(hasil, &driver)
	}

	save = func() error {
		return SaveCekReport(fname, hasil)
	}

	return hasil, save, nil
}
