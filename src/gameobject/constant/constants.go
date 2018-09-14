package constant

// Objects at the top are rendered before
// the objects placed at the bottom of the
// list of constants.

const (
	// BACKGROUND needs to be drawn
	// beneath everything in the game
	BACKGROUND float32 = iota - 1
	// OBJECTS are small items and
	// around the world objects like
	// chests or bags of loot.
	OBJECTS
	// SERVERPLAYER needs to be
	// displayed over the objects, but
	// still underneath the player.
	SERVERPLAYER
	// PLAYER needs to be drawn
	// above everything in the game
	PLAYER
)
