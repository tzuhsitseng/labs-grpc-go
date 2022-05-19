package main

import (
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tzuhsitseng/labs-grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: in.First + in.Second}, nil
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	if in.Number < 0 {
		return nil, status.Error(codes.InvalidArgument, "Input should be positive")
	}
	return &pb.SqrtResponse{Result: math.Sqrt(float64(in.Number))}, nil
}
