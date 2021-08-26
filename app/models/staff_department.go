package models

import (
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
)

type StaffDepartment struct {
	ExtCorpID       string            `json:"ext_corp_id" gorm:"index;type:char(18);uniqueIndex:idx_ext_corp_id_ext_staff_id"`
	ExtStaffID      string            `json:"ext_staff_id" gorm:"type:char(32);index;uniqueIndex:idx_ext_corp_id_ext_staff_id"`
	ExtDepartmentID int64             `json:"ext_department_id" gorm:"type:int unsigned;uniqueIndex:idx_ext_corp_id_ext_staff_id"`
	StaffID         string            `json:"staff_id" gorm:"primaryKey;type:bigint" `
	DepartmentID    string            `json:"department_id" gorm:"primaryKey;type:bigint;" `
	IsLeader        constants.Boolean `json:"is_leader" gorm:"type:tinyint unsigned;comment:是否是所在部门的领导"`
	Order           uint32            `json:"order" gorm:"type:int unsigned;comment:所在部门的排序"`
}

func (s StaffDepartment) Upsert(sd ...StaffDepartment) error {
	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ext_customer_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"order", "is_leader", "staff_id", "department_id"}),
	}).CreateInBatches(&sd, len(sd)).Error
	if err != nil {
		return err
	}

	return err
}

func (s StaffDepartment) Delete(sd ...StaffDepartment) error {
	for _, staffDepartment := range sd {
		err := DB.Model(&StaffDepartment{}).
			Where("ext_corp_id = ? and ext_staff_id = ? and ext_department_id = ?",
				staffDepartment.ExtCorpID, staffDepartment.ExtStaffID, staffDepartment.ExtDepartmentID).Delete(&StaffDepartment{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
