package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) DeleteUser(ctx context.Context, in *pb.UserId) (*emptypb.Empty, error) {
	log.Println("Invoking DeleteUser function...")

	res, err := collection.DeleteOne(ctx, bson.M{"id": in.Id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not delete user: %v\n", err)
	}

	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &emptypb.Empty{}, nil
}
