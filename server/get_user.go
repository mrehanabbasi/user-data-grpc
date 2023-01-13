package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	log.Println("Invoked GetUser function...")

	user := User{}

	err := collection.FindOne(ctx, bson.M{"id": in.Id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v\n", err)
	}

	res, err := convertMongoDocToUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot convert from MongoDB to User object")
	}

	return res, nil
}
