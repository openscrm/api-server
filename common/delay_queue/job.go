package delay_queue

import (
	"context"
	"encoding/json"
	"openscrm/app/constants"
	"openscrm/common/util"
)

type Job struct {
	Topic       constants.Topic `json:"topic"`
	ID          string          `json:"id"`           // job唯一标识ID
	ExecuteAt   int64           `json:"execute_at"`   // 预定执行时间
	TTR         int64           `json:"ttr"`          // 轮询间隔
	FailedCount int64           `json:"failed_count"` // 失败次数
	Body        string          `json:"body"`
}

// 获取Job
func getJob(key string) (job Job, err error) {
	value, err := Rdb.Get(context.TODO(), key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(value), &job)
	if err != nil {
		return
	}

	return
}

// 添加Job
func putJob(key string, job Job) error {
	err := Rdb.Do(context.TODO(), "set", key, util.JsonEncode(job)).Err()
	return err
}

// 更新Job
func setJob(key string, job Job) error {
	err := Rdb.Do(context.TODO(), "setnx", key, util.JsonEncode(job)).Err()
	return err
}

// 删除Job
func removeJob(key string) error {
	return Rdb.Del(context.Background(), key).Err()
}
