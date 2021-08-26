package controller

import (
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/app"
	"openscrm/common/ecode"
)

type Base struct {
}

// GetStaffAdminInfo 从session中获取企业普通管理员信息
func (o *Base) GetStaffAdminInfo(handler *app.Handler) (staffAdmin models.Staff, err error) {
	err = handler.GetSessionVal(constants.StaffAdminSessionName, constants.StaffInfo, &staffAdmin)
	if err != nil {
		err = errors.Wrap(err, "get session failed")
		return
	}

	if staffAdmin.ID == "" {
		err = ecode.InvalidSessionError
		return
	}

	item, err := (&models.Staff{}).Get(staffAdmin.ExtID, "", false)
	if err != nil {
		err = errors.Wrap(err, "CachedGet failed")
		return
	}

	staffAdmin = *item

	return
}

// GetStaffInfo 从session中获取企业成员信息
func (o *Base) GetStaffInfo(handler *app.Handler) (staffAdmin models.Staff, err error) {
	err = handler.GetSessionVal(constants.StaffSessionName, constants.StaffInfo, &staffAdmin)
	if err != nil {
		err = errors.Wrap(err, "get session failed")
		return
	}

	if staffAdmin.ID == "" {
		err = ecode.InvalidSessionError
		return
	}

	item, err := (&models.Staff{}).Get(staffAdmin.ExtID, "", false)
	if err != nil {
		err = errors.Wrap(err, "CachedGet failed")
		return
	}

	staffAdmin = *item

	return
}
