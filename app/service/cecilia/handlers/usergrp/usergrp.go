package usergrp

import (
	"context"
	"fmt"
	"github.com/fkaanoz/cecilia.git/business/core"
	"net/http"
)

type UserGroup struct {
	Core  *core.UserCore
	Redis *core.RedisCore
}

func (ug *UserGroup) GetSID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// decode body or headers and deal with errors

	// read from redis
	cachedSid, err := ug.Redis.ReadSessionID("test-user-id")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("cachedSID:%s", cachedSid)))
		return nil
	}

	// read from db
	dbSid, err := ug.Core.ReadSessionID("test-user-id")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("sid not found"))
		return nil
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("dbSID:%s", dbSid)))

	return nil
}
