package setup

import (
	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	cfg := global.Config.Sqlite
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	db.AutoMigrate()

	slog.Info("successfully connect to db", slog.String("path", cfg.Path))
	return db
}
