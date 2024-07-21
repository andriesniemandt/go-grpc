package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("Doing a sum on the server...")
	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  9212,
		SecondNumber: 8024,
	})

	if err != nil {
		log.Fatalf("Failed to do sum: %\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
