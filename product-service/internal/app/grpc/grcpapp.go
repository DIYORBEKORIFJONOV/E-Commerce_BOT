package grpcapp

import (
	"net"
	productgrpc "productservice/internal/gprc/product"
	productusecase "productservice/internal/usecase/product/productservice"
	logger "productservice/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type App struct {
	grpcServer *grpc.Server
	port string
	logger *logger.Logger
}

func NewApp(port string, product *productusecase.ProductUseCaseIml,log *logger.Logger ) *App {
	grpcServer := grpc.NewServer()
	productgrpc.RegisterProductGrpcService(grpcServer,product)
	reflection.Register(grpcServer)
	return &App{
		grpcServer: grpcServer,
		port: port,
		logger: log,
	}
}

func (a *App)Run() error {
	l,err := net.Listen("tcp",a.port)
	if err != nil {
		return err
	}

	logCtx := a.logger.WithContext(map[string]string{
		"HOST":a.port,
	})
	logCtx.Info("Starting gRpc Server"	)
	err = a.grpcServer.Serve(l)
	if err != nil {
		return err
	}
	
	return nil
}

func (a *App)Stop() {
	logCtx := a.logger.WithContext(map[string]string{
		"HOST":a.port,
	})
	logCtx.Info("stopping gRpc server")
	a.grpcServer.GracefulStop()
}