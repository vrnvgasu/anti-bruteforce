package internalgrpc

import (
	"context"
	"vrnvgasu/anti-bruteforce/internal/app"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	ok, err := s.app.Check(ctx, app.CheckRequest{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
		IP:       req.GetIp(),
	})
	if err != nil {
		s.l.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "server grpc Check app.Check failed")
	}

	return &pb.CheckResponse{
		Result: ok,
	}, nil
}

func (s *Server) Clear(ctx context.Context, req *pb.ClearRequest) (*pb.OK, error) {
	err := s.app.Clear(ctx, app.ClearRequest{
		Login: req.GetLogin(),
		IP:    req.GetIp(),
	})
	if err != nil {
		s.l.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "server grpc Clear app.Clear failed")
	}

	return &pb.OK{}, nil
}

func (s *Server) AddToList(ctx context.Context, req *pb.ListRequest) (*pb.OK, error) {
	err := s.app.AddToList(ctx, app.AddToListRequest{
		ListType: req.GetType(),
		Subnet:   req.GetSubnet(),
	})
	if err != nil {
		s.l.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "server grpc AddToList app.AddToList failed")
	}

	return &pb.OK{}, nil
}

func (s *Server) RemoveFromList(ctx context.Context, req *pb.ListRequest) (*pb.OK, error) {
	err := s.app.RemoveFromWhiteList(ctx, app.RemoveFromListRequest{
		ListType: req.GetType(),
		Subnet:   req.GetSubnet(),
	})
	if err != nil {
		s.l.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "server grpc RemoveFromList app.RemoveFromList failed")
	}

	return &pb.OK{}, nil
}
