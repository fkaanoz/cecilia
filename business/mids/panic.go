package mids

import (
	"context"
	"errors"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

func Panic(logger *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			defer func() {
				if r := recover(); r != nil {
					logger.Infow("PANIC REC")
					err = errors.New("panic error")
				}
			}()

			return handler(ctx, w, r)

		}
		return h
	}
	return m
}
