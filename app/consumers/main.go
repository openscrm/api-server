package consumers

import (
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"math"
	"openscrm/app/constants"
	"openscrm/common/delay_queue"
	"openscrm/common/log"
	"runtime"
	"sync"
	"time"
)

var _handlers = make(map[constants.Topic]func(delay_queue.Job) error)
var _locker = sync.Mutex{}
var _stopSignal = make(chan interface{}, 1)

func registerHandler(topic constants.Topic, handler func(delay_queue.Job) error) {
	_locker.Lock()
	defer _locker.Unlock()
	_handlers[topic] = handler
}

func registerHandlers() {
	registerHandler(constants.RefreshContactWayTopic, contactWay{}.Refresh)
	registerHandler(constants.MassMsgTopic, SendMassMsg)
	registerHandler(constants.SyncCustomerDataTopic, SyncCustomerData)
	registerHandler(constants.RemainderTopic, SendRemainderMsg)
	registerHandler(constants.GroupChatMassMsgTopic, SendGroupChatMassMsg)
	dataExporter := NewDataExporter()
	registerHandler(constants.DataExportTopic, dataExporter.DataExport)

}

func Start() {
	registerHandlers()
	for topic, handler := range _handlers {
		go ProtectedRun(topic, handler, Consume)
	}
}

func Stop() {
	_stopSignal <- 1
}

func Consume(topic constants.Topic, handler func(delay_queue.Job) error) {
	for {
		select {
		case <-_stopSignal:
			log.Sugar.Info("stop consume")
		default:
			job, err := delay_queue.Listen(topic)
			if err == redis.Nil {
				continue
			}
			if err != nil {
				log.TracedError("delay_queue.Listen failed", errors.WithStack(err))
				continue
			}
			// 没有任务，redis阻塞超时
			if job.ID == "" {
				continue
			}

			err = handler(job)
			if err != nil {
				log.TracedError("handle job failed", errors.WithStack(err))
				job.FailedCount++
				// 任务失败时，等待时间指数级增长，最大15分钟间隔
				delay := time.Second * time.Duration(math.Pow(2, float64(job.FailedCount)))
				if delay > time.Minute*15 {
					delay = time.Minute * 15
				}
				job.ExecuteAt = time.Now().Add(delay).Unix()
				err = delay_queue.Update(job)
				if err != nil {
					log.TracedError("delay_queue.Update failed", errors.WithStack(err))
				}
				continue
			}

			err = delay_queue.Remove(job.ID)
			if err != nil {
				log.TracedError("delay_queue.Remove failed", errors.WithStack(err))
			}

		}
	}
}

func ProtectedRun(topic constants.Topic, handler func(delay_queue.Job) error, fn func(topic constants.Topic, handler func(delay_queue.Job) error)) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			log.Sugar.Error("runtime error:", err)
		default: // 非运行时错误
			log.Sugar.Error("error:", err)
		}
	}()
	fn(topic, handler)
}
