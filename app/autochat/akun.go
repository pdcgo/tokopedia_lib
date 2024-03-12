package autochat

import (
	"log"
	"os"
	"strings"
	"sync"

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

func GetAkuns(fname string) (akuns []*Akun, err error) {

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines, err := fileLineSplit(file)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {

		akun := Akun{}
		akun.ParseLine(line)
		akuns = append(akuns, &akun)
	}

	return
}

func SaveAkuns(fname string, akuns []*Akun) error {

	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := []string{}
	for _, akun := range akuns {
		lines = append(lines, akun.ToLine())
	}

	_, err = file.Write([]byte(strings.Join(lines, "\n")))
	return err
}

func (app *Application) IterateAkunSender() (chan *AutochatSender, error) {

	var lock sync.Mutex
	fname := app.base.Path(app.config.AkunLoc)
	akunsendchan := make(chan *AutochatSender, app.config.Concurrent)
	akuns, err := GetAkuns(fname)

	if os.IsNotExist(err) {
		close(akunsendchan)
		return akunsendchan, nil

	} else if err != nil {
		return nil, err
	}

	go func() {
		defer close(akunsendchan)

		for _, akun := range akuns {

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

			sender := NewAutochatSender(api, app.message, app.config)
			sender.OnDone = func() {
				lock.Lock()
				defer lock.Unlock()

				akun.Status = AKUN_STATUS_DONE
				SaveAkuns(fname, akuns)
			}
			akunsendchan <- sender
		}
	}()

	return akunsendchan, nil
}
