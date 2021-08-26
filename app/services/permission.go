package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/app"
)

type Permission struct {
	model models.Permission
}

func NewDefaultPermission() *Permission {
	return &Permission{model: models.Permission{}}
}

func NewPermission(m models.Permission) *Permission {
	return &Permission{model: m}
}

func (o *Permission) Query(req entities.QueryPermissionReq, sorter *app.Sorter, pager *app.Pager) (items []models.Permission, total int64, err error) {
	param := models.Permission{}
	err = copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	return o.model.Query(param, sorter, pager)
}

func (o *Permission) Get(id string) (item models.Permission, err error) {
	return o.model.Get(id)
}
