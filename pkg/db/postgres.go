package db

import (
	"fmt"
	"log"

	"github.com/vincentsandrya/GO-ECommerce/Internal/customer"
	"github.com/vincentsandrya/GO-ECommerce/Internal/order"
	"github.com/vincentsandrya/GO-ECommerce/Internal/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func Connect() (*gorm.DB, error) {
	conn := "user=postgres password=myB@nk88 dbname=GOECommerce sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	// Migrate the schema/Table
	gd := &GormDB{DB: db}
	gd.MigrationDatabase()

	return db, nil
}

func (gd *GormDB) MigrationDatabase() {

	err := gd.DB.AutoMigrate(&product.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate Product table: %v", err)
	}

	err = gd.DB.AutoMigrate(&customer.Customer{})
	if err != nil {
		log.Fatalf("Failed to migrate User table: %v", err)
	}

	err = gd.DB.AutoMigrate(&order.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate User table: %v", err)
	}

	err = gd.DB.AutoMigrate(&order.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate User table: %v", err)
	}

	// err = gd.DB.AutoMigrate(&user.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate User table: %v", err)
	// }

}
