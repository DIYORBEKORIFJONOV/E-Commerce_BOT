package config

import (
	"os"

)

type Config struct {
	AppPort string
	OrderServicePort string
	ProductServicePort string
	UserServicePort string
}



func NewConfig() (*Config, error) {
	var config Config
	config.AppPort =getEnv("APP_PORT","")
	config.OrderServicePort = getEnv("ORDER_PORT","")
	config.ProductServicePort = getEnv("PRODUCT_ORDER","")
	config.UserServicePort = getEnv("USER_SERVICE","")
	return &config, nil
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
