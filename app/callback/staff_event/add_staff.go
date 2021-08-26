package staff_event

import (
	"fmt"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"strings"
)

// EventAddStaffHandler
// Description  企业新增员工回调事件处理, 同步员工数据  todo staff_num refresh in tx
// Param msg gowx.RxMessage 回调消息结构体
// return
func EventAddStaffHandler(msg *gowx.RxMessage) error {
	if msg.MsgType != gowx.MessageTypeEvent ||
		msg.Event != gowx.EventTypeChangeContact ||
		msg.ChangeType != gowx.ChangeTypeCreateUser {
		err := errors.New("wrong handler for the callback event")
		log.Sugar.Error("err", err)
		return err
	}
	eventAddExternalContact, ok := msg.EventCreateUser()
	if !ok {
		err := errors.New("msg.EventCreateUser failed")
		log.Sugar.Errorw("get event msg failed", "err", err)
		return err
	}
	extStaffID := eventAddExternalContact.GetUserID()
	extCorpID := conf.Settings.WeWork.ExtCorpID
	err := SyncStaff(extCorpID, extStaffID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	var departmentIDs constants.Int64ArrayField
	err = models.Department{}.AddStaffNum(1, extCorpID, departmentIDs)
	return errors.WithStack(err)
}

func SyncStaff(extCorpID string, ExtStaffID string) error {
	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	info, err := client.Customer.GetUser(ExtStaffID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	staff := models.Staff{
		Model:     models.Model{ID: id_generator.StringID()},
		ExtCorpID: conf.Settings.WeWork.ExtCorpID,
		RoleID:    string(constants.DefaultCorpStaffRoleID),
		RoleType:  string(constants.RoleTypeStaff),
		ExtID:     info.UserID,
		Name:      info.Name,
		Address:   info.Address,
		Alias:     info.Alias,
		// 有0,40,60数值可选，0代表640640正方形头像
		AvatarURL: fmt.Sprint(strings.Trim(info.AvatarURL, "/0"), "/60"),
		Email:     info.Email,
		Gender:    constants.UserGender(info.Gender),
		Status:    constants.UserStatus(info.Status),
		Mobile:    info.Mobile,
		QRCodeURL: info.QRCodeURL,
		DeptIds:   info.DeptIDs,
	}

	err = models.Staff{}.Upsert(staff)
	if err != nil {
		return err
	}

	newStaff, err := models.Staff{}.Get(ExtStaffID, extCorpID, false)
	if err != nil {
		return err
	}

	// 对于老员工进入新部门，这里会丢失员工-部门对应关系，在同步部门信息时补充
	staffDepts := make([]models.StaffDepartment, 0)
	for _, deptInfo := range info.Departments {
		dept, err := models.Department{}.GetByExtID(deptInfo.DeptID, extCorpID)
		if err != nil {
			return err
		}
		staffDepartment := models.StaffDepartment{
			StaffID:         newStaff.ID,
			DepartmentID:    dept.ID,
			ExtCorpID:       extCorpID,
			ExtStaffID:      info.UserID,
			ExtDepartmentID: deptInfo.DeptID,
			Order:           deptInfo.Order,
			IsLeader:        constants.False,
		}
		if deptInfo.IsLeader {
			staffDepartment.IsLeader = constants.True
		}
		staffDepts = append(staffDepts, staffDepartment)
	}
	err = models.StaffDepartment{}.Upsert(staffDepts...)
	if err != nil {
		return err
	}
	return nil
}
