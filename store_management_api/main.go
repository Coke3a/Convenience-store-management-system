package main

import (
	"fmt"
	"strings"

	handler_customer "github.com/Coke3a/Convenience-store-management-system/store_management/handler/customer"
	handler_product "github.com/Coke3a/Convenience-store-management-system/store_management/handler/product"
	handler_promotion "github.com/Coke3a/Convenience-store-management-system/store_management/handler/promotion"
	handler_user "github.com/Coke3a/Convenience-store-management-system/store_management/handler/user"
	repository_customer "github.com/Coke3a/Convenience-store-management-system/store_management/repository/customer"
	repository_product "github.com/Coke3a/Convenience-store-management-system/store_management/repository/product"
	repository_promotion "github.com/Coke3a/Convenience-store-management-system/store_management/repository/promotion"
	repository_user "github.com/Coke3a/Convenience-store-management-system/store_management/repository/user"
	service_customer "github.com/Coke3a/Convenience-store-management-system/store_management/services/customer"
	service_product "github.com/Coke3a/Convenience-store-management-system/store_management/services/product"
	service_promotion "github.com/Coke3a/Convenience-store-management-system/store_management/services/promotion"
	service_user "github.com/Coke3a/Convenience-store-management-system/store_management/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	promotionRepository := repository_promotion.NewPromotionRepository(db)
	promotionService := service_promotion.NewPromotionService(promotionRepository)
	promotionHandler := handler_promotion.NewPromotionHandler(promotionService)

	customerRepo := repository_customer.NewCustomerRepository(db)
	customerService := service_customer.NewCustomerService(customerRepo)
	customerHandler := handler_customer.NewCustomerHandler(customerService)

	productRepo := repository_product.NewProductRepository(db)
	productService := service_product.NewProductService(productRepo)
	productHandler := handler_product.NewProductHandler(productService)

	userRepo := repository_user.NewUserRepository(db)
	userService := service_user.NewUserService(userRepo)
	userHandler := handler_user.NewUserHandler(userService)


	app := fiber.New()

	app.Post("/store_managment_api/add_customer", customerHandler.CreateNewCustomer)
	app.Get("/store_managment_api/find_customers", customerHandler.FindAllCustomer)
	app.Get("/store_managment_api/find_customer/:id", customerHandler.FindCustomerByid)
	app.Put("/store_managment_api/update_customer/:id", customerHandler.UpdateCustomer)
	app.Delete("/store_managment_api/delete_customer/:id", customerHandler.DeleteCustomer)

	app.Post("/store_managment_api/add_user", userHandler.CreateNewUser)
	app.Get("/store_managment_api/find_users", userHandler.FindAllUser)
	app.Get("/store_managment_api/find_users/:id", userHandler.FindUserByid)
	app.Put("/store_managment_api/update_user/:id", userHandler.UpdateUser)
	app.Delete("/store_managment_api/delete_user/:id", userHandler.DeleteUser)


	app.Post("/store_managment_api/add_product", productHandler.CreateNewProduct)
	app.Get("/store_managment_api/find_products", productHandler.FindAllProducts)
	app.Get("/store_managment_api/find_product/:id", productHandler.FindProductByid)
	app.Put("/store_managment_api/update_product/:id", productHandler.UpdateProduct)
	app.Delete("/store_managment_api/delete_productproduct/:id", productHandler.DeleteProduct)


	app.Post("/store_managment_api/add_promotion_price", promotionHandler.CreateNewPromotionDiscountPrice)
	app.Post("/store_managment_api/add_promotion_percent", promotionHandler.CreateNewPromotionDiscountPercent)
	app.Get("/store_managment_api/find_promotions", promotionHandler.FindAllPromotion)
	app.Get("/store_managment_api/find_promotion/:id", promotionHandler.FindPromotionByid)
	app.Put("/store_managment_api/update_promotion/:id", promotionHandler.UpdatePromotion)
	app.Delete("/store_managment_api/delete_promotion/:id", promotionHandler.DeletePromotion)

		app.Listen(":4000")

}




