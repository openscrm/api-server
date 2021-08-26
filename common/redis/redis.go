package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
	"time"
)

var RedisClient *redis.Client

func Setup(host string, pw string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pw,
		DB:       db,
	})
}

// GetOrSetFunc 获取或设置缓存
// result 接收反序列化的值
func GetOrSetFunc(key string, f func() (interface{}, error), duration time.Duration, result interface{}) error {
	jsonData := RedisClient.Get(context.Background(), key).Val()
	if jsonData == "" {
		value, err := f()
		if err != nil {
			return err
		}
		if value == nil {
			return nil
		}

		err = copier.Copy(result, value)
		if err != nil {
			return err
		}

		jsonData, err = jsoniter.MarshalToString(value)
		if err != nil {
			return err
		}
		return RedisClient.Set(context.Background(), key, jsonData, duration).Err()
	}

	return jsoniter.UnmarshalFromString(jsonData, &result)
}

func Delete(keys ...string) error {
	return RedisClient.Del(context.Background(), keys...).Err()
}
