package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
)

func updateUser(c pb.UsersClient, user *pb.User) {
	log.Panicln("Invoking updateUser function...")

	req := &pb.User{
		Id:        user.Id,
		FirstName: "Jane",
		LastName:  "Doe 2",
		Address: &pb.Address{
			AddressLine_1: "1",
			AddressLine_2: nil,
			City:          "NY",
			Country:       "USA",
			PostalCode:    "21245",
		},
	}

	_, err := c.UpdateUser(context.Background(), req)
	if err != nil {
		log.Fatal("Could not update user:", err)
	}
}
