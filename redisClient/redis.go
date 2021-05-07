package redisClient

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
 * Created by @CaomaoBoy on 2021/5/4.
 *  email:<115882934@qq.com>
 */

var Instance *redis.Client

const (
	REDIS_NETWORK  = "tcp"
	REDIS_HOST     = "127.0.0.1"
	REDIS_PORT     = "6379"
	REDIS_PASSWORD = "123456"
	REDIS_DB       = 0
)

func init() {
	options := redis.Options{
		Network:            REDIS_NETWORK,
		Addr:               fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           REDIS_PASSWORD,
		DB:                 REDIS_DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	}
	// 新建一个client
	Instance = redis.NewClient(&options)
	// close
	// defer Instance.Close()
}
