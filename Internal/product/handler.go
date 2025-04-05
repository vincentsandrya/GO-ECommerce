package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for products
type Handler struct {
	Service *Service
}

// NewHandler creates a new product handler
func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

// GetAllProducts handles GET requests to fetch all products
func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID handles GET requests to fetch a single product by ID
func (h *Handler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("product_id"))
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST requests to create a new product
func (h *Handler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct, err := h.Service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct handles PUT requests to update an existing product
func (h *Handler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("product_id"))
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ProductId = uint(id)
	updatedProduct, err := h.Service.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles DELETE requests to delete a product by ID
func (h *Handler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("product_id"))
	err := h.Service.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
