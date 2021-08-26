package services

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/delay_queue"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"time"
)

type GroupChatMassMsg struct {
	groupChatMassMsgRepo models.GroupChatMassMsg
	massMsgStaffRepo     models.MassMsgStaff
	staffRepo            models.Staff
}

func NewGroupChatMassMsg() *GroupChatMassMsg {
	return &GroupChatMassMsg{
		massMsgStaffRepo:     models.MassMsgStaff{},
		groupChatMassMsgRepo: models.GroupChatMassMsg{},
		staffRepo:            models.Staff{},
	}
}

func (o GroupChatMassMsg) SendMassMsgAsync(req requests.SendGroupChatMassMsgReq, extStaffID string, extCorpID string) (msg models.GroupChatMassMsg, err error) {
	// 发送时间校验
	if req.SendType == constants.Timed {
		if req.SendAt.ToInt64() < time.Now().Unix() {
			err = ecode.EarlierThanNowError
			return
		}
	} else if req.SendType == constants.Instant {
		// 写redis比写DB快，等待2秒后让db写完,执行任务才能读到数据
		req.SendAt = constants.DateTimeFiled(time.Now().Add(2 * time.Second).Format(constants.DateTimeLayout))
	}
	// 推消息
	msgBytes, err := json.Marshal(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	t := req.SendAt.ToInt64()
	missionID := id_generator.StringID()
	job := delay_queue.Job{
		Topic:     constants.GroupChatMassMsgTopic,
		ID:        missionID,
		ExecuteAt: t,
		TTR:       5,
		Body:      string(msgBytes),
	}
	err = delay_queue.Add(job)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	msg = models.GroupChatMassMsg{
		ExtCorpModel:   models.ExtCorpModel{ID: missionID, ExtCorpID: extCorpID, ExtCreatorID: extStaffID},
		SendType:       req.SendType,
		ExtStaffIDs:    req.ExtStaffIDs,
		Msg:            req.Msg,
		MissionStatus:  constants.NotActive, // 默认为及时发送
		UnDeliveredNum: len(req.ExtStaffIDs),
		SendAt:         req.SendAt,
	}
	err = o.groupChatMassMsgRepo.Create(msg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// UpdateGroupMsgSentStatus
// 发送状态：0-未发送 1-已发送 2-因客户不是好友导致发送失败 3-因客户已经收到其他群发消息导致发送失败
func (o GroupChatMassMsg) UpdateGroupMsgSentStatus() (err error) {
	// 找出还没发送的员工和对应extMsgID
	extCorpID := conf.Settings.WeWork.ExtCorpID
	massMsgStaff := models.MassMsgStaff{ExtCorpModel: models.ExtCorpModel{ExtCorpID: extCorpID}, IsSent: 0}
	msgs, err := o.massMsgStaffRepo.QueryNotSentMsg(massMsgStaff)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	for _, msg := range msgs {
		req := gowx.GetGroupMsgSendResultExternalContactReq{
			Msgid:  msg.ExtMsgID,
			Userid: msg.ExtStaffID,
		}
		res, err := client.Customer.GetGroupMsgSendResultExternalContact(req)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		for _, item := range res.SendList {
			// 若为客户群群发，由于用户还未选择群，所以不返回未发送记录，只返回已发送记录
			if item.ChatID != "" {
				if item.Status == 1 {
					err := o.massMsgStaffRepo.Update(
						models.MassMsgStaff{
							ExtCorpModel: models.ExtCorpModel{ID: msg.ID},
							ExtStaffID:   msg.ExtStaffID,
							ExtChatID:    msg.ExtMsgID,
							IsSent:       uint8(item.Status),
							IsDelivered:  uint8(item.Status),
						},
					)
					if err != nil {
						err = errors.WithStack(err)
						return err
					}
				}
			} else {
				// 群发给客户
			}
		}
	}
	return
}

func (o GroupChatMassMsg) DoSendGroupChatMassMsg(jobBody, JobID string) (extMsgID string, err error) {
	log.Sugar.Debug(jobBody)

	req := requests.SendGroupChatMassMsgReq{}
	err = json.Unmarshal([]byte(jobBody), &req)
	if err != nil {
		log.Sugar.Error("unmarshal group msg failed", err)
		return
	}

	// 定时发送可能被删
	if req.SendType == constants.Timed {
		msg, err := o.groupChatMassMsgRepo.Get(JobID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Sugar.Info("msg not found, msgID", JobID)
				return "", err
			}
			log.Sugar.Error("Get failed", err)
			return "", err
		}
		if msg.MissionStatus == constants.Deleted {
			log.Sugar.Info("timed msg has been deleted", req)
			return "", err
		}
	}

	// 发到微信
	template := gowx.AddMsgTemplateReq{}
	err = copier.Copy(&template, req)
	if err != nil {
		log.Sugar.Error("Copy to template from req failed", err)
		return
	}

	template.Text = gowx.Text{Content: req.Msg.Text}
	template.ChatType = string(constants.Group)
	err = copier.CopyWithOption(&template.Attachments, req.Msg.Attachments, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 每个群主发送给他所有的群
	for _, extStaffID := range req.ExtStaffIDs {
		template.Sender = extStaffID
		extMsgID, _, err = client.Customer.AddMsgTemplate(template)
		if err != nil {
			log.Sugar.Error("AddMsgTemplate failed", err)
			return
		}
	}

	return extMsgID, err
}

func (o GroupChatMassMsg) DeleteTimedMassMsg(ids []string) error {
	MassMsgs, err := o.groupChatMassMsgRepo.GetByIDs(ids)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		return err
	}
	for _, msg := range MassMsgs {
		if msg.SendType != constants.Timed {
			err = ecode.UnsupportedMsgError
			return err
		}

		err = o.groupChatMassMsgRepo.Delete(msg.ID)
		if err != nil {
			log.Sugar.Error(err)
			return err
		}

		err = delay_queue.Remove(msg.ID)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}

func (o GroupChatMassMsg) Get(ID string) (resp responses.GroupChatMassMsgDetail, err error) {
	msg, err := o.groupChatMassMsgRepo.Get(ID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	staffMainInfo, err := o.staffRepo.GetMainInfo(msg.ExtCreatorID, msg.ExtCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return responses.GroupChatMassMsgDetail{GroupChatMassMsg: msg, Creator: staffMainInfo}, nil
}

func (o GroupChatMassMsg) Query(extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]models.GroupChatMassMsg, int64, error) {
	return o.groupChatMassMsgRepo.Query(extCorpID, sorter, pager)
}
