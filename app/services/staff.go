package services

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type StaffService struct {
	staffRepo             models.Staff
	staffDepartmentRepo   models.StaffDepartment
	customerStaffRepo     models.CustomerStaff
	customerStaffTagRepo  models.CustomerStaffTag
	welcomeMsgRepo        models.WelcomeMsg
	deptRepo              models.Department
	tagRepo               models.Tag
	csRelationHistoryRepo models.CustomerStaffRelationHistory
	EventRepo             models.CustomerEvent
	CustomerRepo          models.Customer
}

func NewStaffService() *StaffService {
	return &StaffService{
		staffRepo:             models.Staff{},
		customerStaffRepo:     models.CustomerStaff{},
		welcomeMsgRepo:        models.WelcomeMsg{},
		staffDepartmentRepo:   models.StaffDepartment{},
		deptRepo:              models.Department{},
		customerStaffTagRepo:  models.CustomerStaffTag{},
		tagRepo:               models.Tag{},
		csRelationHistoryRepo: models.CustomerStaffRelationHistory{},
		CustomerRepo:          models.Customer{},
		EventRepo:             models.CustomerEvent{},
	}
}

func GetCorpWxClient(extCorpID string) (we_work.Client, error) {
	return we_work.Clients.Get(extCorpID)
}

// Sync 同步数据的顺序: 部门-员工-客户
func (o StaffService) Sync(extCorpID string) (err error) {
	err = o.SyncStaffByCorp(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	err = o.SyncStaffByApplication(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	return nil
}

// SyncStaffByApplication 第三方应用同步员工数据
func (o StaffService) SyncStaffByApplication(extCorpID string) (err error) {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	departments, err := client.MainApp.ListAllDepartments()
	if err != nil {
		log.Sugar.Error("ListAllDepartments failed", err)
		return err
	}
	for _, department := range departments {
		usersInfo, err := client.Customer.ListUsersByDeptID(department.ID, true /*是否递归查询下级部门*/)
		if err != nil {
			log.Sugar.Error("批量获取部门成员信息失败", err)
			return err
		}

		AuthorizedExtStaffIDs := make([]string, 0)
		for _, info := range usersInfo {
			AuthorizedExtStaffIDs = append(AuthorizedExtStaffIDs, info.UserID)
		}
		err = o.staffRepo.UpdateAuthorizedStatus(AuthorizedExtStaffIDs)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}

func (o StaffService) SyncStaffByCorp(extCorpID string) (err error) {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	departments, err := client.Customer.ListAllDepartments()
	if err != nil {
		log.Sugar.Error("ListAllDepartments failed", err)
		return err
	}
	for _, department := range departments {
		usersInfo, err := client.Customer.ListUsersByDeptID(department.ID, true /*是否递归查询下级部门*/)
		if err != nil {
			log.Sugar.Error("批量获取部门成员信息失败", err)
			return err
		}
		staffs := make([]models.Staff, 0)
		extStaffIDs := make([]string, 0)
		extDeptIDs := make([]int64, 0)
		extDepts := make([]gowx.UserDeptInfo, 0)
		staffDepts := make([]models.StaffDepartment, 0)
		for _, info := range usersInfo {
			staff := models.Staff{
				ExtCorpID:    conf.Settings.WeWork.ExtCorpID,
				RoleID:       string(constants.DefaultCorpStaffRoleID),
				RoleType:     string(constants.RoleTypeStaff),
				ExtID:        info.UserID,
				Name:         info.Name,
				Address:      info.Address,
				Alias:        info.Alias,
				AvatarURL:    fmt.Sprint(strings.Trim(info.AvatarURL, "/0"), "/60"), // 有0,40,60数值可选，0代表640640正方形头像
				Email:        info.Email,
				Gender:       constants.UserGender(info.Gender),
				Status:       constants.UserStatus(info.Status),
				Mobile:       info.Mobile,
				QRCodeURL:    info.QRCodeURL,
				DeptIds:      info.DeptIDs,
				IsAuthorized: constants.False,
			}
			staff.ID = id_generator.StringID()
			staffs = append(staffs, staff)
			extStaffIDs = append(extStaffIDs, staff.ExtID)

			for _, deptInfo := range info.Departments {
				extDeptIDs = append(extDeptIDs, deptInfo.DeptID)
				extDepts = append(extDepts, deptInfo)

				staffDepartment := models.StaffDepartment{
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

		}

		err = o.staffRepo.BatchUpsert(staffs)
		if err != nil {
			return err
		}

		staffIDExtIds, err := o.staffRepo.GetIDsByExtIDs(extStaffIDs)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		deptIDExtIDs, err := o.deptRepo.GetIDsByExtIDs(extDeptIDs)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}

		for i, dept := range staffDepts {
			if staffID, ok := staffIDExtIds[dept.ExtStaffID]; ok {
				staffDepts[i].StaffID = staffID
			}
			if deptID, ok := deptIDExtIDs[dept.ExtDepartmentID]; ok {
				staffDepts[i].DepartmentID = deptID
			}

		}

		err = o.staffDepartmentRepo.Upsert(staffDepts...)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}

	err = o.staffRepo.CleanCache(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return nil
}

func (o StaffService) Query(
	req requests.QueryStaffReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]*models.Staff, int64, error) {
	staff := models.Staff{
		ExtCorpID:     extCorpID,
		RoleID:        req.RoleID,
		RoleType:      req.RoleType,
		Name:          req.Name,
		EnableMsgArch: req.EnableMsgArch,
	}
	if len(req.ExtDepartmentIDs) > 0 {
		staff.DeptIds = req.ExtDepartmentIDs
	}

	return o.staffRepo.Query(staff, extCorpID, sorter, pager)
}

func (o StaffService) Get(extStaffID string, extCorpID string) (*models.Staff, error) {
	return o.staffRepo.Get(extStaffID, extCorpID, true)
}

func (o StaffService) Enable(req requests.EnableStaffs, extCorpID string) error {
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		return err
	}
	if len(req.ExtStaffIDs) > 0 {
		for _, extStaffID := range req.ExtStaffIDs {
			reqUpdateUser := gowx.UpdateUserReq{Userid: extStaffID, Enable: 1}
			_, err := client.Contact.UpdateUser(reqUpdateUser)
			if err != nil {
				return err
			}
		}
	}
	if len(req.ExcludeExtStaffIDs) > 0 {
		for _, extStaffID := range req.ExcludeExtStaffIDs {
			reqUpdateUser := gowx.UpdateUserReq{Userid: extStaffID, Enable: 0}
			_, err := client.Contact.UpdateUser(reqUpdateUser)
			if err != nil {
				return err
			}
		}
	}
	return o.staffRepo.EnableInBatches(req.ExtStaffIDs, req.ExcludeExtStaffIDs, extCorpID)
}

// GetDetail 获取员工详情，包括权限和角色信息
func (o StaffService) GetDetail(extStaffID string, extCorpID string) (item responses.StaffDetail, err error) {
	staff, err := o.staffRepo.Get(extStaffID, extCorpID, false)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = copier.Copy(&item, staff)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	err = models.DB.Where("id = ?", staff.RoleID).First(&item.Role).Error
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

const SecondsPerDay = 24 * 60 * 60

func (o StaffService) StaffInfoStatistics(userID string, extCorpID string) (*models.StaffCustomerCount, error) {
	currentTime := time.Now()
	startTime := time.Date(
		currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	endTime := startTime + SecondsPerDay
	return o.customerStaffRepo.GetStaffByStaffID(startTime, endTime, userID)
}

func (o StaffService) UpdateCustomerInternalTag(
	req requests.UpdateCustomerInternalTagsReq, extCorpID string) error {
	customerStaff := models.CustomerStaff{
		ExtCorpModel:  models.ExtCorpModel{ExtCorpID: extCorpID},
		ExtStaffID:    req.ExtStaffID,
		ExtCustomerID: req.ExtCustomerID,
	}
	customerStaff, err := o.customerStaffRepo.Get(customerStaff)
	if err != nil {
		return err
	}
	if len(req.RemoveTags) > 0 {
		r := funk.Subtract(customerStaff.InternalTagIDs, req.RemoveTags)
		customerStaff.InternalTagIDs = r.(constants.StringArrayField)
	}
	if len(req.AddTags) > 0 {
		for _, tagID := range req.AddTags {
			customerStaff.InternalTagIDs = append(customerStaff.InternalTagIDs, tagID)
		}
	}
	err = o.customerStaffRepo.Update(&models.CustomerStaff{InternalTagIDs: customerStaff.InternalTagIDs}, req.ExtStaffID, req.ExtCustomerID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCustomerTag
// Description: 员工更新客户标签
func (o StaffService) UpdateCustomerTag(req requests.UpdateCustomerTagsReq, extStaffID, extCorpID string) error {
	for _, extCustomerID := range req.ExtCustomerIDs {
		// 查所有客户
		relations, err := o.customerStaffRepo.GetRelationsByExtCustomerID(extCustomerID, "", extCorpID)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		staff, err := o.staffRepo.Get(extStaffID, extCorpID, false)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		customer, err := o.CustomerRepo.GetByExtID(extCustomerID, nil, false)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		extTagIDs := make([]string, 0)
		extTagIDs = append(extTagIDs, req.RemoveExtTagIDs...)
		extTagIDs = append(extTagIDs, req.AddExtTagIDs...)
		tags, err := o.tagRepo.Get(extTagIDs)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		tagsMap := map[string]models.Tag{}
		for _, tag := range tags {
			tagsMap[tag.ExtID] = tag
		}

		ces := make([]models.CustomerEvent, 0)
		for _, relation := range relations {
			if req.RemoveExtTagIDs != nil && len(req.RemoveExtTagIDs) != 0 {
				err = o.customerStaffTagRepo.Delete(relation.ID, req.RemoveExtTagIDs, true)
				if err != nil {
					err = errors.WithStack(err)
					return err
				}

				// 记录事件流水
				for _, extTagID := range req.RemoveExtTagIDs {
					content := fmt.Sprintf(constants.RemoveTagEvent, staff.Name, customer.Name, tagsMap[extTagID].Name)
					ce := o.GenEventModel(extStaffID, extCustomerID, content, constants.CustomerEventCustomerAction, constants.EventNameDeleteExternalUser)
					ces = append(ces, ce)
				}
			}

			if req.AddExtTagIDs != nil && len(req.AddExtTagIDs) != 0 {
				tags, err := o.tagRepo.Query(req.AddExtTagIDs)
				if err != nil {
					err = errors.WithStack(err)
					return err
				}
				var csTags []models.CustomerStaffTag
				for _, tag := range tags {
					csTag := models.CustomerStaffTag{
						ExtCorpModel:    models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
						CustomerStaffID: relation.ID,
						ExtTagID:        tag.ExtID,
						GroupName:       tag.GroupName,
						TagName:         tag.Name,
					}
					csTags = append(csTags, csTag)
				}
				err = o.customerStaffTagRepo.Upsert(csTags)
				if err != nil {
					err = errors.WithStack(err)
					return err
				}
				// 记录事件流水
				for _, extTagID := range req.RemoveExtTagIDs {
					content := fmt.Sprintf(constants.AddTagEvent, staff.Name, customer.Name, tagsMap[extTagID].Name)
					ce := o.GenEventModel(extStaffID, extCustomerID, content, constants.CustomerEventCustomerAction, constants.EventNameAddExternalUser)
					ces = append(ces, ce)
				}
			}
			err = o.EventRepo.CreateInBatches(ces)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o StaffService) GenEventModel(
	extStaffID, extCustomerID, content string, eventType constants.EventType, eventName constants.EventName) (ce models.CustomerEvent) {

	ce = models.CustomerEvent{
		ExtCorpModel: models.ExtCorpModel{
			ID: id_generator.StringID(), ExtCorpID: conf.Settings.WeWork.ExtCorpID, ExtCreatorID: extStaffID},
		Content:       content,
		EventType:     string(eventType),
		EventName:     string(eventName),
		ExtCustomerID: extCustomerID,
		ExtStaffID:    extStaffID,
	}

	return
}

// SyncStaffSignatures
// Description: 通过数据签名的方式同步员工数据
func (o StaffService) SyncStaffSignatures() error {
	//departments, err := we_work.App.ListAllDepartments()
	//if err != nil {
	//	log.Sugar.Error("ListAllDepartments failed", err)
	//	return err
	//}
	//hash := sha1.New()
	//for _, dept := range departments {
	//	usersInfo, err := we_work.App.ListUsersByDeptID(dept.ID, true /*是否递归查询下级部门*/)
	//	if err != nil {
	//		log.Sugar.Error("批量获取部门成员信息失败", err)
	//		return err
	//	}
	//	ids := make([]string, 0)
	//	signatures := make([]string, 0)
	//	// get updated staffRepo info
	//	idsSignatureMap := map[string]string{}
	//	for _, info := range usersInfo {
	//		log.Sugar.Debug("usersInfo", util2.JsonEncode(info))
	//		ids = append(ids, info.UserID)
	//		marshal, err := json.Marshal(&info)
	//		if err != nil {
	//			log.Sugar.Error("marshal info failed", err)
	//			return err
	//		}
	//		_, err = hash.Write(marshal)
	//		if err != nil {
	//			log.Sugar.Error("sha1 write info failed", err)
	//			return err
	//		}
	//		bs := hash.Sum(nil)
	//		signatureString := hex.EncodeToString(bs)
	//		signatures = append(signatures, signatureString)
	//		hash.Reset()
	//		idsSignatureMap[info.UserID] = signatureString
	//	}
	//	updatedIDs, err := s.staffRepo.GetStaffByIDSAndSignatures(ids, signatures)
	//	if err != nil {
	//		log.Sugar.Error("get updated staffRepo ids failed", err)
	//		return err
	//	}
	//	for _, id := range updatedIDs {
	//		for _, info := range usersInfo {
	//			if info.UserID == id {
	//				var updateStaff models.Staff
	//				var staffDepartments []models.StaffDepartments
	//				for _, department := range info.Departments {
	//					var staffDepartment models.StaffDepartments
	//					err = copier.Copy(&staffDepartment, department)
	//					if err != nil {
	//						log.Sugar.Error("get staffRepo department failed", err)
	//						return err
	//					}
	//					staffDepartment.ExtStaffID = info.UserID
	//					staffDepartments = append(staffDepartments, staffDepartment)
	//				}
	//				//updateStaff.StaffDepartments = staffDepartments
	//				err = copier.Copy(&updateStaff, info)
	//				if err != nil {
	//					log.Sugar.Error("copy staffRepo from api result failed", err)
	//					return err
	//				}
	//				updateStaff.Signature = idsSignatureMap[id]
	//				err = s.staffRepo.UpdateWelcomeMsg(id, updateStaff)
	//				if err != nil {
	//					log.Sugar.Error("copy staffRepo from api result failed", err)
	//					return err
	//				}
	//			}
	//		}
	//	}
	//
	//	// get not stored users' info
	//	var allUserIds []string
	//	allUserIds, err = s.staffRepo.GetAllStaffIDs()
	//	if err != nil {
	//		log.Sugar.Error("get all user ids failed", err)
	//		return err
	//	}
	//	allIDsMaps := map[string]int{}
	//	for _, id := range allUserIds {
	//		allIDsMaps[id] = 1
	//	}
	//	newStaffs := make([]models.Staff, 0)
	//	for _, info := range usersInfo {
	//		if _, ok := allIDsMaps[info.UserID]; !ok {
	//			staff := models.Staff{}
	//			err := copier.Copy(&staff, info)
	//			if err != nil {
	//				log.Sugar.Error(err)
	//				return err
	//			}
	//			staff.ID = id_generator.StringID()
	//			staff.Signature = idsSignatureMap[staff.ExtID]
	//
	//			var staffDepartments []models.StaffDepartments
	//			for _, department := range info.Departments {
	//				var staffDepartment models.StaffDepartments
	//				err = copier.Copy(&staffDepartment, department)
	//				if err != nil {
	//					log.Sugar.Error("get staffRepo department failed", err)
	//					return err
	//				}
	//				staffDepartment.ExtStaffID = info.UserID
	//				staffDepartments = append(staffDepartments, staffDepartment)
	//			}
	//			//staff_event.StaffDepartments = staffDepartments
	//			staff.ExtID = info.UserID
	//			staff.ExtCorpID = conf.Settings.App.ExtCorpID
	//			newStaffs = append(newStaffs, staff)
	//		}
	//	}
	//	err = s.staffRepo.CreateStaffInBatches(newStaffs)
	//	if err != nil {
	//		log.Sugar.Error("create new staffRepo failed", err)
	//		return err
	//	}
	//}
	return nil
}

// SendWelcomeMsg
// Description: 执行发送欢迎语的流程
// Detail: 查DB中的欢迎语,发送到wx
func (o StaffService) SendWelcomeMsg(
	welcomeMsg constants.AutoReplyField, welcomeCode string, extCorpID string, extStaffID string) error {

	wxClient, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		log.Sugar.Errorw("get wx client failed", "err", err)
		return err
	}

	attachments := make([]gowx.Attachments, len(welcomeMsg.Attachments))
	err = copier.CopyWithOption(&attachments, welcomeMsg.Attachments, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	req := gowx.SendWelcomeMsgReq{
		Attachments: attachments,
		Text:        gowx.Text{Content: welcomeMsg.Text},
		WelcomeCode: welcomeCode,
	}

	_, err = wxClient.Customer.SendWelcomeMsg(req)
	if err != nil {
		log.Sugar.Errorw("send welcome msg failed", "err", err)
		return err
	}
	return nil
}

// SendDefaultWelcomeMsg
// Description: 发送默认欢迎语
// Detail: 渠道码没有欢迎语,使用默认欢迎语
func (o StaffService) SendDefaultWelcomeMsg(
	welcomeCode string, extCorpID string, extStaffID string) error {
	welcomeMsg, err := o.staffRepo.GetWelcomeMsgByExtStaffID(extStaffID, extCorpID)
	if err != nil {
		log.Sugar.Errorw("s.staffRepo.GetWelcomeMsgByExtStaffID failed", "err", err)
		return err
	}
	return o.SendWelcomeMsg(welcomeMsg.WelcomeMsg, welcomeCode, extCorpID, extStaffID)
}

// QueryMainInfo
// Description: 查询员工的简要信息,有缓存
func (o StaffService) QueryMainInfo(
	req requests.QueryMainStaffInfoReq, extCorpID string, pager *app.Pager) (models.StaffsMainInfoCache, error) {
	pager.SetDefault()
	return o.staffRepo.CachedQueryMainInfo(req, extCorpID, pager)
}

// ExportDeleteCustomerList
// Description: 导出删人提醒记录
func (o StaffService) ExportDeleteCustomerList(
	req requests.QueryStaffDeleteCustomerHistoryReq, extCorpID string) (*bytes.Buffer, string, error) {

	log.Sugar.Debugw("ExportDeleteCustomerList", "req", req)

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeDeleteCustomerWarning),
		time.Now().Format(constants.DateLayout),
		extCorpID,
		constants.DataExportDeleteCustomerFilenamePrefix+".xlsx",
	)
	fullPath := filepath.Join(
		conf.Settings.Storage.LocalRootPath,
		filename,
	)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	file.NewSheet(constants.DataExportDeleteCustomerListSheetName)
	file.DeleteSheet("Sheet1")
	titles := []string{"删除客户", "操作人", "删除时间", "添加好友时间", "unionid"}
	err := PrettifySheet(constants.DataExportDeleteCustomerListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return nil, "", err
	}

	// 按page分页写入文件
	pager := &req.Pager
	customerLossInfo, n, err := o.csRelationHistoryRepo.QueryStaffDeleteCustomer(req, extCorpID, pager, &req.Sorter)
	if err != nil {
		log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
		return nil, "", err
	}
	for n > 0 && len(customerLossInfo) > 0 {
		for k, v := range customerLossInfo {
			values := []string{
				v.ExtCustomerName,
				v.StaffName,
				v.RelationDeleteAt.Format(constants.DateTimeLayout),
				v.RelationCreateAt.Format(constants.DateTimeLayout),
				"",
			}
			err = file.SetSheetRow(constants.DataExportDeleteCustomerListSheetName, fmt.Sprint("A", k+3), &values)
			if err != nil {
				log.Sugar.Errorw("write excel failed", "err", err)
				return nil, "", err
			}
		}

		pager.Page += 1
		customerLossInfo, n, err = o.csRelationHistoryRepo.QueryStaffDeleteCustomer(req, extCorpID, pager, &req.Sorter)
		if err != nil {
			log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
			return nil, "", err
		}
	}
	buf, err := file.WriteToBuffer()

	return buf, filename, nil
}

// UpdateStaffMsgArchStatus 每15分钟刷新员工开通会话存档的状态
// 1-会话内容存档办公版 2- 会话内容存档服务版 3-会话内容存档企业版
func (o StaffService) UpdateStaffMsgArchStatus() error {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	// todo 区分版本
	// 1-会话内容存档办公版 2- 会话内容存档服务版 3-会话内容存档企业版
	for i := 0; i < 3; i++ {
		err := o.doUpdateMsgArchStatus(client, extCorpID, i)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}

func (o StaffService) doUpdateMsgArchStatus(client we_work.Client, extCorpID string, i int) error {
	staffs, err := client.Customer.ListMsgAuditPermitUser(gowx.MsgAuditEdition(i))
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	err = o.staffRepo.UpdateStaffMsgArchStatus(extCorpID, staffs, constants.True)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return err
}

// GetCustomerSummary
// Description: 统计 首页员工客户和客户群数量
// Detail: 使用缓存
// Param: extStaffID
// return: models.CustomerSummary 客户数量
func (o StaffService) GetCustomerSummary(extStaffID string, extCorpID string) (models.CustomerSummary, error) {
	return o.staffRepo.CachedGetCustomerSummary(extStaffID, extCorpID)
}
