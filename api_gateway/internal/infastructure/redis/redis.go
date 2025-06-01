package redisCash

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/core"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	"github.com/go-redis/redis/v8"
)

type Service interface {
	SaveAccount(ctx context.Context, account *SaveAccount) error
	GetAccount(ctx context.Context, phone string) (account Account, err error)
}

type redisCash struct {
	redisClient *redis.Client
}

func NewRedis(cfg config.Config) Service {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &redisCash{
		redisClient: client,
	}
}

func (r *redisCash) SaveAccount(ctx context.Context, account *SaveAccount) (err error) {
	key := fmt.Sprintf("%s%s", core.RegisterAccountSubKey, account.Phone)
	accountMarshalled, err := json.Marshal(account)
	if err != nil {
		return
	}
	err = r.redisClient.Set(ctx, key, accountMarshalled, core.RegisterAccountTTl).Err()
	return
}

func (r *redisCash) GetAccount(ctx context.Context, phone string) (account Account, err error) {
	key := fmt.Sprintf("%s%s", core.RegisterAccountSubKey, phone)
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return
	}
	var saveAccount *SaveAccount
	err = json.Unmarshal([]byte(val), &saveAccount)
	if err != nil {
		return
	}
	account.Username = saveAccount.Username
	account.Password = saveAccount.Password
	account.Phone = phone
	account.Name = saveAccount.Name
	account.SenderCode = saveAccount.SenderCode
	return
}
