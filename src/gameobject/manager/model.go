package manager

import (
	"engo.io/ecs"
	"github.com/damienfamed75/engo-xaro/src/communication"
	"github.com/damienfamed75/engo-xaro/src/gameobject/player"
)

// Manager holds all the information about your connection
// along with players and enemies in order to render them.
type Manager struct {
	// Player is the main controllable client
	// that will be uninterrupted and sent data from.
	Player *player.Player
	// Client is the Server's connection information
	// that is utilized to send player data.
	Client *communication.Client

	// Stores the world to make creating and
	// destroying of new game objects easier.
	world *ecs.World

	// ServerPlayers is the map of currently
	// connected players to the server with the
	// index of the player's unique ID.
	ServerPlayers map[uint32]*player.Player
}
