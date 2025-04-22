package config

import (
	"os"
	"time"
)

type Config struct {
	APP      string
	GRPCPort string
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SslMode  string
		CollectionCategory string
		CollectionProductName string
		CollectionModel string
	}
	TelegramBotTokens struct {
		Client string
		Admin string
	}

	RedisURL string
	Token struct {
		Secret    string
		AccessTTL time.Duration
	}
	PayPal struct {
		PaypalClientID string
		PaypalClientSecret string 
		PaypalAPIBase string
		WebhookPort string
		PublicURL string
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "")
	return c.Token.Secret
}

func NewConfig() (*Config, error) {
	var config Config
	config.PayPal.PaypalClientID = getEnv("PAYPAL_CLIENTID","")
	config.PayPal.PaypalClientSecret = getEnv("PAYPAL_SECRET","")
	config.PayPal.PaypalAPIBase = getEnv("PAYPAL_API_BASE","")
	config.PayPal.PublicURL = getEnv("PAYPAL_PUBLIC_URL","")
	config.PayPal.WebhookPort = getEnv("PAYPAL_WEBHOOK","")
	
	config.APP = getEnv("APP", "USER-SERVICE")
	config.GRPCPort = getEnv("RPC_PORT", "")

	config.DB.Host = getEnv("MONGO_HOST", "")
	config.DB.Port = getEnv("MONGO_PORT", "")
	config.DB.Name = getEnv("MONGO_DATABASE", "")
	config.DB.User = getEnv("MONGO_USER", "")
	config.DB.Password = getEnv("MONGO_PASSWORD", "")
	config.DB.SslMode = getEnv("MONGO_SSLMODE", "")
	config.DB.CollectionModel = getEnv("DB_COLLECTION_MODEL", "")
	config.DB.CollectionCategory = getEnv("DB_COLLECTION_CATEGORY","")
	config.DB.CollectionProductName = getEnv("DB_COLLECTION_PRODUCT_NAME","")

	config.RedisURL = getEnv("REDIS_URL", "")

	config.Token.Secret = getEnv("TOKEN_SECRET", "")
	accessTTL, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", ""))
	if err != nil {
		return nil, err
	}
	config.Token.AccessTTL = accessTTL

	config.TelegramBotTokens.Admin = getEnv("TELEGRAM_BOT_TOKEN_ADMIN", "")
    config.TelegramBotTokens.Client = getEnv("TELEGRAM_BOT_TOKEN_CLIENT", "")


	return &config, nil
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
