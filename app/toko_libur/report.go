package toko_libur

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/rs/zerolog"
)

type AkunTokoLibur struct {
	Username string `csv:"username" json:"username"`
	Password string `csv:"password" json:"password"`
	Secret   string `csv:"secret" json:"secret"`
}

type ReportItem struct {
	*AkunTokoLibur
	StatusBefore string `csv:"status_before" json:"status_before"`
	StatusAfter  string `csv:"status_after" json:"status_after"`
	Message      string `csv:"message" json:"message"`
	Error        string `csv:"error" json:"error"`
}

func (r *ReportItem) ReportError(err error) error {
	r.Error = err.Error()
	return pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
		data, _ := json.Marshal(r)
		return event.
			Str("app", "tokopedia_toko_libur").
			Str("username", r.Username).
			RawJSON("data", data)
	})
}

func (item *ReportItem) CreateDriver() (*tokopedia_lib.DriverAccount, error) {
	return tokopedia_lib.NewDriverAccount(
		item.Username,
		item.Password,
		item.Secret,
	)
}

type Report struct {
	sync.Mutex
	fname string
	Data  []*ReportItem
}

func NewReport(config *TokopediaTokoLiburConfig) (*Report, error) {

	fname := config.GetReportPath()
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	report := Report{
		fname: fname,
		Data:  []*ReportItem{},
	}
	err = gocsv.UnmarshalFile(file, &report.Data)
	if errors.Is(err, gocsv.ErrEmptyCSVFile) {
		err = nil
	}

	return &report, err
}

func (r *Report) Save() error {
	r.Lock()
	defer r.Unlock()

	file, err := os.Create(r.fname)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(r.Data, file)
}

func (r *Report) Iterate(
	handler func(
		item *ReportItem,
		updateitem func(handler func(item *ReportItem) error) error,
	) error,
) error {

	for _, item := range r.Data {
		ritem := item
		err := handler(ritem, func(itemhandler func(item *ReportItem) error) error {
			defer func() {
				err := r.Save()
				if err != nil {
					pdc_common.ReportError(err)
				}
			}()
			return itemhandler(ritem)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
