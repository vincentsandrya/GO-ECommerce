package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId      uint      `gorm:"primaryKey;identity"`
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
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
