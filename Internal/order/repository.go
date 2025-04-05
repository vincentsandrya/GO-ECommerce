package order

import (
	"time"

	"github.com/vincentsandrya/GO-ECommerce/Internal/customer"
	"github.com/vincentsandrya/GO-ECommerce/Internal/product"
	"gorm.io/gorm"
)

type Order struct {
	OrderId     int                `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CustomerId  int                `json:"customer_id"`
	Customer    *customer.Customer `json:"customer" gorm:"foreignKey:CustomerID;references:CustomerID"`
	TotalAmount float64            `json:"total_amount" gorm:"not null"`
	Status      string             `json:"status" gorm:"not null"`
	PaymentDate time.Time          `json:"payment_date"`
	CreatedDate time.Time          `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
}

type OrderItem struct {
	OrderItemId int              `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	OrderId     int              `json:"order_id"`
	Order       *Order           `json:"order" gorm:"foreignKey:OrderId;references:OrderId"`
	ProductId   int              `json:"product_id"`
	Product     *product.Product `json:"product" gorm:"foreignKey:ProductId;references:ProductId"`
	TotalAmount float64          `json:"total_amount" gorm:"not null"`
	Status      string           `json:"status" gorm:"not null"`
	PaymentDate time.Time        `json:"payment_date"`
	CreatedDate time.Time        `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
}

// Repository defines methods for accessing products in the database
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new product repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
