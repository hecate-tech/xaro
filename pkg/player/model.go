package player

import (
	"github.com/hecatetech/xaro/general"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Player is the main game object that represents
// a Xaro client. This is a playable struct...
type Player struct {
	// general.Basic contains the essential
	// parts of a Xaro game object such as
	// the scale, spritesheet, and aseprite.
	*general.Basic
	// ShootSpeed controls the speed at which
	// the player can use their weapon.
	ShootSpeed float32
	// MoveSpeed is the variable that affects
	// the movement speed the player can move at.
	MoveSpeed float32
	// IsShooting is a current state booleon
	// to check if the player is actioning.
	IsShooting bool
	// Velocity stores the X, and Y variables
	// At which the player is moving towards.
	Velocity engo.Point
	// ecs.BasicEntity is used for indentifying
	// this specific entity in the game. This
	// is used for deletion as well.
	ecs.BasicEntity
	// common.RenderComponent is the component
	// that allows this struct to render into the game.
	common.RenderComponent
	// common.SpaceComponent contains essential spacial
	// variables such as position, size, and rotation.
	common.SpaceComponent
	// common.MouseComponent allows this game object
	// to read from the mouse input using engo's systems.
	common.MouseComponent
}
