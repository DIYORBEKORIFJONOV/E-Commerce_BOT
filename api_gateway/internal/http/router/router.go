package router

import (
	"api_gateway/internal/http/handler"
	middleware "api_gateway/internal/http/midleware"
	usecaseorder "api_gateway/internal/usecase/order"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func RegisterRouter(orderUC *usecaseorder.OrderUseCaseIml) *gin.Engine {

	orderHandler := handler.NewOrderHandler(orderUC)

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.IPFilterMiddleware([]string{"127.0.0.1"}))
	r.Use(middleware.TimingMiddleware())


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("swagger/doc.json")))


	order := r.Group("/orders")
	{
		order.POST("/create", orderHandler.CreateOrder)
		order.GET("/getall", orderHandler.GetOrders)
		order.PUT("/completed", orderHandler.OrderCompleted)
	}

	cart := r.Group("/carts")
	{
		cart.POST("/add/product",orderHandler.AddProduct2Cart)
		cart.GET("/",orderHandler.GetCart)
		cart.PUT("/",orderHandler.UpdateCart)
		cart.DELETE("/",orderHandler.DeleteCart)
		cart.DELETE("/product",orderHandler.DeleteProductsFromCart)

	}


	

	return r
}
