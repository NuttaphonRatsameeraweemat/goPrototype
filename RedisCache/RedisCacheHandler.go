package rediscache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	env "redis-cache-api/EnvironmentVariable"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     env.GetRedisURL(),
	Password: env.GetRedisPassword(), // no password set
	DB:       0,                      // use default DB
})

// SetValue func()
func SetValue(key string, value interface{}) {
	// encode struct to []byte  && set []byte to redis
	val, errVal := json.Marshal(value)
	if errVal != nil {
		panic(errVal)
	}
	err := rdb.Set(ctx, key, val, getDuration()).Err()
	if err != nil {
		panic(err)
	}
}

// GetValue func(key string)
func GetValue(key string) (interface{}, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println("err ", err)
	}
	return val, err
}

// ClearValue func(key string)
func ClearValue(key string) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Println("err ", err)
	}
}

func getDuration() time.Duration {
	t := time.Now().AddDate(0, 0, 1)
	expireTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return expireTime.Sub(time.Now())
}
