package staff_event

import (
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"openscrm/app/models"
	"openscrm/common/log"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
)

// EventUpdateStaffHandler
// Description: 更新员工回调
// Detail:
//	同步员工数据
//	维护部门人数
// Param: msg 解析后的回调消息
func EventUpdateStaffHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeContact ||
		msg.ChangeType != gowx.ChangeTypeUpdateUser {
		err := errors.New("wrong handler for the callback event")
		log.Sugar.Error("err", err)
		return err
	}

	extras, ok := msg.EventUpdateUser()
	if !ok {
		err := errors.New("msg.EventUpdateUser failed")
		log.Sugar.Errorw("get event msg failed", "err", err)
		return err
	}
	extStaffID := extras.GetUserID()
	extCorpID := conf.Settings.WeWork.ExtCorpID
	staff, err := models.Staff{}.Get(extStaffID, extCorpID, false)
	if err != nil {
		return err
	}

	// 同步数据后，添加所在部门的人数
	err = SyncStaff(extCorpID, extStaffID)
	if err != nil {
		return err
	}

	// 找出不同部门，然后新的部门增加人数，少的部门减少人数
	newStaff, err := models.Staff{}.Get(extStaffID, extCorpID, false)
	if err != nil {
		return err
	}

	log.Sugar.Debug(staff.DeptIds.ToInt64Array(), newStaff.DeptIds.ToInt64Array())

	rmDepartmentIDs := funk.Subtract(staff.DeptIds.ToInt64Array(), newStaff.DeptIds.ToInt64Array()).([]int64)
	if len(rmDepartmentIDs) > 0 {
		// 删除员工不在的部门
		err = RemoveStaffDepartment(rmDepartmentIDs, extStaffID, extCorpID)
		if err != nil {
			return err
		}
		// 维护部门员工数据
		err = models.Department{}.AddStaffNum(-1, extCorpID, rmDepartmentIDs)
		if err != nil {
			return err
		}
	}

	newDepartmentIDs := funk.Subtract(newStaff.DeptIds, staff.DeptIds).([]int64)
	if len(newDepartmentIDs) > 0 {
		err = models.Department{}.AddStaffNum(1, extCorpID, newDepartmentIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveStaffDepartment(rmDepartmentIDs []int64, extStaffID string, extCorpID string) error {
	staffDepartments := make([]models.StaffDepartment, 0)
	for _, extDepartmentID := range rmDepartmentIDs {
		staffDepartment := models.StaffDepartment{
			ExtCorpID:       extCorpID,
			ExtStaffID:      extStaffID,
			ExtDepartmentID: extDepartmentID,
		}
		staffDepartments = append(staffDepartments, staffDepartment)
	}
	err := models.StaffDepartment{}.Delete(staffDepartments...)
	if err != nil {
		return err
	}
	return nil
}
