package cecilia

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia/handlers"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia/handlers/usergrp"
	"github.com/fkaanoz/cecilia.git/business/core"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

type ApiConfig struct {
	Logger      *zap.SugaredLogger
	ServerErr   chan error
	RedisClient *redis.Client
	Database    *sqlx.DB
}

func NewApiServer(config ApiConfig) *web.App {
	a := &web.App{
		ContextMux:  httptreemux.NewContextMux(),
		Logger:      config.Logger,
		ServerErr:   config.ServerErr,
		RedisClient: config.RedisClient,
		Database:    config.Database,
	}

	return v1(a)
}

func v1(app *web.App) *web.App {
	app.Handle(http.MethodGet, "/healthy", handlers.Health)

	ug := usergrp.UserGroup{
		Core:  core.NewUserCore(app.Database),
		Redis: core.NewRedisCore(app.RedisClient, app.Logger),
	}

	app.Handle(http.MethodGet, "/sid", ug.GetSID)

	return app
}
