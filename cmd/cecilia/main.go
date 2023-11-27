package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/fkaanoz/cecilia.git/app/service/cecilia"
	"github.com/fkaanoz/cecilia.git/business/logger"
	"github.com/fkaanoz/cecilia.git/foundation/redis"
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
			ReadTimeout     time.Duration `conf:"default:20s"`
			WriteTimeout    time.Duration `conf:"default:20s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
		}
		Redis struct {
			Username string `conf:"default:fkaanoz"`
			Password string `conf:"default:7733700991bcfd88576227e4f4eb656a5c75e9d151c5d6233ea6d081ef97d78c"`
			Host     string `conf:"default:0.0.0.0"`
			Port     string `conf:"default:6379"`
			Database string `conf:"cecilia"`
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

	// init redis client
	redisClient, err := redis.Connect(redis.Config{
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Database: cfg.Redis.Database,
	})

	if err != nil {
		return err
	}

	server := http.Server{
		Addr: cfg.Web.Addr,
		Handler: cecilia.NewApiServer(cecilia.ApiConfig{
			Logger:      logger,
			ServerErr:   serverErrCh,
			RedisClient: redisClient,
		}),
		ReadTimeout:  cfg.Web.ReadTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
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
