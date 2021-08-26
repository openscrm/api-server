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
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/delay_queue"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"time"
)

type MassMsgService struct {
	massMsgRepo      models.MassMsg
	MassMsgStaffRepo models.MassMsgStaff
	CustomerRepo     models.Customer
	staffRepo        models.Staff
}

func NewDefaultMassMsgService() *MassMsgService {
	return &MassMsgService{
		massMsgRepo:      models.MassMsg{},
		MassMsgStaffRepo: models.MassMsgStaff{},
		CustomerRepo:     models.Customer{},
		staffRepo:        models.Staff{},
	}
}

// Create
// 定时和立即发送都统一异步发送
func (o MassMsgService) Create(req requests.SendMassMsgReq, creator, extCorpID string) (msg models.MassMsg, err error) {
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

	msg = models.MassMsg{
		ExtCorpModel:            models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: creator},
		SendType:                req.SendType,
		ExtStaffIDs:             req.ExtStaffIDs,
		ExtDepartmentIDs:        req.ExtDepartmentIDs,
		Msg:                     req.Msg,
		MissionStatus:           constants.NotActive, // 默认为及时发送
		ExtCustomerFilterEnable: req.ExtCustomerFilterEnable,
		ExtCustomerFilter:       req.ExtCustomerFilter,
		UnDeliveredNum:          len(req.ExtStaffIDs),
		SendAt:                  req.SendAt,
	}

	// 员工:客户列表
	// 找到需要发送的员工和对应的客户，写表。
	staffsCustomers, total, err := o.GetStaffsCustomers(req.ExtStaffIDs, req.ExtCustomerFilterEnable, req.ExtCustomerFilter)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if total == 0 {
		err = ecode.NoMassMsgReceiversErr
		return
	}

	MassMsgStaffs := make([]models.MassMsgStaff, 0)
	for _, staffCustomer := range staffsCustomers {
		MassMsgStaffs = append(MassMsgStaffs,
			models.MassMsgStaff{
				ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: creator, ExtCorpID: extCorpID},
				ExtStaffID:    staffCustomer.ExtStaffID,
				ExtCustomerID: staffCustomer.ExtCustomerID,
				MassMsgID:     msg.ID,
			})
	}

	msg.Staffs = MassMsgStaffs
	// 初始化为全部未送达
	msg.FailedNum = len(MassMsgStaffs)
	err = o.massMsgRepo.Create(msg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 推消息
	msgBytes, err := json.Marshal(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	job := delay_queue.Job{
		Topic:     constants.MassMsgTopic,
		ID:        msg.ID,
		ExecuteAt: req.SendAt.ToInt64(),
		TTR:       5,
		Body:      string(msgBytes),
	}
	err = delay_queue.Add(job)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return o.massMsgRepo.Get(msg.ID)
}

// UpdateMassMsg
// 不支持删除立即发送消息
func (o MassMsgService) UpdateMassMsg(
	req requests.UpdateMassMsgReq, id string, name string, extCorpID string) (msg models.MassMsg, err error) {

	massMsg, err := o.massMsgRepo.Get(id)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if massMsg.SendType == constants.Instant || massMsg.MissionStatus > constants.NotActive {
		err = ecode.UnsupportedMsgError
		err = errors.WithStack(err)
		return
	}
	if req.SendAt.ToInt64() <= time.Now().Unix() {
		err = ecode.EarlierThanNowError
		return
	}
	// 更新延迟发送的消息
	msgBytes, err := json.Marshal(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	job := delay_queue.Job{
		Topic:     constants.MassMsgTopic,
		ID:        id,
		ExecuteAt: req.SendAt.ToInt64(),
		TTR:       5,
		Body:      string(msgBytes),
	}
	err = delay_queue.Add(job)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 接口不能更新消息发送人员统计数据
	msg = models.MassMsg{
		ExtCorpModel:            models.ExtCorpModel{ID: id, ExtCorpID: extCorpID, ExtCreatorID: name},
		SendType:                req.SendType,
		ExtStaffIDs:             req.ExtStaffIDs,
		ExtDepartmentIDs:        req.ExtDepartmentIDs,
		Msg:                     req.Msg,
		MissionStatus:           constants.NotActive,
		ExtCustomerFilter:       req.ExtCustomerFilter,
		ExtCustomerFilterEnable: req.ExtCustomerFilterEnable,
		SendAt:                  req.SendAt,
	}
	if msg.SendType == constants.Instant {
		msg.MissionStatus = constants.Sending
	}

	// 员工:客户列表
	// 找到需要发送的员工和对应的客户，写表。
	staffsCustomers, total, err := o.GetStaffsCustomers(req.ExtStaffIDs, req.ExtCustomerFilterEnable, req.ExtCustomerFilter)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if total != 0 {
		err = ecode.NoMassMsgReceiversErr
		return
	}

	MassMsgStaffs := make([]models.MassMsgStaff, 0)
	for _, staffCustomer := range staffsCustomers {
		MassMsgStaffs = append(MassMsgStaffs,
			models.MassMsgStaff{
				ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: name, ExtCorpID: extCorpID},
				ExtStaffID:    staffCustomer.ExtStaffID,
				ExtCustomerID: staffCustomer.ExtCustomerID,
				MassMsgID:     id,
			})
	}

	msg.Staffs = MassMsgStaffs

	err = o.massMsgRepo.Update(msg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	msg, err = o.massMsgRepo.Get(id)
	return
}

func (o MassMsgService) DeleteTimedMassMsg(ids []string) error {
	MassMsgs, err := o.massMsgRepo.GetMsgs(ids)
	if err != nil {
		err = errors.Wrap(err, "Get failed")
		return err
	}
	for _, msg := range MassMsgs {
		if msg.SendType != constants.Timed {
			err = ecode.UnsupportedMsgError
			return err
		}

		err = o.massMsgRepo.Delete(msg.ID)
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

func (o MassMsgService) Get(msgID, extCorpID string) (res responses.MassMsgDetail, err error) {
	msg, err := o.massMsgRepo.Get(msgID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	staffMainInfo, err := o.staffRepo.GetMainInfo(msg.ExtCreatorID, extCorpID)
	return responses.MassMsgDetail{MassMsg: msg, Creator: staffMainInfo}, nil
}

// GetSendMassMsgResult
// Description:  查询群发结果
// Detail:
func (o MassMsgService) GetSendMassMsgResult(missionID string) (*requests.SendMassMsgResp, error) {
	customer := models.Customer{}
	msg, err := customer.GetMassMsg(missionID)
	resp := &requests.SendMassMsgResp{MissionID: missionID, MissionStatus: constants.Sending}
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return resp, nil
		} else {
			return nil, err
		}
	}
	resp.MissionStatus = msg.MissionStatus
	resp.DeliveredNum = msg.DeliveredNum
	resp.SuccessNum = msg.SuccessNum
	resp.UnDeliveredNum = msg.UnDeliveredNum
	resp.FailedNum = msg.FailedNum
	return resp, nil
}

// SendMassMsgToWx
// Description: 延迟队列发送消息到wx
// Detail:
//	延迟队列的消息->req->we_work request
//  同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息。
func (o MassMsgService) SendMassMsgToWx(body string, msgID string) (extMsgID string, err error) {
	log.Sugar.Debug(body)
	req := requests.SendMassMsgReq{}
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Sugar.Error("unmarshal group msg failed", err)
		return
	}
	template := gowx.AddMsgTemplateReq{}
	err = copier.Copy(&template, req)
	if err != nil {
		log.Sugar.Error("Copy to template from req failed", err)
		return
	}

	template.Text = gowx.Text{Content: req.Msg.Text}
	err = copier.CopyWithOption(&template.Attachments, req.Msg.Attachments, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 定时发送可能被删
	if req.SendType == constants.Timed {
		msg, err := o.massMsgRepo.Get(msgID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Sugar.Info("msg not found, msgID", msgID)
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

	// 创建时保存了谁发给谁，这里只发送.
	staffCustomers, err := o.massMsgRepo.GetExtStaffIDs(msgID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// Map 员工->客户列表
	staffCustomerMap := map[string][]string{}
	for _, StaffCustomer := range staffCustomers {
		extStaffID := StaffCustomer.ExtStaffID
		staffCustomerMap[extStaffID] = append(staffCustomerMap[extStaffID], StaffCustomer.ExtCustomerID)
	}

	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	for extStaffID, extCustomerIDs := range staffCustomerMap {
		if len(extCustomerIDs) <= 0 || extCustomerIDs[0] == "" {
			continue
		}
		template.Sender = extStaffID
		template.ExternalUserid = staffCustomerMap[extStaffID]

		log.Sugar.Debugw(util.JsonEncode(template))

		//同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息
		//接受消息的userid列表中每个id接收者都已收到超过4条消息, 则会返回 no customer to send 错误
		extMsgID, _, err = client.Customer.AddMsgTemplate(template)
		if err != nil {
			log.Sugar.Error("AddMsgTemplate failed", err)
			return
		}
	}

	return extMsgID, err
}

// Notify
// Description: 通知未发送群发的人员发送
func (o MassMsgService) Notify(ids []string, extCorpID string) error {
	// 找到需要发送的该消息的员工id-客户id
	staffCustomerIDs, err := o.MassMsgStaffRepo.QueryExtStaffIds(ids)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	extStaffIDs := make([]string, len(staffCustomerIDs))
	for _, staffCustomer := range staffCustomerIDs {
		extStaffIDs = append(extStaffIDs, staffCustomer.ExtStaffID)

	}
	recipient := gowx.Recipient{UserIDs: extStaffIDs}

	MassMsgs, err := o.massMsgRepo.GetMsgs(ids)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	for _, MassMsg := range MassMsgs {
		customer, err := o.CustomerRepo.GetByExtID(staffCustomerIDs[0].ExtCustomerID, extStaffIDs, false)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		content := fmt.Sprintf(constants.NotifyStaffSendMassMsg, MassMsg.CreatedAt.Format(constants.DateTimeLayout), customer.Name, len(staffCustomerIDs))

		err = client.MainApp.SendTextMessage(&recipient, content, false)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}

func (o MassMsgService) Query(extCorpID string, sorter *app.Sorter, pager *app.Pager) (msgs []models.MassMsg, total int64, err error) {
	msgs = make([]models.MassMsg, 0)
	msgs, total, err = o.massMsgRepo.Query(extCorpID, sorter, pager)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// GetStaffsCustomers
// 员工-客户对应, 一个客户对应多个员工会随机去重
func (o MassMsgService) GetStaffsCustomers(
	extStaffIDs constants.StringArrayField,
	filterEnable constants.Boolean,
	filter constants.ExtCustomerFilter) (res []models.StaffsCustomers, total int64, err error) {

	res = make([]models.StaffsCustomers, 0)
	res, total, err = o.MassMsgStaffRepo.GetStaffsCustomers(extStaffIDs, filterEnable, filter)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
