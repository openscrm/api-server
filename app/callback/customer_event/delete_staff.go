package customer_event

import (
	"fmt"
	"github.com/pkg/errors"
	"openscrm/app/callback/staff_event"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
)

// EventDelFollowUserHandler
// Description: 客户删除员工回调
// Detail:
//	通知员工
//	删除缓存
func EventDelFollowUserHandler(msg *gowx.RxMessage) (err error) {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeExternalContact ||
		msg.ChangeType != gowx.ChangeTypeDelFollowUser {
		return errors.New("wrong handler for the callback event")
	}

	eventDeleteExternalContact, ok := msg.EventDelFollowUser()
	if !ok {
		return errors.New("get EventDelFollowUser data failed")
	}
	extCustomerID := eventDeleteExternalContact.GetExternalUserID()
	extStaffID := eventDeleteExternalContact.GetUserID()

	// 通知员工
	eventNotifyRule, err := models.EventNotify{}.Get(msg.ToUserID)
	if err != nil {
		log.Sugar.Errorw("get eventNotifyRule failed", "err", err, "corpid", msg.ToUserID)
		return err
	}

	var customer models.Customer
	if eventNotifyRule.IsNotifyStaff == constants.EventNotifyStatusOn {
		client, err := we_work.Clients.Get(msg.ToUserID)
		if err != nil {
			log.Sugar.Errorw("get wx clients failed", "err", err)
			return err
		}
		customer, err = models.Customer{}.GetByExtID(extCustomerID, []string{extStaffID}, false)
		if err != nil {
			log.Sugar.Errorw("get wx clients failed", "err", err)
			return err

		}
		recipient := &gowx.Recipient{UserIDs: []string{extStaffID}}
		content := fmt.Sprintf("您已被客户 [%s] 删除", customer.Name)
		err = client.MainApp.SendTextMessage(recipient, content, false)
		if err != nil {
			log.Sugar.Errorw("client.MainApp.SendTextMessage", "err", err)
			return err
		}
	}

	extCorpID := conf.Settings.WeWork.ExtCorpID
	staff, err := models.Staff{}.Get(extStaffID, extCorpID, false)
	if err != nil {
		return err
	}
	err = CreateCustomerDeleteStaffEvent(customer, *staff)
	if err != nil {
		return err
	}
	// 删除员工客户关系, 记录关系流水
	err = models.CustomerStaffRelationHistory{}.CustomerDeleteStaff(extStaffID, extCustomerID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 当前员工和客户不是好友关系，则不用再更新员工客户数。
	should, err := staff_event.ShouldChangeCustomerNum(extStaffID, extCustomerID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if should {
		// 更新员工客户数
		err = models.CustomerStatistic{}.Upsert(extStaffID, -1)
		if err != nil {
			log.Sugar.Errorw("update customer statistics failed",
				"extStaffID", extStaffID, "extCustomerID", extCustomerID)
			return
		}
	}

	//删除首页的员工缓存数据
	err = models.Staff{}.CleanStaffSummaryCache(extStaffID, extCorpID)
	if err != nil {
		return
	}
	return
}

// CreateCustomerDeleteStaffEvent
// Description:  记录事件流水
func CreateCustomerDeleteStaffEvent(customer models.Customer, staff models.Staff) (err error) {
	content := fmt.Sprintf("员工 [%s] 删除了客户 [%s] ", staff.Name, customer.Name)

	err = models.CustomerEvent{}.Create(models.CustomerEvent{
		ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: conf.Settings.WeWork.ExtCorpID, ExtCreatorID: customer.ExtID},
		Content:       content,
		EventType:     constants.CustomerEventCustomerAction,
		EventName:     constants.EventNameDeleteExternalUser,
		ExtCustomerID: customer.ExtID,
		ExtStaffID:    staff.ExtID,
	})
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return err
}
