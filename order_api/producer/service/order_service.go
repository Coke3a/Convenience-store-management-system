package service_producer

import (
	"log"

	"github.com/Coke3a/Convenience-store-management-system/order/producer/commands"
	"github.com/google/uuid"
)

//////
//test

// func (repo orderService) CreateNewOrder() error {
// 	orderID := uuid.NewString()
// 	err := repo.orderRepo.CreateOrder(orderID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }

// func (repo orderService) GetLastOrder() (order commands.Order_Command, err error) {
// 	order, err = repo.orderRepo.GetLastOrder()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	return order, nil
// }

func (repo orderService) CreateNewOrderAndGetLastOrder() (order commands.Order_Command, err error) {
	orderID := uuid.NewString()
	err = repo.orderRepo.CreateOrder(orderID)
	if err != nil {
		log.Println(err)
		return order, err
	}
	order, err = repo.orderRepo.GetLastOrder()
	if err != nil {
		log.Println(err)
		return order, err 
	}
	return order, nil
}
////////

func (repo orderService) AddProduct(product_barcode string, order_id string) (product commands.Product_Command, err error) {

	product, err = repo.orderRepo.GetProductById(product_barcode)
	if err != nil {
		log.Println(err)
		return product, err
	}
	orderDetailID := uuid.NewString()
	order_detail := commands.Order_detail_Command{
		ID:              orderDetailID,
		Product_id:      product.ID,
		Order_id:        order_id,
		Product_barcode: product_barcode,
		Product_name:    product.Name,
		Product_Price:   product.Price,
	}
	err = repo.orderRepo.CreateOrderdetail(order_detail)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo orderService) DeleteOrder_Detail(id string) error {
	err := repo.orderRepo.DeleteOrderDetailByID(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo orderService) Total_Price_Without_Promotion(order_id string) (total_price int, err error) {
	order_detail, err := repo.orderRepo.GetAllOrder_DetailByOrderId(order_id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	total_price = 0
	for _, o := range order_detail {
		total_price += o.Product_Price
	}
	UpdateTotalOrder := commands.Order_Command{Total_without_Promo: total_price}
	repo.orderRepo.UpdateOrder(order_id, UpdateTotalOrder)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return total_price, nil
}

func (repo orderService) FindAllOrderDetailByOrderID(order_id string) (order_details []commands.Order_detail_Command , err error) {
	order_details, err = repo.orderRepo.GetAllOrder_DetailByOrderId(order_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order_details, nil
}



func (repo orderService) FindCustomerByTelephone_number(order_id string, telephone_number string) (customer commands.Customer_command, err error) {
	customer, err = repo.orderRepo.GetCustomerByTelephone_number(telephone_number)
	if err != nil {
		log.Println(err)
	}
	UpdateCustomerIDOrder := commands.Order_Command{Customer_ID: customer.ID, Customer_point: customer.Point}
	repo.orderRepo.UpdateOrder(order_id, UpdateCustomerIDOrder)
	if err != nil {
		log.Println(err)
	}
	return customer, nil
}

func (repo orderService) DeleteCustomer(order_id string) error {

	UpdateCustomerIDOrder := commands.Order_Command{Customer_ID: "", Customer_point: 0}
	err := repo.orderRepo.UpdateOrder(order_id, UpdateCustomerIDOrder)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo orderService) FindPromotionwithConditions(order_id string) (promotions []commands.Promotion_Command, err error) {
	order, err := repo.orderRepo.GetOrderById(order_id)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	promotions, err = repo.orderRepo.GetPromotionwithConditions(order.Total_without_Promo, order.Customer_point)
	if err != nil {
		log.Println(err)
		return
	}
	return promotions, nil
}

func (repo orderService) FindPromotionByIdAndCalculate(order_id string, promotion_id string) (TotalWithPromo int, discount int, err error) {
	promotion, err := repo.orderRepo.GetPromotionbyId(promotion_id)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	order, err := repo.orderRepo.GetOrderById(order_id)
	if err != nil {
		log.Println(err)
		return
	}
	if promotion.Type == 1 {
		TotalWithPromo, discount = repo.CalculatePromotionDiscountPrice(order.Total_without_Promo, promotion)
		UpdateCustomerIDOrder := commands.Order_Command{Total_Price: TotalWithPromo, Discount: discount}
		repo.orderRepo.UpdateOrder(order_id, UpdateCustomerIDOrder)
		if err != nil {
			log.Println(err)
			return 0, 0, err
		}
		return TotalWithPromo, discount, nil
	} else if promotion.Type == 2 {
		TotalWithPromo, discount = repo.CalculatePromotionDiscountPercent(order.Total_without_Promo, promotion)
		UpdateCustomerIDOrder := commands.Order_Command{Total_Price: TotalWithPromo, Discount: discount}
		repo.orderRepo.UpdateOrder(order_id, UpdateCustomerIDOrder)
		if err != nil {
			log.Println(err)
			return 0, 0, err
		}
		return TotalWithPromo, discount, nil

	}
	return TotalWithPromo, discount, nil
}

func (repo orderService) CalculatePromotionDiscountPrice(totalNopromo int, promotion commands.Promotion_Command) (TotalWithPromo int, discount int) {
	discount = promotion.Discount
	TotalWithPromo = totalNopromo - discount
	return TotalWithPromo, discount
}

func (repo orderService) CalculatePromotionDiscountPercent(totalNopromo int, promotion commands.Promotion_Command) (TotalWithPromo int, discount int) {
	discount = totalNopromo * promotion.Discount / 100
	// TotalWithPromo = totalNopromo - discount
	TotalWithPromo = totalNopromo - (totalNopromo * promotion.Discount / 100)
	return TotalWithPromo, discount
}
