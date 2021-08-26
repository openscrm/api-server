package delay_queue

import (
	"context"
	"fmt"
	"openscrm/conf"
	"time"
)

// 添加JobId到队列中
func pushToReadyQueue(queueName string, jobId string) error {
	queueName = fmt.Sprintf(conf.Settings.DelayQueue.QueueName, queueName)
	return Rdb.RPush(context.TODO(), queueName, jobId).Err()
}

// 从队列中阻塞获取JobId
func blockPopFromReadyQueue(queues []string, timeout int) (string, error) {
	var args []string
	for _, queue := range queues {
		queue = fmt.Sprintf(conf.Settings.DelayQueue.QueueName, queue)
		args = append(args, queue)
	}
	value, err := Rdb.BLPop(context.Background(), time.Duration(timeout)*time.Second, args...).Result()
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", nil
	}
	if len(value) == 0 {
		return "", nil
	}
	element := value[1]

	return element, nil
}
