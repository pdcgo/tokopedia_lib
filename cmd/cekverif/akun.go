package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

type VerifDriverAccount struct {
	*tokopedia_lib.DriverAccount
	Pesan  string
	Status string
}

func getakunFromFile(fname string) ([]*VerifDriverAccount, func(), error) {
	hasil := []*VerifDriverAccount{}
	data, _ := os.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

Parent:
	for _, line := range lines {
		if line == "" {
			continue Parent
		}

		dataline := make([]string, 5)

		fixline := strings.ReplaceAll(line, "\r", "")

		for ind, value := range strings.Split(fixline, "|") {
			dataline[ind] = value
		}

		acdriver, err := tokopedia_lib.NewDriverAccount(dataline[0], dataline[1], dataline[2])
		driver := VerifDriverAccount{
			DriverAccount: acdriver,
			Status:        dataline[3],
			Pesan:         dataline[4],
		}

		if err != nil {
			pdc_common.ReportError(err)
		}

		hasil = append(hasil, &driver)
	}

	return hasil, func() {
		f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for _, driver := range hasil {
			f.WriteString(fmt.Sprintf("%s|%s|%s|%s|%s\n", driver.Username, driver.Password, driver.Secret, driver.Status, driver.Pesan))
		}

	}, nil

}
