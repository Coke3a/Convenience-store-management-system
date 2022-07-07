package repository_producer

import "github.com/Coke3a/Convenience-store-management-system/order/producer/commands"



func (p orderRepository) GetProductById(product_code string) (product commands.Product_Command,err error) {
	err = p.db.Table("products").Where("product_barcode=?", product_code).First(&product).Error
	if err != nil {
		return product , err  
		}
	return product,err
}