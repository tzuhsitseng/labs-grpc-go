package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/tzuhsitseng/labs-grpc-go/greet/proto"
)

const addr = ":50051"

type Server struct {
	pb.GreetServiceServer // XXX: maybe embed "UnimplementedGreetServiceServer" is better
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
