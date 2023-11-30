package web

import (
	"context"
)

type contextKey struct {
	key string
}

var valuesKey contextKey = contextKey{
	key: "custom-ctx-key",
}

type Values struct {
	TrackingID string
	StatusCode int
}

func GetTrackingId(ctx context.Context) string {
	values, ok := ctx.Value(valuesKey).(Values)
	if !ok {
		return ""
	}
	return values.TrackingID
}

func GetStatusCode(ctx context.Context) int {
	values, ok := ctx.Value(valuesKey).(Values)
	if !ok {
		return 0
	}
	return values.StatusCode
}

func SetStatusCode(ctx context.Context, statusCode int) context.Context {
	trackingID := GetTrackingId(ctx)

	newCtx := context.WithValue(ctx, valuesKey, Values{
		TrackingID: trackingID,
		StatusCode: statusCode,
	})

	return newCtx
}
