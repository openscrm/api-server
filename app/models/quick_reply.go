package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/app"
)

// QuickReply 话术内容
type QuickReply struct {
	ExtCorpModel
	// 内部企业ID
	CorpId string `gorm:"comment:内部企业id" json:"ext_corp_id"`
	// 可见部门
	//DepartmentList constants.Int64ArrayField `gorm:"type:json;comment:可见部门" json:"department_list"`
	// 话术名
	Name string `gorm:"type:varchar(255);comment:话术名" json:"name"`
	// 话术类型
	QuickReplyType constants.QuickReplyType `gorm:"type:tinyint" json:"quick_reply_type"`
	// 用于搜索的词语，多为标题
	SearchableText constants.StringArrayField `gorm:"type:json;comment: 用于搜索的词语，多为标题" json:"searchable_text"`
	// 发送次数
	SendCount int `gorm:"type:int unsigned;comment:已发送次数" json:"send_count"`
	// 创建者ID
	ExtStaffId string `gorm:"type:char(32);comment:创建人企微ID" json:"staff_ext_id"`
	// 创建人名字
	StaffName string `gorm:"type:varchar(128);comment:创建人名字" json:"staff_name"`
	// 可用范围
	Scope string `gorm:"type:varchar(32)" json:"scope"`
	// 话术条目详情
	ReplyDetails []QuickReplyDetail `gorm:"foreignKey:QuickReplyID" json:"reply_details"`
	// 分组ID
	GroupID string `gorm:"分组ID" json:"group_id"`
	//QuickReplyGroup QuickReplyGroup `gorm:"foreignKey:GroupID" json:"quick_reply_group"`
	Timestamp
}

func (c QuickReplyDetail) TableName() string {
	return "quick_reply_detail"
}

type QuickReplyWithAvatar struct {
	QuickReply
	Avatar string `json:"avatar"`
}

func (q QuickReply) DeleteByGroupIDs(db *gorm.DB, groupIDs []string) error {
	return db.Model(&QuickReply{}).Where("  group_id in (?)", groupIDs).Delete(&QuickReply{}).Error
}

func (q QuickReply) TableName() string {
	return "quick_reply"
}
func (q QuickReply) Delete(ids []string, extCorpId string) (int64, error) {
	result := DB.Model(&QuickReply{}).
		Where("id in (?)", ids).
		Where("ext_corp_id = ?", extCorpId).
		Delete(&QuickReply{})
	err := result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete ContactWayGroup failed")
		return 0, err
	}
	total := result.RowsAffected
	return total, err
}

func (q QuickReply) Get(id string) (QuickReply, error) {
	var r QuickReply
	err := DB.Model(&QuickReply{}).Preload("ReplyDetails").Where("id = ?", id).First(&r).Error
	return r, err
}

func (q QuickReply) Query(
	req requests.QueryQuickReplyReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []QuickReplyWithAvatar, total int64, err error) {

	db := DB.Table("quick_reply").
		Joins(" left join quick_reply_group  on quick_reply.group_id = quick_reply_group.id ").
		Joins("left join staff on staff.ext_id = quick_reply.ext_creator_id").
		Where(" quick_reply.ext_corp_id = ?", extCorpID).Where("quick_reply.deleted_at is null")

	if req.Keyword != "" {
		db = db.Where("quick_reply.name like ?", req.Keyword+"%")
	}

	if req.GroupID != "" {
		db.Where("quick_reply.group_id = ?", req.GroupID)
	}

	// 用闭包支持 and,or 混合查
	if len(req.DepartmentIDs) > 0 {
		db = db.Where(func(db *gorm.DB) *gorm.DB {
			for _, deptID := range req.DepartmentIDs {
				db = db.Or("json_contains(quick_reply_group.departments, json_array(?))", deptID)
			}
			return db
		}(DB))
	}

	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count quick_reply failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{
		Column: clause.Column{Table: QuickReply{}.TableName(), Name: string(sorter.SortField)},
		Desc:   sorter.SortType == constants.SortTypeDesc},
	)

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Select("quick_reply.*, staff.avatar_url as avatar").
		Preload("ReplyDetails").
		//Preload("QuickReplyGroups").
		Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find quick reply failed")
		return
	}
	return
}

func (q QuickReply) Create(group *QuickReply) error {
	return DB.Create(group).Error
}

func (q QuickReply) Update(quickReply QuickReply) error {
	return DB.Model(&quickReply).Where("id = ?", quickReply.ID).Updates(&quickReply).Error
}

func (q QuickReply) QueryByKeyword(keyword string, extCorpID string) (replies []QuickReply, err error) {
	err = DB.Model(&QuickReply{}).
		Where("ext_corp_id = ?", extCorpID).
		Where("name like ?", keyword+"%").Find(&replies).Error
	return
}

//QuickReplyDetail 话术库每条记录内容
type QuickReplyDetail struct {
	ExtCorpModel
	// 话术ID
	QuickReplyID string `json:"quick_reply_id"`
	// 话术内容json 类型
	QuickReplyContent constants.QuickReplyField `gorm:"type:json" json:"quick_reply_content"`
	// 可见范围 corp/staff_event/group
	Scope string `json:"scope"`
	// 单一类型(文字/图片/图文/PDF/视频)
	ContentType constants.QuickReplyType `gorm:"type:tinyint unsigned;comment:单项类型" json:"content_type"`
	// 已使用次数
	SendCount int `json:"send_count"`
	Timestamp
}

func (c QuickReplyDetail) Delete(ids []string) error {
	return DB.Model(&QuickReplyDetail{}).Where("id in (?)", ids).Delete(&QuickReplyDetail{}).Error
}

func (c QuickReplyDetail) Upsert(details []QuickReplyDetail) error {
	return DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{"name", "quick_reply_content", "content_type", "ext_staff_id"},
		)}).CreateInBatches(&details, 1000).Error
}
