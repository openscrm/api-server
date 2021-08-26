package services

import (
	"openscrm/app/models"
	"openscrm/conf"
)

// Syncs 同步所有企业的全部信息
func Syncs(config conf.AppConfig) {
	var err error
	if !config.AutoSyncWeWorkData {
		return
	}

	departmentService := NewDepartment()
	groupChatService := NewGroupChatService()
	staffService := NewStaffService()
	customerService := NewCustomer()

	err = departmentService.Sync(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		panic(err)
	}

	err = staffService.Sync(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		panic(err)
	}

	err = groupChatService.SyncAll(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		panic(err)
	}

	err = customerService.Sync(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		panic(err)
	}

	models.SetupStaffRole() // 初始化超级管理员权限
}
