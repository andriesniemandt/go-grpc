package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/greetings/proto"
	"log"
)

func doGreeting(c pb.GreetingServiceClient) {
	log.Println("Greeting the server...")
	res, err := c.Greet(context.Background(), &pb.GreetingRequest{
		FirstName: "Andries",
	})

	if err != nil {
		log.Fatalf("Failed to do greeting: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
