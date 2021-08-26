package models

import (
	"context"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/common/redis"
)

func SetupIDConverter() {
	//err := setupStaffIDConverter()
	//if err != nil {
	//	log.Sugar.Errorw("setupStaffIDConverter failed", "err", err)
	//	os.Exit(1)
	//	return
	//}
	//return
}

func setupStaffIDConverter() (err error) {
	items := make([]Staff, 0)
	err = DB.Model(&Staff{}).Select("id,ext_id").Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "find Staff failed")
		return
	}

	if len(items) == 0 {
		return nil
	}
	values := make([]interface{}, 0)
	for _, item := range items {
		if item.ExtID == "" {
			continue
		}
		values = append(values, item.ExtCorpID, item.ID)
	}

	err = redis.RedisClient.Del(context.Background(), constants.StaffIDConverterKey).Err()
	if err != nil {
		err = errors.Wrap(err, "Del failed")
		return
	}

	err = redis.RedisClient.HMSet(context.Background(), constants.StaffIDConverterKey, values...).Err()
	if err != nil {
		err = errors.Wrap(err, "HMSet failed")
		return
	}

	return
}
