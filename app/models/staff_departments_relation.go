package models

type StaffDepartments struct {
	StaffID  string `json:"staff_id" gorm:"comment:微信定义的员工ID"`
	DeptID   int64  `json:"dept_id" gorm:"comment:部门ID"`
	Order    uint32 `json:"order" gorm:"comment:"`
	IsLeader bool   `json:"is_leader" gorm:"comment:在所在的部门内是否为上级"`
	Timestamp
}
