package main

import (
	pb "github.com/mrehanabbasi/user-data-grpc/proto"
)

type Server struct {
	pb.UsersServer
}
