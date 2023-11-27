package handlers

import (
	"context"
	"net/http"
)

func Health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

	return nil
}
