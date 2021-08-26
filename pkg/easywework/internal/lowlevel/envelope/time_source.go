package envelope

import (
	"time"
)

type TimeSource interface {
	GetCurrentTimestamp() time.Time
}

type DefaultTimeSource struct{}

var _ TimeSource = DefaultTimeSource{}

func (DefaultTimeSource) GetCurrentTimestamp() time.Time {
	return time.Now()
}
