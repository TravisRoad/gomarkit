package setup

import (
	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	var db *gorm.DB
	slog.Info("connecting to database", slog.String("type", global.Config.Database.Type))
	switch global.Config.Database.Type {
	case "sqlite":
		db = initSqlite()
	case "mysql":
		db = initMysql()
	default:
		slog.Error("unsupported database type", slog.String("type", global.Config.Database.Type))
		panic("unsupported database type")
	}

	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		slog.Error("failed during automigration", "err", err.Error())
		panic(err)
	}

	return db
}

func initSqlite() *gorm.DB {
	cfg := global.Config.Database.Sqlite
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	slog.Info("successfully connect to db", slog.String("path", cfg.Path))
	return db
}

func initMysql() *gorm.DB {
	m := global.Config.Database.Mysql
	dsn := m.Dsn()
	cfg := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(cfg))
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	slog.Info("successfully connect to db", slog.String("dsn", dsn))
	return db
}
