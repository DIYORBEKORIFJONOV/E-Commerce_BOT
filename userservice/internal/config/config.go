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
	c.Database.CollectionOrders= osGetenv("DB_C1", "useraccounts")
	c.Database.DBname = osGetenv("DB_NAME", "account")

	c.User.Host = osGetenv("USER_HOST", "tcp")
	c.User.Port = osGetenv("USER_PORT", "account-service:8083")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
