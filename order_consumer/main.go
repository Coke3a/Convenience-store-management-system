package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	repository_order "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/order"
	repository_customer "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/customer"
	repository_product "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/product"
	service_customer "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/customer"
	service_product "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/product"
	service_consumer "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/consumer"
	event_consumer "github.com/Coke3a/Convenience-store-management-system/order_consumer/services/consumer/events"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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

	db := initDatabase()

	orderRepo := repository_order.NewOrderRepository(db)
	customerRepo := repository_customer.NewCustomerRepository(db)
	customerService := service_customer.NewCustomerService(customerRepo)

	productRepo := repository_product.NewProductRepository(db)
	productService := service_product.NewProductService(productRepo)


	consumerService := service_consumer.NewConsumerService(customerService, productService, orderRepo)
	consumerHandler := service_consumer.NewConsumerHandler(consumerService)

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	fmt.Println("Order consumer started... ")
	for {
		consumer.Consume(context.Background(), event_consumer.Topics, consumerHandler)
	}
}
