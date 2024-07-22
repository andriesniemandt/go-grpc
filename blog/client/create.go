package main

import (
	"context"
	pb "github.com/andriesniemandt/go-grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("call createBlog")
	blog := &pb.Blog{
		Author:  "Andries",
		Title:   "Very first gRPC blog",
		Content: "Content of the first gRPC blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}
	log.Printf("CreateBlog response %s\n", res)
	return res.Id
}
