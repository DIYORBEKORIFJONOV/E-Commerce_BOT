package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/app"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/logger"
	"github.com/joho/godotenv"
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

func main() {

	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("❌ Ошибка получения текущего пути: %v", err)
	}

	envPath := filepath.Join(rootPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("⚠️ .env файл не найден по пути %s, продолжаем с переменными окружения", envPath)
	}

	log.Println("DEBUG: REPLY_TIMEOUT =", os.Getenv("REPLY_TIMEOUT"))

	// Конфигурация
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("❌ Ошибка инициализации конфигурации: %v", err)
	}

	lg := logger.SetupLogger("local")

	application := app.NewApp(lg, cfg)

	go application.HTTPApp.Start()

	select {}
}
