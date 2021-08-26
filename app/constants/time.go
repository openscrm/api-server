package constants

import "time"

var BeiJinTime = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))

var PRCLocation = BeiJinTime

const (
	TimeLayout     = "15:04:05"
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"
)

var WeekdayMap = map[string]int{
	"周一": 1,
	"周二": 2,
	"周三": 3,
	"周四": 4,
	"周五": 5,
	"周六": 6,
	"周日": 0,
}
