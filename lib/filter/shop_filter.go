package filter

import (
	"fmt"
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Shop struct {
	Id     int
	Domain string
}

type ShopFilter struct {
	BaseFilter
	Shop Shop
}

func (filter *ShopFilter) getShopStats() (model_public.ShopStatisticQueryData, error) {
	variable := model_public.ShopStatisticQueryVar{
		ShopID:    filter.Shop.Id,
		ShopIDStr: fmt.Sprintf("%d", filter.Shop.Id),
	}
	stats, err := filter.api.ShopStatisticQuery(&variable)

	return stats.Data, err
}

func (filter *ShopFilter) RatingFilter(rating float64) bool {
	shopStats, err := filter.getShopStats()
	if err != nil {
		fmt.Printf("error [ shop ] : terjadi kesalahan pada toko [ %s ]\n", filter.Shop.Domain)
		return true
	}
	shopRating, _ := strconv.Atoi(shopStats.ShopRating.RatingScore)
	return float64(shopRating) > rating
}

func (filter *ShopFilter) FilterPoint() bool {
	if !filter.GrabTokopedia.LastLoginActive {
		return false
	}
	shopStats, err := filter.getShopStats()
	if err != nil {
		fmt.Printf("error [ shop ] : terjadi kesalahan pada toko [ %s ]\n", filter.Shop.Domain)
		return true
	}
	shopPoint, _ := strconv.Atoi(shopStats.ShopReputation[0].ScoreMap)
	if filter.MinPoint(shopPoint) {
		fmt.Printf("filter [ shop ] : point toko [ %s ] kurang dari filter\n", filter.Shop.Domain)
		return true
	}
	if filter.MaxPoint(shopPoint) {
		fmt.Printf("filter [ shop ] : point toko [ %s ] lebih dari filter\n", filter.Shop.Domain)
		return true
	}
	return false

}

func (filter *ShopFilter) FilterBlacklistUsername() bool {
	if !filter.BaseFilter.GrabBasic.BlacklistUsername.Active {
		return false
	}
	if filter.BlacklistUsername(filter.Shop.Domain) {
		fmt.Printf("filter [ shop ] : toko [ %s ] termasuk blacklist\n", filter.Shop.Domain)
		return true
	}
	return false
}

func (filter *ShopFilter) ApplyFilter() bool {
	filters := []func() bool{
		filter.FilterPoint,
		filter.FilterBlacklistUsername,
	}
	for _, filter := range filters {
		if filter() {
			return true
		}
	}
	return false
}

func CreateShopFilter(base BaseFilter, shop Shop) *ShopFilter {
	return &ShopFilter{
		base,
		shop,
	}
}
