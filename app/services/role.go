package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/id_generator"
)

type Role struct {
	model models.Role
}

func NewRole() *Role {
	return &Role{model: models.Role{}}
}

func (o *Role) Query(req requests.QueryRoleReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []responses.Role, total int64, err error) {
	param := models.Role{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	var roles []models.Role
	roles, total, err = o.model.Query(param, extCorpID, sorter, pager)
	if err != nil {
		err = errors.Wrap(err, "o.model.Query failed")
		return
	}

	counts := make([]struct {
		RoleID string
		Total  int64
	}, 0)

	err = models.DB.Model(&models.Staff{}).Where("ext_corp_id = ?", extCorpID).Group("role_id").Select("count(*) as total,role_id").Scan(&counts).Error
	if err != nil {
		err = errors.Wrap(err, "count Staff failed")
		return
	}

	countMap := make(map[string]int64, 0)
	for _, count := range counts {
		countMap[count.RoleID] = count.Total
	}

	for _, role := range roles {
		item := responses.Role{
			Role:  role,
			Count: countMap[role.ID],
		}
		items = append(items, item)
	}

	return
}

func (o *Role) Get(id string) (item responses.Role, err error) {
	role, err := o.model.Get(id)
	if err != nil {
		err = errors.Wrap(err, "get role failed")
		return
	}

	item.Role = role
	item.Permissions, err = (&models.Permission{}).BatchGetByIdentities(role.PermissionIDs...)
	if err != nil {
		err = errors.Wrap(err, "get Permission failed")
		return
	}

	return
}

func (o *Role) Create(param requests.CreateRoleReq, extCorpID string) (item models.Role, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}
	item.ID = id_generator.StringID()

	return o.model.Create(item, extCorpID)
}

func (o *Role) Update(id string, param requests.UpdateRoleReq) (item models.Role, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	return o.model.Update(id, item)
}

func (o *Role) AssignToStaffs(extStaffIDs []string, roleID string) (total int64, err error) {
	return o.model.AssignToStaffs(extStaffIDs, roleID)
}

func (o *Role) QueryStaffs(req requests.QueryRoleStaffsReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (items []models.Staff, total int64, err error) {
	param := models.Staff{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	param.ExtID = req.ExtStaffID
	param.ID = req.StaffID

	return o.model.QueryStaffs(param, extCorpID, sorter, pager)
}
