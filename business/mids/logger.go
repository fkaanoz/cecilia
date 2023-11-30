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

			logger.Infow("REQUEST", "status", "started", "trackingID", web.GetTrackingId(ctx))

			if err := handler(ctx, w, r); err != nil {
				logger.Infow("REQUEST", "status", "error", "ERROR", err, "trackingID", web.GetTrackingId(ctx))
				return err
			} else {
				logger.Infow("REQUEST", "status", "successful", "trackingID", web.GetTrackingId(ctx))
			}

			return nil
		}
		return h
	}
	return m
}
