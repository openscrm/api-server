package tasks

import (
	"context"
	"openscrm/common/log"
	"openscrm/common/redis"
	"time"
)

type Base struct {
}

// Lock 获取分布式锁
func (o Base) Lock(key string, ttl time.Duration) (bool, error) {
	return redis.RedisClient.SetNX(context.Background(), key, time.Now().Unix(), ttl).Result()
}

// Unlock 解除分布式锁
func (o Base) Unlock(key string) {
	err := redis.RedisClient.Del(context.Background(), key).Err()
	if err != nil {
		log.Sugar.Errorw("Unlock failed", "err", err)
	}
}
