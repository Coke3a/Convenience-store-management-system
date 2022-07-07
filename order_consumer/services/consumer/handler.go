package service_consumer

import (
	"encoding/json"
	"log"
	"reflect"
	repository_order "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/order"
	service_consumer "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/consumer/events"
	service_customer "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/customer"
	service_product "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/product"
	
)


type consumerService struct {
	customerService service_customer.CustomerService
	productService    service_product.ProductService
	orderRepo       repository_order.OrderRepository
}

func NewConsumerService(customerService service_customer.CustomerService, productService service_product.ProductService, orderRepo repository_order.OrderRepository) EventHandler {
	return consumerService{customerService, productService, orderRepo}
}

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

func (r consumerService) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(service_consumer.SubmitOrderEvent{}).Name():
		event := &service_consumer.SubmitOrderEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		newOrder := repository_order.Order{
			ID:                  event.ID,
			Customer_ID:         event.Customer_ID,
			Customer_point:      event.Customer_point,
			Promotion_ID:        event.Promotion_ID,
			Use_point:           event.Use_point,
			Total_without_Promo: event.Total_without_Promo,
			Discount:            event.Discount,
			Total_Price:         event.Total_Price,
		}
		
		err = r.orderRepo.Create_order(newOrder)
		if err != nil {
			log.Println(err)
			return
		}
		if newOrder.Use_point != 0 {
		newOrder.Customer_point -= newOrder.Use_point
		} else { 
			uppoint := newOrder.Total_Price/20 
			newOrder.Customer_point += uppoint
		}
		r.customerService.UsePoints(newOrder.Customer_ID, newOrder.Customer_point)
		log.Printf("%#v", event)
		
	case reflect.TypeOf(service_consumer.SubmitOrder_Order_DetailEvent{}).Name():
		event := &service_consumer.SubmitOrder_Order_DetailEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		newOrderDetail := repository_order.Order_detail{
			ID:              event.ID,
			Product_id:      event.Product_id,
			Order_id:        event.Order_id,
			Product_barcode: event.Product_barcode,
			Product_name:    event.Product_name,
			Product_Price:   event.Product_Price,
		}
		err = r.orderRepo.Create_order_detail(newOrderDetail)
		if err != nil {
			log.Println(err)
			return
		}

		err = r.productService.IncreaseSold(newOrderDetail.Product_id)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%#v", event)

	}

}
