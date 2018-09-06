package communication

import (
	"context"
	"errors"
	"log"

	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/util"
)

// Server is used to implement engoxaro.GreeterServer
type Server struct {
	clients   map[uint32]*pb.Player
	idManager *util.IDManager
}

// NewServer returns a new pointer to valid server
func NewServer() *Server {
	return &Server{
		clients:   make(map[uint32]*pb.Player),
		idManager: util.NewIDManager(),
	}
}

// UserJoined returns a printed message and replies to user
func (s *Server) UserJoined(ctx context.Context, in *pb.Player) (*pb.JoinMessage, error) {
	newID := s.idManager.NextPlayerID()
	s.clients[newID] = in
	s.clients[newID].AnimName = "downidle"

	log.Printf("%v (%v) has joined the game with IP: %v", in.Username, newID, in.IP)
	log.Printf("%v players connected.\n", len(s.clients))

	return &pb.JoinMessage{Message: "You have connected to server", Newid: newID}, nil
}

// SendPlayerData you send your current position and are returned with the other player's positions
func (s *Server) SendPlayerData(ctx context.Context, in *pb.Player) (*pb.Players, error) {
	if _, ok := s.clients[in.ID]; !ok {
		return nil, errors.New("in SendPlayerData 'map[uint32]*pb.Player clients' did not have player {" + in.Username + "," + string(in.ID) + "} inside")
	}

	// Send this information to the players.
	s.clients[in.ID].Position = in.Position
	s.clients[in.ID].AnimName = in.AnimName

	// Assure you're not sending the player who requested.
	sPlayers := &pb.Players{Players: make(map[uint32]*pb.Player)}
	for i, c := range s.clients {
		if c.ID != in.ID {
			sPlayers.Players[i] = c
		}
	}

	return sPlayers, nil
}

// UserLeft removes player once they leave
func (s *Server) UserLeft(ctx context.Context, in *pb.Player) (*pb.ServerMessage, error) {
	if _, ok := s.clients[in.ID]; ok {
		delete(s.clients, in.ID)
	}

	log.Printf("%v (%v) has left the game with IP: %v\n", in.Username, in.ID, in.IP)
	return &pb.ServerMessage{Message: "You have disconnected from server"}, nil
}

// UsersConnected gets the user count of the server
func (s *Server) UsersConnected(ctx context.Context, in *pb.Empty) (*pb.PlayerCount, error) {
	return &pb.PlayerCount{Count: int32(len(s.clients))}, nil
}
