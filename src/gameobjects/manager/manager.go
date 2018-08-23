package manager

import (
	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"
)

// Manager handles all connected players and renders them
type Manager struct {
	players map[uint32]*player.Player
}

// New returns a new manager
func New() *Manager {
	return &Manager{}
}
