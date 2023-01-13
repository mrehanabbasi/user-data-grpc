package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) DeleteUser(ctx context.Context, in *pb.UserId) (*emptypb.Empty, error) {
	log.Println("Invoking DeleteUser function...")

	_, err := collection.DeleteOne(ctx, bson.M{"id": in.Id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "Could not delete user: %v\n", err)
	}

	return &emptypb.Empty{}, nil
}
