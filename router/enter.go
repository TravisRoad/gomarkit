package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.RouterGroup)
}

func Register(r *gin.Engine) {
	rr := r.Group("/api")

	rts := []Router{
		&AuthRouter{},
		&AdminRouter{},
	}

	for _, rt := range rts {
		rt.Register(rr)
	}
}
