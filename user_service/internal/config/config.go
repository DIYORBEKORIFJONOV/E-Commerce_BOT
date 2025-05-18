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
	}
	RedisUrl string
	Token struct {
		Secret    string
		AccessTTL time.Duration
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "D1YOR")
	return c.Token.Secret
}

func NewConfig() (*Config, error) {
	var config Config

	config.APP = getEnv("APP", "USER_SERVICE")
	config.GRPCPort = getEnv("RPC_PORT", "8083")

	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "+_+diyor2005+_+")
	config.DB.SslMode = getEnv("POSTGRES_SSLMODE", "disable")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "chat_service_user")

	config.Token.Secret = getEnv("TOKEN_SECRET", "D1YOR")
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil, err
	}
	config.Token.AccessTTL = accessTTl
	config.RedisUrl = getEnv("REDIS_URL","localhost:6379")
	return &config, nil
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
