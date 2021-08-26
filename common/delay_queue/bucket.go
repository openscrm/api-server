package delay_queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// BucketItem bucket中的元素
type BucketItem struct {
	timestamp int64
	jobID     string
}

// 添加JobId到bucket中
func pushToBucket(key string, timestamp int64, jobId string) error {
	z := redis.Z{
		Score:  float64(timestamp),
		Member: jobId,
	}
	return Rdb.ZAdd(context.TODO(), key, &z).Err()
}

// 从bucket中获取延迟时间最小的JobId
func getFromBucket(key string) (*BucketItem, error) {
	value, err := Rdb.ZRangeWithScores(context.Background(), key, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	if value == nil || len(value) == 0 {
		return nil, nil
	}

	item := &BucketItem{}
	item.timestamp = int64(value[0].Score)
	item.jobID = (value[0].Member).(string)
	return item, nil
}

// 从bucket中删除JobId
func removeFromBucket(bucket string, jobId string) error {
	return Rdb.ZRem(context.TODO(), bucket, jobId).Err()
}
