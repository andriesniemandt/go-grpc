package main

import (
	pb "github.com/andriesniemandt/go-grpc/greetings/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:8080"

type Server struct {
	pb.GreetingServiceServer
}

func main() {
	srv, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start server on: %v", err)
	}

	log.Printf("Listening on %v", srv.Addr())
	defer srv.Close()

	gs := grpc.NewServer()
	if err = gs.Serve(srv); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
