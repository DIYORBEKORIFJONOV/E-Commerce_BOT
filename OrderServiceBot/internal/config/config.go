package config

import "os"

type Config struct {
	Database struct {
		CollectionOrders string
		CollectionCart string
		DBname   string
	}
	User struct {
		Host string
		Port string
	}
}

func Configuration() *Config {
	c := &Config{}
	c.Database.CollectionCart= osGetenv("DB_C2", "carts")
	c.Database.CollectionOrders= osGetenv("DB_C1", "order")
	c.Database.DBname = osGetenv("DB_NAME", "orders")

	c.User.Host = osGetenv("USER_HOST", "tcp")
	c.User.Port = osGetenv("USER_PORT", ":8080")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
