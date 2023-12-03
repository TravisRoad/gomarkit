package setup

import (
	"path/filepath"
	"runtime"

	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/helper"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	DEV_CONFIG_FILE  = "config.dev.yaml"
	PROD_CONFIG_FILE = "config.prod.yaml"
)

func InitViper() {
	mode := helper.Mode()

	v := viper.New()
	configFile := DEV_CONFIG_FILE

	switch mode {
	case global.DEV:
		gin.SetMode(gin.DebugMode)
		configFile = DEV_CONFIG_FILE
	case global.TEST:
		gin.SetMode(gin.TestMode)
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(filepath.Dir(b))
		configFile = filepath.Join(path, DEV_CONFIG_FILE)
	case global.PROD:
		gin.SetMode(gin.ReleaseMode)
		configFile = PROD_CONFIG_FILE
	}

	slog.Info("config file", slog.String("path", configFile), slog.String("mode", mode))

	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
