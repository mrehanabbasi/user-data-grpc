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

func (*Server) UpdateUser(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	log.Println("Invoking Updateuser function...")

	filter := bson.M{"id": in.Id}
	update := bson.M{"$set": in}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v\n", err)
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &emptypb.Empty{}, nil
}
