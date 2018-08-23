package player

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	goasperite "github.com/SolarLune/GoAseprite"
	"github.com/damienfamed75/engo-xaro/src/communication"
)

// Player is the object that represents the client
type Player struct {
	Scale       float32
	Spritesheet *common.Spritesheet
	Ase         goasperite.AsepriteFile
	ShootSpeed  float32
	MoveSpeed   float32
	Velocity    engo.Point
	Username    string

	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent

	Updater engo.Updater

	diff     engo.Point
	shooting bool
	Client   *communication.Client
}

var (
	loadedSprite = common.LoadedSprite

	imagePath = "/graphics/player.png"
)
