package services

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/delay_queue"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	gowx "openscrm/pkg/easywework"
)

type Remainder struct {
	reminderRepo      models.Remainder
	customerEventRepo models.CustomerEvent
}

func NewRemainder() *Remainder {
	return &Remainder{
		reminderRepo:      models.Remainder{},
		customerEventRepo: models.CustomerEvent{},
	}
}

// Create todo 添加到日历
func (o Remainder) Create(req requests.CreateRemainderReq, extCorpID string, extStaffID string) (models.CustomerEvent, error) {

	ce := models.CustomerEvent{
		ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
		ExtStaffID:    req.ExtStaffID,
		EventType:     string(constants.CustomerEventRemainder),
		Content:       req.Content,
		SendAt:        req.SendAt,
		ExtCustomerID: req.ExtCustomerID,
	}

	log.Sugar.Debugw("event", "ce", ce)

	err := copier.CopyWithOption(&ce, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return ce, err
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		log.Sugar.Errorw("marshal req failed", "req", req)
		return ce, err
	}

	job := delay_queue.Job{
		Topic:     constants.RemainderTopic,
		ID:        ce.ID,
		ExecuteAt: (req.SendAt).ToInt64(),
		TTR:       10,
		Body:      string(reqBytes),
	}
	err = delay_queue.Add(job)
	if err != nil {
		log.Sugar.Errorw("add remainder job failed", "err", err)
		return ce, err
	}

	err = o.customerEventRepo.Create(ce)
	if err != nil {
		log.Sugar.Errorw("create customer_event failed", "CustomerEvent", ce)
	}
	return ce, err
}

// Update 只更新内容
func (o Remainder) Update(id string, req requests.UpdateRemainderReq, extCorpID string) (models.CustomerEvent, error) {
	r := models.CustomerEvent{
		ExtCorpModel: models.ExtCorpModel{ID: id, ExtCorpID: extCorpID},
		Content:      req.Content,
	}
	res, err := o.customerEventRepo.Update(r)
	if err != nil {
		log.Sugar.Errorw("update remainder failed", "id", id)
		err = errors.WithStack(err)
	}
	return res, err
}

func (o Remainder) Delete(id string, extCorpID string) (int64, error) {
	err := delay_queue.Remove(id)
	if err != nil {
		log.Sugar.Errorw("remove remainder msg failed", "id", id)
		return 0, err
	}
	return o.customerEventRepo.Delete([]string{id}, extCorpID)
}

// SendRemainderMsg 执行提醒任务
func (o Remainder) SendRemainderMsg(job delay_queue.Job) error {
	remainder, err := o.customerEventRepo.Get(job.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Sugar.Errorw("get remainder failed")
			return nil
		}
		return errors.Wrap(err, "find remainder event failed")
	}

	client, err := we_work.Clients.Get(remainder.ExtCorpID)
	if err != nil {
		log.Sugar.Errorw("get remainder failed", "extCorpID", remainder.ExtCorpID)
		return err
	}
	recipient := &gowx.Recipient{UserIDs: []string{remainder.ExtStaffID}}
	req := requests.CreateRemainderReq{}
	err = json.Unmarshal([]byte(job.Body), &req)
	if err != nil {
		return err
	}
	content := fmt.Sprintf(constants.RemainderContent, req.ExtStaffID, req.CustomerName, remainder.Content)
	err = client.MainApp.SendTextMessage(recipient, content, false)
	if err != nil {
		log.Sugar.Errorw("send msg to staff failed", "err", err)
		return err
	}
	return nil
}
