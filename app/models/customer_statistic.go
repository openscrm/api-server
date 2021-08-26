package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	log2 "openscrm/common/log"
	"openscrm/conf"
	"time"
)

// CustomerStatistic 按天统计客户数量
// unique_index: ext_staff_id - date
type CustomerStatistic struct {
	ExtCorpModel
	ExtStaffID string `json:"ext_staff_id" gorm:"type:char(48);unique_index:ext_staff_id_date;comment:外部员工ID"`
	// 总客户数
	TotalCustomerNum int64 `json:"total_customer_num" gorm:"type:bigint unsigned;comment:客户总数"`
	// 新增客户数
	IncreaseCustomerNum int64 `json:"increase_customer_num"  gorm:"type:bigint unsigned;comment:新增客户总数"`
	// 流失客户数
	DecreaseCustomerNum int64 `json:"decrease_customer_num"  gorm:"type:bigint unsigned;comment:流失客户总数"`
	// 净增加客户数 = IncreaseCustomerNum -  DecreaseCustomerNum
	// 日期
	Date constants.DateField `json:"date" gorm:"unique_index:ext_staff_id_date;comment:日期"`
	Timestamp
}

// CustomerTrend   客户统计结果
type CustomerTrend struct {
	Number int64  `json:"number"`
	Date   string `json:"date"`
}

func (s CustomerStatistic) Query(req requests.QueryCustomerStatisticReq, extCorpID string) (res []CustomerTrend, err error) {

	db := DB.Model(&CustomerStatistic{}).
		Where("ext_corp_id= ?", extCorpID).Where("date between ? and ?", req.StartTime, req.EndTime)

	if len(req.ExtStaffIDs) == 0 {
		switch req.StatisticType {
		case constants.StatisticTypeTotal:
			err = db.Select("sum(total_customer_num) as number, date").Group("ext_staff_id,date").Order("date").Find(&res).Error
		case constants.StatisticTypeIncrease:
			err = db.Select("sum(increase_customer_num) as number, date").Group("ext_staff_id,date").Order("date").Find(&res).Error
		case constants.StatisticTypeDecrease:
			err = db.Select("sum(decrease_customer_num) as number, date").Group("ext_staff_id,date").Order("date").Find(&res).Error
		case constants.StatisticTypeNetIncrease:
			err = db.Select("(sum(increase_customer_num) - sum(decrease_customer_num)) as number, date").Group("ext_staff_id,date").Order("date").Find(&res).Error
		default:
			return
		}
	} else {
		db = db.Where("ext_staff_id in (?)", req.ExtStaffIDs)
		switch req.StatisticType {
		case constants.StatisticTypeTotal:
			err = db.Select("(total_customer_num) as number, date").Order("date").Find(&res).Error
		case constants.StatisticTypeIncrease:
			err = db.Select("(increase_customer_num) as number, date").Order("date").Find(&res).Error
		case constants.StatisticTypeDecrease:
			err = db.Select("(decrease_customer_num) as number, date").Order("date").Find(&res).Error
		case constants.StatisticTypeNetIncrease:
			err = db.Select("increase_customer_num - decrease_customer_num as number, date").Order("date").Find(&res).Error
		default:
			return
		}
	}
	return
}

// Upsert 更新员工客户数
// num 为负数表示流失客户，为正数表示新增客户
func (s CustomerStatistic) Upsert(extStaffID string, num int64) (err error) {
	if num == 0 {
		return ecode.CustomerNumErr
	}
	// 方法一
	//用统计表中前一天的方式
	//return s.UpsertByStatistic(extStaffID, num)

	// 方法二
	// 用count customer-staff 关系表的方式
	return s.UpsertByCount(extStaffID, num)
}

// UpsertByStatistic 用统计表中前一天的方式
func (s CustomerStatistic) UpsertByStatistic(extStaffID string, num int64) (err error) {
	return DB.Transaction(func(tx *gorm.DB) error {
		date := time.Now().Format(constants.DateLayout)
		customerStatistic := CustomerStatistic{}
		// 查当天的员工客户数记录
		err = tx.Model(&CustomerStatistic{}).Where("ext_staff_id = ? and date = ?", extStaffID, date).First(&customerStatistic).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 当天还没有记录，则需要找上一条数据
				err = tx.Model(&CustomerStatistic{}).Where("ext_staff_id = ?", extStaffID).Order("date desc").Limit(1).Find(&customerStatistic).Error
				var cs CustomerStatistic
				if err == gorm.ErrRecordNotFound {
					// 没有找到员工的客户记录，则只可能是新增。
					if num <= 0 {
						return ecode.CustomerNumErr
					}
					cs = CustomerStatistic{
						ExtCorpModel:     ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID},
						ExtStaffID:       extStaffID,
						TotalCustomerNum: num,
						Date:             constants.DateField(date),
					}
				} else if err != nil {
					log2.Sugar.Errorw("query CustomerStatistic failed", "extStaffID", extStaffID)
					return err
				} else {
					// 在前一条数据上累计TotalCustomerNum，避免全表扫描
					cs = CustomerStatistic{
						ExtCorpModel:     ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID},
						ExtStaffID:       extStaffID,
						TotalCustomerNum: customerStatistic.TotalCustomerNum + num,
						Date:             constants.DateField(date),
					}
					if num < 0 {
						cs.DecreaseCustomerNum = -num
					} else {
						cs.IncreaseCustomerNum = num
					}
				}
				log2.Sugar.Debugw("create new customer statistic record", "CustomerStatistic", cs)

				err = tx.Create(&cs).Error
				return err
			}
			return err
		}
		// 当天已有数据，则更新
		customerStatistic.TotalCustomerNum += num
		if num > 0 {
			customerStatistic.IncreaseCustomerNum += num
		} else if num < 0 {
			customerStatistic.DecreaseCustomerNum += -num
		}

		err = tx.Model(&CustomerStatistic{}).Omit("id").
			Where("ext_staff_id = ?", extStaffID).
			Where("date = ?", date).
			Updates(&customerStatistic).Error
		if err != nil {
			log2.Sugar.Errorw("update CustomerStatistic failed", "ext_staff_id", extStaffID)
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// UpsertByCount 用count员工-客户关系表的方式
func (s CustomerStatistic) UpsertByCount(extStaffID string, num int64) (err error) {
	extCorpID := conf.Settings.WeWork.ExtCorpID
	date := time.Now().Format(constants.DateLayout)
	return DB.Transaction(func(tx *gorm.DB) error {
		var total int64
		err = tx.Model(&CustomerStaff{}).
			Where("ext_corp_id = ? and ext_staff_id = ?", extCorpID, extStaffID).Count(&total).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
		cs := CustomerStatistic{
			ExtCorpModel:     ExtCorpModel{ID: id_generator.StringID(), ExtCreatorID: extStaffID, ExtCorpID: extCorpID},
			ExtStaffID:       extStaffID,
			TotalCustomerNum: total,
			Date:             constants.DateField(date),
		}

		// 当天已有数据则使用当天数据,否则新建.
		var data CustomerStatistic
		err = tx.Where("ext_corp_id = ? and ext_staff_id = ? and date = ?", extCorpID, extStaffID, date).
			Attrs(&cs).FirstOrInit(&data).Error
		if err != nil {
			err = errors.WithStack(err)
			return err
		}

		data.TotalCustomerNum = total
		if num < 0 {
			data.DecreaseCustomerNum += 1
		} else {
			data.IncreaseCustomerNum += 1
		}

		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ext_staff_id"}, {Name: "date"}},
			DoUpdates: clause.AssignmentColumns([]string{"total_customer_num", "increase_customer_num", "decrease_customer_num"})},
		).Create(&data).Error
		if err != nil {
			return err
		}

		return nil
	})
}
