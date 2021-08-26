package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
	"strings"
)

type ContactWay struct {
	model models.ContactWay
}

func NewContactWay() *ContactWay {
	return &ContactWay{model: models.ContactWay{}}
}

func (o *ContactWay) Query(req requests.QueryContactWayReq, extCorpID string) (items []responses.ContactWay, total int64, err error) {
	param := models.ContactWay{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	contactWays := make([]models.ContactWay, 0)
	contactWays, total, err = o.model.Query(req, extCorpID)
	if err != nil {
		err = errors.Wrap(err, "Query ContactWay failed")
		return
	}

	tagExtIDs := make([]string, 0)
	for _, contactWay := range contactWays {
		tagExtIDs = append(tagExtIDs, contactWay.CustomerTagExtIDs...)
	}
	tagExtIDs = funk.UniqString(tagExtIDs)
	tags := make([]models.Tag, 0)
	err = models.DB.Model(&models.Tag{}).Where("ext_id in (?)", tagExtIDs).Find(&tags).Error
	if err != nil {
		err = errors.Wrap(err, "get Tags failed")
		return
	}

	tagMap := make(map[string]models.Tag, 0)
	for _, tag := range tags {
		tagMap[tag.ExtID] = tag
	}

	for _, contactWay := range contactWays {
		item := responses.ContactWay{}
		item.ContactWay = contactWay
		for _, tagExtID := range item.CustomerTagExtIDs {
			if tag, ok := tagMap[tagExtID]; ok {
				item.CustomerTags = append(item.CustomerTags, tag)
			}
		}
		items = append(items, item)
	}

	return
}

func (o *ContactWay) Get(id string, extCorpID string) (item responses.ContactWay, err error) {
	contactWay, err := o.model.Get(id, extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get contactWay failed")
		return
	}

	item.ContactWay = contactWay

	err = models.DB.Model(&models.Tag{}).Where("ext_id in (?)", []string(contactWay.CustomerTagExtIDs)).Find(&item.CustomerTags).Error
	if err != nil {
		err = errors.Wrap(err, "get Tag failed")
		return
	}

	err = models.DB.Model(&models.ContactWayGroup{}).Where("id = ?", contactWay.GroupID).First(&item.Group).Error
	if err != nil {
		err = errors.Wrap(err, "get ContactWayGroup failed")
		return
	}

	return
}

type preprocessMode string

const preprocessModeCreate = "create"
const preprocessModeUpdate = "update"

// Preprocess 创建和更新前对参数进行预处理
func (o *ContactWay) Preprocess(mode preprocessMode, item *models.ContactWay, extCorpID string, extCreatorID string) (err error) {
	// 公共部分
	item.ExtCorpID = extCorpID
	item.ExtCreatorID = extCreatorID
	if item.ScheduleEnable == constants.False {
		item.Schedules = nil
	} else {
		item.Staffs = nil
	}
	extStaffIDs := make([]string, 0)
	for _, staff := range item.Staffs {
		extStaffIDs = append(extStaffIDs, staff.ExtStaffID)
	}
	for _, backupStaff := range item.BackupStaffs {
		extStaffIDs = append(extStaffIDs, backupStaff.ExtStaffID)
	}
	for _, schedule := range item.Schedules {
		for _, staff := range schedule.Staffs {
			extStaffIDs = append(extStaffIDs, staff.ExtStaffID)
		}
	}
	staffs := make([]models.Staff, 0)
	err = models.DB.Select("name,avatar_url,ext_id").Where("ext_id in (?)", extStaffIDs).Find(&staffs).Error
	if err != nil {
		err = errors.Wrap(err, "Find staffs failed")
		return
	}
	staffMap := make(map[string]models.Staff, 0)
	for _, staff := range staffs {
		staffMap[staff.ExtID] = staff
	}

	// 创建渠道码
	if mode == preprocessModeCreate {
		item.ID = id_generator.StringID()
		// 设置关联数据主键
		for i := range item.Staffs {
			item.Staffs[i].ID = id_generator.StringID()
		}
		for i := range item.BackupStaffs {
			item.BackupStaffs[i].ID = id_generator.StringID()
		}
		for i, schedule := range item.Schedules {
			item.Schedules[i].ID = id_generator.StringID()
			for j := range schedule.Staffs {
				item.Schedules[i].Staffs[j].ID = id_generator.StringID()
			}
		}
	}

	// 更新渠道码
	if mode == preprocessModeUpdate {
		// 设置关联数据主键
		for i := range item.Staffs {
			if len(item.Staffs[i].ID) == 0 {
				item.Staffs[i].ID = id_generator.StringID()
			}
		}
		for i := range item.BackupStaffs {
			if len(item.BackupStaffs[i].ID) == 0 {
				item.BackupStaffs[i].ID = id_generator.StringID()
			}
		}
		for i, schedule := range item.Schedules {
			if len(item.Schedules[i].ID) == 0 {
				item.Schedules[i].ID = id_generator.StringID()
			}
			for j := range schedule.Staffs {
				if len(item.Schedules[i].Staffs[j].ID) == 0 {
					item.Schedules[i].Staffs[j].ID = id_generator.StringID()
				}
			}
		}
	}

	// 设置关联数据主键
	for i := range item.Staffs {
		staff, ok := staffMap[item.Staffs[i].ExtStaffID]
		if ok {
			item.Staffs[i].Name = staff.Name
			item.Staffs[i].AvatarURL = staff.AvatarURL
		}
		item.Staffs[i].ContactWayID = item.ID
		item.Staffs[i].ExtCorpID = extCorpID
		item.Staffs[i].ExtCreatorID = extCreatorID
	}
	for i := range item.BackupStaffs {
		staff, ok := staffMap[item.BackupStaffs[i].ExtStaffID]
		if ok {
			item.BackupStaffs[i].Name = staff.Name
			item.BackupStaffs[i].AvatarURL = staff.AvatarURL
		}
		item.BackupStaffs[i].ContactWayID = item.ID
		item.BackupStaffs[i].ExtCorpID = extCorpID
		item.BackupStaffs[i].ExtCreatorID = extCreatorID
	}
	for i, schedule := range item.Schedules {
		item.Schedules[i].ContactWayID = item.ID
		item.Schedules[i].ExtCorpID = extCorpID
		item.Schedules[i].ExtCreatorID = extCreatorID
		for j := range schedule.Staffs {
			staff, ok := staffMap[item.Schedules[i].Staffs[j].ExtStaffID]
			if ok {
				item.Schedules[i].Staffs[j].Name = staff.Name
				item.Schedules[i].Staffs[j].AvatarURL = staff.AvatarURL
			}
			item.Schedules[i].Staffs[j].ContactWayID = item.ID
			item.Schedules[i].Staffs[j].ContactWayScheduleID = item.Schedules[i].ID
			item.Schedules[i].Staffs[j].ExtCorpID = extCorpID
			item.Schedules[i].Staffs[j].ExtCreatorID = extCreatorID
		}
	}

	return
}

func (o *ContactWay) Create(req requests.CreateContactWayReq, extCorpID string, extCreatorID string) (item models.ContactWay, err error) {
	err = copier.Copy(&item, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	item.ID = id_generator.StringID()

	err = o.Preprocess(preprocessModeCreate, &item, extCorpID, extCreatorID)
	if err != nil {
		err = errors.Wrap(err, "Preprocess failed")
		return
	}

	return o.model.Create(item, extCorpID)
}

func (o *ContactWay) Update(id string, req requests.UpdateContactWayReq, extCorpID string) (item models.ContactWay, err error) {
	err = copier.Copy(&item, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	item.ID = id

	err = o.Preprocess(preprocessModeUpdate, &item, extCorpID, "")
	if err != nil {
		err = errors.Wrap(err, "Preprocess failed")
		return
	}

	return o.model.Update(id, item, extCorpID)
}

func (o *ContactWay) Delete(ids []string, extCorpID string) (total int64, err error) {
	return o.model.Delete(ids, extCorpID)
}

func (o *ContactWay) BatchUpdate(req requests.BatchUpdateContactWayReq, extCorpID string) (total int64, err error) {
	param := models.ContactWay{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	return o.model.BatchUpdate(req.IDs, extCorpID, param)
}

// DealAddCustomerEvent 处理添加客户时，渠道码应该做的操作
func (o *ContactWay) DealAddCustomerEvent(tx *gorm.DB, event workwx.EventAddExternalContact) (shouldSendWelcomeMsg bool, err error) {
	if !strings.HasPrefix(event.GetState(), constants.ContactWayStatePrefix) {
		return
	}
	shouldSendWelcomeMsg = true
	staffSrv := NewStaffService()
	extStaffID := event.GetUserID()
	extCustomerID := event.GetExternalUserID()
	contactWayID := strings.TrimPrefix(event.GetState(), constants.ContactWayStatePrefix)
	contactWay := models.ContactWay{}
	err = tx.Where("id = ?", contactWayID).First(&contactWay).Error
	if err == gorm.ErrRecordNotFound {
		log.Sugar.Debugw("contactWay not found", "contactWayID", contactWayID)
		return shouldSendWelcomeMsg, nil
	}
	if err != nil {
		err = errors.Wrap(err, "get contactWay failed")
		return
	}

	client, err := we_work.Clients.Get(contactWay.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return
	}

	if contactWay.AutoTagEnable.Bool() {
		updateCustomerTagErr := staffSrv.UpdateCustomerTag(requests.UpdateCustomerTagsReq{
			ExtCustomerIDs: []string{event.GetExternalUserID()},
			AddExtTagIDs:   contactWay.CustomerTagExtIDs,
		}, contactWay.ExtCreatorID, contactWay.ExtCorpID)
		if updateCustomerTagErr != nil {
			log.Sugar.Errorw("UpdateCustomerTag failed", "contactWay", contactWay, "err", updateCustomerTagErr)
		}
	}

	// 修改客户备注，客户描述
	if contactWay.CustomerRemarkEnable.Bool() || contactWay.CustomerDescEnable.Bool() {
		err = client.Customer.RemarkExternalContact(&workwx.ExternalContactRemark{
			Userid:         extStaffID,
			ExternalUserid: extCustomerID,
			Remark:         contactWay.Remark,
			Description:    contactWay.CustomerDesc,
		})
		if err != nil {
			err = errors.Wrap(err, "wx RemarkExternalContact failed")
			return
		}
	}

	// 欢迎语屏蔽
	shouldBlockAutoReply := false
	if contactWay.NicknameBlockEnable.Bool() {
		// 获取客户详情
		customerInfo, err := client.Customer.GetExternalContact(extCustomerID)
		if err != nil {
			err = errors.Wrap(err, "wx GetExternalContact failed")
			return shouldSendWelcomeMsg, err
		}
		if contactWay.NicknameBlockList.Match(customerInfo.ExternalContact.Name) {
			shouldBlockAutoReply = true
		}

		shouldSendWelcomeMsg = false
	}

	if !shouldBlockAutoReply && contactWay.AutoReplyType == constants.ContactWayAutoReplyTypeCustom {
		shouldSendWelcomeMsg = false
		sendErr := staffSrv.SendWelcomeMsg(contactWay.AutoReply, event.GetWelcomeCode(), contactWay.ExtCorpID, contactWay.ExtCreatorID)
		if sendErr != nil {
			log.Sugar.Errorw("SendWelcomeMsg failed", "contactWay", contactWay, "err", sendErr)
		}
	}

	if !shouldBlockAutoReply && contactWay.AutoReplyType == constants.ContactWayAutoReplyTypeDefault {
		shouldSendWelcomeMsg = true
	}

	if !shouldBlockAutoReply && contactWay.AutoReplyType == constants.ContactWayAutoReplyTypeDisable {
		shouldSendWelcomeMsg = false
	}

	// 员工添加人数统计
	contactWayStaff := models.ContactWayStaff{}
	err = tx.Where("contact_way_id = ?", contactWayID).Where("ext_staff_id = ?", extStaffID).First(&contactWayStaff).Error
	if err == gorm.ErrRecordNotFound {
		// 如果刚好修改了渠道码，把绑定员工取消了，这里可能找不到，不算错误
		log.Sugar.Warnw("contactWayStaff not found", "contactWayID", contactWayID, "extStaffID", extStaffID)
		return shouldSendWelcomeMsg, nil
	}
	if err != nil {
		err = errors.Wrap(err, "get contactWayStaff failed")
		return
	}

	err = tx.Model(&contactWayStaff).Updates(map[string]interface{}{
		"add_customer_count":     gorm.Expr("add_customer_count + 1"),
		"day_add_customer_count": gorm.Expr("day_add_customer_count + 1"),
	}).Error
	if err != nil {
		err = errors.Wrap(err, "update contactWayStaff failed")
		return
	}

	if contactWayStaff.DailyAddCustomerCount >= contactWayStaff.DailyAddCustomerLimit {
		log.Sugar.Infow("[渠道码][员工添加客户数量超过限制]", "contactWayID", contactWayID, "extStaffID", extStaffID)
		contactWay, err = models.ContactWay{}.Refresh(tx, contactWayID)
		if err != nil {
			err = errors.Wrap(err, "ContactWay Refresh failed")
			return
		}
	}

	return
}
