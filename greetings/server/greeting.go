package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/greetings/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, r *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Greeting invoked with %v\n", r)

	return &pb.GreetingResponse{
		Result: "Hello, " + r.FirstName,
	}, nil
}
