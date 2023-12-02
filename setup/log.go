package setup

import (
	"os"

	"go.uber.org/zap"
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
	logger, err := zap.NewDevelopment(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	return logger
}
