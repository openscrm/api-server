package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

// GroupChatWelcomeMsg
// 入群欢迎语
type GroupChatWelcomeMsg struct {
	ExtCorpModel
	// 文字内容
	Content string `json:"content"`
	// 附件类型
	AttachmentType string `json:"attachment_type"`
	// 附件内容
	Attachment constants.GroupChatWelcomeMsgField `json:"attachment"`
	Timestamp
}

func (o GroupChatWelcomeMsg) Delete(ids []string) (int64, error) {
	res := DB.Model(&GroupChatWelcomeMsg{}).
		Where("id in (?)", ids).
		Delete(&GroupChatWelcomeMsg{})
	return res.RowsAffected, res.Error
}

func (o GroupChatWelcomeMsg) Create(msg GroupChatWelcomeMsg) error {
	return DB.Create(&msg).Error
}

func (o GroupChatWelcomeMsg) Get(id string, extCorpID string) (msg GroupChatWelcomeMsg, err error) {
	err = DB.Model(&GroupChatWelcomeMsg{}).
		Where("id = ?", id).
		Where("ext_corp_id = ?", extCorpID).
		First(&msg).Error
	return
}

func (o GroupChatWelcomeMsg) GetByIDs(ids []string, extCorpID string) (msg GroupChatWelcomeMsg, err error) {
	err = DB.Model(&GroupChatWelcomeMsg{}).
		Where("id in (?)", ids).
		Where("ext_corp_id = ?", extCorpID).
		First(&msg).Error
	return
}

func (o GroupChatWelcomeMsg) Update(msg GroupChatWelcomeMsg) (err error) {
	err = DB.Model(&GroupChatWelcomeMsg{}).
		Where("id = ?", msg.ID).
		Updates(&msg).Error
	return
}
func (o GroupChatWelcomeMsg) Query(
	msg GroupChatWelcomeMsg,
	extCorpID string,
	pager *app.Pager,
	sorter *app.Sorter) (msgs []GroupChatWelcomeMsg, total int64, err error) {

	db := DB.Model(&GroupChatWelcomeMsg{}).Where("ext_corp_id = ?", extCorpID)
	if msg.Content != "" {
		db = db.Where("content like ?", msg.Content+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		err = errors.Wrap(err, "Count GroupChatWelcomeMsg failed")
		return
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&msgs).Error
	if err != nil {
		err = errors.Wrap(err, "Find GroupChatWelcomeMsg failed")
		return
	}

	return
}
