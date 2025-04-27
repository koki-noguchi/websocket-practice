package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	SugaredLogger *zap.SugaredLogger
)

func Init() {
	writer := zapcore.AddSync(os.Stdout)

	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "massage",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level := zapcore.DebugLevel

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		writer,
		level,
	)

	rawLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	SugaredLogger = rawLogger.Sugar()
}

func S() *zap.SugaredLogger {
	return SugaredLogger
}
