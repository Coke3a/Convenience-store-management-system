package service_product

import (
	"log"

	repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/product"
	"github.com/google/uuid"
)

func (r productService) AddNewProduct(product_barcode string, name string, price int) error {
	product := repository.Product{
		ID:              uuid.NewString(),
		Product_barcode: product_barcode,
		Name:            name,
		Price:           price,
	}
	err := r.productRepo.Create(product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (r productService) UpdateProduct(id string, product_barcode string, name string, price int) error {
	product := repository.Product{
		ID:              id,
		Product_barcode: product_barcode,
		Name:            name,
		Price:           price,
	}
	err := r.productRepo.Update(product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r productService) DeleteProduct(id string) error {
	err := r.productRepo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r productService) FindAllProducts() (products []ProductResponse, err error) {
	productDB, err := r.productRepo.FindAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, p := range productDB {
		products = append(products, ProductResponse{
			ID:              p.ID,
			Product_barcode: p.Product_barcode,
			Name:            p.Name,
			Price:           p.Price,
			Sold:            p.Sold,
		})
	}
	return products, nil
}

func (r productService) FindProductById(id string) (product ProductResponse, err error) {
	productDB, err := r.productRepo.FindbyId(id)
	if err != nil {
		log.Println(err)
		return
	}
	product = ProductResponse{
		ID:              productDB.ID,
		Product_barcode: productDB.Product_barcode,
		Name:            productDB.Name,
		Price:           productDB.Price,
		Sold:            productDB.Sold,
	}
	return product, nil
}

