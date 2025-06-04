package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client  *redis.Client
	Context context.Context
}

func NewCache() *Cache {
	ctx := context.Background()
	return &Cache{
		Client:  client,
		Context: ctx,
	}
}

// type RedisClient struct {
// 	Client  *redis.Client
// 	Context context.Context
// }

// func NewRedisClient(config configs.RedisInfo) *RedisClient {
// 	ctx := context.Background()
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     config.Host + ":" + config.Port,
// 		Password: config.Password,
// 		DB:       config.Db,
// 	})
// 	return &RedisClient{
// 		Client:  client,
// 		Context: ctx,
// 	}
// }

// func (r *RedisClient) Ping() (string, error) {
// 	return r.Client.Ping(r.Context).Result()
// }

// func (r *RedisClient) Close() error {
// 	return r.Client.Close()
// }

func (r *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(r.Context, key, value, expiration).Err()
}

func (r *Cache) Get(key string) (string, error) {
	return r.Client.Get(r.Context, key).Result()
}

func (r *Cache) Del(key string) error {
	return r.Client.Del(r.Context, key).Err()
}
