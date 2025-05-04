// file: internal/http/router/router.go
package router

import (
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/http/handler"
	middleware "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/http/midleware"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(
    orderUC *usecaseorder.OrderUseCaseIml,
    productUC *productusecase.ProductUseCaseIml,
    minioClient *minao1.FileStorage,
) *gin.Engine {
    r := gin.Default()
    r.Use(middleware.CorsMiddleware())
    r.Use(middleware.IPFilterMiddleware([]string{"127.0.0.1","172.18.0.1"}))
    r.Use(middleware.TimingMiddleware())

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

    orderHandler := handler.NewOrderHandler(orderUC)
    productHandler := handler.NewProductHandler(productUC, minioClient)

    // Orders
    orders := r.Group("/orders")
    {
        orders.POST("/create", orderHandler.CreateOrder)
        orders.GET("/getall", orderHandler.GetOrders)
        orders.PUT("/completed", orderHandler.OrderCompleted)
    }

    // Cart
    cart := r.Group("/carts")
    {
        cart.POST("/add/product", orderHandler.AddProduct2Cart)
        cart.GET("/", orderHandler.GetCart)
        cart.PUT("/", orderHandler.UpdateCart)
        cart.DELETE("/", orderHandler.DeleteCart)
        cart.DELETE("/product", orderHandler.DeleteProductsFromCart)
    }

    // Products
    products := r.Group("/products")
    {
        products.POST("/", productHandler.CreateProduct)
        products.POST("/addmodel", productHandler.AddModel)
        products.PUT("/name", productHandler.UpdateProductName)
        products.GET("/main", productHandler.GetMainProduct)
        products.GET("/", productHandler.GetAllProduct)
        products.PATCH("/", productHandler.UpdateProduct)
        products.DELETE("/", productHandler.DeleteProduct)
    }

    // Categories
    categories := r.Group("/products/categories")
    {
        categories.POST("/", productHandler.CreateCategory)
        categories.GET("/", productHandler.GetAllCategory)
        categories.PUT("/", productHandler.UpdateCategory)
        categories.DELETE("/", productHandler.DeleteCategory)
    }

    return r
}
