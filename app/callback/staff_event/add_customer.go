package staff_event

import (
	"encoding/json"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/services"
	"openscrm/common/delay_queue"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	gowx "openscrm/pkg/easywework"
	"time"
)

// EventAddExternalContactHandler
// Description: 添加客户回调事件处理
// Detail:
// 	1. 同步客户数据
//	2. 发送欢迎语
//	3. 更新客户数
// Param: msg  wx回调参数
func EventAddExternalContactHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalContact ||
		msg.ChangeType != gowx.ChangeTypeAddExternalContact {
		err := errors.New("wrong handler for the callback event")
		log.Sugar.Error("err", err)
		return err
	}
	eventAddExternalContact, ok := msg.EventAddExternalContact()
	if !ok {
		err := errors.New("msg.EventEditExternalContact failed")
		log.Sugar.Errorw("get event msg failed", "err", err)
		return err
	}
	extStaffID := eventAddExternalContact.GetUserID()
	extCustomerID := eventAddExternalContact.GetExternalUserID()
	welcomeCode := eventAddExternalContact.GetWelcomeCode()
	log.Sugar.Debugw("call back biz info:",
		"customerID", extCustomerID, "staffID", extStaffID, "welcomeCode", welcomeCode)

	err := SyncExternalContact(extStaffID, extCustomerID)
	if err != nil {
		log.Sugar.Errorw("add sync job failed", "err", err)
		return err
	}

	shouldSendWelcomeMsg, err := (&services.ContactWay{}).DealAddCustomerEvent(models.DB, eventAddExternalContact)
	if err != nil {
		log.Sugar.Errorw("ContactWay.DealAddCustomerEvent failed",
			"err", err, "eventAddExternalContact", eventAddExternalContact)
		return err
	}

	// 渠道码已发过欢迎语，这里就不再发了
	if shouldSendWelcomeMsg {
		welcomeCode := eventAddExternalContact.GetWelcomeCode()
		if welcomeCode != "" {
			log.Sugar.Infow("send default welcome msg")
			staff := &models.Staff{}
			err := models.DB.Model(&models.Staff{}).Where("ext_id = ?", extStaffID).First(&staff).Error
			if err != nil {
				log.Sugar.Infow("get staff info failed", "err", err)
				return err
			}

			// send default welcome msg
			staffService := services.NewStaffService()
			return staffService.SendDefaultWelcomeMsg(welcomeCode, staff.ExtCorpID, staff.ExtID)
		}
	}

	return nil
}

// SyncExternalContact
// Description: 生成同步客户数据的异步任务
// Detail: 异步处理，出错后可重试
// Param: extStaffID 员工外部ID
// Param: extCustomerID 客户外部ID
func SyncExternalContact(extStaffID, extCustomerID string) error {
	customerStaffRelation := models.CustomerStaffRelation{
		ExtStaffID:    extStaffID,
		ExtCustomerID: extCustomerID,
	}
	relation, err := json.Marshal(customerStaffRelation)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	job := delay_queue.Job{
		Topic:     constants.SyncCustomerDataTopic,
		ID:        id_generator.StringID(),
		ExecuteAt: time.Now().Unix(),
		TTR:       10,
		Body:      string(relation),
	}
	return delay_queue.Add(job)
}
