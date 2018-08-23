package communication

import (
	"context"
	"log"

	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/utils"
)

// Server is used to implement engoxaro.GreeterServer
type Server struct {
	clients   map[uint32]*pb.Player
	idManager *utils.IDManager
}

// NewServer returns a new pointer to valid server
func NewServer() *Server {
	return &Server{
		clients:   make(map[uint32]*pb.Player),
		idManager: utils.NewIDManager(),
	}
}

// UserJoined returns a printed message and replies to user
func (s *Server) UserJoined(ctx context.Context, in *pb.Player) (*pb.JoinMessage, error) {
	newID := s.idManager.NextPlayerID()
	in.ID = newID
	s.clients[in.ID] = in
	log.Println("IN:", in.ID)
	log.Println("NEW:", newID)

	log.Printf("%v (%v) has joined the game with IP: %v", in.Username, in.ID, in.IP)
	log.Printf("%v players connected.\n", len(s.clients))
	return &pb.JoinMessage{Message: "You have connected to server...", Newid: newID}, nil
}

// SendPlayerData you send your current position and are returned with the other player's positions
func (s *Server) SendPlayerData(ctx context.Context, in *pb.Player) (*pb.Players, error) {
	s.clients[in.ID] = in

	var players = &pb.Players{
		Players: s.clients,
	}

	for _, p := range players.Players {
		log.Printf("Player %v (%v): {%v, %v}", p.Username, p.ID, p.Position.GetX(), p.Position.GetY())
	}

	return players, nil
}

// UserLeft removes player once they leave
func (s *Server) UserLeft(ctx context.Context, in *pb.Player) (*pb.ServerMessage, error) {
	_, ok := s.clients[in.ID]
	if ok {
		delete(s.clients, in.ID)
	}
	log.Printf("%v (%v) has left the game with IP: %v\n", in.Username, in.ID, in.IP)
	return &pb.ServerMessage{Message: "You have disconnected from server..."}, nil
}

// UsersConnected gets the user count of the server
func (s *Server) UsersConnected(ctx context.Context, in *pb.Empty) (*pb.PlayerCount, error) {
	return &pb.PlayerCount{Count: int32(len(s.clients))}, nil
}
