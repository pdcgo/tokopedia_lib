package autochat

import (
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/pdcgo/common_conf/pdc_application"
)

type ShopStatus string

const SHOP_STATUS_DONE ShopStatus = "done"

type Shop struct {
	ShopName string
	Status   ShopStatus
}

func (s *Shop) ParseLine(line string) {
	dataline := make([]string, 2)
	linesplit := strings.Split(line, "|")
	for ind, value := range linesplit {
		if ind < 2 {
			dataline[ind] = value
		}
	}

	s.ShopName = dataline[0]
	s.Status = ShopStatus(dataline[1])
}

func (s *Shop) ToLine() string {
	return strings.Join([]string{
		s.ShopName,
		string(s.Status),
	}, "|")
}

type ShopData struct {
	sync.Mutex

	fname string
	index int
	Data  []*Shop
}

func NewShopData(base pdc_application.BaseApplication, config *AutochatConfig) (*ShopData, error) {

	fname := base.Path(config.ShopLoc)
	shopdata := ShopData{
		fname: fname,
	}

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return &shopdata, err
	}

	lines, err := fileLineSplit(file)
	if err != nil {
		return &shopdata, err
	}

	for _, line := range lines {
		shop := Shop{}
		shop.ParseLine(line)
		shopdata.Data = append(shopdata.Data, &shop)
	}

	return &shopdata, nil
}

func (ad *ShopData) Save() error {

	file, err := os.Create(ad.fname)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := []string{}
	for _, shop := range ad.Data {
		lines = append(lines, shop.ToLine())
	}

	_, err = file.Write([]byte(strings.Join(lines, "\n")))
	return err
}

func (ad *ShopData) Iterate(handler func(shop *Shop) error) error {

	for _, shop := range ad.Data {

		if shop.Status == SHOP_STATUS_DONE {
			log.Printf("[ %s ] shop status done...", shop.ShopName)
			continue
		}

		err := handler(shop)
		if err != nil {
			return err
		}
	}

	return nil
}

var ErrNoShop = errors.New("shop kosong")
var ErrNoShopMore = errors.New("shop habis")

func (ad *ShopData) Get() (*Shop, error) {

	ad.Lock()
	defer ad.Unlock()

	if len(ad.Data) == 0 {
		return nil, ErrNoShop
	}

	if len(ad.Data) <= ad.index {
		return nil, ErrNoShopMore
	}

	defer func() {
		ad.index++
	}()

	return ad.Data[ad.index], nil
}
