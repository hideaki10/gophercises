package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-exercises/helloworld"

	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:5050"
)

func main() {
	//
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//defer
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	name := "grpc"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	replay, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", replay.Message)

}
