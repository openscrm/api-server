package consumers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/services"
	"openscrm/common/delay_queue"
	"openscrm/common/log"
)

// SyncCustomerData
// Description: 同步客户数据
// Detail: 员工-客户关系表中该关系可能被删除，这里要在upsert时将deleted_at设置为null
// Param: job 异步任务数据
func SyncCustomerData(job delay_queue.Job) (err error) {

	log.Sugar.Info("job info:", job)

	if job.Topic == constants.SyncCustomerDataTopic {
		customerService := services.NewCustomer()

		customerStaffRelation := models.CustomerStaffRelation{}
		err = json.Unmarshal([]byte(job.Body), &customerStaffRelation)
		if err != nil {
			err = errors.WithStack(err)
			return
		}

		extStaffID := customerStaffRelation.ExtStaffID
		extCustomerID := customerStaffRelation.ExtCustomerID
		err = customerService.SyncSingleCustomerData(extStaffID, extCustomerID)
		if err != nil {
			log.Sugar.Error(err)
			return
		}

		// 更新员工的客户数
		err = models.CustomerStatistic{}.Upsert(extStaffID, 1)
		if err != nil {
			log.Sugar.Errorw("update customer statistics failed",
				"extStaffID", extStaffID, "extCustomerID", extCustomerID)
			return err
		}

	}
	return
}
