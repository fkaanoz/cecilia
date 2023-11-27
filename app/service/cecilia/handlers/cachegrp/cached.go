package cachegrp

import (
	"context"
	"github.com/fkaanoz/cecilia.git/business/core"
	"net/http"
)

type CacheGrp struct {
	Core *core.RedisCore
}

func (cg *CacheGrp) CachedResultHn(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	sid, err := cg.Core.ReadSessionID("BC8B826B-DC2E-48C5-9D55-8B5041A49378")
	if err != nil {
		w.Write([]byte("no cached result"))
		return nil
	}

	w.Write([]byte(sid))
	return nil
}
