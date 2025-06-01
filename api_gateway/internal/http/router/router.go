package router

import (
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/http/handler"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/http/middleware"
	authusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/auth"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host hurmomarketshop.duckdns.org
// @BasePath        /
// @schemes         https
// @securityDefinitions.apiKey ApiKeyAuth
// @in              header
// @name            Authorization

func RegisterRouter(
	orderUC *usecaseorder.OrderUseCaseIml,
	productUC *productusecase.ProductUseCaseIml,
	minioClient *minao1.FileStorage,
	authIML authusecase.AuthUseCaseIml,
) *gin.Engine {
	r := gin.Default()

	// Timing Middleware (можно оставить на все)
	r.Use(middleware.TimingMiddleware)

	// Swagger UI — БЕЗ Casbin и rate limiting middleware
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Подключаем основное middleware (после swagger)
	r.Use(middleware.Middleware)

	// Инициализация хендлеров
	orderHandler := handler.NewOrderHandler(orderUC)
	productHandler := handler.NewProductHandler(productUC, minioClient)
	authHandler := handler.NewAuthHandler(authIML)

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("register", authHandler.RegisterAccountHandler)
		auth.POST("verify", authHandler.VerifyAccountHandler)
		auth.POST("login", authHandler.LoginHandler)
	}

	account := r.Group("/account")
	{
		account.PUT("update-password", authHandler.ChangePasswordHandler)
	}

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
		products.POST("", productHandler.CreateProduct)
		products.POST("/addmodel", productHandler.AddModel)
		products.PUT("/name", productHandler.UpdateProductName)
		products.GET("/main", productHandler.GetMainProduct)
		products.GET("", productHandler.GetAllProduct)
		products.PATCH("", productHandler.UpdateProduct)
		products.DELETE("", productHandler.DeleteProduct)
	}

	// Categories
	categories := r.Group("/products/categories")
	{
		categories.POST("", productHandler.CreateCategory)
		categories.GET("", productHandler.GetAllCategory)
		categories.PUT("", productHandler.UpdateCategory)
		categories.DELETE("", productHandler.DeleteCategory)
	}

	return r
}
