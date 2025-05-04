package htppapp

import (
	"log/slog"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/http/router"
	minao1 "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/minao"
	usecaseorder "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/order"
	productusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, botIml *usecaseorder.OrderUseCaseIml,minioPhoto *minao1.FileStorage,productUsceIml *productusecase.ProductUseCaseIml) *App {
	sever := router.RegisterRouter(botIml,productUsceIml,minioPhoto)
	return &App{
		Port:   port,
		Server: sever,
		Logger: logger,
	}
}

func (app *App) Start() {
	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.SetTrustedProxies(nil)
	if err != nil {
		log.Error("Error setting trusted proxies", "error", err)
		return
	}
	err = app.Server.Run(app.Port)
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
