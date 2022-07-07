package service_product

import repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/product"

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
	AddNewProduct(product_barcode string, name string, price int) error
	UpdateProduct(id string, product_barcode string, name string, price int) error
	DeleteProduct(id string) error
	FindAllProducts() (products []ProductResponse, err error)
	FindProductById(id string) (product ProductResponse, err error)
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}
