package player

import (
	"context"
	"path/filepath"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	goasperite "github.com/SolarLune/GoAseprite"
	comm "github.com/damienfamed75/engo-xaro/src/common"
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

// New is used to create a new player
func New(w *ecs.World) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		ShootSpeed:  0.5,
		MoveSpeed:   120.0,
		Scale:       4.0,
		Username:    "Damien",
	}

	jsonPath, err := filepath.Abs("assets/graphics/player.json")
	comm.ErrorCheck(err)
	err = engo.Files.Load(imagePath)
	comm.ErrorCheck(err)

	//// Setting Player Vars /////
	p.setupConnection("98.144.164.154:8081")
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	p.Position.Add(*p.Velocity.MultiplyScalar(dt))
	p.Client.Player.Position.X, p.Client.Player.Position.Y = p.Position.X, p.Position.Y
	p.Client.Conn.SendPlayerData(ctx, p.Client.Player)
	p.Drawable = p.Spritesheet.Drawable(int(p.Ase.CurrentFrame))
}

// Remove deletes the player and player systems
func (p *Player) Remove(ecs.BasicEntity) {
	p.Drawable.Close()
}
