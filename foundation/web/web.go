package web

import (
	"context"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
type Middleware func(handler Handler) Handler

type App struct {
	*httptreemux.ContextMux
	Logger      *zap.SugaredLogger
	ServerErr   chan error
	RedisClient *redis.Client
	Database    *sqlx.DB
	Middlewares []Middleware
}

func (a *App) Handle(method string, path string, handler Handler, middlewares ...Middleware) {

	// wrapping with application wise middlewares
	handler = wrapMiddleware(handler, a.Middlewares...)

	// wrapping with handler wise middlewares
	handler = wrapMiddleware(handler, middlewares...)

	h := func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), valuesKey, Values{
			TrackingID: uuid.New().String(),
		})

		if err := handler(ctx, w, r); err != nil {
			a.ServerErr <- err
		}
	}

	a.ContextMux.Handle(method, path, h)
}
