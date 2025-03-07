package myRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisI interface {
	Get(ctx context.Context, key string, val interface{}) error
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Del(ctx context.Context, key ...string) error
}

func MyReds(r redis.UniversalClient) (RedisI, error) {
	return &myRedis{
		client: r,
	}, nil
}

type myRedis struct {
	client redis.UniversalClient
}

// Get 将redis中的值取出，并用val来接收，返回函数的执行情况
func (m *myRedis) Get(ctx context.Context, key string, val interface{}) error {
	return m.client.Get(ctx, key).Scan(val)
}

// Set 将新的信息存入redis，并设置一个到期时间
func (m *myRedis) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return m.client.Set(ctx, key, val, expiration).Err()
}

// Del 根据key删除信息
func (m *myRedis) Del(ctx context.Context, key ...string) error {
	return m.client.Del(ctx, key...).Err()
}
