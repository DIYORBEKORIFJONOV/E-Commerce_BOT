package app

import (
	"log"
	"log/slog"

	htppapp "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/app/htpp"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	clientgrpcserver "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/client_grpc_server"
	minao1 "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	orderservice "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustrequest"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/order/adjust/adjustresponse"
	product_service "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/product"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	cleintGrpc,err := clientgrpcserver.NewSerevice(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	addJustRes := adjustresponse.NewAddJsutResponse()
	addJustReq := adjustrequest.NewAddReqeust(cleintGrpc.OrderService(),cleintGrpc,addJustRes)

	serviceOrder := orderservice.NewOrderService(addJustReq,addJustRes,&cleintGrpc)

	orderServiceIml := usecaseorder.NewOrderService(serviceOrder)


	minClien,err := minao1.NewClient("localhost:9005", "minioadmin", "minioadmin", "products",false)
	if err != nil {
		log.Fatal(err)
	}
	

	serviceProduct:= product_service.NewProductReqService(&cleintGrpc)
	serviceProductIml :=productusecase.NewProductUsage(serviceProduct)

	server := htppapp.NewApp(logger,cfg.AppPort,orderServiceIml,minClien,serviceProductIml)
	return &App{
		HTTPApp: server,
	}
}
