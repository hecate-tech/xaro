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

	diff engo.Point
}

var (
	loadedSprite = common.LoadedSprite

	imagePath = "/graphics/player.png"
)

// New is used to create a new player
func New(w *ecs.World) *Player {
	p := &Player{}

	jsonPath, err := filepath.Abs("assets/graphics/player.json")
	comm.ErrorCheck(err)
	err = engo.Files.Load(imagePath)
	comm.ErrorCheck(err)

	//// Setting Player Vars /////
	p.Ase = goasperite.New(jsonPath, "player")
	p.Velocity = engo.Point{X: 0, Y: 0}
	p.BasicEntity = ecs.NewBasic()
	p.ShootSpeed = 0.5
	p.MoveSpeed = 120
	p.Scale = 4.0
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

	//// Registering Buttons /////
	//engo.Input.RegisterButton("attack", engo.KeySpace)
	engo.Input.RegisterButton("left", engo.KeyA)
	engo.Input.RegisterButton("right", engo.KeyD)
	engo.Input.RegisterButton("up", engo.KeyW)
	engo.Input.RegisterButton("down", engo.KeyS)

	w.AddSystem(p)
	p.Ase.Play("right") // Queues starting animation

	return p
}

// Update gets called every frame
func (p *Player) Update(dt float32) {
	p.Ase.Update(dt)

	p.Velocity.X, p.Velocity.Y = 0, 0

	p.updateMovement()
	p.updateAnimation()
	p.updateIdleAnimation()

	p.diff.X, p.diff.Y = p.Velocity.X*dt, p.Velocity.Y*dt

	p.Position.X += p.diff.X
	p.Position.Y += p.diff.Y

	p.Drawable = p.Spritesheet.Drawable(int(p.Ase.CurrentFrame))
}

// Remove deletes the player and player systems
func (p *Player) Remove(ecs.BasicEntity) {

}
