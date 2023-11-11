package deleter_product

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type DeleteReportItem struct {
	Username string              `csv:"username"`
	Judul    string              `csv:"judul"`
	Url      string              `csv:"url"`
	Status   model.ProductStatus `csv:"status"`
	Reason   string              `csv:"reason"`
}

func NewReport() chan *DeleteReportItem {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.FieldsPerRecord = -1
		r.LazyQuotes = true
		return r
	})

	report := make(chan *DeleteReportItem, 10)
	return report
}

func SaveReport(fname string, reports []*DeleteReportItem) error {

	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(reports, file)
	return err
}
