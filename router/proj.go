package router

import (
	"github.com/TravisRoad/gomarkit/api"
	"github.com/TravisRoad/gomarkit/middleware"
	"github.com/gin-gonic/gin"
)

type ProjRouter struct{}

func (pr *ProjRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/proj")
	projApi := new(api.ProjApi)

	r.Use(middleware.Auth())

	r.GET("/", projApi.GetProj)
	r.POST("/", projApi.AddProj)
	r.POST("/:sqid", projApi.UpdateProj)
	r.DELETE("/", projApi.RemoveProj)
}
