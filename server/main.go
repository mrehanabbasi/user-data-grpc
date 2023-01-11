package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/mrehanabbasi/user-data-grpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Loading environment variables from .env.local file
	godotenv.Load(".env.local")

	// MongoDB Connection
	mongoUri := os.Getenv("MONGODB_URI")

	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB not reachable:", err)
	}

	// Creating/Defining MongoDB database and collection
	collection = client.Database("userdb").Collection("user_data")

	// Creating gRPC server listener
	hostAddress := net.JoinHostPort(os.Getenv("HOST_URL"), os.Getenv("HOST_PORT"))
	fmt.Println("HostAddress:", hostAddress)

	lis, err := net.Listen("tcp", hostAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v\n", hostAddress, err)
	}
	defer lis.Close()
	log.Println("Listening at", hostAddress)

	opts := []grpc.ServerOption{}

	// TLS Certification Authentication functionality
	tls := true // change to false for insecure connection
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

	s := grpc.NewServer(opts...)
	pb.RegisterUsersServer(s, &Server{})
	reflection.Register(s)
	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
