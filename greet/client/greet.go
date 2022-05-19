package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/tzuhsitseng/labs-grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Chris"})
	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}
	log.Printf("Greeting: %s", res.Result)
}

func doGreetWithDeadline(c pb.GreetServiceClient, t time.Duration) {
	ctx, cancel := context.WithTimeout(context.TODO(), t)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "Chris"})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.Canceled {
				log.Println("The request is timeout")
				return
			}
		} else {
			log.Fatalf("Get non-grpc error: %v\n", err)
		}
	}
	log.Printf("Greeting: %s", res.Result)
}
