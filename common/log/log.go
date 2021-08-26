package log

import (
	"fmt"
	"go.uber.org/zap"
	"openscrm/app/constants"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger
var Env string

func SetupLogger(env string) {
	Env = env
	Logger, _ = zap.NewDevelopment()
	if env == constants.PROD {
		Logger, _ = zap.NewProduction()
	}
	defer Logger.Sync() // flushes buffer, if any
	Sugar = Logger.Sugar()
}

// TracedError 打印错误，线上环境固定打Json格式，其他环境打Console格式
func TracedError(msg string, err error) {
	if Env == constants.PROD {
		Sugar.Errorw(msg, "err", err)
		return
	} else {
		fmt.Printf("%s %+v", msg, err)
		return
	}
}
