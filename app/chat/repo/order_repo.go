package repo

import (
	"time"

	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

type ListOrderFilter struct {
	*Pagination[*model.Order] `form:",inline" schema:",inline"`
	Status                    string    `form:"status" schema:"status" json:"status"`
	BuyerName                 string    `form:"buyer_name" schema:"buyer_name" json:"buyer_name"`
	ProductName               string    `form:"product_name" schema:"product_name" json:"product_name"`
	PriceMin                  int       `form:"price_min" schema:"price_min" json:"price_min"`
	PriceMax                  int       `form:"price_max" schema:"price_max" json:"price_max"`
	DateMin                   time.Time `form:"date_min" schema:"date_min" json:"date_min"`
	DateMax                   time.Time `form:"date_max" schema:"date_max" json:"date_max"`
	TypeDate                  string    `form:"type_date" schema:"type_date" json:"type_date"`
}

func (repo *OrderRepo) Paginate(filter *ListOrderFilter) (res *PaginationResult[*model.Order], err error) {

	tx := repo.db.
		Joins("Account").
		Joins("OrderItems").
		Joins("OrderSheet")

	items := []*model.Order{}
	tx = tx.Scopes(filter.Paginate(items, res, tx)).Find(&items)
	if err = tx.Error; err != nil {
		return
	}

	res.Items = items
	return
}

func (repo *OrderRepo) CreateOrUpdateOrder(orderid int, handler func(order *model.Order) error) error {

	order := model.Order{
		ID: orderid,
	}
	tx := repo.db.
		Preload("Account").
		Preload("OrderItems").
		Preload("OrderSheet").
		First(&order)

	if tx.Error != nil {
		return tx.Error
	}
	if order.ID == 0 {
		order.ID = orderid
	}

	err := handler(&order)
	if err != nil {
		return tx.Error
	}

	tx = repo.db.Save(&order)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Save(&order.Account)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Save(&order.OrderItems)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Save(&order.OrderSheet)
	return tx.Error
}
