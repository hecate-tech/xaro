package manager

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"

	"github.com/hecategames/xaro/src/gameobject/player"
	"github.com/hecategames/xaro/src/report"
)

// New returns a new manager
func New(w *ecs.World) *Manager {
	report.Status("Loading Player")

	m := &Manager{
		Player: player.New(w),
		world:  w,
	}

	w.AddSystem(m)

	return m
}

// Update gets called every frame.
func (m *Manager) Update(dt float32) {

	if engo.Input.Button("quit").Down() {
		// Pause menu
	}
}

// Remove deletes all the players and main player
func (m *Manager) Remove(ecs.BasicEntity) {

}
