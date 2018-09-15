package manager

import (
	"context"
	"time"

	"engo.io/engo"
	"github.com/damienfamed75/engo-xaro/src/communication"
	"github.com/damienfamed75/engo-xaro/src/gameobject/player"
	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/report"
	"github.com/damienfamed75/engo-xaro/src/system"
	"google.golang.org/grpc"
)

// EstablishConnection connects to selected server
func (m *Manager) EstablishConnection() {
	report.Status("Establishing connection")
	_, config := system.LoadViperConfig()

	conn, err := grpc.Dial(config.Connection.GetAddress(), grpc.WithInsecure())
	report.Error("error dialing with grpc:", err)

	// Creating new client for server
	c := pb.NewXaroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// Set the Player's Client object to a new Client
	m.Client = communication.NewClient(c, m.Player.Username)

	// Send the server a message that I've joined and receive a new ID
	r, err := c.UserJoined(ctx, m.Client.Player)
	report.Error("error joining server:", err)

	m.Client.Player.ID = r.Newid // Set new ID to current client.

	report.Success(r.Message)
}

// TerminateConnection disconnects player from server
func (m *Manager) TerminateConnection() {
	report.Status("Terminating connection")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send server that player has disconnected.
	_, err := m.Client.Conn.UserLeft(ctx, m.Client.Player)
	report.Error("error leaving server:", err)

	// Delete all server players.
	if len(m.ServerPlayers) > 0 {
		for i := range m.ServerPlayers {
			delete(m.ServerPlayers, i)
		}
	}

	engo.Exit() // Currently exits program once disconnected
}

// GetPlayers returns all recorded clients on the server.
func (m *Manager) GetPlayers() map[uint32]*player.Player {
	return m.ServerPlayers
}

func (m *Manager) retrievePlayers() {

}
