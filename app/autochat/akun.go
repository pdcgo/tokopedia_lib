package autochat

import (
	"os"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

func (app *Application) IterateAkunSender() (chan *AutochatSender, error) {

	fname := app.base.Path(app.config.AkunLoc)
	akunsendchan := make(chan *AutochatSender, app.config.Concurrent)
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if os.IsNotExist(err) {
		close(akunsendchan)
		return akunsendchan, nil

	} else if err != nil {
		return nil, err
	}

	lines, err := fileLineSplit(file)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(akunsendchan)

		for _, line := range lines {

			dataline := make([]string, 3)
			linesplit := strings.Split(line, "|")
			for ind, value := range linesplit {
				if ind < 3 {
					dataline[ind] = value
				}
			}

			driver, err := tokopedia_lib.NewDriverAccount(dataline[0], dataline[1], dataline[2])
			if err != nil {
				pdc_common.ReportError(err)
				continue
			}

			api, _, err := driver.CreateApi()
			if err != nil {
				pdc_common.ReportError(err)
				continue
			}

			sender := NewAutochatSender(api, app.message, app.config)
			akunsendchan <- sender
		}
	}()

	return akunsendchan, nil
}
