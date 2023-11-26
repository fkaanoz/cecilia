package mids

import (
	"context"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

func Logger(logger *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// logging stuff should be done in here

			return nil
		}

		return h
	}

	return m
}
