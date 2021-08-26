package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/common/app"
)

type Department struct {
	Model
	// 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"index;uniqueIndex:idx_ext_corp_id_ext_dept_id;type:char(18);comment:外部企业ID"`
	// 企业微信部门id
	ExtID int64 `gorm:"type:int;uniqueIndex:idx_ext_corp_id_ext_dept_id;comment:企微定义的部门ID" json:"ext_id"`
	// 部门名称
	Name string `gorm:"type:varchar(255);comment:部门名称" json:"name"`
	// 上级部门id
	ExtParentID int64 `gorm:"type:int unsigned;comment:上级部门ID,根部门为1" json:"ext_parent_id"`
	// 在上级部门中的排序
	Order uint32 `gorm:"type:int unsigned;comment:在父部门中的次序值" json:"order"`
	// 欢迎语id
	WelcomeMsgID *string `gorm:"type:bigint;comment:部门使用的欢迎语" json:"welcome_msg_id"`
	// 直属下级部门, 不为空时前端可请求获取其子部门信息
	SubDepartments []Department `gorm:"-" json:"sub_departments"`
	StaffNum       int64        `gorm:"type:int;comment:成员数量;default:0" json:"staff_num"`
	Timestamp
}

// GetMainInfoByMsgID
// Description: 查给定欢迎语的部门
func (o Department) GetMainInfoByMsgID(msgID string) ([]Dept, error) {
	res := make([]Dept, 0)
	err := DB.Model(&Department{}).Where("welcome_msg_id = ?", msgID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RemoveOriginalWelcomeMsg 移除旧的欢迎语
func (o Department) RemoveOriginalWelcomeMsg(tx *gorm.DB, extCorpID, msgID string) error {
	return tx.Model(&Department{}).
		Where("ext_corp_id = ?", extCorpID).
		Where("welcome_msg_id = ?", msgID).
		Update("welcome_msg_id", nil).Error

}

func (o Department) UpdateWelcomeMsg(tx *gorm.DB, msgID, extCorpID string, extDeptID []int64) error {
	return tx.Model(&Department{}).
		Omit("id", "ext_id").
		Where("ext_corp_id = ?", extCorpID).
		Where("ext_id in (?)", extDeptID).
		Update("welcome_msg_id", msgID).Error
}

func (o Department) GetByExtID(extID int64, extCorpID string) (Department, error) {
	var dept Department
	err := DB.Model(&Department{}).Where("ext_id = ?", extID).Where("ext_corp_id = ?", extCorpID).First(&dept).Error
	if err != nil {
		return Department{}, err
	}
	return dept, nil
}

func (o Department) GetByExtIDs(extIDs []int64, extCorpID string) (depts []Department, err error) {
	err = DB.Model(&Department{}).Where("ext_id in (?)", extIDs).Where("ext_corp_id = ?", extCorpID).First(&depts).Error
	if err != nil {
		return
	}
	return
}

func (o Department) Get(ExtDeptID int64, extCorpID string) (Department, error) {
	var rootDept Department
	db := DB.Model(&Department{}).
		Where("ext_corp_id =?", extCorpID)

	if ExtDeptID != 0 {
		db = db.Where("ext_id = ?", ExtDeptID)
	}

	//parentDepts := []*Department{rootDept}
	//subDepts := make([]Department, 0)
	//
	//for len(parentDepts) != 0 {
	//	department := parentDepts[0]
	//	parentDepts = parentDepts[1:]
	//
	//	err = DB.Model(&Department{}).
	//		Where("ext_corp_id = ? ", extCorpID).
	//		Where("ext_parent_id in (?)", department.ExtID).
	//		Find(&subDepts).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//	department.SubDepartments = subDepts
	//
	//	for i := range subDepts {
	//		parentDepts = append(parentDepts, &subDepts[i])
	//	}
	//}
	return rootDept, nil
}

func (o Department) Upsert(departments ...Department) error {
	tx := DB.Begin()
	defer tx.Rollback()
	for _, department := range departments {
		err := DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ext_corp_id"}, {Name: "ext_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "order", "ext_parent_id"}),
		}).Create(&department).Error
		if err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func (o Department) Query(extPid int64, extCorpID string, IDs []int64, sorter *app.Sorter, pager *app.Pager) (res []Department, total int64, err error) {

	db := DB.Model(&Department{}).Where("ext_corp_id = ?", extCorpID)
	if extPid != 0 {
		db = db.Where("ext_parent_id = ?", extPid)
	}
	if len(IDs) != 0 {
		db = db.Where("ext_id in (?)", IDs)
	}

	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count Department failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(sorter.SortField)}, Desc: sorter.SortType == constants.SortTypeDesc})

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	err = db.Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

func (o Department) AddStaffNum(num int, ExtCorpID string, ExtDepartmentIDs []int64) error {
	return DB.Model(&Department{}).
		Where("ext_corp_id = ? and ext_id in (?)", ExtCorpID, ExtDepartmentIDs).
		Update("staff_num", gorm.Expr("staff_num + (?)", num)).Error
}

func (o Department) GetIDsByExtIDs(extDeptIDs []int64) (res map[int64]string, err error) {
	type IDExtID struct {
		ID    string `json:"id"`
		ExtID int64  `json:"ext_id"`
	}
	var ids []IDExtID
	res = make(map[int64]string, 0)
	err = DB.Model(&Department{}).Where("ext_id in (?)", extDeptIDs).Find(&ids).Error
	if err != nil {
		return
	}
	for _, id := range ids {
		res[id.ExtID] = id.ID
	}
	return
}
