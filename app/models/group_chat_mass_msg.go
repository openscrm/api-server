package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

// GroupChatMassMsg 客户群群发消息内容
// 消息内容不可修改
type GroupChatMassMsg struct {
	ExtCorpModel
	// 消息类型
	SendType constants.SendMassMsgType `gorm:"type:tinyint unsigned;comment:1-立即发送,2-定时发送" json:"send_type" `
	// 员工ID
	ExtStaffIDs constants.StringArrayField `gorm:"type:JSON" json:"ext_staff_ids"`
	// 消息内容
	Msg constants.AutoReplyField `gorm:"type:json;comment:消息内容" json:"msg"`
	// wx消息ID
	ExtMsgID string `gorm:"type:varchar(33);comment:微信消息ID;index" json:"ext_msg_id"`
	// 任务状态 1-预约发送,2-发送中,3-发送成功,4-发送失败,5-已取消; <=1  可修改,其余不可改
	MissionStatus constants.SendMassMsgStatus `gorm:"comment:创建企业群发消息的状态,1-预约发送,2-发送中,3-发送成功,4-发送失败,5-已取消;type:tinyint unsigned;" json:"mission_status"`
	// 已发送群主计数
	DeliveredNum int `gorm:"comment:已发送群主计数;type:int unsigned" json:"delivered_num"`
	// 已送达群聊数
	SuccessNum int `gorm:"comment:已送达群聊数;type:int unsigned" json:"success_num"`
	// 未发送群主计数
	UnDeliveredNum int `gorm:"comment:未发送群主计数;type:int unsigned"  json:"undelivered_num"`
	// 	未送达群聊
	FailedNum int `gorm:"comment:	未送达群聊计数;type:int" json:"failed_num"`
	// 需要发送的员工
	//Staffs []MassMsgStaff `gorm:"foreignKey:MassMsgID;references:ID" json:"staffs"`
	// 定时发送时间
	SendAt constants.DateTimeFiled `json:"send_at" validate:"omitempty,gt=0"`
	Timestamp
}

func (m GroupChatMassMsg) Create(msg GroupChatMassMsg) error {
	return DB.Create(&msg).Error
}

func (m GroupChatMassMsg) Get(id string) (msg GroupChatMassMsg, err error) {
	err = DB.Model(&GroupChatMassMsg{}).Where("id = ?", id).First(&msg).Error
	return
}

func (m GroupChatMassMsg) GetByIDs(ids []string) (msgs []GroupChatMassMsg, err error) {
	err = DB.Model(&GroupChatMassMsg{}).Where("id in (?)", ids).Find(&msgs).Error
	return
}

func (m GroupChatMassMsg) Delete(id string) error {
	return DB.Where("id = ?", id).Delete(&GroupChatMassMsg{}).Error
}

func (m GroupChatMassMsg) Query(extCorpID string, sorter *app.Sorter, pager *app.Pager) (res []GroupChatMassMsg, total int64, err error) {
	db := DB.Model(&GroupChatMassMsg{}).Where("ext_corp_id = ?", extCorpID)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count GroupChatMassMsg failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&res).Error
	if err != nil {
		err = errors.Wrap(err, "Find MassMsg failed")
		return
	}
	return
}
