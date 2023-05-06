package main

import (
	"context"
	pb "github/erastusk/grpc_overview/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "127.0.0.1:5051"

// Generated Greetservice server is only defined as an interface with no implementation of the required
// methods. That's why handle it here, otherwise that could've been used.
type Server struct {
	pb.GreetServiceServer
}

// Implement GreetServiceServer Interface requirements Greet for Server
func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function invoked with:", req)
	return &pb.GreetResponse{
		Result: "Hello : " + req.FirstName,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Could not start server")
	}
	log.Println("Listening on port:", addr)
	s := grpc.NewServer()                       // Server registrar
	pb.RegisterGreetServiceServer(s, &Server{}) //Register greet server with new server
	err = s.Serve(listen)                       //~conn.Accept() Serve accepts incoming connections on the listener listen
	if err != nil {
		log.Fatalf("Could not server")
	}
}
