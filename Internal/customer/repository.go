package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerId  uint      `json:"customer_id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"unique;not null"`
	Address     string    `json:"address" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	CreatedDate time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
}

// Repository defines methods for accessing products in the database
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new product repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
