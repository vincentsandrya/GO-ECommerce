package main

import (
	"log"

	"github.com/vincentsandrya/GO-ECommerce/middleware"

	"github.com/gin-gonic/gin"
	"github.com/vincentsandrya/GO-ECommerce/Internal/product"
	"github.com/vincentsandrya/GO-ECommerce/pkg/db"
)

func main() {
	// Connect to PostgreSQL database using GORM
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize repository, service, and handler
	repoProd := product.NewRepository(db)
	servProd := product.NewService(repoProd)
	handProd := product.NewHandler(servProd)

	// repoUser := user.NewRepository(db)
	// servUser := user.NewService(repoUser)
	// handUser := user.NewHandler(servUser)

	// repoCust := customer.NewRepository(db)
	// servCust := customer.NewService(repoCust)
	// handCust := customer.NewHandler(servCust)

	// Initialize Gin router
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	// Define routes
	r.GET("/products", handProd.GetAllProducts)
	r.GET("/product/:id", handProd.GetProductByID)
	r.POST("/products", handProd.CreateProduct)
	r.PUT("/product/:id", handProd.UpdateProduct)
	r.DELETE("/product/:id", handProd.DeleteProduct)
	// r.GET("/customer", handCust.GetAllCustomers)

	// Start the server
	r.Run(":8080")
}
