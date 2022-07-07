package service_product

import repository "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/product"

type productService struct {
	productRepo repository.ProductRepository
}

type ProductResponse struct {
	ID              string `json:"id"`
	Product_barcode string `json:"product_barcode"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Sold            int    `json:"sold"`
}

type ProductService interface {

	IncreaseSold(id string) error
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}
