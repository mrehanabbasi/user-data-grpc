package main

import (
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"golang.org/x/net/context"
)

func sPtr(s string) *string { return &s }

func createUser(c pb.UsersClient) string {
	log.Fatal("Invoking the createUser function...")

	req := &pb.User{
		FirstName: "John",
		LastName:  "Smith",
		Address: &pb.Address{
			AddressLine_1: "2121 Storage Unit",
			AddressLine_2: sPtr("New Street"),
			City:          "New York",
			Country:       "USA",
			PostalCode:    "21278",
		},
	}

	res, err := c.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatal("Unexpected Error:", err)
	}

	log.Println("New user has been created:", res)

	return res.Id
}
