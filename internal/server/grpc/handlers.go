package internalgrpc

import (
	"context"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"
)

func (s *Server) Check(_ context.Context, _ *pb.CheckRequest) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{
		Result: true,
	}, nil
}

func (s *Server) Clear(_ context.Context, _ *pb.ClearRequest) (*pb.OK, error) {
	return &pb.OK{}, nil
}

func (s *Server) AddToList(_ context.Context, _ *pb.ListRequest) (*pb.OK, error) {
	return &pb.OK{}, nil
}

func (s *Server) RemoveFromList(_ context.Context, _ *pb.IP) (*pb.OK, error) {
	return &pb.OK{}, nil
}
