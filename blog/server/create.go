package main

import (
	"context"
	"fmt"
	pb "github.com/andriesniemandt/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create blog invoked %v", in)
	data := BlogItem{
		Author:  in.Author,
		Title:   in.Title,
		Content: in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Error(codes.Internal, "Cannot convert to ObjectID")
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
