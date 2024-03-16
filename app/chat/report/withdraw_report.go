package report

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
)

type WitdrawReport struct {
	fname string
	items []*withdraw.WithdrawReport
}

func NewWitdrawReport(fname string) *WitdrawReport {
	return &WitdrawReport{
		fname: fname,
	}
}

func (r *WitdrawReport) Add(items ...*withdraw.WithdrawReport) {
	r.items = append(r.items, items...)
}

func (r *WitdrawReport) Save() {
	file, err := os.Create(r.fname)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}
	defer file.Close()

	err = gocsv.MarshalFile(r.items, file)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}
}
