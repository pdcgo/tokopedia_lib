package repo

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

var (
	PAGINATION_DEFAULT_PAGE      = 1
	PAGINATION_DEFAULT_PAGE_SIZE = 50
)

type PaginationResult[T any] struct {
	Size  int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Items []T `json:"items"`
	Total int `json:"total"`
}

type Pagination[T any] struct {
	Size     int    `json:"limit" form:"limit" schema:"limit"`
	Page     int    `json:"page" form:"page" schema:"page"`
	SortBy   string `json:"sort_by" form:"sort_by" schema:"sort_by"`
	SortType string `json:"sort_type" form:"sort_type" schema:"sort_type"`
}

func (p *Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination[T]) GetLimit() int {
	if p.Size == 0 {
		p.Size = PAGINATION_DEFAULT_PAGE_SIZE
	}
	return p.Size
}

func (p *Pagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = PAGINATION_DEFAULT_PAGE
	}
	return p.Page
}

func (p *Pagination[T]) GetSort() string {
	if p.SortBy == "" {
		p.SortBy = fmt.Sprintf("%s %s", p.SortBy, p.SortType)
	}
	return p.SortBy
}

func (p *Pagination[T]) Paginate(value any, res *PaginationResult[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	var totalRows int64
	db.Model(value).Count(&totalRows)

	res.Page = p.Page
	res.Size = p.Size
	res.Total = int(totalRows)
	res.Pages = int(math.Ceil(float64(totalRows) / float64(p.Size)))

	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(p.GetOffset()).
			Limit(p.GetLimit()).
			Order(p.GetSort())
	}
}
