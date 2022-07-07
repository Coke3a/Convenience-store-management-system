package repository_product

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID           string `gorm:"primary_key;not_null"`
	Product_barcode string
	Name         string
	Price        int
	Sold         int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(product Product) error
	Update(product Product) error
	Delete(id string) error
	FindAll() (products []Product, err error)
	FindbyId(id string) (product Product, err error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.Table("products").AutoMigrate(&Product{})
	return productRepository{db}
}
