package setup

import (
	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/router"
	"github.com/boj/redistore"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const (
	HEALTH_PATH = "/api/health"
	SECRET_KEY  = "iV6pNvjdHVUVc5Q*Wi4S&" // random
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// session
	store, err := redis.NewStore(
		10,
		"tcp",
		global.Config.Redis.Addr,
		global.Config.Redis.Password,
		[]byte(SECRET_KEY),
	)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	_, rs := redis.GetRedisStore(store)
	rs.SetSerializer(redistore.JSONSerializer{})

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: false,
	})
	r.Use(sessions.Sessions("session", store))

	r.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{HEALTH_PATH}}),
		gin.Recovery(),
	)

	r.GET(HEALTH_PATH, func(ctx *gin.Context) {})

	router.Register(r)

	return r
}
