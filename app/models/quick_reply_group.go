package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

//QuickReplyGroup 话术库分组
type QuickReplyGroup struct {
	ExtCorpModel
	// 分组名称
	Name string `gorm:"type:varchar(255);comment:分组名称" json:"name"`
	// 上级分组id
	ParentID *string `gorm:"type:bigint;comment:父级id" json:"parent_id"`
	// 分组 可见部门ids
	Departments constants.Int64ArrayField `gorm:"type:json;comment:可见部门ids" json:"departments"`
	// 是否为顶级分组
	IsTopGroup constants.Boolean `gorm:"type:tinyint;comment:是否是顶级分组" json:"is_top_group"`
	// 下级分组
	SubGroups []QuickReplyGroup `gorm:"foreignKey:ParentID;comment:二级分组" json:"sub_groups"`
	// 话术条目
	QuickReplies []QuickReply `gorm:"foreignKey:GroupID" json:"quick_replies"`
	Order        int64        `gorm:"type:int" json:"order"`
	Timestamp
}

func (qg QuickReplyGroup) Update(groups []QuickReplyGroup) error {
	for _, group := range groups {
		err := DB.Model(&QuickReplyGroup{ExtCorpModel: ExtCorpModel{ID: group.ID}}).Updates(group).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (qg QuickReplyGroup) Delete(tx *gorm.DB, ids []string, extCorpID string) error {
	err := tx.Where("ext_corp_id = ?", extCorpID).Where("parent_id in (?)", ids).Delete(&QuickReplyGroup{}).Error
	if err != nil {
		err = errors.Wrap(err, "Delete QuickReplyGroup failed")
		return err
	}
	err = tx.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&QuickReplyGroup{}).Error
	if err != nil {
		err = errors.Wrap(err, "Delete QuickReplyGroup failed")
		return err
	}
	//err = tx.Model(&QuickReply{}).Where("ext_corp_id = ?", extCorpID).Where("quick_reply_group_id in (?)", ids).Update("quick_reply_group_id", nil).Error
	return nil
}

func (qg QuickReplyGroup) Get(id string, extCorpID string) ([]QuickReplyGroup, error) {
	var group []QuickReplyGroup
	db := DB.Model(&QuickReplyGroup{})
	if id != "" {
		db.Where("id = ?", id)
	}
	if extCorpID != "" {
		db.Where("ext_corp_id = ?", extCorpID)
	}
	//err := db.Preload("SubGroups").Find(&group).Error
	// v1 暂时没有使用SubGroups
	err := db.Find(&group).Error
	return group, err
}

func (qg QuickReplyGroup) Query(
	extCorpID string, sorter *app.Sorter, pager *app.Pager) (groups []QuickReplyGroup, total int64, err error) {

	db := DB.Model(&QuickReplyGroup{}).Where("ext_corp_id = ?", extCorpID)

	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count QuickReplyGroup failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	//err = db.Preload("SubGroups").Preload("QuickReplies").Preload("QuickReplies.ReplyDetails").Find(&groups).Error
	// v1 暂时没有使用SubGroups
	err = db.Preload("QuickReplies").Preload("QuickReplies.ReplyDetails").Find(&groups).Error
	if err != nil {
		err = errors.Wrap(err, "Find QuickReplyGroup failed")
		return
	}
	return
}

func (qg QuickReplyGroup) Create(group QuickReplyGroup) error {
	return DB.Create(&group).Error
}

func (qg QuickReplyGroup) CreateInBatches(groups []QuickReplyGroup) error {
	return DB.Create(groups).Error
}

func (qg QuickReplyGroup) Upsert(groups []QuickReplyGroup) error {
	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "departments"}),
	}).Create(&groups).Error

	return err
}

func (qg QuickReplyGroup) QueryByKeyword(
	keyword string, extCorpID string, pager *app.Pager) (groups []QuickReplyGroup, err error) {

	pager.SetDefault()
	err = DB.Model(&QuickReplyGroup{}).
		Where("ext_corp_id  = ?", extCorpID).
		Where("Name like ?", keyword+"%").
		Preload("QuickReplies").Preload("QuickReplies.ReplyDetails").
		Offset(pager.GetOffset()).Limit(pager.GetLimit()).
		Find(&groups).Error
	return
}

func (qg QuickReplyGroup) QueryByID(extCorpID string, groupIDs []string, repliesIDs []string) (groups []QuickReplyGroup, err error) {
	err = DB.Model(&QuickReplyGroup{}).
		Where("ext_corp_id  = ?", extCorpID).
		Preload("QuickReplies", " id in (?)", repliesIDs).Preload("QuickReplies.ReplyDetails").
		Where("id in (?)", groupIDs).
		Find(&groups).Error
	return

}

func (qg QuickReplyGroup) GetByIDs(ids []string) (groups []QuickReplyGroup, err error) {
	err = DB.Model(&QuickReplyGroup{}).Where("id in (?)", ids).Find(&groups).Error
	return
}
