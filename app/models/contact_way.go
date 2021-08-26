package models

import (
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/delay_queue"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
	"time"
)

// ContactWay 渠道码
type ContactWay struct {
	ExtCorpModel
	Name string `gorm:"type:varchar(255);index;comment:'渠道码名称'" json:"name"`
	// ConfigID 渠道码配置ID
	ConfigID string `json:"config_id" gorm:"index;comment:渠道码配置ID" validate:"alpha"`
	// GroupID 渠道码分组ID
	GroupID string `json:"group_id" gorm:"index;type:bigint;comment:活码分组ID"`
	// QrCode 联系二维码的URL，仅在scene为2时返回
	QrCode string `json:"qr_code" gorm:"comment:联系二维码的URL"`
	// Remark 渠道码的备注信息，用于助记
	Remark string `json:"remark" gorm:"comment:渠道码的备注信息"`
	// SkipVerify 外部客户添加时是否无需验证，假布尔类型
	SkipVerify constants.Boolean `json:"skip_verify" gorm:"default:1;comment:外部客户添加时是否无需验证，假布尔类型" validate:"oneof=1 2"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“<a href="#13878">获取外部联系人详情</a>”时会返回该参数值
	State string `json:"state" gorm:"comment:企业自定义的state参数"`
	// AddCustomerCount 扫码添加人次
	AddCustomerCount int `json:"add_customer_count" gorm:"->;index;default:0;comment:扫码添加人次"`
	// AutoReplyType 欢迎语类型：1，渠道欢迎语；2, 渠道默认欢迎语；3，不送欢迎语；
	AutoReplyType constants.ContactWayAutoReplyType `json:"auto_reply_type" gorm:"index;default:1;comment:'欢迎语类型：1，渠道欢迎语；2, 渠道默认欢迎语；3，不送欢迎语；'"`
	// AutoReply 欢迎语策略
	AutoReply constants.AutoReplyField `json:"auto_reply" gorm:"comment:欢迎语策略"`
	// CustomerDesc 客户描述
	CustomerDesc string `json:"customer_desc" gorm:"comment:客户描述"`
	// CustomerDescEnable 是否开启客户描述
	CustomerDescEnable constants.Boolean `json:"customer_desc_enable" gorm:"comment:是否开启客户描述" validate:"oneof=1 2"`
	// CustomerRemark 客户备注
	CustomerRemark string `json:"customer_remark" gorm:"comment:客户备注"`
	// CustomerRemarkEnable 是否开启客户备注
	CustomerRemarkEnable constants.Boolean `json:"customer_remark_enable" gorm:"comment:是否开启客户备注" validate:"oneof=1 2"`
	// DailyAddCustomerLimitEnable 是否开启员工每日添加上限
	DailyAddCustomerLimitEnable constants.Boolean `json:"daily_add_customer_limit_enable" gorm:"comment:是否开启员工每日添加上限" validate:"oneof=1 2"`
	// 员工每日添加上限
	DailyAddCustomerLimit int64 `json:"daily_add_customer_limit" gorm:"comment:员工每日添加上限" validate:"gte=0"`
	// ScheduleEnable 是否开启工作日调度
	ScheduleEnable constants.Boolean `json:"schedule_enable" gorm:"comment:是否开启工作日调度" validate:"oneof=1 2"`
	// StaffControlEnable 是否开启员工自行上下线
	StaffControlEnable constants.Boolean `json:"staff_control_enable" gorm:"comment:是否开启员工自行上下线" validate:"oneof=1 2"`
	// Staffs 绑定员工
	Staffs []ContactWayStaff `json:"staffs" gorm:"foreignKey:ContactWayID;"`
	// BackupStaffs 绑定备份员工
	BackupStaffs []ContactWayBackupStaff `json:"backup_staffs" gorm:"foreignKey:ContactWayID;"`
	// Schedules
	Schedules []ContactWaySchedule `json:"schedules" gorm:"foreignKey:ContactWayID;"`
	// AutoTagEnable 是否自动打标签
	AutoTagEnable constants.Boolean `json:"auto_tag_enable" gorm:"comment:'是否自动打标签'" validate:"oneof=1 2"`
	// CustomerTagExtIDs 自动打标签绑定的标签ExtID数组
	CustomerTagExtIDs constants.StringArrayField `json:"customer_tag_ext_ids" gorm:"type:json;comment:'自动打标签绑定的标签ExtID数组'" validate:"omitempty,dive,ext_id"`
	// AutoSkipVerifyEnable 是否开启自动通过好友时段控制
	AutoSkipVerifyEnable constants.Boolean `json:"auto_skip_verify_enable" gorm:"default:1;comment:是否开启自动通过好友时段控制" validate:"oneof=1 2"`
	// SkipVerifyStartTime 自动通过好友开启时刻
	SkipVerifyStartTime constants.TimeField `json:"skip_verify_start_time" gorm:"comment:自动通过好友开启时刻" validate:"required_if=AutoSkipVerifyEnable 1,time"`
	// SkipVerifyEndTime 自动通过好友结束时刻
	SkipVerifyEndTime constants.TimeField `json:"skip_verify_end_time" gorm:"comment:自动通过好友结束时刻" validate:"required_if=AutoSkipVerifyEnable 1,time"`
	// ExtStaffIDs 实时关联的外部员工ID
	ExtStaffIDs constants.StringArrayField `json:"ext_staff_ids" gorm:"type:json;comment:'实时关联的外部员工ID'" validate:"omitempty,dive,int64"`
	// NicknameBlockEnable 是否开启客户昵称屏蔽欢迎语
	NicknameBlockEnable constants.Boolean `json:"nickname_block_enable" gorm:"comment:是否开启客户昵称屏蔽欢迎语" validate:"oneof=1 2"`
	// NicknameBlockList 客户昵称屏蔽欢迎语列表
	NicknameBlockList constants.StringArrayField `json:"nickname_block_list" gorm:"type:json;comment:'客户昵称屏蔽欢迎语列表'" validate:"omitempty,dive,int64"`
	Timestamp
}

func (o ContactWay) Query(param requests.QueryContactWayReq, extCorpID string) (items []ContactWay, total int64, err error) {
	items = make([]ContactWay, 0)
	db := DB.Model(&ContactWay{}).Preload(clause.Associations).Preload("Schedules.Staffs").Where("ext_corp_id = ?", extCorpID)

	if len(param.ExtStaffIDs) > 0 {
		ids := make([]string, 0)
		tmpIDs := make([]string, 0)
		DB.Model(&ContactWayStaff{}).Where("contact_way_id is not null ").Where("ext_staff_id in (?)", param.ExtStaffIDs).Pluck("contact_way_id", &tmpIDs)
		ids = append(ids, tmpIDs...)
		DB.Model(&ContactWayBackupStaff{}).Where("contact_way_id is not null ").Where("ext_staff_id in (?)", param.ExtStaffIDs).Pluck("contact_way_id", &tmpIDs)
		ids = append(ids, tmpIDs...)
		DB.Model(&ContactWayScheduleStaff{}).Where("contact_way_id is not null ").Where("contact_way_schedule_id is not null ").Where("ext_staff_id in (?)", param.ExtStaffIDs).Pluck("contact_way_id", &tmpIDs)
		ids = append(ids, tmpIDs...)
		db = db.Where("id in (?)", ids)
		param.ExtStaffIDs = nil
	}

	if param.Name != "" {
		db = db.Where("name like ?", param.Name+"%")
		param.Name = ""
	}

	if !param.CreatedAtStart.MustTime().IsZero() {
		db = db.Where("created_at > ?", param.CreatedAtStart.MustTime())
	}

	if !param.CreatedAtEnd.MustTime().IsZero() {
		db = db.Where("created_at < ?", param.CreatedAtEnd.MustTime())
	}

	queryParam := &ContactWay{}
	copier.Copy(queryParam, param)

	db = db.Where(queryParam)
	err = db.Count(&total).Error
	if err != nil {
		err = errors.Wrap(err, "Count ContactWay failed")
		return
	}

	param.Sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(param.Sorter.SortField)}, Desc: param.Sorter.SortType == constants.SortTypeDesc})

	param.Pager.SetDefault()
	db = db.Offset(param.Pager.GetOffset()).Limit(param.Pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find ContactWay failed")
		return
	}

	return
}

func (o ContactWay) Get(id string, extCorpID string) (item ContactWay, err error) {
	err = DB.Model(&ContactWay{}).Preload(clause.Associations).Preload("Schedules.Staffs").Where("ext_corp_id = ?", extCorpID).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First ContactWay failed")
		return
	}

	return
}

// makeJobID 生成任务ID
func (o ContactWay) makeJobID(id string, executeAt time.Time, state string) string {
	return constants.ContactWayJobPrefix.String() + id + gmd5.MustEncryptString(fmt.Sprintf("%s-%d-%s", id, executeAt.Unix(), state))
}

// AddRefreshJob 添加刷新任务，指定时间计算最新的渠道码状态并更新
func (o ContactWay) AddRefreshJob(id string, executeAt time.Time, state string) (err error) {
	defer util.FuncTracer("id", id, "executeAt", executeAt, "state", state)()
	// 任务执行时间是过去的时间，自动跳过
	if executeAt.Before(util.Now()) {
		log.Sugar.Debugw("任务执行时间是过去的时间，自动跳过", "executeAt", executeAt, "now", util.Now())
		return
	}
	err = delay_queue.Add(delay_queue.Job{
		Topic:     constants.RefreshContactWayTopic,
		ID:        o.makeJobID(id, executeAt, state),
		ExecuteAt: executeAt.Unix(),
		TTR:       5,
		Body:      id,
	})
	if err != nil {
		err = errors.Wrap(err, "add job failed")
		return
	}
	return
}

// CalLatestStatus 计算最新的渠道码状态，并自动添加触发任务
func (o ContactWay) CalLatestStatus(item *ContactWay) (err error) {
	//自动通过好友时段控制
	if item.AutoSkipVerifyEnable == constants.True {
		if time.Now().Unix() > util.Today().Unix()+item.SkipVerifyStartTime.Seconds() &&
			time.Now().Unix() < util.Today().Unix()+item.SkipVerifyEndTime.Seconds() {
			item.SkipVerify = constants.True
		} else {
			item.SkipVerify = constants.False
		}

		// 开始时段触发
		err = o.AddRefreshJob(item.ID, util.Today().Add(item.SkipVerifyStartTime.Duration()), "SkipVerifyStartTime")
		if err != nil {
			err = errors.Wrap(err, "AddRefreshJob failed")
			return
		}

		// 结束时段触发
		err = o.AddRefreshJob(item.ID, util.Today().Add(item.SkipVerifyEndTime.Duration()), "SkipVerifyEndTime")
		if err != nil {
			err = errors.Wrap(err, "AddRefreshJob failed")
			return
		}
	} else {
		item.SkipVerify = constants.True
	}

	item.ExtStaffIDs = make([]string, 0)

	// 计算调度设置绑定的员工
	if item.ScheduleEnable == constants.Enable {
		// 每天凌晨触发工作时段控制
		err = o.AddRefreshJob(item.ID, util.Today().Add(time.Hour*24+time.Second), "ScheduleEnable")
		if err != nil {
			err = errors.Wrap(err, "AddRefreshJob failed")
			return
		}

		for _, schedule := range item.Schedules {
			// 工作周期开始工作时触发
			err = o.AddRefreshJob(item.ID, util.Today().Add(schedule.StartTime.Duration()), "staff.StartTime")
			if err != nil {
				err = errors.Wrap(err, "AddRefreshJob failed")
				return
			}

			//工作周期结束工作时触发
			err = o.AddRefreshJob(item.ID, util.Today().Add(schedule.EndTime.Duration()), "staff.EndTime")
			if err != nil {
				err = errors.Wrap(err, "AddRefreshJob failed")
				return
			}

			// 工作时间段以外，休息
			if time.Now().Unix() < util.Today().Unix()+schedule.StartTime.Seconds() ||
				time.Now().Unix() > util.Today().Unix()+schedule.EndTime.Seconds() {
				continue
			}

			// 绑定调度计划设置的员工
			for _, staff := range schedule.Staffs {
				// 每日添加人数限制
				if item.DailyAddCustomerLimitEnable.Bool() {
					if staff.DailyAddCustomerLimit > 0 && staff.DailyAddCustomerLimit <= staff.DailyAddCustomerCount {
						continue
					}
				}

				// 员工自行上下线
				if item.StaffControlEnable.Bool() {
					if staff.Online == constants.False {
						continue
					}
				}

				item.ExtStaffIDs = append(item.ExtStaffIDs, staff.ExtStaffID)
			}
		}
	}

	// 计算普通绑定员工
	if item.ScheduleEnable == constants.Disable {
		for _, staff := range item.Staffs {
			// 每日添加人数限制
			if item.DailyAddCustomerLimitEnable.Bool() {
				if staff.DailyAddCustomerLimit > 0 && staff.DailyAddCustomerLimit <= staff.DailyAddCustomerCount {
					continue
				}
			}

			// 员工自行上下线
			if item.StaffControlEnable.Bool() {
				if staff.Online == constants.False {
					continue
				}
			}

			item.ExtStaffIDs = append(item.ExtStaffIDs, staff.ExtStaffID)
		}
	}

	// 当没有有效关联员工时，使用备份员工
	if len(item.ExtStaffIDs) == 0 {
		for _, staff := range item.BackupStaffs {
			item.ExtStaffIDs = append(item.ExtStaffIDs, staff.ExtStaffID)
		}
	}

	item.ExtStaffIDs = funk.UniqString(item.ExtStaffIDs)

	return
}

func (o ContactWay) Create(param ContactWay, extCorpID string) (item ContactWay, err error) {
	err = copier.Copy(&item, param)
	if err != nil {
		err = errors.Wrap(err, "copier.Copy failed")
		return
	}
	item.ExtCorpID = extCorpID
	if item.ID == "" {
		item.ID = id_generator.StringID()
	}
	item.State = constants.ContactWayStatePrefix + item.ID

	err = o.CalLatestStatus(&item)
	if err != nil {
		err = errors.Wrap(err, "CalLatestStatus failed")
		return
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return
	}

	item.ConfigID, err = client.Customer.AddContactWay(workwx.AddContactWay{
		IsTemp:     false,
		Remark:     item.Remark,
		Scene:      2,
		SkipVerify: item.SkipVerify == constants.True,
		State:      item.State,
		Type:       workwx.ContactWayTypeMultiple,
		User:       item.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx AddContactWay failed")
		return
	}

	wxContactWay, err := client.Customer.GetContactWay(item.ConfigID)
	if err != nil {
		err = errors.Wrap(err, "wx GetContactWay failed")
		return
	}

	item.QrCode = wxContactWay.QrCode

	err = DB.Create(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Create ContactWay failed")
		return
	}

	return
}

func (o ContactWay) Update(id string, param ContactWay, extCorpID string) (item ContactWay, err error) {
	tx := DB.Begin()
	defer tx.Rollback()
	//origin := &ContactWay{}
	err = tx.Preload(clause.Associations).Preload("Schedules.Staffs").Where("id = ?", id).Where("ext_corp_id = ?", extCorpID).First(&item).Error
	if err != nil {
		err = errors.Wrap(err, "get ContactWay failed")
		return
	}

	// 对于关联数据的处理行为是，带主键ID的进行修改，没有带主键ID进行添加，数据库里存在但请求没有带的记录进行删除
	//err = DeleteRefRecord(tx, &ContactWayStaff{}, item.Staffs, param.Staffs, "ID")
	//if err != nil {
	//	err = errors.Wrap(err, "delete ContactWayStaff failed")
	//	return
	//}
	//
	//err = DeleteRefRecord(tx, &ContactWayBackupStaff{}, item.BackupStaffs, param.BackupStaffs, "ID")
	//if err != nil {
	//	err = errors.Wrap(err, "delete ContactWayBackupStaff failed")
	//	return
	//}
	//
	//err = DeleteRefRecord(tx, &ContactWaySchedule{}, item.Schedules, param.Schedules, "ID")
	//if err != nil {
	//	err = errors.Wrap(err, "delete ContactWaySchedule failed")
	//	return
	//}

	err = copier.CopyWithOption(&item, param, copier.Option{IgnoreEmpty: true})
	if err != nil {
		err = errors.Wrap(err, "copy param failed")
		return
	}

	item.Staffs = param.Staffs
	item.BackupStaffs = param.BackupStaffs

	err = o.CalLatestStatus(&item)
	if err != nil {
		err = errors.Wrap(err, "CalLatestStatus failed")
		return
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return
	}

	_, err = client.Customer.UpdateContactWay(workwx.UpdateContactWay{
		ConfigID:   item.ConfigID,
		Remark:     item.Remark,
		SkipVerify: item.SkipVerify == constants.True,
		State:      item.State,
		User:       item.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx UpdateContactWay failed")
		return
	}

	//err = tx.Save(&item).Error
	//if err != nil {
	//	err = errors.Wrap(err, "save ContactWay failed")
	//	return
	//}

	err = tx.Omit(clause.Associations).Updates(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Update ContactWay failed")
		return
	}

	err = tx.Model(&item).Association("Staffs").Replace(&item.Staffs)
	if err != nil {
		err = errors.Wrap(err, "Update Staffs failed")
		return
	}

	err = tx.Model(&item).Association("BackupStaffs").Replace(&item.BackupStaffs)
	if err != nil {
		err = errors.Wrap(err, "Update BackupStaffs failed")
		return
	}

	for i := range item.Schedules {
		err = tx.Model(&item.Schedules[i]).Association("Staffs").Replace(&item.Schedules[i].Staffs)
		if err != nil {
			err = errors.Wrap(err, "Update Schedules.Staffs failed")
			return
		}
	}

	err = tx.Model(&item).Association("Schedules").Replace(&item.Schedules)
	if err != nil {
		err = errors.Wrap(err, "Update Schedules failed")
		return
	}

	err = tx.Commit().Error
	if err != nil {
		err = errors.Wrap(err, "tx.Commit() failed")
		return
	}

	return
}

// Refresh 计算渠道码最新状态并保存
func (o ContactWay) Refresh(tx *gorm.DB, id string) (item ContactWay, err error) {
	err = tx.Where("id = ?", id).First(&item).Error
	// 如果找不到数据，不返回错误，直接跳过任务，重试没意义
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First ContactWay failed")
		return
	}

	err = o.CalLatestStatus(&item)
	if err != nil {
		err = errors.Wrap(err, "CalLatestStatus failed")
		return
	}

	client, err := we_work.Clients.Get(item.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "get Client failed")
		return
	}

	_, err = client.Customer.UpdateContactWay(workwx.UpdateContactWay{
		ConfigID:   item.ConfigID,
		Remark:     item.Remark,
		SkipVerify: item.SkipVerify == constants.True,
		State:      item.State,
		User:       item.ExtStaffIDs,
	})
	if err != nil {
		err = errors.Wrap(err, "wx UpdateContactWay failed")
		return
	}

	err = tx.Omit(clause.Associations).Save(&item).Error
	if err != nil {
		err = errors.Wrap(err, "Save ContactWay failed")
		return
	}

	return
}

func (o ContactWay) Delete(ids []string, extCorpID string) (total int64, err error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&ContactWay{})
	err = result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete ContactWay failed")
		return
	}
	total = result.RowsAffected

	return
}

// BatchUpdate 批量修改
func (o ContactWay) BatchUpdate(ids []string, extCorpID string, param ContactWay) (total int64, err error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Updates(param)
	err = result.Error
	if err != nil {
		err = errors.Wrap(err, "BatchUpdate ContactWay failed")
		return
	}
	total = result.RowsAffected

	return
}

// StaffOnline 员工自主控制上下线
func (o ContactWay) StaffOnline(id string, extCorpID string, extStaffID string, online constants.Boolean) (err error) {
	item := ContactWay{}
	err = DB.Where("ext_corp_id = ?", extCorpID).Where("id = ?", id).First(&item).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First ContactWay failed")
		return
	}

	if item.StaffControlEnable == constants.False {
		err = errors.WithStack(ecode.ForbiddenError)
		return
	}

	contactWayStaff := ContactWayStaff{}
	err = DB.Where("contact_way_id = ?", item.ID).Where("ext_staff_id = ?", extStaffID).First(&contactWayStaff).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.WithStack(ecode.ItemNotFoundError)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "First contactWayStaff failed")
		return
	}

	err = DB.Model(&contactWayStaff).Update("online", online).Error
	if err != nil {
		err = errors.Wrap(err, "update ContactWay failed")
		return
	}

	return

}
