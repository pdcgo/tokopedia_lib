package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

type CekVerifReport struct {
	*tokopedia_lib.DriverAccount
	Pesan  string
	Status string
}

func SaveCekVerifReport(fname string, akuns []*CekVerifReport) error {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("username,password,secret,status,pesan\n")
	for _, driver := range akuns {
		f.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s\n", driver.Username, driver.Password, driver.Secret, driver.Status, driver.Pesan))
	}
	return nil
}

func NewCekVerifReport(fname string) (akuns []*CekVerifReport, save func() error, err error) {
	hasil := []*CekVerifReport{}
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

		dataline := make([]string, 5)

		fixline := strings.ReplaceAll(line, "\r", "")

		for ind, value := range strings.Split(fixline, ",") {
			dataline[ind] = value
		}

		acdriver, err := tokopedia_lib.NewDriverAccount(dataline[0], dataline[1], dataline[2])
		driver := CekVerifReport{
			DriverAccount: acdriver,
			Status:        dataline[3],
			Pesan:         dataline[4],
		}

		if err != nil {
			pdc_common.ReportError(err)
		}

		hasil = append(hasil, &driver)
	}

	save = func() error {
		return SaveCekVerifReport(fname, hasil)
	}

	return hasil, save, nil
}
