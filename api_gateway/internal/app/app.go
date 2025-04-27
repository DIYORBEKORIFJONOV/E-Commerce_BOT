package app

import (
	htppapp "api_gateway/internal/app/htpp"
	"api_gateway/internal/config"
	clientgrpcserver "api_gateway/internal/infastructure/client_grpc_server"
	orderservice "api_gateway/internal/service/order"
	"api_gateway/internal/service/order/adjust/adjustrequest"
	"api_gateway/internal/service/order/adjust/adjustresponse"
	usecaseorder "api_gateway/internal/usecase/order"
	"log"
	"log/slog"
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

	server := htppapp.NewApp(logger,cfg.AppPort,orderServiceIml)
	return &App{
		HTTPApp: server,
	}
}
