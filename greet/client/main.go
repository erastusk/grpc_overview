package main

import (
	"context"
	pb "github/erastusk/grpc_overview/greet/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "127.0.0.1:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not dial")
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn) //create client
	resp, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Erastus"})
	if err != nil {
		log.Fatalf("Couldn't send greeting %+v %+v", err, resp)
	}
	log.Println(resp)

}
