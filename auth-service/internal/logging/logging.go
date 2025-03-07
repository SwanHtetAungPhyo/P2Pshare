package logging

import (
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var colorFunc func(format string, a ...interface{}) string
	switch level {
	case zapcore.DebugLevel:
		colorFunc = color.New(color.FgBlue).SprintfFunc()
	case zapcore.InfoLevel:
		colorFunc = color.New(color.FgGreen).SprintfFunc()
	case zapcore.WarnLevel:
		colorFunc = color.New(color.FgYellow).SprintfFunc()
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		colorFunc = color.New(color.FgRed).SprintfFunc()
	default:
		colorFunc = fmt.Sprintf
	}
	enc.AppendString(colorFunc(level.CapitalString()))
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = CustomLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	Logger = logger
}

func CloseLogger() {
	if Logger != nil {
		Logger.Sync()
	}
}
