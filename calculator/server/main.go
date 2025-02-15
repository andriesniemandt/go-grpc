package main

import (
	pb "github.com/andriesniemandt/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:8080"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	srv, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start server on: %v", err)
	}

	log.Printf("Listening on %v", srv.Addr())

	gs := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(gs, &Server{})

	if err = gs.Serve(srv); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
