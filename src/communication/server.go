package communication

import (
	"context"
	"log"

	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/util"
)

// Server is used to receive queries from the clients
// and give them responses of the proper information.
type Server struct {
	clients   map[uint32]*pb.Player
	idManager *util.IDManager
}

// NewServer returns a new pointer to a
// ready server to use for the server executable.
func NewServer() *Server {
	return &Server{
		clients:   make(map[uint32]*pb.Player),
		idManager: util.NewIDManager(),
	}
}

// UserJoined returns a join message to the client after creating
// a user ID and adding the new client to the map of clients.
func (s *Server) UserJoined(ctx context.Context, in *pb.Player) (*pb.JoinMessage, error) {
	newID := s.idManager.NextPlayerID()    // generates a new ID to be used for a player.
	s.clients[newID] = in                  // saves the new client to a map of clients.
	s.clients[newID].AnimName = "downidle" // sets the default animation of server player.

	// logs for the server client to keep track of what's going on.
	log.Printf("%v (%v) has joined the game with IP: %v", in.Username, newID, in.IP)
	log.Printf("%v players connected.\n", len(s.clients))

	return &pb.JoinMessage{Message: "You have connected to server", Newid: newID}, nil
}

// SendPlayerData is used from the client to query the server
// their current data and the server will save it and return the
// other players' data in order for the query to render them.
func (s *Server) SendPlayerData(ctx context.Context, in *pb.Player) (*pb.Players, error) {

	// Save the important information about the player.
	s.clients[in.ID].Position = in.Position
	s.clients[in.ID].AnimName = in.AnimName

	// Make sure that the information you send back does not
	// belong to the client who queried initially.
	sPlayers := &pb.Players{Players: make(map[uint32]*pb.Player)}
	for i, c := range s.clients {
		if c.ID != in.ID {
			sPlayers.Players[i] = c
		}
	}

	return sPlayers, nil
}

// UserLeft removes the client's saved information from the map
// of currently connected clients and returns a leave message.
func (s *Server) UserLeft(ctx context.Context, in *pb.Player) (*pb.ServerMessage, error) {
	if _, ok := s.clients[in.ID]; ok {
		delete(s.clients, in.ID)
	}

	log.Printf("%v (%v) has left the game with IP: %v\n", in.Username, in.ID, in.IP)
	return &pb.ServerMessage{Message: "You have disconnected from server"}, nil
}

// UsersConnected returns the number of currently
// connected clients. (length of the map of clients)
func (s *Server) UsersConnected(ctx context.Context, in *pb.Empty) (*pb.PlayerCount, error) {
	return &pb.PlayerCount{Count: int32(len(s.clients))}, nil
}
