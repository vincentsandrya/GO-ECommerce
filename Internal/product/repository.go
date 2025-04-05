package product

import (
	"time"

	"gorm.io/gorm"
)

// Product represents a product in the system
type Product struct {
	ProductId   uint      `json:"product_id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	Stock       int64     `json:"stock"`
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

// GetAllProducts retrieves all products from the database
func (r *Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID
func (r *Repository) GetProductByID(id uint) (Product, error) {
	var product Product
	err := r.DB.First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

// CreateProduct adds a new product to the database
func (r *Repository) CreateProduct(p Product) (Product, error) {
	err := r.DB.Create(&p).Error
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// UpdateProduct updates an existing product in the database
func (r *Repository) UpdateProduct(p Product) (Product, error) {
	err := r.DB.Save(&p).Error
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// DeleteProduct deletes a product from the database by ID
func (r *Repository) DeleteProduct(id uint) error {
	err := r.DB.Delete(&Product{}, id).Error
	return err
}
