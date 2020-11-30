package redis

import (
	"errors"
	"time"
	//"github.com/gomodule/redigo/redis"
	"github.com/go-redis/redis"
	"newhope/config"
)

type RedisClient struct {
	c *redis.Client
}

var redisdb *RedisClient

func InitRedis() error {
	redisdb = new(RedisClient)
	redisdb.c = newLoginRedisClient(100)
	if redisdb.c == nil {
		return errors.New("init redis err")
	}
	return nil
}

func newLoginRedisClient(poolSize int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:         config.GetLoginRedisAddr(),
		Password:     config.GetLoginRedisPassword(),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     poolSize,
		PoolTimeout:  30 * time.Second,
	})
	_, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	return c
}

func Set(key string, value interface{}) string {
	err := redisdb.c.Set(key, value, 0).Err()
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func Get(key string) (string, error) {
	res, err := redisdb.c.Get(key).Result()
	return res, err
}

func Append(key string, value string) string {
	err := redisdb.c.Append(key, value).Err()
	if err != nil {
		return err.Error()
	}
	return "ok"
}
