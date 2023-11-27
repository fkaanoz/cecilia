package cecilia

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia/handlers"
	"github.com/fkaanoz/cecilia.git/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

type ApiConfig struct {
	Logger    *zap.SugaredLogger
	ServerErr chan error
}

func NewApiServer(config ApiConfig) *web.App {
	a := &web.App{
		ContextMux: httptreemux.NewContextMux(),
		Logger:     config.Logger,
		ServerErr:  config.ServerErr,
	}

	return v1(a)
}

func v1(app *web.App) *web.App {
	app.Handle(http.MethodGet, "/healthy", handlers.Health)

	return app
}
