package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/log"
)

// CustomerEvent 事件类型
type CustomerEvent struct {
	ExtCorpModel
	Content string `gorm:"type:text;comment:事件内容" json:"content"`
	// 筛选
	EventType         string `gorm:"type:char(32);index;comment:事件类型" json:"event_type"`
	EventName         string `gorm:"type:char(32);comment:事件名称" json:"event_name"`
	ExtCustomerID     string `gorm:"type:char(32);comment:企微定义的客户ID" json:"ext_customer_id"`
	ExtStaffID        string `gorm:"type:varchar(32);index;comment:微信定义的员工ID" json:"ext_staff_id"`
	RelateStaffAvatar string `gorm:"type:varchar(128);comment:员工头像" json:"relate_staff_avatar"`
	RelateStaffName   string `gorm:"type:varchar(255);comment:员工名字" json:"relate_staff_name"`
	// 提醒类型事件的发送时间
	SendAt constants.DateTimeFiled `gorm:"comment:提醒类型事件的发送时间" json:"send_at"`
	Timestamp
}

func (c CustomerEvent) Query(ce CustomerEvent, pager *app.Pager, sorter *app.Sorter) ([]CustomerEvent, int64, error) {
	db := DB.Model(&CustomerEvent{}).Where("ext_corp_id = ?", ce.ExtCorpID)

	if ce.ExtStaffID != "" {
		db.Where("ext_staff_id = ?", ce.ExtStaffID)
	}
	if ce.ExtCustomerID != "" {
		db.Where("ext_customer_id = ?", ce.ExtCustomerID)
	}

	if ce.EventType != "" {
		db.Where("event_type = ?", ce.EventType)
	}
	items := make([]CustomerEvent, 0)
	total := int64(0)
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count customer_event failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find customer event failed")
		return nil, 0, err
	}

	return items, total, nil
}

// Create 支持前端创建 模板记录、跟进记录和提醒事件
func (c CustomerEvent) Create(ce CustomerEvent) error {
	return DB.Model(&CustomerEvent{}).Create(&ce).Error
}

func (c CustomerEvent) CreateInBatches(ce []CustomerEvent) error {
	return DB.Model(&CustomerEvent{}).CreateInBatches(&ce, 100).Error
}

func (c CustomerEvent) Update(ce CustomerEvent) (CustomerEvent, error) {
	customerEvent := CustomerEvent{}
	err := DB.Model(CustomerEvent{ExtCorpModel: ExtCorpModel{ID: ce.ID}}).First(&customerEvent).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return customerEvent, err
	}
	if err != nil {
		err = errors.Wrap(err, "find CustomerEvent failed")
		return customerEvent, err
	}

	err = DB.Model(CustomerEvent{ExtCorpModel: ExtCorpModel{ID: ce.ID}}).Updates(&ce).Error
	if err != nil {
		err = errors.Wrap(err, "Update CustomerEvent failed")
		return customerEvent, err
	}

	return ce, err
}

func (c CustomerEvent) Delete(ids []string, extCorpID string) (int64, error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&ContactWayGroup{})
	err := result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete ContactWayGroup failed")
		return 0, err
	}

	return result.RowsAffected, err
}

func (c CustomerEvent) Get(id string) (CustomerEvent, error) {
	var ce CustomerEvent
	err := DB.Model(&CustomerEvent{}).Where("id = ?", id).First(&ce).Error
	if err != nil {
		log.Sugar.Errorw("get customer event failed", "id", id)
		return ce, err
	}
	return ce, err
}
