package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

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

func doSqrt(c pb.CalculatorServiceClient, num int32) {
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: num,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Printf("Get grpc error code: %s\n", e.Code())
			log.Printf("Get grpc error msg: %s\n", e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("Get non-grpc error: %v\n", err)
		}

	}
	log.Printf("Sqrt: %f", res.Result)
}
