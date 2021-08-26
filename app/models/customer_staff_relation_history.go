package models

import (
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/id_generator"
	"openscrm/conf"
	"time"
)

// CustomerStaffRelationHistory
// 员工客户关系的历史数据（流水记录）。
// 员工删除客户/客户删除员工时 新增一条数据，写入 customer_delete_staff_at/staff_delete_customer_at, 同时软删除原有记录。
type CustomerStaffRelationHistory struct {
	ExtCorpModel
	// 企微员工ID
	ExtStaffID string `gorm:"type:char(32);index;comment:员工ID" json:"ext_staff_id"`
	// 企微客户ID
	ExtCustomerID string `gorm:"type:char(32);index;comment:客户ID" json:"ext_customer_id"`
	// 员工添加客户的时间,与wx返回的一致，以便使用copier
	Createtime time.Time `gorm:"comment:员工添加客户的时间" json:"createtime"`
	// 客户删除员工的时间
	CustomerDeleteStaffAt null.Time `gorm:"comment:客户删除员工的时间;" json:"customer_delete_staff_at"`
	// 员工删除客户的时间
	StaffDeleteCustomerAt null.Time `gorm:"comment:员工删除客户的时间;" json:"staff_delete_customer_at"`
	Timestamp
}

// CustomerDeleteStaff
// StaffDeleteCustomer
// Description: 员工删除客户
// Detail: 删除关系记录,新增关系流水
func (o CustomerStaffRelationHistory) CustomerDeleteStaff(extStaffID string, extCustomerID string) (err error) {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 删除客户员工关系
		var cs CustomerStaff
		err = tx.Where("ext_staff_id = ? and ext_customer_id = ?", extStaffID, extCustomerID).First(&cs).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		err = tx.Where("ext_staff_id = ? and ext_customer_id = ?", extStaffID, extCustomerID).
			Delete(&CustomerStaff{}).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}

		// 新增客户删除员工的记录
		extCorpID := conf.Settings.WeWork.ExtCorpID
		customerStaffRelationHistory := CustomerStaffRelationHistory{
			ExtCorpModel:          ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: extCustomerID},
			ExtStaffID:            extStaffID,
			ExtCustomerID:         extCustomerID,
			Createtime:            cs.Createtime,
			CustomerDeleteStaffAt: null.TimeFrom(time.Now()),
		}
		err = tx.Model(&CustomerStaffRelationHistory{}).Create(&customerStaffRelationHistory).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		return nil
	})
}

// StaffDeleteCustomer
// Description: 员工删除客户
// Detail: 删除关系记录,新增关系流水
func (o CustomerStaffRelationHistory) StaffDeleteCustomer(extStaffID string, extCustomerID string) (err error) {
	return DB.Transaction(func(tx *gorm.DB) error {
		var cs CustomerStaff
		err = tx.Where("ext_staff_id = ? and ext_customer_id = ?", extStaffID, extCustomerID).First(&cs).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		// 删除客户员工关系
		err = tx.Where("ext_staff_id = ? and ext_customer_id = ?", extStaffID, extCustomerID).
			Delete(&CustomerStaff{}).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}

		// 新增客户删除员工的记录
		extCorpID := conf.Settings.WeWork.ExtCorpID
		customerStaffRelationHistory := CustomerStaffRelationHistory{
			ExtCorpModel:          ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: extCustomerID},
			ExtStaffID:            extStaffID,
			ExtCustomerID:         extCustomerID,
			Createtime:            cs.CreatedAt,
			StaffDeleteCustomerAt: null.TimeFrom(time.Now()),
		}
		err = tx.Model(&CustomerStaffRelationHistory{}).Create(&customerStaffRelationHistory).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		return nil
	})
}

func (o CustomerStaffRelationHistory) QueryStaffDeleteCustomer(
	req requests.QueryStaffDeleteCustomerHistoryReq, extCorpID string, pager *app.Pager, sorter *app.Sorter) ([]StaffDeleteCustomer, int64, error) {
	db := DB.Table("customer_staff_relation_history").
		Joins("join customer_staff on customer_staff.ext_staff_id = customer_staff_relation_history.ext_staff_id and customer_staff.ext_customer_id = customer_staff_relation_history.ext_customer_id").
		Joins("join customer  on customer.ext_id = customer_staff.ext_customer_id").
		Joins("join staff on customer_staff.ext_staff_id = staff.ext_id").
		Select("customer_staff.id as id," +
			" customer.ext_id as ext_customer_id, " +
			" customer.avatar as ext_customer_avatar, " +
			" customer.name as ext_customer_name, " +
			" customer.type as customer_type, " +
			" customer.corp_name as customer_corp_name, " +
			" customer_staff_relation_history.createtime as  relation_create_at, " +
			" customer_staff_relation_history.staff_delete_customer_at as relation_delete_at, " +
			" staff.name as staff_name, " +
			" staff.ext_staff_id as ext_staff_id, " +
			" staff.id as staff_id, " +
			" staff.avatar_url as ext_staff_avatar ")

	if req.ExtDepartmentID != 0 {
		db = db.Where("  json_contains(staff.dept_ids, json_array(?) )", req.ExtDepartmentID)
	}

	if len(req.ExtStaffIDs) > 0 {
		db = db.Where(" staff.ext_id in (?) ", req.ExtStaffIDs)
	}

	if req.ConnectionCreateStart != "" && req.ConnectionCreateEnd != "" {
		db = db.Where("customer_staff.createtime between ? and ?", req.ConnectionCreateStart, req.ConnectionCreateEnd)
	}

	if req.DeleteCustomerStart != "" && req.DeleteCustomerEnd != "" {
		db = db.Where("customer_staff_relation_history.staff_delete_customer_at between ? and ?", req.DeleteCustomerStart, req.DeleteCustomerEnd)
	} else {
		db = db.Where("customer_staff_relation_history.staff_delete_customer_at is not null")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count StaffDeleteCustomer failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: string(sorter.SortField)},
		Desc:   sorter.SortType == constants.SortTypeDesc},
	)

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	var items []StaffDeleteCustomer
	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find StaffDeleteCustomer failed")
		return nil, 0, err
	}
	return items, total, err
}

func (o CustomerStaffRelationHistory) QueryCustomerDeleteStaff(
	req requests.QueryCustomerLossesReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) ([]CustomerLossInfo, int64, error) {

	db := DB.Table("customer_staff_relation_history").
		Joins("join customer_staff on customer_staff.ext_staff_id = customer_staff_relation_history.ext_staff_id and customer_staff.ext_customer_id = customer_staff_relation_history.ext_customer_id ").
		Joins("join customer on customer.ext_id = customer_staff.ext_customer_id").
		Joins("join staff s on customer_staff.ext_staff_id = s.ext_id").
		Select("customer_staff.id as id, " +
			" customer.ext_id as ext_customer_id, " +
			" customer.avatar as customer_avatar, " +
			" customer.name as ext_customer_name, " +
			" customer.type as customer_type, " +
			" customer.corp_name as customer_corp_name, " +
			" customer_staff_relation_history.createtime as  relation_create_at, " +
			" customer_staff_relation_history.customer_delete_staff_at as customer_delete_staff_at, " +
			" s.name as staff_name, " +
			" s.ext_id as ext_staff_id, " +
			" s.id as staff_id, " +
			" s.avatar_url as staff_avatar, " +
			" timestampdiff(day, customer_staff_relation_history.createtime ,customer_staff_relation_history.customer_delete_staff_at) as in_connection_time_range, " +
			" customer_staff.ext_tag_ids as ext_tag_ids ").
		Where("customer_staff_relation_history.customer_delete_staff_at is not null")

	if extCorpID != "" {
		db = db.Where("s.ext_corp_id = ? ", extCorpID)
	}

	if len(req.ExtStaffIDs) > 0 {
		db = db.Where("s.ext_id in(?) ", req.ExtStaffIDs)
	}

	if req.LossStart != "" {
		db = db.Where("customer_staff_relation_history.customer_delete_staff_at between ? and ?", req.LossStart, req.LossEnd)
	}

	if req.ConnectionCreateStart != "" {
		db = db.Where("customer_staff.createtime between ? and ?", req.ConnectionCreateStart, req.ConnectionCreateEnd)
	}

	if req.TimeSpanLowerLimit > 0 {
		db = db.Where(" timestampdiff(day,  customer_staff_relation_history.createtime, customer_staff_relation_history.customer_delete_staff_at)  > ?", req.TimeSpanLowerLimit)
	}

	if req.TimeSpanUpperLimit > 0 {
		db = db.Where(" timestampdiff(day,  customer_staff_relation_history.createtime, customer_staff_relation_history.customer_delete_staff_at) < ?", req.TimeSpanUpperLimit)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count customer_delete_staff failed")
		return nil, 0, err
	}

	sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: string(sorter.SortField)},
		Desc:   sorter.SortType == constants.SortTypeDesc},
	)

	pager.SetDefault()
	db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

	items := make([]CustomerLossInfo, 0)
	err = db.Preload("Tags").Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find QueryCustomerDeleteStaff failed")
		return nil, 0, err
	}
	return items, total, err
}
