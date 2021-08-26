package delay_queue

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
	"openscrm/app/constants"
	"openscrm/conf"
	"time"
)

var (
	// 每个定时器对应一个bucket
	timers []*time.Ticker
	// bucket名称chan
	bucketNameChan <-chan string
)

// SetupDelayQueue 初始化延时队列
func SetupDelayQueue() {
	NewRedisClient()
	initTimers()
	bucketNameChan = generateBucketName()
}

// Get 查询Job
func Get(jobID string) (job Job, err error) {
	job, err = getJob(jobID)
	if err != nil {
		return
	}

	// 消息不存在, 可能已被删除
	if job.ID == "" {
		return
	}

	return
}

// Add 添加一个Job到队列中
func Add(job Job) error {
	if job.ID == "" || job.Topic == "" || job.ExecuteAt < 0 || job.TTR <= 0 {
		return errors.New("invalid job")
	}

	err := putJob(job.ID, job)
	if err != nil {
		//log.Printf("添加job到job pool失败# putJob job-%+v#%s", job, err.Error())
		return err
	}
	err = pushToBucket(<-bucketNameChan, job.ExecuteAt, job.ID)
	if err != nil {
		//log.Printf("添加job到bucket失败# pushToBucket job-%+v#%s", job, err.Error())
		return err
	}

	return nil
}

// Update 更新一个Job
func Update(job Job) (err error) {
	if job.ID == "" || job.Topic == "" || job.ExecuteAt < 0 || job.TTR <= 0 {
		return errors.New("invalid job")
	}

	err = Remove(job.ID)
	if err != nil {
		err = errors.Wrap(err, "Remove job failed")
		return err
	}

	err = Add(job)
	if err != nil {
		err = errors.Wrap(err, "Remove job failed")
		return err
	}

	return
}

// Listen 轮询获取Job
func Listen(topics ...constants.Topic) (job Job, err error) {
	jobID, err := blockPopFromReadyQueue(topicsToStrings(topics), conf.Settings.DelayQueue.QueueBlockTimeout)
	if err != nil {
		return
	}

	// 队列为空
	if jobID == "" {
		return
	}

	// 获取job元信息
	job, err = getJob(jobID)
	if err != nil {
		return
	}

	// 消息不存在, 可能已被删除
	if job.ID == "" {
		return
	}

	timestamp := time.Now().In(constants.PRCLocation).Unix() + job.TTR
	err = pushToBucket(<-bucketNameChan, timestamp, job.ID)

	return job, err
}

// Remove 删除Job
func Remove(jobID string) error {
	return removeJob(jobID)
}

// 轮询获取bucket名称, 使job分布到不同bucket中, 提高扫描速度
func generateBucketName() <-chan string {
	c := make(chan string)
	go func() {
		i := 1
		for {
			c <- fmt.Sprintf(conf.Settings.DelayQueue.BucketName, i)
			if i >= conf.Settings.DelayQueue.BucketSize {
				i = 1
			} else {
				i++
			}
		}
	}()

	return c
}

// 初始化定时器
func initTimers() {
	timers = make([]*time.Ticker, conf.Settings.DelayQueue.BucketSize)
	var bucketName string
	for i := 0; i < conf.Settings.DelayQueue.BucketSize; i++ {
		timers[i] = time.NewTicker(2 * time.Second)
		bucketName = fmt.Sprintf(conf.Settings.DelayQueue.BucketName, i+1)
		go waitTicker(timers[i], bucketName)
	}
}

func waitTicker(timer *time.Ticker, bucketName string) {
	for {
		select {
		case t := <-timer.C:
			tickHandler(t, bucketName)
		}
	}
}

// 扫描bucket, 取出延迟时间小于当前时间的Job
func tickHandler(t time.Time, bucketName string) {
	t = t.In(constants.PRCLocation)
	for {
		bucketItem, err := getFromBucket(bucketName)
		if err != nil {
			log.Printf("扫描bucket错误#bucket-%s#%s", bucketName, err.Error())
			return
		}

		// 集合为空
		if bucketItem == nil {
			return
		}

		// 延迟时间未到
		if bucketItem.timestamp > t.Unix() {
			//log.Printf("%s not now,expected timestamp %d, now %d", bucketItem.jobID, bucketItem.timestamp, t)
			return
		}

		// 延迟时间小于等于当前时间, 取出Job元信息并放入ready queue
		job, err := getJob(bucketItem.jobID)
		if err != nil && err != redis.Nil {
			log.Printf("获取Job元信息失败#jobID%s#bucket-%s#%s", bucketItem.jobID, bucketName, err.Error())
			continue
		}

		// job元信息不存在, 从bucket中删除
		if err == redis.Nil || job.ID == "" {
			removeFromBucket(bucketName, bucketItem.jobID)
			continue
		}

		// 再次确认元信息中delay是否小于等于当前时间
		if job.ExecuteAt > t.Unix() {
			// 从bucket中删除旧的jobID
			removeFromBucket(bucketName, bucketItem.jobID)
			// 重新计算delay时间并放入bucket中
			pushToBucket(<-bucketNameChan, job.ExecuteAt, bucketItem.jobID)
			continue
		}

		err = pushToReadyQueue(string(job.Topic), bucketItem.jobID)
		if err != nil {
			log.Printf("jobID放入ready queue失败#bucket-%s#job-%+v#%s",
				bucketName, job, err.Error())
			continue
		}

		// 从bucket中删除
		removeFromBucket(bucketName, bucketItem.jobID)
	}
}
