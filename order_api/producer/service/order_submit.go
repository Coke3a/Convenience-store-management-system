package service_producer

import (
	"log"

	"github.com/Coke3a/Convenience-store-management-system/order/events"
	"github.com/Coke3a/Convenience-store-management-system/order/producer/commands"
	repository_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/repository"
)


type produceService struct {
	eventProducer	EventProducer
	orderRepo	repository_producer.OrderRepository
}


type ProduceService interface {
	SubmitOrder(orderId string) error 
	SubmitOrder_Order_Detail(order_detail commands.Order_detail_Command) error	
	LoopForSubmitOrder_Detail(orderID string ) error
	DeleteOrderIdAndOrder_detailByOrderID(orderID string) error
}

func NewProduceService (eventProducer EventProducer, orderRepo	repository_producer.OrderRepository) ProduceService {
	return produceService{eventProducer, orderRepo}
}


func (p produceService) SubmitOrder(orderId string) error  {
	order, err  := p.orderRepo.GetOrderById(orderId)
	if err != nil {
		log.Println(err)
		return err 
	}
	event := events.SubmitOrderEvent{
		ID: order.ID,
		Customer_ID: order.Customer_ID,
		Customer_point: order.Customer_point,
		Promotion_ID: order.Promotion_ID,
		Use_point: order.Use_point,
		Total_without_Promo: order.Total_without_Promo,
		Discount: order.Discount,
		Total_Price: order.Total_Price,
	}
	log.Printf("%#v", event)

	return p.eventProducer.Produce(event)
}


func (p produceService) SubmitOrder_Order_Detail(order_detail commands.Order_detail_Command) error {
		
	event := events.SubmitOrder_Order_DetailEvent{
		ID: order_detail.ID,
		Product_id: order_detail.Product_id,
		Order_id: order_detail.Order_id,
		Product_barcode: order_detail.Product_barcode,
		Product_name: order_detail.Product_name,
		Product_Price: order_detail.Product_Price,
	}
	log.Printf("%#v", event)
	return p.eventProducer.Produce(event)
}





func (p produceService) LoopForSubmitOrder_Detail (orderID string ) error {
	order_details, err := p.orderRepo.GetAllOrder_DetailByOrderId(orderID)
	if err!= nil {
		log.Println(err)
		return err 
	}
	for _ , o := range order_details {
		err = p.SubmitOrder_Order_Detail(o)
		if err!= nil {
			log.Println(err)
			return err 
		}
	}
	return nil 
} 




func (p produceService) DeleteOrderIdAndOrder_detailByOrderID(orderID string) error {
	err := p.orderRepo.Delete_temp_order_by_id(orderID)
	if err!= nil {
		log.Println(err)
		return err 
	}
	err = p.orderRepo.Delete_temp_order_details_by_order_id(orderID)
	if err!= nil {
		log.Println(err)
		return err 
	}
	return nil 
} 