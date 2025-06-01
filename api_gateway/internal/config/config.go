package config

import (
	"os"
	"time"
)

type Config struct {
	AppPort            string
	OrderServicePort   string
	RedisURL           string
	ProductServicePort string
	AccountServicePort string
	EskizEmail         string
	EskizPassword      string
	EskizSenderId      string
	Token              struct {
		Secret     string
		AccessTTL  time.Duration
		RefreshTTL time.Duration
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "")
	return c.Token.Secret
}

func NewConfig() (*Config, error) {
	var config Config
	config.AppPort = getEnv("APP_PORT", "")
	config.OrderServicePort = getEnv("ORDER_PORT", "")
	config.ProductServicePort = getEnv("PRODUCT_ORDER", "")
	config.AccountServicePort = getEnv("ACCOUNT_SERVICE", "")
	config.Token.Secret = getEnv("TOKEN_SECRET", "D1YORTOP4EEK")

	refreshTTL, err := time.ParseDuration(getEnv("TOKEN_REFRESH_TTL", "24h"))
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil, err
	}
	config.Token.AccessTTL = accessTTl
	config.Token.RefreshTTL = refreshTTL

	config.EskizEmail = getEnv("ESKIZ_EMAIL", "")
	config.EskizPassword = getEnv("ESKIZ_PASSWORD", "")
	config.EskizSenderId = getEnv("ESKIZ_SENDER_ID", "")

	config.RedisURL = getEnv("REDIS_URL", "")
	return &config, nil
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
