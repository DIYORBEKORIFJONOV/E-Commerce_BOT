package grpcapp

import (
	"fmt"
	"net"
	userserver "user_service/internal/gprc/user"
	usecaseuser "user_service/internal/usecase/user"
	logger "user_service/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type App struct {
	grpcServer *grpc.Server
	port string
	logger *logger.Loggger 
}

func NewApp(port string, user usecaseuser.UserUseCaseIml,log *logger.Loggger ) *App {
	grpcServer := grpc.NewServer()
	userserver.RegisterUserServer(grpcServer,user)
	reflection.Register(grpcServer)
	return &App{
		grpcServer: grpcServer,
		port: port,
		logger: log,
	}
}

func (a *App)Run() error {
	l,err := net.Listen("tcp",fmt.Sprintf(":%s",a.port))
	if err != nil {
		return err
	}
	a.logger.Console.Info().
	Str("HOST","localhost").
	Str("PORT",a.port).
	Msg("starting gRPC server on ")
	err = a.grpcServer.Serve(l)
	if err != nil {
		return err
	}
	
	return nil
}

func (a *App)Stop() {
	a.logger.Console.Info().
	Str("HOST","localhost").
	Str("PORT",a.port).
	Msg("stopping gRPC server on")
	a.grpcServer.GracefulStop()
}