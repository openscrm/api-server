package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
)

type GroupChatAutoJoin struct {
	groupChatAutoCreateRepo models.GroupChatAutoJoinCode
	groupChatGroupRepo      models.GroupChatGroup
}

func NewGroupChatAutoJoin() *GroupChatAutoJoin {
	return &GroupChatAutoJoin{
		groupChatAutoCreateRepo: models.GroupChatAutoJoinCode{},
		groupChatGroupRepo:      models.GroupChatGroup{},
	}
}

func (o *GroupChatAutoJoin) Create(
	req requests.CreateGroupChatAutoJoinCodeReq, extCorpID string) (autoCreateCode models.GroupChatAutoJoinCode, err error) {

	err = copier.Copy(&autoCreateCode, req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	autoCreateCode.ExtCorpID = extCorpID
	autoCreateCode.ID = id_generator.StringID()
	autoCreateCode.State = constants.GroupChatAutoCreateCodeStatePrefix + autoCreateCode.ID

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return
	}

	autoCreateCode.ConfigID, err = client.Customer.AddContactWay(workwx.AddContactWay{
		IsTemp:     false,
		Remark:     autoCreateCode.Remark,
		Scene:      2,
		SkipVerify: autoCreateCode.SkipVerify == constants.True,
		State:      autoCreateCode.State,
		Type:       workwx.ContactWayTypeMultiple,
		User:       autoCreateCode.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx AddContactWay failed")
		return
	}

	autoCreateCode.Staffs, err = o.genAutoJoinCodeStaffs(autoCreateCode.ID, req.Staffs, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	autoCreateCode.BackupStaffs, err = o.genAutoJoinCodeBackupStaffs(autoCreateCode.ID, req.BackupStaffs, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	wxContactWay, err := client.Customer.GetContactWay(autoCreateCode.ConfigID)
	if err != nil {
		err = errors.Wrap(err, "wx GetContactWay failed")
		err = errors.WithStack(err)
		return
	}

	autoCreateCode.QrCode = wxContactWay.QrCode
	err = o.groupChatAutoCreateRepo.Create(autoCreateCode)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (o *GroupChatAutoJoin) Delete(ids []string, extCorpID string) (total int64, err error) {
	total, err = o.groupChatAutoCreateRepo.Delete(ids, extCorpID)
	return
}

func (o *GroupChatAutoJoin) Query(
	req requests.QueryGroupChatAutoJoinReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]models.GroupChatAutoJoinCode, int64, error) {

	param := models.GroupChatAutoJoinCode{}
	err := copier.Copy(&param, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		err = errors.WithStack(err)
		return nil, 0, err
	}

	return o.groupChatAutoCreateRepo.Query(param, extCorpID, sorter, pager)
}

func (o *GroupChatAutoJoin) Update(
	id string, req requests.UpdateGroupChatAutoJoinQrCodeReq, extCorpID string) (autoJoinCode models.GroupChatAutoJoinCode, err error) {

	err = copier.Copy(&autoJoinCode, req)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		err = errors.WithStack(err)
		return
	}
	autoJoinCode.ExtCorpID = extCorpID
	autoJoinCode.ID = id
	autoJoinCode.State = constants.GroupChatAutoCreateCodeStatePrefix + autoJoinCode.ID

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		err = errors.WithStack(err)
		return
	}

	autoJoinCode.ConfigID, err = client.Customer.AddContactWay(workwx.AddContactWay{
		IsTemp:     false,
		Remark:     autoJoinCode.Remark,
		Scene:      2,
		SkipVerify: autoJoinCode.SkipVerify == constants.True,
		State:      autoJoinCode.State,
		Type:       workwx.ContactWayTypeMultiple,
		User:       autoJoinCode.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx AddContactWay failed")
		err = errors.WithStack(err)
		return
	}

	autoJoinCode.Staffs, err = o.genAutoJoinCodeStaffs(autoJoinCode.ID, req.Staffs, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	autoJoinCode.BackupStaffs, err = o.genAutoJoinCodeBackupStaffs(autoJoinCode.ID, req.BackupStaffs, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	wxContactWay, err := client.Customer.GetContactWay(autoJoinCode.ConfigID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	autoJoinCode.QrCode = wxContactWay.QrCode
	newAutoCreateCode, err := o.groupChatAutoCreateRepo.Update(id, autoJoinCode, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return *newAutoCreateCode, err
}

func (o *GroupChatAutoJoin) genAutoJoinCodeStaffs(
	autoJoinQRCodeID string, req []requests.GroupChatAutoJoinCodeStaffParam, extCorpID string) ([]models.GroupChatAutoJoinCodeStaff, error) {
	// 处理绑定员工
	staffs := make([]models.Staff, 0)
	extStaffIDs := make([]string, 0)
	paramMap := make(map[string]requests.GroupChatAutoJoinCodeStaffParam, 0)
	for _, param := range req {
		extStaffIDs = append(extStaffIDs, param.ExtStaffID)
		paramMap[param.ExtStaffID] = param
	}

	err := models.DB.Where("ext_id in (?)", extStaffIDs).Find(&staffs).Error
	if err != nil {
		err = errors.Wrap(err, "Find staffs failed")
		return nil, err
	}

	items := make([]models.GroupChatAutoJoinCodeStaff, 0)
	for _, staff := range staffs {
		param, ok := paramMap[staff.ID]
		if !ok {
			err = errors.New("invalid params")
			return nil, err
		}

		contactWayStaff := models.GroupChatAutoJoinCodeStaff{
			ExtCorpModel: models.ExtCorpModel{
				ID:        id_generator.StringID(),
				ExtCorpID: extCorpID,
			},
			GroupChatAutoJoinCodeID: autoJoinQRCodeID,
			Avatar:                  staff.AvatarURL,
			DayAddCustomerCount:     0,
			DayAddCustomerLimit:     param.DayAddCustomerLimit,
			StaffID:                 staff.ID,
			ExtStaffID:              staff.ExtID,
			Name:                    staff.Name,
		}

		items = append(items, contactWayStaff)
	}

	return items, nil

}

func (o *GroupChatAutoJoin) genAutoJoinCodeBackupStaffs(
	autoJoinQRCodeID string, params []requests.GroupChatAutoJoinCodeStaffParam, extCorpID string) (items []models.GroupChatAutoJoinBackupStaff, err error) {
	contactWayStaffs, err := o.genAutoJoinCodeStaffs(autoJoinQRCodeID, params, extCorpID)
	if err != nil {
		err = errors.Wrap(err, "makeRefStaffs failed")
		return
	}

	err = copier.Copy(&items, contactWayStaffs)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}

	return
}

func (o *GroupChatAutoJoin) BatchRegroup(req requests.BatchRegroupReq, extCorpID string) error {
	_, err := o.groupChatGroupRepo.Get(req.NewGroupID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Sugar.Errorw("new group not exists", "groupID", req.NewGroupID)
			return ecode.GroupChatNotExistsError
		}
		return err
	}
	err = o.groupChatAutoCreateRepo.UpdateGroupID(req.IDs, req.NewGroupID)
	if err != nil {
		log.Sugar.Errorw(" update group chat group in batches failed ", "err", err)
		return err
	}
	return nil
}
