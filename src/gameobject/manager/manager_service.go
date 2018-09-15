package manager

import (
	"context"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/damienfamed75/engo-xaro/src/gameobject/player"
	"github.com/damienfamed75/engo-xaro/src/report"
)

// New returns a new manager
func New(w *ecs.World) *Manager {
	report.Status("Loading Player")

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
	report.Error("error sending player data to server:", err)

	m.updateConnectedPlayers(sPlayers, dt)

	if engo.Input.Button("quit").Down() {
		m.TerminateConnection()
	}
}

// Remove deletes all the players and main player
func (m *Manager) Remove(ecs.BasicEntity) {

}
