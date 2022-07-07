package repository_producer

import "github.com/Coke3a/Convenience-store-management-system/order/producer/commands"







func (db orderRepository) CreateOrder(id string) (err error) {
	order := commands.Order_Command{ID: id}
	err = db.db.Table("tempdb_orders").Create(&order).Error
	if err != nil {
		return err
	}
	return nil	
}

func (db orderRepository) UpdateOrder(id string , order commands.Order_Command) error {
		err := db.db.Table("tempdb_orders").Model(&commands.Order_Command{}).Where("id=?", id).Updates(order).Error
		if err != nil {
			return err
		}
		return nil
}

func (db orderRepository) GetOrderById(id string) (order commands.Order_Command,err error) {
	err = db.db.Table("tempdb_orders").Where("id=?", id).First(&order).Error
	if err != nil {
		return order , err  
		}
	return order,nil
}



func (db orderRepository) GetLastOrder()(order commands.Order_Command, err error) {
	err = db.db.Table("tempdb_orders").Last(&order).Error
	if err != nil {
		return order, err  
	}
	return order, nil 
}





func (db orderRepository) CreateOrderdetail(order_detail commands.Order_detail_Command) (err error) {
	err = db.db.Table("tempdb_order_details").Create(&order_detail).Error
	if err != nil {
		return err 
	}
	return nil	
}


////

func (db orderRepository) GetAllOrder_DetailByOrderId(order_id string) (order_details []commands.Order_detail_Command, err error) {
	err = db.db.Table("tempdb_order_details").Where("order_id=?",order_id).Find(&order_details).Error
	if err != nil {
		return nil, err
	}
	return order_details, nil
}


func (db orderRepository) DeleteOrderDetailByID(id string) error {
	err := db.db.Table("tempdb_order_details").Where("id=?",id).Delete(&commands.Order_detail_Command{}).Error
	if err != nil {
		return err
	}
	return nil 
}



func (db orderRepository)Delete_temp_order_by_id(orderID string) error{
	err := db.db.Table("tempdb_orders").Where("id=?", orderID).Delete(&commands.Order_Command{}).Error
	if err != nil {
		return err
	}
	return nil

}

func (db orderRepository) Delete_temp_order_details_by_order_id(orderID string) error{
	err := db.db.Table("tempdb_order_details").Where("order_id=?", orderID).Delete(&commands.Order_detail_Command{}).Error
	if err != nil {
		return err
	}
	return nil

} 