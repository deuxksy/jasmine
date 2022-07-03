package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	yyMMdd := time.Now().Local().Format("060101")
	var err error
	config := zap.NewProductionConfig()
	if config.Level, err = zap.ParseAtomicLevel("debug"); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(fmt.Sprintf("logs/%s", yyMMdd), os.ModePerm); err != nil {
		panic(err)
	}
	config.OutputPaths = append(config.OutputPaths, fmt.Sprintf("logs/%s/out.log", yyMMdd))
	config.ErrorOutputPaths = append(config.ErrorOutputPaths, fmt.Sprintf("logs/%s/error.log", yyMMdd))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig
	if log, err = config.Build(zap.AddCallerSkip(1)); err != nil {
		panic(err)
	}
}

func Debug(template string, args ...interface{}) {
	log.Sugar().Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	log.Sugar().Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	log.Sugar().Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	log.Sugar().Errorf(template, args...)
}

// func DPanic(template string, args ...interface{}) {
// 	log.Sugar().DPanicf(template, args...)
// }

func Panic(template string, args ...interface{}) {
	log.Sugar().Panicf(template, args...)
}
