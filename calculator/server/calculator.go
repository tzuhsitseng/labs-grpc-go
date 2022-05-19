package main

import (
	"context"

	pb "github.com/tzuhsitseng/labs-grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: in.First + in.Second}, nil
}
