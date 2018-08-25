package manager

import (
	"context"
	"log"
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
func (m *Manager) Update(float32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	m.Client.Player.Position.X, m.Client.Player.Position.Y = m.Player.Position.X, m.Player.Position.Y
	m.Client.Conn.SendPlayerData(ctx, m.Client.GetPlayer())

	if engo.Input.Button("quit").Down() {
		m.TerminateConnection()
	}
}

// EstablishConnection connects to selected server
func (m *Manager) EstablishConnection() {
	_, config := system.LoadViperConfig("/config/")

	conn, err := grpc.Dial(config.Connection.GetAddress())
	common.ErrorCheck(err)

	// Creating new client for server
	c := pb.NewXaroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// Set the Player's Client object to a new Client
	m.Client = communication.NewClient(c, m.Player.Username)

	// Send the server a message that I've joined and reeive a new ID
	r, _ := c.UserJoined(ctx, m.Client.Player)
	m.Client.Player.ID = r.Newid // Set new ID to current client.

	log.Println(r.Message)
}

// TerminateConnection disconnects player from server
func (m *Manager) TerminateConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	m.Client.Conn.UserLeft(ctx, m.Client.Player)
	engo.Exit()
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
