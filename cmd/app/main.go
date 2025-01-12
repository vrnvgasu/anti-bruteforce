package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vrnvgasu/anti-bruteforce/internal/app"
	"vrnvgasu/anti-bruteforce/internal/config"
	"vrnvgasu/anti-bruteforce/internal/core/bucket"
	"vrnvgasu/anti-bruteforce/internal/logger"
	internalgrpc "vrnvgasu/anti-bruteforce/internal/server/grpc"
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

	antiBruteforce := app.New(lg, st, bucket.NewBucket(st))
	grpcServer := internalgrpc.NewServer(lg, antiBruteforce)

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			lg.Error("failed to start grpc server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	stopStorage(st, lg)
	stopGrpc(grpcServer, lg)
	lg.Info("shutting down")
}

func stopGrpc(s *internalgrpc.Server, lg *logger.Logger) {
	if err := s.Stop(); err != nil {
		lg.Error("failed to stop grpc server: " + err.Error())
	}
	lg.Info("stopped grpc server")
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
