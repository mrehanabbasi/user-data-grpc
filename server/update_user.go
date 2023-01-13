package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) UpdateUser(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	log.Println("Invoking Updateuser function...")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot parse ID")
	}

	filter := bson.M{"_id": oid}
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
