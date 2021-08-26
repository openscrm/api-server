package consumers

import (
	"openscrm/app/constants"
	"openscrm/app/services"
	"openscrm/common/delay_queue"
	"openscrm/common/log"
)

func SendRemainderMsg(job delay_queue.Job) error {
	log.Sugar.Info("job info:", job)
	if job.Topic != constants.RemainderTopic {
		log.Sugar.Infow("job.topic not match", "job.Topic", job.Topic)
		return nil
	}
	remainder := services.NewRemainder()
	err := remainder.SendRemainderMsg(job)
	if err != nil {
		log.Sugar.Errorw("send remainder msg failed", "job.id", job.ID, "job.body", job.Body)
		return err
	}

	err = delay_queue.Remove(job.ID)
	if err != nil {
		log.Sugar.Errorw("send remainder msg failed", "job.id", job.ID, "job.body", job.Body)
		return err
	}
	return nil
}
