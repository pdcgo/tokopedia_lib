package withdraw

import (
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/app_config"
)

func RunWithdraw(payload []*tokopedia_lib.DriverAccount) error {
	prox, cancel := app_config.GetProxy(app_config.WithdrawProxyKey)
	defer cancel()

	for _, driver := range payload {
		driver.Proxy = prox.Addr

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
