package services

import (
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
)

type Department struct {
	model models.Department
}

// Query
// Description: 查部门列表
func (d Department) Query(req entities.GetSubDepartmentReq, extCorpID string) (res []models.Department, total int64, err error) {
	res = make([]models.Department, 0)
	res, total, err = d.model.Query(req.ExtParentId, extCorpID, req.ExtIDs, &req.Sorter, &req.Pager)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (d Department) Sync(extCorpID string) error {
	wxClient, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		return err
	}
	depts, err := wxClient.Customer.ListAllDepartments()
	if err != nil {
		log.Sugar.Error("get all departments from wx failed", err)
		return err
	}

	// 只维护企业微信的部门关系（企业内部）
	departments := make([]models.Department, 0)

	for _, dept := range depts {
		var department models.Department
		department.ID = id_generator.StringID()
		department.ExtCorpID = extCorpID
		department.ExtID = dept.ID
		department.Name = dept.Name
		department.ExtParentID = dept.ParentID
		department.Order = dept.Order
		departments = append(departments, department)
	}

	err = d.model.Upsert(departments...)
	if err != nil {
		log.Sugar.Error("create department failed", err)
		return err
	}

	return nil
}

func (d Department) Get(extDeptID int64, extCorpID string) (models.Department, error) {
	return d.model.Get(extDeptID, extCorpID)
}

func NewDepartment() *Department {
	return &Department{model: models.Department{}}
}
