package repo

import (
	"fmt"
	"log"
	"math"

	"gorm.io/gorm"
)

var (
	PAGINATION_DEFAULT_PAGE      = 1
	PAGINATION_DEFAULT_PAGE_SIZE = 50
)

type PaginationResult[T any] struct {
	SortBy   string `json:"-"`
	SortType string `json:"-"`
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Pages    int    `json:"pages"`
	Items    []T    `json:"items"`
	Total    int    `json:"total"`
}

func (p *PaginationResult[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationResult[T]) GetLimit() int {
	if p.Size == 0 {
		p.Size = PAGINATION_DEFAULT_PAGE_SIZE
	}
	return p.Size
}

func (p *PaginationResult[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = PAGINATION_DEFAULT_PAGE
	}
	return p.Page
}

func (p *PaginationResult[T]) GetSort() string {
	if p.SortBy == "" && p.SortType != "" {
		p.SortBy = fmt.Sprintf("%s %s", p.SortBy, p.SortType)
	}
	return p.SortBy
}

func (p *PaginationResult[T]) Paginate(db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	var totalRows int64
	db.Model(p.Items).Count(&totalRows)

	log.Println(float64(totalRows), float64(p.Size))

	p.Total = int(totalRows)
	p.Pages = int(math.Ceil(float64(totalRows) / float64(p.Size)))

	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(p.GetOffset()).
			Limit(p.GetLimit()).
			Order(p.GetSort())
	}
}
