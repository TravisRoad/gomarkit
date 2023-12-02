package setup

import (
	"context"
	"log/slog"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/redis/go-redis/v9"
)

func initRedis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	slog.Info("redis connect ping response:", slog.String("pong", pong))
	return client
}
