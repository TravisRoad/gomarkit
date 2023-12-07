package setup

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func initZap() *zap.Logger {
	// logger, err := zap.NewDevelopment(zap.AddCaller())
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 30,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		zap.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller())

	// if err != nil {
	// 	panic(err)
	// }

	return logger
}
