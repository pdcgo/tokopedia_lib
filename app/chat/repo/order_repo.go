package repo

import (
	"errors"
	"fmt"
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
	Size        int               `json:"size" form:"size" schema:"size"`
	Page        int               `json:"page" form:"page" schema:"page"`
	SortBy      string            `json:"sort_by" form:"sort_by" schema:"sort_by"`
	SortType    string            `json:"sort_type" form:"sort_type" schema:"sort_type"`
	Status      model.OrderStatus `form:"status" schema:"status" json:"status"`
	BuyerName   string            `form:"buyer_name" schema:"buyer_name" json:"buyer_name"`
	ProductName string            `form:"product_name" schema:"product_name" json:"product_name"`
	PriceMin    int               `form:"price_min" schema:"price_min" json:"price_min"`
	PriceMax    int               `form:"price_max" schema:"price_max" json:"price_max"`
	DateMin     time.Time         `form:"date_min" schema:"date_min" json:"date_min"`
	DateMax     time.Time         `form:"date_max" schema:"date_max" json:"date_max"`
	TypeDate    string            `form:"type_date" schema:"type_date" json:"type_date"`
}

func (repo *OrderRepo) Paginate(filter *ListOrderFilter) (*PaginationResult[*model.Order], error) {

	tx := repo.db.
		Preload("Account").
		Preload("OrderItems").
		Preload("OrderSheet")

	if filter.ProductName != "" {
		tx = tx.Where(
			`id IN (SELECT order_id FROM order_items WHERE product_name LIKE ?)`,
			"%"+filter.ProductName+"%",
		)
	}

	if filter.PriceMin > 0 {
		tx = tx.Where(
			`id IN (SELECT order_id FROM order_items WHERE price >= ?)`,
			filter.PriceMin,
		)
	}
	if filter.PriceMax > 0 {
		tx = tx.Where(
			`id IN (SELECT order_id FROM order_items WHERE price <= ?)`,
			filter.PriceMax,
		)
	}

	if filter.BuyerName != "" {
		tx = tx.Where("buyer_name LIKE ?", "%"+filter.BuyerName+"%")
	}

	if !filter.DateMin.IsZero() {
		tx = tx.Where(fmt.Sprintf("%s >= ?", filter.TypeDate), filter.DateMin)
	}
	if !filter.DateMax.IsZero() {
		tx = tx.Where(fmt.Sprintf("%s <= ?", filter.TypeDate), filter.DateMax)
	}

	statuses := model.OrderStatuses.Batch(
		model.NewOrder,
		model.ConfirmShipping,
		model.InShipping,
		model.Complaint,
	)
	if filter.Status != "" {
		statuses = model.OrderStatuses[filter.Status]
	}
	tx = tx.Where("status_id IN ?", statuses)

	res := &PaginationResult[*model.Order]{
		Page:     filter.Page,
		Size:     filter.Size,
		SortBy:   filter.SortBy,
		SortType: filter.SortType,
		Items:    []*model.Order{},
	}
	err := tx.Scopes(res.Paginate(tx)).Find(&res.Items).Error

	return res, err
}

func (repo *OrderRepo) GetOrder(orderid int) (*model.Order, error) {
	order := model.Order{
		ID: orderid,
	}
	err := repo.db.
		Preload("Account").
		Preload("OrderItems").
		Preload("OrderSheet").
		First(&order).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return &order, err
}

func (repo *OrderRepo) CreateOrUpdateOrderItem(prodid int, handler func(item *model.OrderItem) error) error {
	orderItem := model.OrderItem{}
	err := repo.db.
		Where("product_id = ?", prodid).
		First(&orderItem).
		Error
	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !notFound {
		return err
	}

	if err := handler(&orderItem); err != nil {
		return err
	}

	if notFound {
		orderItem.ProductID = prodid
		return repo.db.Create(&orderItem).Error
	}
	return repo.db.
		Where("product_id = ?", prodid).
		Updates(&orderItem).
		Error
}

func (repo *OrderRepo) CreateOrUpdateOrder(orderid int, handler func(order *model.Order) error) (err error) {

	order, err := repo.GetOrder(orderid)
	if err != nil {
		return
	}
	if order.ID == 0 {
		order.ID = orderid
	}

	if err = handler(order); err != nil {
		return
	}

	if err = repo.db.Save(&order).Error; err != nil {
		return
	}

	if order.Account != nil {
		if err = repo.db.Save(&order.Account).Error; err != nil {
			return
		}
	}

	if order.OrderSheet != nil {
		if err = repo.db.Save(&order.OrderSheet).Error; err != nil {
			return
		}
	}
	return
}
