package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/greetings/proto"
	"log"
)

func (s *Server) Greeting(ctx context.Context, r *pb.GreetingRequest) (pb.GreetingResponse, error) {
	log.Printf("Greeting called by server")
	return pb.GreetingResponse{
		Result: "Hello, " + r.FirstName,
	}, nil
}
