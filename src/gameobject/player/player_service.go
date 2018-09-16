package player

import (
	"path/filepath"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	goasperite "github.com/damienfamed75/GoAseprite"
	"github.com/hecategames/xaro/src/gameobject/constant"
	"github.com/hecategames/xaro/src/report"
	"github.com/hecategames/xaro/src/system"
)

// New is used to create a new player
func New(w *ecs.World) *Player {
	_, config := system.LoadViperConfig()

	// retreives the json file for the spritesheet.
	jsonPath, err := filepath.Abs("assets/graphics/player.json")
	report.Error("filepath.Abs failed to output dir:", err)

	// default sprite is set to archer for the time being
	err = engo.Files.Load(constant.PATHARCHER)
	report.Error("failed to load image in engo:", err)

	// Creating Player object...
	p := &Player{}

	// Setting Up Player Variables...
	p.BasicEntity = ecs.NewBasic()          // create new entity ID
	p.ShootSpeed = constant.DEFSHOOTSPEED   // default the shoot speed
	p.MoveSpeed = constant.DEFMOVESPEED     // default the move speed
	p.EntityScale = constant.DEFENTITYSCALE // default the entity scale
	p.Username = config.PlayerData.Username // load username from cfg file
	p.IsPlaying = true                      // is object playing?
	p.Ase = goasperite.New(jsonPath)        // load spritesheet's json
	p.Spritesheet = common.NewSpritesheetFromFile(constant.PATHARCHER, int(p.Ase.FrameWidth), int(p.Ase.FrameHeight))

	p.RenderComponent = common.RenderComponent{
		Drawable: p.Spritesheet.Drawable(0), // the current drawing sprite
		Scale:    engo.Point{X: p.EntityScale, Y: p.EntityScale},
	}
	p.SetZIndex(constant.PLAYER)

	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0}, // Default position for player
		Width:    p.Spritesheet.Width() * p.RenderComponent.Scale.X,
		Height:   p.Spritesheet.Height() * p.RenderComponent.Scale.Y,
	}

	w.AddSystem(p)
	p.Ase.Play("downidle") // Queues starting animation

	return p
}

// Update gets called every frame
func (p *Player) Update(dt float32) {
	if !p.IsPlaying {
		return
	}
	// keep playing the animation accordingly
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
