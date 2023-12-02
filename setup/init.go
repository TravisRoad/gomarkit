package setup

import "github.com/TravisRoad/gomarkit/global"

func Setup() {
	InitViper()
	global.DB = initDB()
	global.Logger = initZap()
	global.Redis = initRedis()
}
