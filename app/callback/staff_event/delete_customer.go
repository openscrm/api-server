package staff_event

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/delay_queue"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"time"
)

// EventDelExternalContactHandler 员工删除客户回调处理
func EventDelExternalContactHandler(msg *gowx.RxMessage) (err error) {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalContact ||
		msg.ChangeType != gowx.ChangeTypeDelExternalContact {
		return errors.New("wrong handler for the callback event")
	}

	eventDeleteExternalContact, ok := msg.EventDelExternalContact()
	if !ok {
		return errors.New("get EventDelExternalContact data failed")
	}
	extCustomerID := eventDeleteExternalContact.GetExternalUserID()
	extStaffID := eventDeleteExternalContact.GetUserID()
	extCorpID := msg.ToUserID
	log.Sugar.Debugw("receive staff delete customer msg", "extStaffID", extStaffID, "extCustomerID", extCustomerID)

	content, err := notifyAdmin(extStaffID, extCustomerID, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 记录事件流水
	err = models.CustomerEvent{}.Create(models.CustomerEvent{
		ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: conf.Settings.WeWork.ExtCorpID, ExtCreatorID: extStaffID},
		Content:       content,
		EventType:     constants.CustomerEventCustomerAction,
		EventName:     constants.EventNameDeleteExternalUser,
		ExtCustomerID: extCustomerID,
		ExtStaffID:    extStaffID,
	})
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 当前员工和客户不是好友关系，则不用再更新员工客户数。
	should, err := ShouldChangeCustomerNum(extStaffID, extCustomerID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	if should {
		// 更新员工客户数
		err = models.CustomerStatistic{}.Upsert(extStaffID, -1)
		if err != nil {
			log.Sugar.Errorw("update customer statistics failed",
				"extStaffID", extStaffID, "extCustomerID", extCustomerID)
			return err
		}
	}

	err = models.CustomerStaffRelationHistory{}.StaffDeleteCustomer(extStaffID, extCustomerID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return
}

// ShouldChangeCustomerNum 是否需要更新员工客户数
func ShouldChangeCustomerNum(extStaffID string, extCustomerID string) (bool, error) {
	_, err := models.CustomerStaff{}.GetCurrentCustomerStaffRelation(extStaffID, extCustomerID)
	if err == gorm.ErrRecordNotFound {
		// 已经不是好友关系
		return false, nil
	} else if err != nil {
		err = errors.WithStack(err)
		return false, err
	}
	return true, nil
}

// 通知管理员客户删除了员工
func notifyAdmin(extStaffID, extCustomerID, extCorpID string) (content string, err error) {
	eventNotifyRule, err := models.EventNotify{}.Get(extCorpID)
	if err != nil {
		log.Sugar.Errorw("get eventNotifyRule failed", "ext_corp_id", extCorpID, "err", err)
		return
	}
	if eventNotifyRule.IsNotifyAdmins == constants.EventNotifyStatusOff {
		return
	}

	staff, err := models.Staff{}.Get(extStaffID, extCorpID, false)
	if err != nil {
		log.Sugar.Errorw("get staff failed", "ext_corp_id", extCorpID, "err", err)
		return
	}

	customer, err := models.Customer{}.GetByExtID(extCustomerID, nil, false)
	if err != nil {
		log.Sugar.Errorw("get customer failed", "ext_corp_id", extCorpID, "err", err)
		return
	}

	content = fmt.Sprintf("员工 [%s] 删除了客户 [%s] ", staff.Name, customer.Name)
	if eventNotifyRule.NotifyType == constants.EventNotifyTimeRealTime {
		client, err := we_work.Clients.Get(extCorpID)
		if err != nil {
			log.Sugar.Errorw("get wx clients failed", "err", err)
			return content, err
		}
		recipient := &gowx.Recipient{UserIDs: eventNotifyRule.ExtStaffIDs}
		err = client.MainApp.SendTextMessage(recipient, content, false)
		if err != nil {
			log.Sugar.Errorw("client.MainApp.SendTextMessage", "err", err)
			return content, err
		}
	} else if eventNotifyRule.NotifyType == constants.EventNotifyTimeTimed {
		notifyAdminMsg := constants.TimedNotifyAdminMsg{
			ExtStaffID:    extStaffID,
			ExtCustomerID: extCustomerID,
			Content:       content,
			AdminIDs:      eventNotifyRule.ExtStaffIDs,
		}
		notifyAdminMsgBytes, err := json.Marshal(notifyAdminMsg)
		if err != nil {
			err = errors.WithStack(err)
			return content, err
		}

		// 每天八点通知
		executeAt := int64(util.Today().Second() + (24+8)*60*60)
		if time.Now().Hour() < 8 {
			executeAt = int64(util.Today().Second() + 8*3600)
		}

		job := delay_queue.Job{
			Topic:     constants.EventNameStaffDeleteCustomer,
			ID:        id_generator.StringID(),
			ExecuteAt: executeAt,
			TTR:       10,
			Body:      string(notifyAdminMsgBytes),
		}
		err = delay_queue.Add(job)
		if err != nil {
			err = errors.WithStack(err)
			return content, err
		}
	}
	return content, err
}
