package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, r *pb.CalculatorRequest) (*pb.CalculatorResponse,
	error) {
	log.Printf("Calc invoked with %v\n", r)

	return &pb.CalculatorResponse{
		Result: r.FirstNumber + r.SecondNumber,
	}, nil
}
