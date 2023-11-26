package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia"
	"github.com/fkaanoz/cecilia.git/business/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var release string
var service = "cecilia"
var build = "dev"

func main() {
	zapLogger, err := logger.InitLogger(service)
	if err != nil {
		log.Println("init logger err ", err)
		os.Exit(1)
	}

	if err := run(zapLogger); err != nil {
		zapLogger.Errorw("RUN", "ERROR", err)
		os.Exit(1)
	}
}

func run(logger *zap.SugaredLogger) error {
	// init conf
	cfg := struct {
		conf.Version
		Web struct {
			Addr            string        `conf:"default:0.0.0.0:8080"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "cecilia service",
		},
	}

	prefix := "CECILIA"
	helpMsg, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(helpMsg)
			return nil
		}
		return err
	}

	serverErrCh := make(chan error, 1)
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGTERM, syscall.SIGINT)

	server := http.Server{
		Addr: cfg.Web.Addr,
		Handler: cecilia.NewApiServer(cecilia.ApiConfig{
			Logger:    logger,
			ServerErr: serverErrCh,
		}),
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}

	go func() {
		logger.Infow("SERVER", "status", "starting...")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorw("SERVER", "ERROR", err)
		}
	}()

	select {
	case <-shutdownCh:
		logger.Infow("SHUTDOWN", "status", "gracefully shutdown is starting")

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Infow("SHUTDOWN", "status", "server is forced to close")
			server.Close()
			return err
		}

		logger.Infow("SHUTDOWN", "status", "gracefully shutdown is finished")
	}

	return nil
}
