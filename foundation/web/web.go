package web

import (
	"context"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
type Middleware func(handler Handler) Handler

type App struct {
	*httptreemux.ContextMux
	Logger      *zap.SugaredLogger
	ServerErr   chan error
	RedisClient *redis.Client
	Database    *sqlx.DB
}

func (a *App) Handle(method string, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		if err := handler(ctx, w, r); err != nil {
			a.ServerErr <- err
		}
	}

	a.ContextMux.Handle(method, path, h)
}
