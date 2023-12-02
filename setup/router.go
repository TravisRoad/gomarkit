package setup

import (
	"github.com/TravisRoad/gomarkit/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	router.Register(r)

	return r
}
