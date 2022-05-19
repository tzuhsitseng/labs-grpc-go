package main

import (
	"log"
	"time"

	pb "github.com/tzuhsitseng/labs-grpc-go/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)

	//doGreetWithDeadline(c, 10*time.Second)
	doGreetWithDeadline(c, 1*time.Second)
}
