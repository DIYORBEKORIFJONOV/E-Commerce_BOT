package main

import (
	"api_gateway/internal/app"
	"api_gateway/internal/config"
	"api_gateway/logger"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

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

	// Логгер
	lg := logger.SetupLogger("local")

	// Приложение
	application := app.NewApp(lg, cfg)

	// HTTP-сервер
	go application.HTTPApp.Start()

	select {} // бесконечное ожидание
}
