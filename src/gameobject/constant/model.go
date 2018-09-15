package constant

import (
	"engo.io/engo"
	"engo.io/engo/common"
	goaseprite "github.com/damienfamed75/GoAseprite"
)

// Playable is used to import essential
// variables for a playable entity.
type Playable struct {
	// ShootSpeed controls the speed at which
	// the player can use their weapon.
	ShootSpeed float32
	// MoveSpeed is the variable that affects
	// the movement speed the player can move at.
	MoveSpeed float32
	// Username is the player's alias sent
	// to the server and other clients.
	Username string
	// IsPlaying is used to check if this is
	// an NPC or actual player object.
	IsPlaying bool
	// IsShooting is a current state booleon
	// to check if the player is actioning.
	IsShooting bool
	// Velocity stores the X, and Y variables
	// At which the player is moving towards.
	Velocity engo.Point
}

// Basic contains essential pieces
// to render an object in the game.
type Basic struct {
	// EntityScale controls the size at which
	// the object's entity is rendered at. (multiplier)
	EntityScale float32
	// Spritesheet contains all the drawables used by
	// the engo RenderComponent.
	Spritesheet *common.Spritesheet
	// Ase is used to keep track of what frame
	// is being played animation by animation using jsons.
	Ase goaseprite.AsepriteFile
}
