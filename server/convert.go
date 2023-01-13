package main

import (
	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Id        string  `bson:"id,omitempty"`
	FirstName string  `bson:"first_name"`
	LastName  string  `bson:"last_name"`
	Address   Address `bson:"address"`
}

type Address struct {
	AddressLine_1 string `bson:"address_line_1"`
	AddressLine_2 string `bson:"address_line_2,omitempty"`
	Street        string `bson:"street"`
	City          string `bson:"city"`
	Country       string `bson:"country"`
	PostalCode    string `bson:"postal_code"`
}

func convertMongoDocToUser(mongoUser User) (*pb.User, error) {
	mongoData, err := bson.Marshal(mongoUser)
	if err != nil {
		return nil, err
	}

	user := &pb.User{}
	err = bson.Unmarshal(mongoData, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
