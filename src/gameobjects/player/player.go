package player

import (
	"path/filepath"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	goasperite "github.com/damienfamed75/GoAseprite"
	comm "github.com/damienfamed75/engo-xaro/src/common"
	"github.com/damienfamed75/engo-xaro/src/system"
)

var (
	loadedSprite = common.LoadedSprite

	imagePath = "/graphics/player.png"
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
	IsPlaying   bool

	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent

	diff     engo.Point
	shooting bool
}

// New is used to create a new player
func New(w *ecs.World) *Player {
	_, config := system.LoadViperConfig()

	p := &Player{
		BasicEntity: ecs.NewBasic(),
		ShootSpeed:  0.5,
		MoveSpeed:   120.0,
		Scale:       4.0,
		Username:    config.PlayerData.Username,
		IsPlaying:   true,
	}

	jsonPath, err := filepath.Abs("assets/graphics/player.json")
	comm.ErrorCheck("filepath.Abs failed to output dir:", err)
	err = engo.Files.Load(imagePath)
	comm.ErrorCheck("failed to load image in engo:", err)

	// Setting Up Player Variables...
	p.Ase = goasperite.New(jsonPath)
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
	p.Ase.Play("downidle") // Queues starting animation

	return p
}

// Update gets called every frame
func (p *Player) Update(dt float32) {
	p.Ase.Update(dt)

	if !p.IsPlaying {
		return
	}

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
