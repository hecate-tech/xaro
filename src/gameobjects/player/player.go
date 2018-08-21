package player

import (
	"path/filepath"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	goasperite "github.com/SolarLune/GoAseprite"
	comm "github.com/damienfamed75/engo-xaro/src/common"
)

// Player is the object that represents the client
type Player struct {
	Scale       float32
	Spritesheet *common.Spritesheet
	Ase         goasperite.AsepriteFile
	ShootSpeed  float32
	MoveSpeed   float32
	Velocity    engo.Point

	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent

	Updater engo.Updater

	diff     engo.Point
	shooting bool
}

var (
	loadedSprite = common.LoadedSprite

	imagePath = "/graphics/player.png"
)

// New is used to create a new player
func New(w *ecs.World) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		ShootSpeed:  0.5,
		MoveSpeed:   120.0,
		Scale:       4.0,
	}

	jsonPath, err := filepath.Abs("assets/graphics/player.json")
	comm.ErrorCheck(err)
	err = engo.Files.Load(imagePath)
	comm.ErrorCheck(err)

	//// Setting Player Vars /////
	p.Ase = goasperite.New(jsonPath, "player")
	p.Spritesheet = common.NewSpritesheetFromFile(imagePath, int(p.Ase.FrameWidth), int(p.Ase.FrameHeight))
	p.RenderComponent = common.RenderComponent{
		Drawable: p.Spritesheet.Drawable(0),
		Scale:    engo.Point{X: p.Scale, Y: p.Scale},
	}
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    p.Spritesheet.Width() * p.RenderComponent.Scale.X,
		Height:   p.Spritesheet.Height() * p.RenderComponent.Scale.Y,
	}

	w.AddSystem(p)
	p.Ase.Play("right") // Queues starting animation

	return p
}

// Update gets called every frame
func (p *Player) Update(dt float32) {
	p.Ase.Update(dt)
	p.Velocity.Set(0, 0)

	p.updateMovement()
	if p.inAction() {
		p.updateAction(dt)
	} else {
		p.updateAnimation()
		p.updateIdleAnimation()
	}

	p.Position.Add(*p.Velocity.MultiplyScalar(dt))
	p.Drawable = p.Spritesheet.Drawable(int(p.Ase.CurrentFrame))
}

// Remove deletes the player and player systems
func (p *Player) Remove(ecs.BasicEntity) {
	p.Drawable.Close()
}
