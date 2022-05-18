package main

import (
	"context"
	"log"

	pb "github.com/tzuhsitseng/labs-grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Chris"})
	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}
	log.Printf("Greeting: %s", res.Result)
}
