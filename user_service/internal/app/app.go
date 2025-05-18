package app

import (
	"log"
	grpcapp "user_service/internal/app/grpc"
	"user_service/internal/config"
	"user_service/internal/infastructure/postgres"
	postgresuser "user_service/internal/infastructure/repository/postgresql/user"
	userRedis "user_service/internal/infastructure/repository/redis/user"
	userservice "user_service/internal/services/user"
	usecaseuser "user_service/internal/usecase/user"
	logger "user_service/log"
)


type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(cfg *config.Config,logger *logger.Loggger) *App{
	postgres,err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	kash,err := userRedis.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := postgresuser.NewUserRepository(postgres)

	userCaseRepoIml := usecaseuser.NewUserRepoUseCase(db)

	service := userservice.NewUserService(userCaseRepoIml,logger,kash)

	userServiceIml := usecaseuser.NewUserUseCase(service)

	server := grpcapp.NewApp(cfg.GRPCPort,*userServiceIml,logger)
	
	return &App{
		GRPCServer: server,
	}

}