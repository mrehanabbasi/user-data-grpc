package main

import (
	"context"
	"log"

	pb "github.com/mrehanabbasi/user-data-grpc/proto"
)

func deleteUser(c pb.UsersClient, id string) {
	log.Println("Invoking deleteUser function...")

	_, err := c.DeleteUser(context.Background(), &pb.UserId{Id: id})
	if err != nil {
		log.Fatal("Could not delete user:", err)
	}
}
