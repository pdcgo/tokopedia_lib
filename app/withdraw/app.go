package withdraw

import (
	"github.com/pdcgo/tokopedia_lib"
)

func RunWithdraw(payload []*tokopedia_lib.DriverAccount) error {
	for _, driver := range payload {
		err := run(driver)
		if err != nil {
			return err
		}
	}

	return nil
}

func run(driver *tokopedia_lib.DriverAccount) error {
	withdraw, err := NewWithdraw(driver)
	if err != nil {
		return err
	}

	return withdraw.Run()
}
