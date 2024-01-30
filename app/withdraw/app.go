package withdraw

import (
	"github.com/pdcgo/tokopedia_lib"
)

func RunWithdraw(payload []*tokopedia_lib.DriverAccount) error {

	for _, driver := range payload {
		withdraw := NewWithdraw(driver)

		err := withdraw.RunWithDriver()
		if err != nil {
			return err
		}

	}

	return nil
}
