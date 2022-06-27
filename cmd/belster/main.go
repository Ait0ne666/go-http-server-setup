package main

import (
	"belster/internal/api"
	"belster/internal/config"
	"belster/internal/server"
	"belster/pkg/database/postgres"
	"belster/pkg/logger"
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func main() {
	configPath := new(string)

	flag.StringVar(configPath, "config-path", "config/config-local.yaml", "specify path to yaml")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with os.Open config"))
	}

	cfg := config.Config{}
	if err := yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Decode config"))
	}

	if err = logger.NewLogger(cfg.Telegram); err != nil {
		logger.LogFatal(err)
	}

	postgresClient, err := postgres.NewPostgres(cfg.PostgresDsn)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewPostgres"))
	}

	_, err = postgresClient.Database()
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Gorm"))
	}

	endpoints := api.NewHandlers(&cfg)

	srv := server.NewServer(&cfg, endpoints)

	go func() {

		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
		}

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Shutdown(ctx); err != nil {
		logger.LogFatal(errors.Wrap(err, "failed to stop server"))
	}

	if err = postgresClient.Close(); err != nil {
		logger.LogFatal(errors.Wrap(err, "failed to stop db"))
	}
}
