package setup

import (
	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	config := global.Config.Sqlite
	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	db.AutoMigrate()
	return db
}
