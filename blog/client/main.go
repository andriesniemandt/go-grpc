package main

import (
	pb "github.com/andriesniemandt/go-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:8080"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	createBlog(c)
}
