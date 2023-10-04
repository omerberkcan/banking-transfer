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

// type Sessions interface {
// 	JWTSession() JWTSession
// }

// type jwtSession struct {
// 	client *redis.Client
// }

// func NewJWTSessions(client *redis.Client) *jwtSession {
// 	return &jwtSession{client: client}
// }

type Redis struct {
	client *redis.Client
}

type JWTSession interface {
	SetToken(token model.TokenDetails, expireDuration time.Duration) error
	DeleteTokenByUserID(userID int) error
	FindTokenByUserID(userID int) (model.TokenDetails, error)
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
	} else {
		log.Printf("cannot connect redis %s ", redisCfg.Addr)
	}
	return client, err
}

func New(r *redis.Client) *Redis {
	return &Redis{
		client: r,
	}
}

func (r *Redis) SetToken(token model.TokenDetails, expireDuration time.Duration) error {
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

func (r *Redis) FindTokenByUserID(userID int) (model.TokenDetails, error) {
	var token model.TokenDetails

	val, err := r.client.Get(context.Background(), strconv.Itoa(userID)).Result()
	if err != nil {
		return token, err
	}

	err = json.Unmarshal([]byte(val), &token)
	if err != nil {
		return token, err
	}

	return token, nil
}
