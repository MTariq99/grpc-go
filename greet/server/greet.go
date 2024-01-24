package main

import (
	"context"
	pb "grpcgo/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet funtion was invoked by : %v", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil

}
