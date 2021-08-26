package services

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"net/url"
	"openscrm/app/constants"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/common/we_work"
	"openscrm/conf"
	"openscrm/pkg/easywework"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CustomerService struct {
	db                   *gorm.DB
	customerInfoRepo     models.CustomerInfo
	customerRepo         models.Customer
	customerStaffRepo    models.CustomerStaff
	customerStaffTagRepo models.CustomerStaffTag
	eventNotifyRepo      models.EventNotify
	innerTag             models.InternalTag
	ceRepo               models.CustomerEvent
	csRepo               models.CustomerStatistic
	relationHistory      models.CustomerStaffRelationHistory
}

func NewCustomer() *CustomerService {
	return &CustomerService{
		customerRepo:         models.Customer{},
		customerInfoRepo:     models.CustomerInfo{},
		customerStaffRepo:    models.CustomerStaff{},
		eventNotifyRepo:      models.EventNotify{},
		customerStaffTagRepo: models.CustomerStaffTag{},
		innerTag:             models.InternalTag{},
		csRepo:               models.CustomerStatistic{},
		relationHistory:      models.CustomerStaffRelationHistory{},
	}
}

// Sync
// Description: 同步企业所有员工的客户数据
// Detail:
//	客户数据：客户信息，客户员工关系，员工给客户的标签
//	客户表,员工-客户关系表,员工-客户标签表,员工-客户画像表
// Param: extCorpID 外部企业ID
// return: err
func (o CustomerService) Sync(extCorpID string) error {
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	staffs, err := client.Customer.ListUsersByDeptID(1, true)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	totalExtCustomerIDs, err := o.BatchFetchCustomers(staffs)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	log.Sugar.Debug(totalExtCustomerIDs)

	// 客户外部IDs 列表去重
	uniqueExtCustomerIDs, err := Deduplicate(totalExtCustomerIDs)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 员工-客户关系
	contactInfoArr, err := o.BatchFetchContactInfo(uniqueExtCustomerIDs)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 批量upsert客户
	customerModels := make([]models.Customer, 0)
	for _, contactInfo := range contactInfoArr {
		customerModel, err := NewCustomerModel(contactInfo)
		if err == ecode.EmptyExternalContactInfoErr {
			continue
		}
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		customerModels = append(customerModels, customerModel)
	}
	err = o.customerRepo.BatchUpsert(customerModels)
	if err != nil {
		return err
	}

	// 批量upsert 客户员工关系
	csRelationModels := make([]models.CustomerStaff, 0)
	tags := make([]models.CustomerStaffTag, 0)
	for _, contactInfo := range contactInfoArr {
		csRelationModel, err := NewCustomerStaffModel(contactInfo)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		csRelationModels = append(csRelationModels, csRelationModel...)

	}
	_, err = o.customerStaffRepo.BatchUpsert(csRelationModels)
	if err != nil {
		return err
	}

	// 客户-员工外部ID hash
	signatures := make([]string, 0)
	for _, relation := range csRelationModels {
		signatures = append(signatures, relation.Signature)
	}
	idSigns, err := o.customerStaffRepo.QueryIDs(signatures)
	if err != nil {
		return err
	}

	for _, idSign := range idSigns {
		for k, relation := range csRelationModels {
			if relation.Signature == idSign.Signature {
				for i, _ := range relation.CustomerStaffTags {
					csRelationModels[k].CustomerStaffTags[i].CustomerStaffID = idSign.ID
				}

				tags = append(tags, relation.CustomerStaffTags...)
			}
		}
	}

	// 员工-客户标签
	err = o.customerStaffTagRepo.Upsert(tags)
	if err != nil {
		return err
	}

	return nil
}

func Deduplicate(extCustomerIDs []string) (uniqueExtCustomerIDs []string, err error) {
	if len(extCustomerIDs) == 0 {
		err = ecode.EmptyExternalContactInfoErr
		return
	}
	mapSet := make(map[string]struct{}, 0)
	for _, id := range extCustomerIDs {
		mapSet[id] = struct{}{}
	}
	for k, _ := range mapSet {
		uniqueExtCustomerIDs = append(uniqueExtCustomerIDs, k)
	}

	return
}

// NewCustomerStaffModel
// Description: wx返回的客户数据构造DB记录
func NewCustomerStaffModel(contactInfo *workwx.ExternalContactInfo) (csRelations []models.CustomerStaff, err error) {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	for _, staff := range contactInfo.FollowUser {
		csRelation := models.CustomerStaff{
			ExtCustomerID:  contactInfo.ExternalContact.ExternalUserid,
			ExtStaffID:     staff.UserID,
			Remark:         staff.Remark,
			Description:    staff.Description,
			Createtime:     time.Unix(int64(staff.Createtime), 0),
			RemarkCorpName: staff.RemarkCorpName,
			RemarkMobiles:  constants.StringArrayField(staff.RemarkMobiles),
			AddWay:         constants.FollowUserAddWay(staff.AddWay),
			OperUserID:     staff.OperUserID,
			State:          staff.State,
			ExtCorpModel:   models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: staff.UserID},
		}

		var encrypt string
		encrypt, err = gmd5.Encrypt(csRelation.ExtCustomerID + csRelation.ExtStaffID)
		if err != nil {
			return
		}
		csRelation.Signature = encrypt

		tags := make([]models.CustomerStaffTag, 0)
		for _, tag := range staff.Tags {
			customerStaffTag := models.CustomerStaffTag{
				ExtCorpModel:    models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
				CustomerStaffID: csRelation.ID, // 如果relation是更新,这里的CustomerStaffID值需要被替换,否则找不到外键报错.
				ExtTagID:        tag.TagID,
				GroupName:       tag.GroupName,
				TagName:         tag.TagName,
				Type:            constants.FollowUserTagType(tag.Type),
			}
			tags = append(tags, customerStaffTag)
		}

		csRelation.CustomerStaffTags = tags
		csRelations = append(csRelations, csRelation)
	}
	return
}

// SyncSingleCustomerData
// Description: 同步单个客户数据
// Detail: 用客户-员工外部ID 同步客户数据
func (o CustomerService) SyncSingleCustomerData(extStaffID, extCustomerID string) error {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 拉取外部联系人数据
	contactInfo, err := client.Customer.GetExternalContact(extCustomerID)
	if err != nil {
		log.Sugar.Error("GetExternalContact failed", err)
		return err
	}
	log.Sugar.Info(util.JsonEncode(contactInfo))

	// 新建外部联系人DB记录
	customer, err := NewCustomerModel(contactInfo)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 更新外部联系人记录
	err = o.customerRepo.Upsert(customer)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 更新关联的员工客户关系记录
	// followUser: 与该客户是好友关系的员工
	for _, followUser := range contactInfo.FollowUser {
		if followUser.UserID == extStaffID {
			// 更新员工-客户关系
			customerStaffRelation, err := o.UpdateCustomerStaffRelation(followUser, extCustomerID)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}

			// 更新客户标签
			err = o.UpdateCustomerTags(customerStaffRelation.ID, followUser.Tags)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}

			// upsert 客户画像
			err = o.UpsertCustomerPortrait(extCustomerID, followUser.UserID)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		}
	}
	return nil
}

// thumbAvatar
// Description: 缩小原头像
// Detail: 有0、46、64、96、132数值可选，0代表640640正方形头像
// Param: 原头像url
// return: 新头像url
func thumbAvatar(avatar string) (newAvatar string, err error) {
	URLParsed, err := url.Parse(avatar)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if URLParsed.Host == constants.WeWorkPicHost {
		newAvatar = fmt.Sprint(strings.Trim(avatar, "/0"), "/60")
	} else if URLParsed.Host == constants.WxPicHost {
		newAvatar = fmt.Sprint(strings.Trim(avatar, "/0"), "/64")
	} else {
		err = errors.New("unknown wx host")
	}
	return
}

// Query
// Description: 查询客户条件
func (o CustomerService) Query(
	req requests.QueryCustomerReq, extCorpID string, pager *app.Pager) (customers []*models.Customer, total int64, err error) {

	customers = make([]*models.Customer, 0)
	customers, total, err = o.customerRepo.Query(req, extCorpID, pager)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// Get
// Description: 客户详情
// Detail: 用外部客户id查详情
func (o CustomerService) Get(extID string, extCorpID string, withStaff bool) (models.Customer, error) {
	return o.customerRepo.GetByExtID(extID, nil, true)
}

// UpdateCustomerStaffRelation
// Description: 更新客户员工关系
// Detail:	已经删除的关系记录,须将 deleted_at 置空
func (o CustomerService) UpdateCustomerStaffRelation(staff workwx.FollowUser, extCustomerID string) (relation models.CustomerStaff, err error) {

	updateFields := models.CustomerStaff{
		ExtCustomerID:  extCustomerID,
		ExtStaffID:     staff.UserID,
		Remark:         staff.Remark,
		Description:    staff.Description,
		Createtime:     time.Unix(int64(staff.Createtime), 0),
		RemarkCorpName: staff.RemarkCorpName,
		RemarkMobiles:  constants.StringArrayField(staff.RemarkMobiles),
		AddWay:         constants.FollowUserAddWay(staff.AddWay),
		OperUserID:     staff.OperUserID,
		State:          staff.State,
	}
	relation = updateFields
	extCorpID := conf.Settings.WeWork.ExtCorpID
	relation.ExtCorpModel = models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: staff.UserID}
	relation.DeletedAt = gorm.DeletedAt{}

	n := models.DB.Model(&models.CustomerStaff{}).Unscoped().
		Where("ext_customer_id = ? and ext_staff_id = ?", extCustomerID, staff.UserID).
		Updates(&updateFields).RowsAffected
	if n == 0 {
		if err = o.customerStaffRepo.Create(&relation); err != nil {
			log.Sugar.Error("create customer_event failed", err)
			return
		}
	} else {
		relation, err = o.customerStaffRepo.Get(relation)
		if err != nil {
			return
		}
	}
	return
}

// NewCustomerModel
// Description: 根据微信客户数据生成models.Customer
func NewCustomerModel(externalContactInfo *workwx.ExternalContactInfo) (customer models.Customer, err error) {
	if externalContactInfo == nil {
		err = ecode.EmptyExternalContactInfoErr
		return
	}

	customerInfo := externalContactInfo.ExternalContact
	newAvatar, err := thumbAvatar(customerInfo.Avatar)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	extCorpID := conf.Settings.WeWork.ExtCorpID
	customer = models.Customer{
		ExtCorpModel: models.ExtCorpModel{
			ID:           id_generator.StringID(),
			ExtCorpID:    extCorpID,
			ExtCreatorID: customerInfo.ExternalUserid,
		},
		ExtID:    customerInfo.ExternalUserid,
		Name:     customerInfo.Name,
		Position: customerInfo.Position,
		CorpName: customerInfo.CorpName,
		Avatar:   newAvatar,
		Type:     int(customerInfo.Type),
		Gender:   int(customerInfo.Gender),
		Unionid:  customerInfo.Unionid,
	}

	err = copier.Copy(&customer.ExternalProfile, customerInfo.ExternalProfile)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// isNoCustomerErr
// Description: 检查wx返回错误是否为 该员工没有客户
// Param: err wrap过的错误
func (o CustomerService) isNoCustomerErr(err error) bool {

	//检查wrap过的错误
	//获取根错误
	rootErr := errors.Cause(err)

	// 微信返回的错误转换为ecode错误码
	if clientError, ok := rootErr.(*workwx.ClientError); ok {
		return int(clientError.Code) == ecode.ErrCode84061.Code()
	}
	return false
}

// UpdateDeleteEventNotify
// Description: 设置流失客户是否通知员工
func (o CustomerService) UpdateDeleteEventNotify(req entities.UpdateCustomerDeleteStaffNotifierReq, extCorpID string) error {
	notify := models.EventNotify{
		ID:            id_generator.StringID(),
		ExtCorpID:     extCorpID,
		IsNotifyStaff: req.IsNotifyStaff,
	}
	return o.eventNotifyRepo.UpdateDeleteEventNotify(notify)
}

// GetNotifyStaffRule
// Description: 获取流失客户是否通知员工选项
// Param: extCorpID 外部企业ID
// return: 流失提醒开关状态
func (o CustomerService) GetNotifyStaffRule(extCorpID string) (*entities.UpdateCustomerDeleteStaffNotifierResp, error) {
	eventNotify, err := o.eventNotifyRepo.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return &entities.UpdateCustomerDeleteStaffNotifierResp{IsNotifyStaff: eventNotify.IsNotifyStaff}, err
}

// GetCustomerLosses
// Description: 客户流失记录
// Detail: 关系流水表查流失记录
func (o CustomerService) GetCustomerLosses(
	req requests.QueryCustomerLossesReq, extCorpID string, pager *app.Pager, sorter *app.Sorter) ([]models.CustomerLossInfo, int64, error) {
	log.Sugar.Debugw("req", "req", req)
	deletedStaffs, total, err := o.relationHistory.QueryCustomerDeleteStaff(req, extCorpID, sorter, pager)
	if err != nil {
		err = errors.WithStack(err)
		return nil, 0, err
	}

	return deletedStaffs, total, nil
}

// ExportDeleteStaffWarningList
// Description: 流失提醒列表下载
// Detail: 查询客户流失记录,写xlsx
// Param: 查询流失记录请求 req
func (o CustomerService) ExportDeleteStaffWarningList(req requests.QueryCustomerLossesReq, extCorpID string) (*bytes.Buffer, string, error) {

	log.Sugar.Debugw("ExportDeleteStaffWarningList", "req", req)

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeDeleteStaffWarning),
		time.Now().Format(constants.DateLayout),
		extCorpID,
		constants.DataExportDeleteStaffFilenamePrefix+".xlsx",
	)
	fullPath := filepath.Join(
		conf.Settings.Storage.LocalRootPath,
		filename,
	)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	file.NewSheet(constants.DataExportDeleteStaffListSheetName)
	file.DeleteSheet("Sheet1")

	titles := []string{"流失客户", "所属客服", "标签", "流失时间", "添加时间", "添加员工时长/天"}
	err := PrettifySheet(constants.DataExportDeleteStaffListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return nil, "", err
	}

	pager := &req.Pager
	customerLossInfos, n, err := o.relationHistory.QueryCustomerDeleteStaff(req, extCorpID, &req.Sorter, pager)
	if err != nil {
		log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
		return nil, "", err
	}
	for n > 0 && len(customerLossInfos) > 0 {
		for k, v := range customerLossInfos {
			values := []string{
				v.ExtStaffID,
				v.StaffName,
				strings.Join(v.ExtTagIDs, ","),
				v.CustomerDeleteStaffAt.Format(constants.DateTimeLayout),
				v.RelationCreateAt.Format(constants.DateTimeLayout),
				strconv.Itoa(int(v.InConnectionTimeRange)),
			}
			err = file.SetSheetRow(constants.DataExportDeleteStaffListSheetName, fmt.Sprint("A", k+3), &values)
			if err != nil {
				log.Sugar.Errorw("write excel failed", "err", err)
				return nil, "", err
			}
		}
		pager.Page += 1
		customerLossInfos, n, err = o.relationHistory.QueryCustomerDeleteStaff(req, extCorpID, &req.Sorter, pager)
		if err != nil {
			log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
			return nil, "", err
		}
	}

	buffer, err := file.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		return nil, "", err
	}

	return buffer, filename, nil
}

// Export
// Description: 导出客户列表
func (o CustomerService) Export(req requests.QueryCustomerReq, extCorpID string, pager *app.Pager) (buf *bytes.Buffer, filename string, err error) {

	log.Sugar.Infow("Export", "req", req)

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename = "/" + path.Join(
		string(constants.DataExportTypeCustomer),
		time.Now().Format(constants.DateLayout),
		extCorpID,
		constants.DataExportGroupCustomerListPrefix+".xlsx",
	)
	fullPath := filepath.Join(conf.Settings.Storage.LocalRootPath, filename)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	file.NewSheet(constants.DataExportCustomerListSheetName)
	file.DeleteSheet("Sheet1")

	titles := []string{
		"客户名称", "客户备注", "客户描述", "流失状态", "企业名称", "添加人", "添加时间",
		"企业标签", "添加渠道", "性别", "电话", "年龄", "生日",
	}
	err = PrettifySheet(constants.DataExportCustomerListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	// 查询并写入xlsx
	customers, n, err := o.customerRepo.ExportQuery(req, extCorpID, pager)
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	log.Sugar.Debug(util.JsonEncode(customers))

	//"客户名称", "客户备注", "客户描述", "流失状态", "企业名称", "添加人", "添加时间", "企业标签", "添加渠道", "性别", "电话", "年龄", "生日",

	//0-未知来源 1-扫描二维码 2-搜索手机号 3-名片分享 4-群聊 5-手机通讯录 6-微信联系人 7-来自微信的添加好友申请 8-安装第三方应用时自动添加的客服人员 9-搜索邮箱 201-内部成员共享 202-管理员/负责人分配
	mapAddWay := map[int]string{0: "未知来源", 1: "扫描二维码", 2: "搜索手机号", 3: "名片分享", 4: "群聊",
		5: "手机通讯录", 6: "微信联系人", 7: "来自微信的添加好友申请", 8: "安装第三方应用时自动添加的客服人员",
		9: "搜索邮箱", 201: "内部成员共享", 202: "管理员/负责人分配"}
	for n > 0 && len(customers) > 0 {
		for k, customer := range customers {
			values := []string{
				customer.CustomerName,
				customer.Remark,
				customer.Description,
				customer.Status,
				customer.CustomerCorpName,
				customer.StaffName,
				customer.Createtime.String(),
				"",
				strconv.FormatInt(customer.AddWay, 10),
				strconv.FormatInt(customer.Gender, 10),
				customer.PhoneNumber,
				strconv.FormatInt(customer.Age, 10),
				customer.Birthday,
			}

			var tagName string
			for _, staff := range customer.Staffs {
				for _, tag := range staff.CustomerStaffTags {
					tagName += tag.TagName + ","
				}
			}
			values[7] = strings.TrimRight(tagName, ",")

			if addWay, ok := mapAddWay[int(customer.AddWay)]; ok {
				values[8] = addWay
			}

			err = file.SetSheetRow(constants.DataExportCustomerListSheetName, fmt.Sprint("A", k+3), &values)
			if err != nil {
				log.Sugar.Errorw("write excel failed", "err", err)
				return nil, "", err
			}
		}

		pager.Page += 1
		customers, n, err = o.customerRepo.ExportQuery(req, extCorpID, pager)
		if err != nil {
			log.Sugar.Error(err)
			return
		}
	}

	buf, err = file.WriteToBuffer()
	if err != nil {
		return
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		return
	}

	return
}

// GetFullCustomerInfo
// Description: H5 客户详情页
// Detail: 查询客户所有信息
// Param:
//	extCustomerID 客户外部ID
//	extStaffID 员工外部ID
//	extCorpID 企业外部ID
// return: 客户基本信息，员工对客户的标记信息，客户与员工的时间流水
func (o CustomerService) GetFullCustomerInfo(extCustomerID, extStaffID, extCorpID string) (responses.FullCustomerInfo, error) {
	var res responses.FullCustomerInfo
	customer, err := o.customerRepo.GetByExtID(extCustomerID, []string{extStaffID}, true)
	if err != nil {
		err = errors.WithStack(err)
		return res, err
	}
	res.Customer = customer

	customerInfo, err := o.customerInfoRepo.Get(
		models.CustomerInfo{
			ExtCorpModel:  models.ExtCorpModel{ExtCorpID: extCorpID},
			ExtCustomerID: extCustomerID,
			ExtStaffID:    extStaffID},
	)
	if err != nil {
		err = errors.WithStack(err)
		return res, err
	}

	res.CustomerInfo = customerInfo

	// 当前员工给其添加的inner tag
	if len(customer.Staffs) > 0 {
		internalTags, err := o.innerTag.GetByIDs(customer.Staffs[0].InternalTagIDs.ToStringArray())
		if err != nil {
			err = errors.WithStack(err)
			return res, err
		}
		res.Customer.Staffs[0].InternalTags = internalTags
	}

	// 事件列表
	customerEvents, total, err := o.ceRepo.Query(
		models.CustomerEvent{ExtCorpModel: models.ExtCorpModel{ExtCorpID: extCorpID}}, &app.Pager{}, &app.Sorter{})
	if err != nil {
		return res, nil
	}
	res.CustomerEvents = responses.CustomerEvents{
		Events: customerEvents,
		Total:  total,
	}

	return res, nil
}

// CustomerStatistic
// Description: 查询员工客户数每日流水
// Detail:
//	1. statistic_type 为total 时,统计所有员工
//	2. 需要返回的数据样式如 -----xxxxx----xxxx---  每个字符表示一天，-表示DB没有该数据，x表示DB有数据
//	此方法实现:
//		当查询类型为全部客户数：
//		a. 将左侧-置零。
//		b. 将中间-部分与第一个x部分最后一天数据设为相同值。
//		c. 将右侧-部分与第二个x部分最后一天数据设为相同值。
//		当查询类型非全部客户数：
//		a. 将左侧-置零。
//		b. 将中间-置零。
//		c. 将右侧-置零。
// Param:
// return: 返回起止时间范围中的员工数量number
func (o CustomerService) CustomerStatistic(
	req requests.QueryCustomerStatisticReq, extCorpID string) (resp []models.CustomerTrend, err error) {

	res, err := o.csRepo.Query(req, extCorpID)

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	end, err := time.ParseInLocation(constants.DateLayout, string(req.EndTime), location)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	start, err := time.ParseInLocation(constants.DateLayout, string(req.StartTime), location)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 返回列表中元素个数==天数
	daysNum := int(end.Sub(start).Hours()/24 + 1)
	resp = make([]models.CustomerTrend, daysNum)
	if len(res) == 0 {
		// 没有查到,返回全为零的数组
		for i := 0; i < daysNum; i++ {
			d := fmt.Sprintf("%dh", i*24)
			duration, err := time.ParseDuration(d)
			if err != nil {
				return res, err
			}
			resp[i].Date = start.Add(duration).Format(constants.DateLayout)
		}
		return resp, err
	} else {
		if req.StatisticType == constants.StatisticTypeTotal {
			for i, j := 0, 0; i < daysNum; i++ {
				d := fmt.Sprintf("%dh", i*24)
				duration, err := time.ParseDuration(d)
				if err != nil {
					return res, err
				}

				curDay := start.Add(duration)
				resp[i].Date = curDay.Format(constants.DateLayout)

				if j < len(res) {
					// 查询结果中的第j个
					NumJDate, err := time.Parse(time.RFC3339, res[j].Date)
					if err != nil {
						err = errors.WithStack(err)
						return res, err
					}
					curTimeSecond := curDay.Unix()
					numJTImeSecond := NumJDate.Unix()
					if curTimeSecond == numJTImeSecond {
						resp[i].Number = res[j].Number
						j++
					} else if curTimeSecond > numJTImeSecond {
						resp[i].Number = resp[i-1].Number
					} else {
						if j == 0 {
							resp[i].Number = res[j].Number
						}
					}
				} else {
					// 右侧 补齐
					resp[i].Number = resp[i-1].Number
				}
			}
		} else {
			//	其余查询类型 DB中若没有数据，补齐零即可
			for i, j := 0, 0; i < daysNum; i++ {
				d := fmt.Sprintf("%dh", i*24)
				duration, err := time.ParseDuration(d)
				if err != nil {
					return res, err
				}

				curDay := start.Add(duration)
				resp[i].Date = curDay.Format(constants.DateLayout)

				if j < len(res) {
					// 查询结果中的第j个
					NumJDate, err := time.Parse(time.RFC3339, res[j].Date)
					if err != nil {
						err = errors.WithStack(err)
						return res, err
					}
					curTimeSecond := curDay.Unix()
					numJTImeSecond := NumJDate.Unix()
					if curTimeSecond == numJTImeSecond {
						resp[i].Number = res[j].Number
						j++
					}
				}
			}
		}

	}
	return resp, err
}

// UpdateCustomerTags
// Description: 同步数据时更新员工给客户打的标签
// Detail:
// Param: relationID CustomerStaff 关系表的ID
// Param: extTags wx返回的外部标签信息
func (o CustomerService) UpdateCustomerTags(relationID string, extTags []workwx.FollowUserTag) (err error) {
	tags := make([]models.CustomerStaffTag, 0)
	extCorpID := conf.Settings.WeWork.ExtCorpID
	for _, tag := range extTags {
		customerStaffTag := models.CustomerStaffTag{
			ExtCorpModel:    models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
			CustomerStaffID: relationID,
			ExtTagID:        tag.TagID,
			GroupName:       tag.GroupName,
			TagName:         tag.TagName,
			Type:            constants.FollowUserTagType(tag.Type),
		}
		customerStaffTag.DeletedAt = gorm.DeletedAt{} // 恢复被删除的
		tags = append(tags, customerStaffTag)
	}
	err = o.customerStaffTagRepo.Upsert(tags)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 删除不存在的tag
	// 若客户的标签被全部删除,则 extTags 为空,需要将客户-员工标签表的数据全部删除
	extTagIDs := make([]string, 0)
	for _, tag := range extTags {
		extTagIDs = append(extTagIDs, tag.TagID)
	}
	err = o.customerStaffTagRepo.Delete(relationID, extTagIDs, false)
	if err != nil {
		return err
	}

	return
}

// UpsertCustomerPortrait
// Description: <新建客户数据时>新增一条客户画像
// Detail: 员工对客户的备注，侧边栏可以更改
// Param: extCustomerID 外部客户ID
// Param: extStaffID 外部员工ID
func (o CustomerService) UpsertCustomerPortrait(extCustomerID string, extStaffID string) (err error) {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	customerInfo := models.CustomerInfo{
		ExtCorpModel:  models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID},
		ExtCustomerID: extCustomerID,
		ExtStaffID:    extStaffID,
	}
	err = o.customerInfoRepo.Upsert(customerInfo)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// BatchFetchCustomers
// Description: 并发获取员工的外部客户ID
// Detail: 不用将员工与客户id数组对应
func (o CustomerService) BatchFetchCustomers(extStaff []*workwx.UserInfo) (totalExtCustomerIDs []string, err error) {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(extStaff))

	responseChannel := make(chan []string, 200)
	doneChannel := make(chan struct{})

	//处理结果协程
	go func() {
		wg.Add(1)
		for response := range responseChannel {
			totalExtCustomerIDs = append(totalExtCustomerIDs, response...)
		}
		// 当 responseChannel被关闭时且channel中所有的值都已经被处理完毕后
		wg.Done()
	}()

	var g errgroup.Group
	for i, _ := range extStaff {
		extStaffID := extStaff[i].UserID
		g.Go(func() error {
			defer wg.Done()
			extCustomerIDs, err := client.Customer.ListExternalContact(extStaffID)
			if err != nil {
				//检查wrap过的错误
				rootErr := errors.Cause(err)
				// 微信返回的错误转换为ecode错误码
				if clientError, ok := rootErr.(*workwx.ClientError); ok {
					rootErr = ecode.Int(int(clientError.Code))
					log.Sugar.Infof("biz error：%v", clientError)
				} else {
					log.Sugar.Errorf("ListExternalContact failed:%s, %#v", extStaffID, err)
				}
				log.Sugar.Errorf("ListExternalContact failed:%s, %#v", extStaffID, err)
				if o.isNoCustomerErr(err) {
					return nil
				}
				return err
			}
			responseChannel <- extCustomerIDs
			return nil
		})
	}

	// wg.Wait()此时也要go出去,防止在wg.Wait()出堵住
	go func() {
		wg.Wait()
		close(doneChannel)
	}()

	select {
	// 正常结束完成
	case <-doneChannel:
	// 超时
	case <-time.After(2 * time.Second):
	}

	return
}

// BatchFetchContactInfo
// Description: 批量请求员工-客户关系
func (o CustomerService) BatchFetchContactInfo(extCustomerIDs []string) (cs []*workwx.ExternalContactInfo, err error) {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	client, err := GetCorpWxClient(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(extCustomerIDs))

	responseChannel := make(chan *workwx.ExternalContactInfo, 200)
	doneChannel := make(chan struct{})

	//处理结果协程
	go func() {
		wg.Add(1)
		for response := range responseChannel {
			cs = append(cs, response)
		}
		// 当 responseChannel被关闭时且channel中所有的值都已经被处理完毕后
		wg.Done()
	}()

	var g errgroup.Group
	for i, _ := range extCustomerIDs {
		extCustomerID := extCustomerIDs[i]
		g.Go(func() error {
			defer wg.Done()
			contactInfo, err := client.Customer.GetExternalContact(extCustomerID)
			if err != nil {
				log.Sugar.Errorf("GetExternalContact failed:%s, %#v", extCustomerID, err)
				if o.isNoCustomerErr(err) {
					return nil
				}
				return err
			}
			responseChannel <- contactInfo
			return nil
		})
	}

	// wg.Wait()此时也要go出去,防止在wg.Wait()出堵住
	go func() {
		wg.Wait()
		close(doneChannel)
	}()

	select {
	// 正常结束完成
	case <-doneChannel:
	// 超时
	case <-time.After(2 * time.Second):
	}

	return
}
