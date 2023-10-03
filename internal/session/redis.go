package session

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/model"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

var (
	errRedisSet     = errors.New("set key and value in redis")
	errJsonMarshall = errors.New("json encoding error ")
)

type Redis struct {
	client *redis.Client
}

type Session interface {
	CreateToken(token model.TokenDetails, expireDuration time.Duration) error
	DeleteTokenByUserID(userID int) error
}

func RedisConnect(redisCfg *config.RedisConfiguration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err == nil {
		log.Printf("%s \n", "Redis Connection Succesful \n")
	}
	return client, err
}

func New(r *redis.Client) *Redis {
	return &Redis{
		client: r,
	}
}

func (r *Redis) CreateToken(token model.TokenDetails, expireDuration time.Duration) error {
	tokenBytes, err := json.Marshal(&token)
	if err != nil {
		return err
	}

	if err = r.client.Set(context.Background(), strconv.Itoa(token.UserID), tokenBytes, expireDuration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Redis) DeleteTokenByUserID(userID int) error {
	if err := r.client.Del(context.Background(), strconv.Itoa(userID)).Err(); err != nil {
		return err
	}

	return nil
}
