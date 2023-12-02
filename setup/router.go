package setup

import (
	"github.com/TravisRoad/gomarkit/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	router := new(router.Router)
	router.Register(r)

	return r
}
