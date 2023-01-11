package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) CreateUser(ctx context.Context, in *pb.User) (*pb.UserId, error) {
	log.Fatalln("Invoked CreateUser function...")

	data := &pb.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Address:   in.Address,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Internal error:", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Error(codes.Internal, "Cannot convert to OID")
	}

	return &pb.UserId{
		Id: oid.Hex(),
	}, nil
}
