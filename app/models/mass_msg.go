package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

func (o MassMsg) Update(msg MassMsg) error {
	return DB.Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", msg.ID).Updates(&msg).Error
}

// MassMsg 企业群发消息内容
// 消息内容不可修改
type MassMsg struct {
	ExtCorpModel
	// 消息类型
	SendType constants.SendMassMsgType `gorm:"type:tinyint unsigned;comment:1-立即发送,2-定时发送" json:"send_type" `
	// 员工ID
	ExtStaffIDs constants.StringArrayField `gorm:"type:JSON" json:"ext_staff_ids"`
	// 可用部门ID
	ExtDepartmentIDs constants.Int64ArrayField `gorm:"type:JSON" json:"ext_department_ids"`
	// 消息内容
	Msg constants.AutoReplyField `gorm:"type:json;comment:消息内容" json:"msg"`
	// wx消息ID
	ExtMsgID string `gorm:"type:varchar(33);comment:微信消息ID;index" json:"ext_msg_id"`
	// 任务状态 1-预约发送,2-发送中,3-发送成功,4-发送失败,5-已取消; <=1  可修改,其余不可改
	MissionStatus constants.SendMassMsgStatus `gorm:"comment:创建企业群发消息的状态,1-预约发送,2-发送中,3-发送成功,4-发送失败,5-已取消;type:tinyint unsigned;" json:"mission_status"`
	// 是否有筛选条件
	ExtCustomerFilterEnable constants.Boolean `gorm:"type:tinyint unsigned;comment:是否有筛选条件" json:"ext_customer_filter_enable" form:"ext_customer_filter_enable"`
	// 接受客户的筛选条件
	ExtCustomerFilter constants.ExtCustomerFilter `gorm:"type:json;comment:发送客户的筛选条件" json:"ext_customer_filter"`
	// 已发员工计数
	DeliveredNum int `gorm:"comment:已发送消息的员工数;type:int" json:"delivered_num"`
	// 已送达客户计数
	SuccessNum int `gorm:"comment:成功送达消息的员工数;type:int" json:"success_num"`
	// 未发送员工计数
	UnDeliveredNum int `gorm:"comment:需要发送消息的员工总数;type:int"  json:"undelivered_num"`
	// 未送达客户计数
	FailedNum int `gorm:"comment:未送达客户计数;type:int" json:"failed_num"`
	// 需要发送的员工
	Staffs []MassMsgStaff `gorm:"foreignKey:MassMsgID;references:ID" json:"staffs"`
	// 定时发送时间
	SendAt constants.DateTimeFiled `json:"send_at" validate:"omitempty,gt=0"`
	Timestamp
}

// GetExtStaffIDs
// Description: 群发消息任务表MassMsgStaff 找到谁发给谁
// Detail: 	 todo opt 分页
// Param:  群发任务ID
// return: 员工-客户
func (o MassMsg) GetExtStaffIDs(msgID string) (staffsCustomers []StaffsCustomers, err error) {
	err = DB.Model(&MassMsgStaff{}).Where("mass_msg_id = ?", msgID).
		Where("is_sent = ?", constants.False).
		Select("ext_staff_id, ext_customer_id").
		Find(&staffsCustomers).Error
	return
}

func (o MassMsg) Create(msg MassMsg) error {
	return DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&msg).Error
}

func (o MassMsg) Delete(id string) error {
	return DB.Model(&MassMsg{}).Where("id=?", id).Delete(&MassMsg{}).Error
}

func (o MassMsg) Get(id string) (msg MassMsg, err error) {
	err = DB.Model(&MassMsg{}).Preload("Staffs").Where("id=?", id).First(&msg).Error
	if err != nil {
		return
	}
	return
}

func (o MassMsg) GetMsgs(ids []string) (msgs []MassMsg, err error) {
	err = DB.Model(&MassMsg{}).Preload("Staffs").Where("id in (?)", ids).Take(&msgs).Error
	if err != nil {
		return nil, err
	}
	return
}
func (o MassMsg) Query(extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]MassMsg, int64, error) {
	items := make([]MassMsg, 0)
	db := DB.Model(&MassMsg{}).Where("ext_corp_id = ?", extCorpID)
	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count MassMsg failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find MassMsg failed")
		return nil, 0, err
	}
	return items, total, nil
}
