package autochat

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
)

type AkunStatus string

const AKUN_STATUS_DONE AkunStatus = "done"

type Akun struct {
	Username string
	Password string
	Secret   string
	Status   AkunStatus
}

func (a *Akun) ParseLine(line string) {
	dataline := make([]string, 4)
	linesplit := strings.Split(line, "|")
	for ind, value := range linesplit {
		if ind < 4 {
			dataline[ind] = value
		}
	}

	a.Username = dataline[0]
	a.Password = dataline[1]
	a.Secret = dataline[2]
	a.Status = AkunStatus(dataline[3])
}

func (a *Akun) ToLine() string {
	return strings.Join([]string{
		a.Username,
		a.Password,
		a.Secret,
		string(a.Status),
	}, "|")
}

type AkunData struct {
	sync.Mutex

	config *AutochatConfig
	fname  string
	Data   []*Akun
}

func NewAkunData(base pdc_application.BaseApplication, config *AutochatConfig) (*AkunData, error) {

	fname := base.Path(config.AkunLoc)
	akundata := AkunData{
		config: config,
		fname:  fname,
	}

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return &akundata, err
	}
	defer file.Close()

	lines, err := fileLineSplit(file)
	if err != nil {
		return &akundata, err
	}

	for _, line := range lines {

		akun := Akun{}
		akun.ParseLine(line)
		akundata.Data = append(akundata.Data, &akun)
	}

	return &akundata, nil
}

func (ad *AkunData) Save() error {

	file, err := os.Create(ad.fname)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := []string{}
	for _, akun := range ad.Data {
		lines = append(lines, akun.ToLine())
	}

	_, err = file.Write([]byte(strings.Join(lines, "\n")))
	return err
}

func (ad *AkunData) IterateAkunSender(message *AutochatMessage, handler func(akun *Akun, sender *AutochatSender) error) error {

	for _, akun := range ad.Data {

		if akun.Status == AKUN_STATUS_DONE {
			log.Printf("[ %s ] akun status done...", akun.Username)
			continue
		}

		driver, err := tokopedia_lib.NewDriverAccount(
			akun.Username,
			akun.Password,
			akun.Secret,
		)
		if err != nil {
			pdc_common.ReportError(err)
			continue
		}

		api, _, err := driver.CreateApi()
		if err != nil {
			pdc_common.ReportError(err)
			continue
		}

		sender := NewAutochatSender(api, message, ad.config)
		err = handler(akun, sender)
		if err != nil {
			return err
		}
	}

	return nil
}
