package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

// ContactWayStaff 渠道码绑定的员工
type ContactWayStaff struct {
	ExtCorpModel
	// 渠道码ID
	ContactWayID          string `json:"contact_way_id" gorm:"index;unique:ContactWayID_ExtStaffID;type:bigint;comment:'渠道码ID'"`
	AddCustomerCount      int    `json:"add_customer_count" gorm:"->;default:0;comment:'员工累计添加客户计数'"`
	DailyAddCustomerCount int    `json:"daily_add_customer_count" gorm:"->;default:0;comment:'员工每日添加客户计数'"`
	DailyAddCustomerLimit int    `json:"daily_add_customer_limit" gorm:"comment:'员工每日添加客户上限'"`
	ExtStaffID            string `json:"ext_staff_id" gorm:"index;unique:ContactWayID_ExtStaffID;comment:'外部员工ID'"`
	// 员工名称
	Name string `gorm:"type:varchar(255);comment:员工名" json:"name"`
	// 头像url
	AvatarURL string            `gorm:"type:varchar(128);comment:头像地址" json:"avatar_url"`
	Online    constants.Boolean `json:"online" gorm:"index;default:1;comment:'员工是否在线'"`
	Timestamp
}

// DailyClean 每日清空添加数统计
func (o ContactWayStaff) DailyClean() (err error) {
	err = DB.Model(&ContactWayStaff{}).Update("day_add_customer_count", 0).Error
	if err != nil {
		err = errors.Wrap(err, "update failed")
		return
	}

	return
}

// BeforeCreate 此表作为关联表时支持字段权限配置，没有更新权限的字段，不会被upsert更新
func (o ContactWayStaff) BeforeCreate(tx *gorm.DB) (err error) {
	var colsNames []string
	// 设置冲突时需要更新的字段
	for _, field := range tx.Statement.Schema.Fields {
		if field.Updatable {
			colsNames = append(colsNames, field.DBName)
		}
	}
	tx.Statement.AddClause(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(colsNames),
	})
	return nil
}
