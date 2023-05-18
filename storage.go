package tokopedia_lib

import (
	"fmt"
	"os"
	"strings"
)

type Account struct {
	Username string
	Pass     string
	Secret   string
	Ktp      string
	Status   string
}

func GetVerificationAkun(fname string) ([]*Account, func(), error) {
	hasil := []*Account{}
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

		hasil = append(hasil, &Account{
			Username: dataline[0],
			Pass:     dataline[1],
			Secret:   dataline[2],
			Ktp:      dataline[3],
			Status:   dataline[4],
		})
	}

	return hasil, func() {
		f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for _, driver := range hasil {
			f.WriteString(fmt.Sprintf("%s|%s|%s|%s|%s\n", driver.Username, driver.Pass, driver.Secret, driver.Ktp, driver.Status))
		}

	}, nil

}
