package report

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
)

type AutoreplyReportItem struct {
	Username         string `csv:"username"`
	ChatFrom         string `csv:"from"`
	ReceiveChatCount int    `csv:"receive_chat_count"`
	SendChatCount    int    `csv:"send_chat_count"`
	Error            string `csv:"error"`
}

type EditAutoreplyReportItem func(handler func(item *AutoreplyReportItem) error) error

var AutoreplyReportName = "autochat_autoreply_report.csv"

type AutoreplyReport struct {
	base  pdc_application.BaseApplication
	Items []*AutoreplyReportItem
}

func SaveAutoreplyReport(report *AutoreplyReport) error {

	fname := report.base.Path(AutoreplyReportName)
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(report.Items, file)
	return err
}

func NewAutoreplyReport(base pdc_application.BaseApplication) *AutoreplyReport {
	return &AutoreplyReport{
		base: base,
	}
}

func (r *AutoreplyReport) CreateItem(username string, from string) (*AutoreplyReportItem, EditAutoreplyReportItem) {

	defer func() {
		err := SaveAutoreplyReport(r)
		if err != nil {
			pdc_common.ReportError(err)
		}
	}()

	item := &AutoreplyReportItem{
		Username: username,
		ChatFrom: from,
	}
	r.Items = append(r.Items, item)

	return item, func(handler func(item *AutoreplyReportItem) error) error {

		err := handler(item)
		if err != nil {
			return pdc_common.ReportError(err)
		}

		err = SaveAutoreplyReport(r)
		if err != nil {
			return pdc_common.ReportError(err)
		}

		return nil
	}
}
