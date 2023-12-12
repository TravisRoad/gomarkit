package setup

import (
	"fmt"
	"os"

	"github.com/TravisRoad/gomarkit/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initZap() *zap.Logger {
	cfg := global.Config.Log
	level := logLevel(cfg.Level)

	stdoutCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		level,
	)
	cores := []zapcore.Core{stdoutCore}

	if cfg.Rotate {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/log.log", cfg.Dir),
			MaxSize:    cfg.MaxSize, // megabytes
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge, // days
		})
		rotateCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
			w,
			level,
		)
		cores = append(cores, rotateCore)
	}

	var core zapcore.Core
	if len(cores) == 1 {
		core = cores[0]
	} else {
		core = zapcore.NewTee(cores...)
	}

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func logLevel(s string) (level zapcore.Level) {
	switch s {
	case "DPanic":
		level = zap.DPanicLevel
	case "Panic":
		level = zap.PanicLevel
	case "Fatal":
		level = zap.FatalLevel
	case "Error":
		level = zap.ErrorLevel
	case "Info":
		level = zap.InfoLevel
	case "Warn":
		level = zap.WarnLevel
	case "Debug":
		level = zap.DebugLevel
	default:
		level = zap.DebugLevel
	}
	return level
}
