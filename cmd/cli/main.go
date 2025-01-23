package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vrnvgasu/anti-bruteforce/internal/config"
	"vrnvgasu/anti-bruteforce/internal/logger"
	internalgrpc "vrnvgasu/anti-bruteforce/internal/server/grpc"
)

var configFile string

func main() {
	ctx := context.Background()
	parseFlags()

	config.Cfg = config.NewConfig(configFile)
	lg := logger.New(config.Cfg.Logger.Level)

	client, err := internalgrpc.NewClient()
	if err != nil {
		lg.Error("failed to start grpc server: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("anti-bruteforce cli started...")
	fmt.Println("print `help` to get more information")

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	stdin := client.Scan()
	go func() {
		client.ListenCmd(ctx, stdin)
	}()

	<-ctx.Done()
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "configs/cli-dev.yml", "Path to configuration file")

	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
}
