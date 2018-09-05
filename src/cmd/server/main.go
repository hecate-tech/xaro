package main

import (
	"fmt"
	"log"
	"net"

	"github.com/damienfamed75/engo-xaro/src/communication"
	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("► Server Created! Listening for clients...")
	s := grpc.NewServer()
	pb.RegisterXaroServer(s, communication.NewServer())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
