package filter

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"golang.org/x/exp/slices"
)

type BaseFilter struct {
	api           *api_public.TokopediaApiPublic
	base          *legacy_source.BaseConfig
	GrabBasic     *legacy.GrabBasic
	GrabTokopedia *legacy.GrabTokopedia
}

func (filter *BaseFilter) getDataBlacklistUsername() []string {
	if filter.GrabBasic.BlacklistUsername.Tokopedia.Filename == "" {
		return filter.GrabBasic.BlacklistUsername.Tokopedia.Data
	}

	path := filter.base.Path(filter.GrabBasic.BlacklistUsername.Tokopedia.Filename)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error [ shop ] : file filter blacklist username tidak ditemukan di [ %s ]", path)
		return []string{}
	}

	strData := string(data)
	blUsername := strings.Split(strData, "\n")
	return blUsername
}

func (filter *BaseFilter) MinPoint(point int) bool {
	pointMin := filter.GrabTokopedia.Point[0]
	return point < pointMin
}

func (filter *BaseFilter) MaxPoint(point int) bool {
	pointMax := filter.GrabTokopedia.Point[1]
	return point > pointMax
}

func (filter *BaseFilter) LastLogin(lastLogin int64) bool {
	if !filter.GrabTokopedia.LastLoginActive {
		return false
	}
	t := time.Now()
	filterLastLogin := t.AddDate(0, 0, -filter.GrabTokopedia.LastLoginDays)
	return filterLastLogin.Unix() > lastLogin
}

func (filter *BaseFilter) SoldPercentage(soldPercentage int) bool {
	if filter.GrabBasic.Prosentase == 0 {
		return false
	}
	return filter.GrabBasic.Prosentase > soldPercentage
}

func (filter *BaseFilter) Sold(sold int) bool {
	if filter.GrabBasic.Penjualan == 0 {
		return false
	}
	return filter.GrabBasic.Penjualan > sold
}

func (filter *BaseFilter) Stock(stock int) bool {
	if filter.GrabBasic.Stock == 0 {
		return false
	}
	return filter.GrabBasic.Stock > stock
}

func (filter *BaseFilter) LastReview(lastReview int64) bool {
	t := time.Now()
	filterLastReview := t.AddDate(0, 0, -filter.GrabTokopedia.LastLoginDays)
	return lastReview < filterLastReview.Unix()
}

func (filter *BaseFilter) BlacklistUsername(username string) bool {
	if !filter.GrabBasic.BlacklistUsername.Active {
		return false
	}
	blUsername := filter.getDataBlacklistUsername()
	return slices.Contains(blUsername, username)
}

func CreateBaseFilter(api *api_public.TokopediaApiPublic, base *legacy_source.BaseConfig) *BaseFilter {
	filter := &BaseFilter{
		api:           api,
		base:          base,
		GrabBasic:     legacy.NewGrabBasic(base),
		GrabTokopedia: legacy.NewGrabTokopedia(base),
	}
	return filter
}
