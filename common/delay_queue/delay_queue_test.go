package delay_queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"openscrm/app/constants"
	setting "openscrm/conf"
	"testing"
	"time"
)

const topic = "test"
const delayTime = 3

func TestPush(t *testing.T) {
	setting.SetupSetting()
	NewRedisClient()
	SetupDelayQueue()
	job := Job{
		Topic:     topic,
		ID:        "hao123",
		ExecuteAt: delayTime + time.Now().In(constants.PRCLocation).Unix(),
		TTR:       3,
		Body:      "hi",
	}
	fmt.Println("b")
	//job.ExecuteAt = job.ExecuteAt + time.Now().Seconds()
	err := Add(job)
	fmt.Println("a")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)

	receivedJob, err := Listen(topic)
	if err != nil {
		log.Fatal(err)
	}
	assert.ObjectsAreEqual(job, receivedJob)
}
