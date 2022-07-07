package service_product

import (
	"log"
)


func (r productService) IncreaseSold(id string) error {
	product, err := r.productRepo.FindbyId(id)
	if err != nil {
		log.Panicln(err)
		return err 
	}
	product.Sold++

	err = r.productRepo.Update(product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}