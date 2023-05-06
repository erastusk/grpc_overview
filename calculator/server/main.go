package main

import (
	"context"
	pb "github/erastusk/grpc_overview/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		A: req.A + req.B,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("Could not listen %+v", err)
	}
	log.Println("Listening on 127.0.0.1:9090")
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Could not serve %+v", err)
	}
	log.Println("Serving on 127.0.0.1:9090")
}
