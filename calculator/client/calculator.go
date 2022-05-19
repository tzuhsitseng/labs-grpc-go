package main

import (
	"context"
	"log"

	pb "github.com/tzuhsitseng/labs-grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		First:  1,
		Second: 1,
	})
	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}
	log.Printf("Sum: %d", res.Result)
}
