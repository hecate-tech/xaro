package player

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/gameobject/constant"
)

// Player is the main game object that represents
// a Xaro client. This is a playable struct...
type Player struct {
	// constant.Basic contains the essential
	// parts of a Xaro game object such as
	// the scale, spritesheet, and aseprite.
	constant.Basic
	// constant.Playable contains the variables
	// needed to have a gameobject that contains
	// user control such as velocity and movespeed.
	constant.Playable

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
