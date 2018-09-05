package manager

import (
	"context"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/damienfamed75/engo-xaro/src/common"
	"github.com/damienfamed75/engo-xaro/src/communication"
	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"
	pb "github.com/damienfamed75/engo-xaro/src/proto"
	"github.com/damienfamed75/engo-xaro/src/system"
	"google.golang.org/grpc"
)

// Manager holds all the information about your connection
// along with players and enemies in order to render them.
type Manager struct {
	Player *player.Player
	Client *communication.Client

	world *ecs.World

	ServerPlayers map[uint32]*player.Player
}

// New returns a new manager
func New(w *ecs.World) *Manager {
	common.StatusPrint("Loading Player")

	m := &Manager{
		Player:        player.New(w),
		ServerPlayers: make(map[uint32]*player.Player),
		world:         w,
	}

	m.EstablishConnection()
	w.AddSystem(m)

	return m
}

// Update gets called every frame.
func (m *Manager) Update(dt float32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Update sent data to server.
	m.Client.UpdatePosition(m.Player.Position.X, m.Player.Position.Y)
	m.Client.Player.AnimName = m.Player.Ase.CurrentAnimation.Name

	// Send the information to server and receive other players.
	sPlayers, err := m.Client.Conn.SendPlayerData(ctx, m.Client.GetPlayer())
	common.ErrorCheck("error sending player data to server:", err)

	m.updateConnectedPlayers(sPlayers, dt)

	if engo.Input.Button("quit").Down() {
		m.TerminateConnection()
	}
}

// EstablishConnection connects to selected server
func (m *Manager) EstablishConnection() {
	common.StatusPrint("Establishing connection")
	_, config := system.LoadViperConfig()

	conn, err := grpc.Dial(config.Connection.GetAddress(), grpc.WithInsecure())
	common.ErrorCheck("error dialing with grpc:", err)

	// Creating new client for server
	c := pb.NewXaroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// Set the Player's Client object to a new Client
	m.Client = communication.NewClient(c, m.Player.Username)

	// Send the server a message that I've joined and reeive a new ID
	r, err := c.UserJoined(ctx, m.Client.Player)
	common.ErrorCheck("error joining server:", err)

	m.Client.Player.ID = r.Newid // Set new ID to current client.

	common.SuccessPrint(r.Message)
}

// TerminateConnection disconnects player from server
func (m *Manager) TerminateConnection() {
	common.StatusPrint("Terminating connection")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send server that player has disconnected.
	_, err := m.Client.Conn.UserLeft(ctx, m.Client.Player)
	common.ErrorCheck("error leaving server:", err)

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

// Remove deletes all the players and main player
func (m *Manager) Remove(ecs.BasicEntity) {

}

func (m *Manager) retrievePlayers() {

}
