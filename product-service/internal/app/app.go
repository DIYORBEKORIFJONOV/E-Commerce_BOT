package app

import (
	"os"
	grpcapp "productservice/internal/app/grpc"
	"productservice/internal/config"
	mongodb "productservice/internal/infastructure/repository/mongo"

	productservice "productservice/internal/service/product"
	productuserepository "productservice/internal/usecase/product/productrepository"
	productusecase "productservice/internal/usecase/product/productservice"
	logger "productservice/log"
)

type App struct {
	GRPCServer *grpcapp.App
}


func NewApp(cfg *config.Config,logger *logger.Logger) *App {
	logCtx := logger.WithContext(map[string]string{
		"aperation":"app",
	})
	mongo,err := mongodb.NewMongoDB(cfg)
	if err != nil {
		logCtx.Error(err,err.Error())
		os.Exit(1)
	}

	mongoIml:= productuserepository.NewProductUsageRepository(mongo)

	serviceIml := productservice.NewProductService(logger,mongoIml)

	service := productusecase.NewProductUsage(serviceIml)

	grpcService := grpcapp.NewApp(cfg.GRPCPort,service,logger)

	return &App{
		GRPCServer: grpcService,
	}

}