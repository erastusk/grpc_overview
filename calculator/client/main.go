package main

import (
	"context"
	pb "github/erastusk/grpc_overview/calculator/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not Dial %+v", err)
	}
	s := pb.NewCalculatorServiceClient(conn)
	resp, err := s.Sum(context.Background(), &pb.CalculatorRequest{
		A: 12,
		B: 8,
	})
	if err != nil {
		log.Fatalf("Could not call service %+v", err)
	}
	log.Println(resp)

}
