package main

import (
	"context"
	pb "github/erastusk/grpc_overview/blog/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
	col *mongo.Collection
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Blog Service Started")

	collection := client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("Could not listen %+v", err)
	}
	log.Println("Listening on 127.0.0.1:9090")

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{col: collection})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Could not serve %+v", err)
	}
	log.Println("Serving on 127.0.0.1:9090")
}
