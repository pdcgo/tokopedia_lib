package report

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/common_conf/pdc_application"
	"github.com/pdcgo/common_conf/pdc_common"
)

type AutosendReportItem struct {
	Username            string `csv:"username"`
	ProductUrl          string `csv:"product_url"`
	SellerChatCount     int    `csv:"seller_chat_count"`
	SellerChatProcessed int    `csv:"seller_chat_processed"`
	SellerChatDone      int    `csv:"seller_chat_done"`
	SendChatCount       int    `csv:"send_chat_count"`
	SellerReplyCount    int    `csv:"seller_reply_count"`
	Error               string `csv:"error"`
}

type EditAutosendReportItem func(handler func(item *AutosendReportItem) error) error

var AutosendReportName = "autochat_autosend_report.csv"

type AutosendReport struct {
	base  pdc_application.BaseApplication
	Items []*AutosendReportItem
}

func SaveAutosendReport(report *AutosendReport) error {

	fname := report.base.Path(AutosendReportName)
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(report.Items, file)
	return err
}

func NewAutosendReport(base pdc_application.BaseApplication) *AutosendReport {
	return &AutosendReport{
		base: base,
	}
}

func (r *AutosendReport) CreateItem(username string, sellerCount int) (*AutosendReportItem, EditAutosendReportItem) {

	defer func() {
		err := SaveAutosendReport(r)
		if err != nil {
			pdc_common.ReportError(err)
		}
	}()

	item := &AutosendReportItem{
		Username:        username,
		SellerChatCount: sellerCount,
	}
	r.Items = append(r.Items, item)

	return item, func(handler func(item *AutosendReportItem) error) error {

		err := handler(item)
		if err != nil {
			return pdc_common.ReportError(err)
		}

		err = SaveAutosendReport(r)
		if err != nil {
			return pdc_common.ReportError(err)
		}

		return nil
	}
}
