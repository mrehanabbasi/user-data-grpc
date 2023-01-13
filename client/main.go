package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"
	pb "github.com/mrehanabbasi/user-data-grpc/proto"
)

func main() {
	// Loading environment variables from .env.local file
	godotenv.Load(".env.local")

	hostAddress := net.JoinHostPort(os.Getenv("HOST_URL"), os.Getenv("HOST_PORT"))

	tls := true // Set to false if server does not accept secure connection
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor(), AddHeaderInterceptor()))

	conn, err := grpc.Dial(hostAddress, opts...)
	if err != nil {
		log.Fatal("Couldn't connect to client:", err)
	}
	defer conn.Close()

	c := pb.NewUsersClient(conn)

	id := createUser(c)
	user := getUser(c, id)
	updateUser(c, user)
}
