package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/app"
)

// WelcomeMsg 主欢迎语-多个分时欢迎语
// 主欢迎语维护可用员工和部门id列表
type WelcomeMsg struct {
	ExtCorpModel
	// 标题
	Name string `json:"name" gorm:"type:char(128);comment:标题"`
	// 欢迎语内容
	WelcomeMsg constants.AutoReplyField `gorm:"type:json" json:"welcome_msg"`
	// 主欢迎语id
	MainWelcomeMsgID *string `gorm:"type:bigint,comment:主欢迎语id" json:"main_welcome_msg_id"`
	// 启用分时欢迎语
	EnableTimePeriodMsg constants.Boolean `json:"enable_time_period_msg" validate:"oneof=1 2"`
	// 分时段欢迎语
	TimePeriodMsg []WelcomeMsg `gorm:"foreignKey:MainWelcomeMsgID;references:ID" json:"time_period_msg"`
	// 生效时间,n-星期n
	EffectiveAt constants.Int64ArrayField `gorm:"type:json" json:"effective_at"`
	// 分时段欢迎语-开始时间
	StartTime constants.TimeField `gorm:"comment:分时段欢迎语-开始时间" json:"start_time"`
	// 分时段欢迎语-结束时间
	EndTime constants.TimeField `gorm:"comment:分时段欢迎语-结束时间" json:"end_time"`
	Timestamp
}

func (w WelcomeMsg) DeleteTimePeriodMsg(tx *gorm.DB, id string) error {
	return tx.Model(&WelcomeMsg{}).Where(" main_welcome_msg_id = ?", id).Delete(&WelcomeMsg{}).Error
}

type WelcomeMsgWithDeptAndStaff struct {
	WelcomeMsg
	Department []Dept           `json:"department"`
	Staffs     []WelcomeMsgUser `json:"staffs"`
}

type Dept struct {
	ID    string `json:"id"`
	ExtID int64  `json:"ext_id"`
	Name  string `json:"name"`
}

type WelcomeMsgUser struct {
	ID        string `json:"id"`
	AvatarURL string `json:"avatar_url"`
	ExtID     string `json:"ext_id"`
	Name      string `json:"name"`
}

func (w WelcomeMsg) Update(tx *gorm.DB, msg WelcomeMsg) error {
	if len(msg.TimePeriodMsg) > 0 {
		err := tx.Unscoped().Where("main_welcome_msg_id = ?", msg.TimePeriodMsg[0].MainWelcomeMsgID).Delete(&WelcomeMsg{}).Error
		if err != nil {
			return err
		}
	}

	err := tx.Model(&msg).Where("id = ?", msg.ID).Updates(&msg).Error
	if err != nil {
		return err

	}

	return nil
}

func (w WelcomeMsg) Query(req requests.QueryWelcomeMsgReq, id string, sorter *app.Sorter, pager *app.Pager) (msgs []WelcomeMsg, total int64, err error) {
	if len(req.ExtStaffIDs) > 0 {
		rawSQL := `select * from welcome_msg
where id in (
    select staff.welcome_msg_id
    from staff
    where staff.ext_id in (?)
    union
    select d.welcome_msg_id
    from staff
             left join staff_department sd on staff.id = sd.staff_id
             left join department d on sd.department_id = d.id
    where staff.ext_id in (?)
      and staff.welcome_msg_id is not null
      and d.welcome_msg_id is not null
);`
		err = DB.Exec(rawSQL, req.ExtStaffIDs, req.ExtStaffIDs).Take(&msgs).Error
		total = int64(len(msgs))
	} else {
		db := DB.Model(&WelcomeMsg{}).Where("ext_corp_id = ?", id)

		if req.Name != "" {
			db = db.Where("name like ?", req.Name+"%")
		}

		err = db.Count(&total).Error
		if err != nil || total == 0 {
			err = errors.Wrap(err, "Count welcome msg failed")
			return nil, 0, err
		}

		sorter.SetDefault()
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

		pager.SetDefault()
		db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

		err = db.Preload("TimePeriodMsg").Find(&msgs).Error
		if err != nil {
			err = errors.Wrap(err, "Find welcome msg failed")
			return nil, 0, err
		}
	}
	return msgs, total, nil
}

func (w WelcomeMsg) Delete(ids []string, extCorpModel string) error {
	return DB.Model(&WelcomeMsg{}).
		Where("id in (?)", ids).
		Where("ext_corp_id =?", extCorpModel).
		Delete(&WelcomeMsg{}).Error
}

func (w WelcomeMsg) Create(db *gorm.DB, msg WelcomeMsg) error {
	return db.Create(&msg).Error
}

func (w WelcomeMsg) Get(id string, extCorpID string) (msg WelcomeMsg, err error) {
	err = DB.Model(&WelcomeMsg{}).
		Preload("TimePeriodMsg").
		Where("id = ?", id).
		Where("ext_corp_id = ?", extCorpID).
		First(&msg).Error
	return
}

func (m WelcomeMsg) GetByIDs(ids []string, extCorpID string) (msgs []WelcomeMsg, err error) {
	err = DB.Model(&WelcomeMsg{}).
		Preload("TimePeriodMsg").
		Where("id in (?)", ids).
		Where("ext_corp_id = ?", extCorpID).
		First(&msgs).Error
	return
}
