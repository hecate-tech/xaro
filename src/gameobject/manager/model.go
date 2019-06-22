package manager

import (
	"github.com/EngoEngine/ecs"
	"github.com/hecategames/xaro/src/gameobject/player"
)

// Manager holds all the information about your connection
// along with players and enemies in order to render them.
type Manager struct {
	// Player is the main controllable client
	// that will be uninterrupted and sent data from.
	Player *player.Player

	// Stores the world to make creating and
	// destroying of new game objects easier.
	world *ecs.World
}
