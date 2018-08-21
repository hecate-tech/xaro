package player

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Player is the object that represents the client
type Player struct {
	Scale float32

	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent

	Updater engo.Updater
}

var (
	loadedSprite = common.LoadedSprite
	logFatalln   = log.Fatalln
)

// New is used to create a new player
func New(w *ecs.World) *Player {
	texture, err := loadedSprite("graphics/player.png")
	if err != nil {
		logFatalln(err)
	}

	p := &Player{BasicEntity: ecs.NewBasic()}
	p.Scale = 4

	p.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: p.Scale, Y: p.Scale},
	}
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    texture.Width() * p.RenderComponent.Scale.X,
		Height:   texture.Height() * p.RenderComponent.Scale.Y,
	}

	w.AddSystem(p)

	return p
}

// Update gets called every frame
func (p *Player) Update(dt float32) {
	p.Position.X++
}

// Remove deletes the player and player systems
func (p *Player) Remove(ecs.BasicEntity) {

}
