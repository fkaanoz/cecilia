package cecilia

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia/handlers"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia/handlers/cachegrp"
	"github.com/fkaanoz/cecilia.git/business/core"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

type ApiConfig struct {
	Logger      *zap.SugaredLogger
	ServerErr   chan error
	RedisClient *redis.Client
}

func NewApiServer(config ApiConfig) *web.App {
	a := &web.App{
		ContextMux:  httptreemux.NewContextMux(),
		Logger:      config.Logger,
		ServerErr:   config.ServerErr,
		RedisClient: config.RedisClient,
	}

	return v1(a)
}

func v1(app *web.App) *web.App {
	app.Handle(http.MethodGet, "/healthy", handlers.Health)

	// initiate redis grp
	redisgrp := cachegrp.CacheGrp{Core: core.NewRedisCore(app.RedisClient, app.Logger)}
	app.Handle(http.MethodGet, "/cached/sid", redisgrp.CachedResultHn)

	return app
}
