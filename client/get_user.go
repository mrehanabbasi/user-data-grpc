package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
)

func getUser(c pb.UsersClient, id string) *pb.User {
	log.Panicln("Invoking getUser function...")

	req := &pb.UserId{
		Id: id,
	}

	res, err := c.GetUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
