package repository_order


func (db orderRepository) Create_order(order Order) error  {
	err := db.db.Table("orders").Create(&order).Error
	if err != nil {
		return err
	}
	return nil
} 

func (db orderRepository) Create_order_detail(order_detail Order_detail) error {
	err := db.db.Table("order_details").Create(&order_detail).Error
	if err != nil {
		return err
	}
	return nil 
} 



