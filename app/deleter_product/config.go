package deleter_product

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type AkunDeleteItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type SoldConfig struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (soldf *SoldConfig) GenerateFilter() FilterHandler {
	return func(product *model.SellerProductItem) bool {
		return soldf.Min <= product.TxStats.Sold && product.TxStats.Sold <= soldf.Max
	}
}

type ViewConfig struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (view *ViewConfig) GenerateFilter() FilterHandler {
	return func(product *model.SellerProductItem) bool {
		return view.Min <= product.Stats.CountView && product.Stats.CountView <= view.Max
	}
}

type PriceConfig struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (price *PriceConfig) GenerateFilter() FilterHandler {
	return func(product *model.SellerProductItem) bool {
		return price.Min <= product.Price.Max && product.Price.Max <= price.Max
	}
}

type TokopediaDeleteConfig struct {
	LimitConcurent int                 `json:"limit_concurent"`
	LimitProduct   int                 `json:"limit_product"`
	Title          []string            `json:"title"`
	StatusProduct  model.ProductStatus `json:"product_status"`
	CategoryID     string              `json:"category_id"`
	StartTime      int64               `json:"start_time"`
	EndTime        int64               `json:"end_time"`
	TStartTime     time.Time           `json:"-"`
	TEndTime       time.Time           `json:"-"`
	Akuns          []*AkunDeleteItem   `json:"akuns"`
	SoldFilter     *SoldConfig         `json:"sold_filter,omitempty"`
	ViewFilter     *ViewConfig         `json:"view_filter,omitempty"`
	PriceFilter    *PriceConfig        `json:"price_filter,omitempty"`
}

func NewDeleteConfig(fname string) (*TokopediaDeleteConfig, error) {
	var cfg TokopediaDeleteConfig

	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)

	return &cfg, err
}

func SaveDeleteConfig(fname string, cfg *TokopediaDeleteConfig) error {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *TokopediaDeleteConfig) UnmarshalJSON(data []byte) error {
	type Alias TokopediaDeleteConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(cfg),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	cfg.TStartTime = time.Unix(aux.StartTime, 0)
	cfg.TEndTime = time.Unix(aux.EndTime, 0)

	return nil
}

type FilterHandler func(product *model.SellerProductItem) bool

func (cfg *TokopediaDeleteConfig) GenerateFilter() func(product *model.SellerProductItem) (bool, string) {

	handlers := map[string]FilterHandler{}
	handlers["time"] = cfg.GenerateFilterTime()
	handlers["title"] = cfg.GenerateFilterTitle()

	if cfg.SoldFilter != nil {
		handlers["sold"] = cfg.SoldFilter.GenerateFilter()
	}

	if cfg.ViewFilter != nil {
		handlers["view"] = cfg.ViewFilter.GenerateFilter()
	}

	if cfg.PriceFilter != nil {
		handlers["price"] = cfg.PriceFilter.GenerateFilter()
	}

	return func(product *model.SellerProductItem) (bool, string) {
		for key, handler := range handlers {
			if !handler(product) {

				return false, key
			}
		}

		return true, ""
	}
}

func (cfg *TokopediaDeleteConfig) GenerateFilterTime() FilterHandler {
	return func(product *model.SellerProductItem) bool {
		data := product.CreateTime

		return data.After(cfg.TStartTime) && data.Before(cfg.TEndTime)
	}
}

func (cfg *TokopediaDeleteConfig) GenerateFilterTitle() FilterHandler {
	fstring := []string{}
	fregex := []*regexp.Regexp{}

	if len(cfg.Title) == 0 {
		return func(product *model.SellerProductItem) bool {
			return true
		}
	}

	regpola := `regex-->`

	for _, title := range cfg.Title {
		if strings.HasPrefix(title, regpola) {
			title = strings.Replace(title, regpola, "", 1)
			reitem := regexp.MustCompile(title)
			fregex = append(fregex, reitem)
		} else {
			fstring = append(fstring, title)
		}
	}

	return func(product *model.SellerProductItem) bool {
		datatitle := product.Name
		datatitle = strings.ToLower(datatitle)

		for _, fstr := range fstring {
			if strings.Contains(datatitle, fstr) {
				return true
			}
		}

		for _, freg := range fregex {
			if freg.MatchString(datatitle) {
				return true
			}
		}
		return false
	}
}
