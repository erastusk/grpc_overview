package main

import (
	"context"
	pb "github/erastusk/grpc_overview/blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not Dial %+v", err)
	}
	s := pb.NewBlogServiceClient(conn)
	t := pb.Blog{
		AuthorId: "222",
		Title:    "Golang",
		Content:  "Go Stuff",
	}
	resp, err := s.CreateBlog(context.Background(), &t)
	if err != nil {
		log.Printf("Could not create blog %+v", err)
	}
	log.Println(resp)

}
