package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
)

type MassMsgStaff struct {
	ExtCorpModel
	MassMsgID string `json:"mass_msg_id" gorm:"index;type:bigint"`
	// 发送人
	ExtStaffID string `json:"ext_staff_id"`
	// 接收人
	ExtCustomerID string `json:"ext_customer_id"`
	// 接收消息群ID
	ExtChatID string `json:"ext_chat_id"`
	// 是否投递
	IsSent uint8 `json:"is_sent" gorm:"type:tinyint unsigned;default:2"`
	// 是否送达
	IsDelivered uint8 `json:"is_delivered" gorm:"type:tinyint unsigned;default:2"`
	// 失败原因
	FailedReason uint8 `json:"failed_reason" gorm:"type:tinyint unsigned"`
	Timestamp
}

func (g MassMsgStaff) QueryExtStaffIds(ids []string) (res []MassMsgStaff, err error) {
	err = DB.Model(&MassMsgStaff{}).
		Where("mass_msg_id in (?)", ids).
		Where("is_sent = ?", constants.False).
		Find(&res).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

type StaffsCustomers struct {
	ExtStaffID    string `json:"ext_staff_id"`
	ExtCustomerID string `json:"ext_customer_id"`
}

func (g MassMsgStaff) GetStaffsCustomers(
	extStaffIDs constants.StringArrayField,
	filterEnable constants.Boolean,
	filter constants.ExtCustomerFilter) (cs []StaffsCustomers, total int64, err error) {

	db := DB.Table("customer").
		Joins(" left join customer_staff cs on customer.ext_id = cs.ext_customer_id").
		Joins(" left join group_chat_member gcm on customer.ext_id = gcm.userid")

	if len(extStaffIDs) > 0 {
		db = db.Where("cs.ext_staff_id in (?)", extStaffIDs.ToStringArray())
	}
	if filterEnable != constants.False {
		if len(filter.ExtGroupChatIDs) > 0 && filter.ExtGroupChatIDs[0] != "" {
			db = db.Where("gcm.ext_chat_id in (?)", filter.ExtGroupChatIDs.ToStringArray())
		}

		if filter.StartTime != "" {
			db = db.Where("cs.createtime between ? and ?", filter.StartTime, filter.EndTime)
		}

		if filter.Gender != 0 {
			db = db.Where("customer.gender = ?", filter.Gender)
		}
		if len(filter.ExtTagIDs) > 0 || len(filter.ExcludeExtTagIDs) > 0 {
			db = db.Joins(" left join customer_staff_tag cst on cs.id = cst.customer_staff_id ")

			if len(filter.ExcludeExtTagIDs) > 0 && filter.ExcludeExtTagIDs[0] != "" {
				db = db.Where("cst.ext_tag_id not in (?)", filter.ExcludeExtTagIDs.ToStringArray())
			}
			if len(filter.ExtTagIDs) > 0 && filter.ExtTagIDs[0] != "" {
				if filter.TagLogicalCondition == constants.LogicalConditionAND {
					db = db.Where(func(db *gorm.DB) *gorm.DB {
						for _, tagID := range filter.ExtTagIDs {
							db = db.Where("cst.ext_tag_id = ?", tagID)
						}
						return db
					}(DB))
				} else if filter.TagLogicalCondition == constants.LogicalConditionOR {
					db = db.Where("cst.ext_tag_id in (?)", filter.ExtTagIDs.ToStringArray())
				} else if filter.TagLogicalCondition == constants.LogicalConditionNone {
					db = db.Where("cst.ext_tag_id = ?", nil)
				}
			}
		}
	}

	db = db.Group("cs.ext_staff_id").Group("cs.ext_customer_id")

	err = db.Count(&total).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = db.Select("cs.ext_staff_id, cs.ext_customer_id").Find(&cs).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

type MsgMainInfo struct {
	ID         string `json:"id"`
	ExtMsgID   string `json:"ext_msg_id"`
	ExtStaffID string `json:"ext_staff_id"`
}

// QueryNotSentMsg 还没发送消息的员工和外部消息ID
func (g MassMsgStaff) QueryNotSentMsg(massMsgStaff MassMsgStaff) (res []MsgMainInfo, err error) {
	err = DB.Table("mass_msg").
		Joins("join mass_msg_staff on  mass_msg_staff.mass_msg_id = mass_msg.id").
		Where("mass_msg.ext_corp_id = ?", massMsgStaff.ExtCorpID).
		Where("is_sent = ?", massMsgStaff.IsSent).
		Select("mass_msg.ext_msg_id as ext_msg_id, mass_msg_staff.ext_staff_id as ext_staff_id, mass_msg_staff.id as id").
		Find(&res).Error

	return
}

func (g MassMsgStaff) Update(staff MassMsgStaff) (err error) {
	return DB.Model(&MassMsgStaff{}).Where("id = ?", staff.ID).Updates(&staff).Error
}
