package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	yyMMdd := time.Now().Local().Format("060101")
	var err error
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, fmt.Sprintf("logs/out-%s.log", yyMMdd))
	config.ErrorOutputPaths = append(config.ErrorOutputPaths, fmt.Sprintf("logs/error-%s.log", yyMMdd))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}