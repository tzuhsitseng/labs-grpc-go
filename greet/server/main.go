package main

import (
	"log"
	"net"

	"google.golang.org/grpc/credentials"

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

	opts := make([]grpc.ServerOption, 0)
	tls := true // ff to control whether app use tls

	if tls {
		creds, err := credentials.NewServerTLSFromFile(
			"ssl/server.crt",
			"ssl/server.pem")
		if err != nil {
			log.Fatalf("Failed to load certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
