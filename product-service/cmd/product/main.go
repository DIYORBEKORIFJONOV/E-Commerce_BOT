package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"productservice/internal/app"
	"productservice/internal/config"
	logger "productservice/log"
	"syscall"

	"github.com/joho/godotenv"
)


func main() {
	logg, _, err := logger.New(logger.Config{
		LogFilePath: "service.log",
		ServiceName: "payment-service",
	})
	if err != nil {
		log.Fatal(err)
	}

	logCtx := logg.WithContext(map[string]string{
		"operation": "product-service",
	})

	err = godotenv.Load()
    if err != nil {
        logCtx.Error(err,"⚠️ .env not found, using system environment variables")
    }


	cfg,err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	application  := app.NewApp(cfg,logg)
	logCtx.Info("Product server started")
	go application.GRPCServer.Run()
	stop := make(chan os.Signal,1)
	signal.Notify(stop,syscall.SIGTERM,syscall.SIGINT)
	sig  := <-stop

	
	logCtx.Warn(fmt.Sprintf("received shutdown signal %s",sig.String()))
	logCtx.Info("shuttin down server")
}