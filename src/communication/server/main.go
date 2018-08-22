package main

import (
	"context"
	"log"
	"net"

	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

var (
	clients   = make([]pb.Player, 8)
	positions = make([]*pb.Point, 8)
)

// server is used to implement engoxaro.GreeterServer
type server struct {
	iter    int
	clients map[int]*pb.Player
}

/* Pseudo code - EstablishConnection
if the client's uid doesn't exist in my slice
then add client to slice of clients.
*/

// *pb.PositionReply
// slice of positions

/* Pseudo code - SendPositions(ctx context.Context, in *pb.PositionRequest) (*pb.PositionReply, error)
for every client connected {
	add client's position to tempSlice
}
return &pb.PositionReply{positions: tempSlice}
*/

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hola " + in.Name}, nil
}

func (s *server) UserJoined(ctx context.Context, in *pb.Player) (*pb.HelloReply, error) {
	s.clients[s.iter] = in
	// s.clients = append(clients, in)
	// positions = append(positions, in.Position)
	msg := in.Username + " has joined the game."
	log.Println(msg)
	return &pb.HelloReply{Message: msg}, nil
}

func (s *server) SendPositions(ctx context.Context, in *pb.Player) (*pb.PlayerPositions, error) {
	var tempPositions = &pb.PlayerPositions{
		Positions: make([]*pb.Point, len(s.clients)),
	}
	for i, c := range s.clients {
		if c.ID == in.ID {
			c = in
		}
		tempPositions.Positions[i] = c.Position
	}

	return tempPositions, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{
		iter:    0,
		clients: make(map[int]*pb.Player),
	})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
