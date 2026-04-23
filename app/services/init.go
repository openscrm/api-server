package services

import (
	"openscrm/app/models"
	"openscrm/common/log"
	"openscrm/conf"
)

// Syncs 同步所有企业的全部信息
func Syncs() {
	var err error
	departmentService := NewDepartment()
	groupChatService := NewGroupChatService()
	staffService := NewStaffService()
	//customerService := NewCustomer()
	//tagService := NewTag()

	err = departmentService.Sync(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		log.Errorf("[ Wechat][ Error]: sync departments failed: %v", err)
		return
	}

	err = staffService.Sync(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		log.Errorf("[ Wechat][ Error]: sync staff failed: %v", err)
		return
	}

	err = groupChatService.SyncAll(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		log.Errorf("[ Wechat][ Error]: sync group chats failed: %v", err)
		return
	}

	//err = customerService.Sync(conf.Settings.WeWork.ExtCorpID)
	//if err != nil {
	//	panic(err)
	//}

	//err = tagService.Sync(conf.Settings.WeWork.ExtCorpID)
	//if err != nil {
	//	panic(err)
	//}

	models.SetupStaffRole() // 初始化超级管理员权限
}
