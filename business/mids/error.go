package mids

import (
	"context"
	"github.com/fkaanoz/cecilia.git/foundation/validate"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

func Error(logger *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			if err := handler(ctx, w, r); err != nil {
				//response := validate.ErrorResponse{}

				// TODO:  Unwrap errors and respond accordingly....

				switch validate.Unwrap(err) {

				}
			}

			return nil
		}
		return h
	}
	return m
}
