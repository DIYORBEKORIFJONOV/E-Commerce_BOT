package userRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user_service/internal/config"
	entityuser "user_service/internal/entity/user"

	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
	redisClient *redis.Client
}


func NewRedis(cfg *config.Config) (*RedisUserRepository,error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisUrl,
		Password: "",
		DB: 0,
	})
	_,err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil,err
	}

	return &RedisUserRepository{
		redisClient: client,
	},nil
}

func (r *RedisUserRepository)SaveToKash(ctx context.Context ,user *entityuser.User) error {
	userJson,err :=json.Marshal(user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf(":%s:",user.UserID)
	err  = r.redisClient.Set(ctx,key,string(userJson),time.Hour).Err()
	if err != nil {
		return  err
	}

	return nil
}


func (r *RedisUserRepository)GetFromKash(ctx context.Context,userId string)(*entityuser.User,error) {
	key := fmt.Sprintf(":%s:",userId)

	val,err := r.redisClient.Get(ctx,key).Result()
	if err != nil {
		if err ==  redis.Nil {
			return nil,nil 
		}
		return nil,err
	}

	var user entityuser.User
	err = json.Unmarshal([]byte(val),&user)
	if err != nil {
		return nil,err
	}

	return &user,nil
}

func (r *RedisUserRepository)DeleteFromKash(ctx context.Context, userId string)error {
	key := fmt.Sprintf(":%s:",userId)

	err := r.redisClient.Del(ctx,key).Err()
	if err != nil {
		return err
	}
	return nil
}