package internalgrpc

import (
	"context"
	"fmt"
	"net"
	"vrnvgasu/anti-bruteforce/internal/app"
	"vrnvgasu/anti-bruteforce/internal/config"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"google.golang.org/grpc"
)

type Logger interface {
	Warn(msg string)
	Info(msg string)
	Error(msg string)
}

type Server struct {
	pb.UnimplementedAntiBruteforceServer
	app    *app.App
	l      Logger
	server *grpc.Server
}

func NewServer(logger Logger, app *app.App) *Server {
	return &Server{
		app:    app,
		l:      logger,
		server: grpc.NewServer(),
	}
}

func (s *Server) Start(ctx context.Context) error {
	lsn, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", config.Cfg.GRPSServer.Host, config.Cfg.GRPSServer.Port))
	if err != nil {
		return fmt.Errorf("start grps server Listen: %w", err)
	}

	pb.RegisterAntiBruteforceServer(s.server, s)

	s.l.Info(fmt.Sprintf("starting server on %s", lsn.Addr().String()))
	if err = s.server.Serve(lsn); err != nil {
		return fmt.Errorf("start grpc server Serve: %w", err)
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop() error {
	s.server.GracefulStop()

	return nil
}
