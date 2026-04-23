package log

import (
	"fmt"
	"openscrm/app/constants"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger
var Env string

func SetupLogger(env string) {
	Env = env
	var err error

	// 开发环境使用彩色日志
	if env == constants.PROD {
		Logger, err = zap.NewProduction()
	} else {
		// 开发环境配置彩色日志
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色级别
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		encoder := zapcore.NewConsoleEncoder(encoderConfig)
		Logger = zap.New(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	}

	if err != nil {
		panic(err)
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

// Error 打印错误日志，支持彩色显示
func Error(args ...interface{}) {
	Sugar.Error(args...)
}

// Errorf 打印格式化错误日志，支持彩色显示
func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

// Info 打印信息日志，支持彩色显示
func Info(args ...interface{}) {
	Sugar.Info(args...)
}

// Infof 打印格式化信息日志，支持彩色显示
func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}
