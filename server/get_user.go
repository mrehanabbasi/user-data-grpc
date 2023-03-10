package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	log.Println("Invoked GetUser function...")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot parse ID")
	}

	user := User{}
	filter := bson.M{"_id": oid}

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v\n", err)
	}

	return convertMongoDocToUser(&user), nil
}
