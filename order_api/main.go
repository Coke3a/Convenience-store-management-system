package main

import (
	"fmt"
	"strings"

	handler_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/handler"
	repository_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/repository"
	service_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/service"
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func init(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	return db
} 

func main() {
	producer, err  := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}  
	defer producer.Close()   
	db := initDatabase()
	orderRepository := repository_producer.NewOrderRepository(db)
	orderService := service_producer.NewOrderService(orderRepository)
	eventService := service_producer.NewEventProducer(producer)
	produceService := service_producer.NewProduceService(eventService, orderRepository)
	orderHandler := handler_producer.NewOrderHandler(orderService,produceService)

	

	



	app := fiber.New()
	
	
	app.Get("/order_api/create_order",orderHandler.CreateNewOrderAndGetLastOrder)
	app.Post("/order_api/:id/product", orderHandler.CreateOrder_DetailAndAddProduct)
	app.Get("/order_api/:id/product",orderHandler.FindAllProductsByOrderID)
	app.Delete("/order_api/:id/product", orderHandler.DeleteOrder_Detail)
	app.Get("/order_api/:id/total_without_promotion", orderHandler.Total_Price_Without_Promotion)
	app.Post("/order_api/:id/customer", orderHandler.FindCustomerByTelephone_number)
	app.Put("/order_api/:id/customer", orderHandler.DeleteCustomer)

	app.Get("/order_api/:id/promotion", orderHandler.FindPromotionwithConditions)
	app.Put("/order_api/:id/promotion/:promotion_id", orderHandler.FindPromotionByIdAndCalculate)
	app.Put("/order_api/:id/submit", orderHandler.SubmitOrderAndOrderDetail)


	app.Listen(":8000")


	

}