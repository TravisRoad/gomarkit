package router

import (
	"github.com/TravisRoad/gomarkit/api"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (*AuthRouter) Register(r *gin.RouterGroup) {
	rt := r.Group("/auth")
	authApi := &api.AuthApi{}

	{
		rt.POST("/login", authApi.Login)
		rt.POST("/logout", authApi.Logout)
		rt.GET("/islogin", authApi.IsLogin)
	}
}
