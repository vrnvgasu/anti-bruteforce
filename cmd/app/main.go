package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"vrnvgasu/anti-bruteforce/internal/config"
	"vrnvgasu/anti-bruteforce/internal/logger"
	"vrnvgasu/anti-bruteforce/internal/storage"
	"vrnvgasu/anti-bruteforce/internal/storage/postgres"
	"vrnvgasu/anti-bruteforce/internal/storage/redis"
)

var configFile string

func main() {
	ctx := context.Background()
	parseFlags()

	config.Cfg = config.NewConfig(configFile)
	lg := logger.New(config.Cfg.Logger.Level)

	st := mustStartStorage(ctx, lg)
	_ = st

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	<-ctx.Done()

	stopStorage(st, lg)
	lg.Info("shutting down")
}

func stopStorage(st *storage.Storage, lg *logger.Logger) {
	if err := st.ShutDown(); err != nil {
		lg.Error("failed to stop storage" + err.Error())
	}
	lg.Info("stopped storage")
}

func mustStartStorage(ctx context.Context, lg *logger.Logger) *storage.Storage {
	st := storage.NewStorage(postgres.NewPostgresService(), redis.NewRedisService())
	if err := st.Start(ctx); err != nil {
		msg := "failed to start storage" + err.Error()
		lg.Error(msg)
		log.Fatal(msg)
	}
	lg.Info("storage started")

	return st
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "configs/app-dev.yml", "Path to configuration file")

	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
}
