package main

import (
	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Address   Address            `bson:"address"`
}

type Address struct {
	AddressLine_1 string `bson:"address_line_1"`
	AddressLine_2 string `bson:"address_line_2,omitempty"`
	City          string `bson:"city"`
	Country       string `bson:"country"`
	PostalCode    string `bson:"postal_code"`
}

func convertMongoDocToUser(data *User) *pb.User {
	return &pb.User{
		Id:        data.Id.Hex(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address: &pb.Address{
			AddressLine_1: data.Address.AddressLine_1,
			AddressLine_2: &data.Address.AddressLine_2,
			City:          data.Address.City,
			Country:       data.Address.Country,
			PostalCode:    data.Address.PostalCode,
		},
	}
}
