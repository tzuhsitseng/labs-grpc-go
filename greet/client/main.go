package main

import (
	"log"

	pb "github.com/tzuhsitseng/labs-grpc-go/greet/proto"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	tls := true
	opts := make([]grpc.DialOption, 0)

	// go1.18 has by default started rejecting SHA-1 certs (https://go.dev/doc/go1.18#sha1)
	// so this example would be worked by add runtime env "GODEBUG=x509sha1=1"
	if tls {
		creds, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatalf("Failed to load ca cert: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Printf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	//doGreetWithDeadline(c, 10*time.Second)
	//doGreetWithDeadline(c, 1*time.Second)
}
