package redis

import (
	"api/configs"
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

type Client struct {
	*redis.Client
	Context context.Context
}

func Setup() {
	config := configs.App.RedisInfo
	client = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.Database,
	})

	// Test the connection
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic("Redis连接失败: " + err.Error())
	}
}

func Close() {
	if client != nil {
		if err := client.Close(); err != nil {
			panic("Redis关闭失败: " + err.Error())
		}
	}
}

func GetClient() *redis.Client {
	if client == nil {
		panic("Redis client is not initialized")
	}
	return client
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

// func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
// 	return r.Client.Set(r.Context, key, value, expiration).Err()
// }

// func (r *RedisClient) Get(key string) (string, error) {
// 	return r.Client.Get(r.Context, key).Result()
// }

// func (r *RedisClient) Del(key string) error {
// 	return r.Client.Del(r.Context, key).Err()
// }
