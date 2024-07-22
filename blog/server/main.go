package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var collection *mongo.Collection
var addr = "0.0.0.0:8080"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blog").Collection("blog")

	srv, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start server on: %v", err)
	}

	log.Printf("Listening on %v", srv.Addr())

	gs := grpc.NewServer()
	pb.RegisterBlogServiceServer(gs, &Server{})

	if err = gs.Serve(srv); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
