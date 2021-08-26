package util

import (
	"openscrm/app/constants"
	"time"
)

// Today 获取今天0点时间
func Today() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, constants.PRCLocation)
}

func Now() time.Time {
	return time.Now().In(constants.PRCLocation)
}
