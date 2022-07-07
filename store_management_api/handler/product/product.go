package handler_product

import service "github.com/Coke3a/Convenience-store-management-system/store_management/services/product"


type productHandler struct {
	productSev service.ProductService
}

func NewProductHandler (productSev service.ProductService) productHandler {
	return productHandler{productSev: productSev}
}


type ProductRequest struct {
	ID              string `json:"id"`
	Product_barcode string `json:"product_barcode"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Sold            int    `json:"sold"`
}
