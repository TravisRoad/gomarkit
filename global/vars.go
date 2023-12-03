package global

import (
	"github.com/TravisRoad/gomarkit/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	TEST = "TEST"
	DEV  = "DEV"
	PROD = "PROD"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
	Redis  *redis.Client
)

const (
	USER_INFO_KEY = "user_info"
)

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)
