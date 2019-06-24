/*
Package player stores the playable character that can be used in all game scenes.
*/
package player

import (
	"image/color"

	"github.com/hecatetech/xaro/general"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

func setupPlayer() *Player {
	return &Player{
		BasicEntity: ecs.NewBasic(),
		ShootSpeed:  shootSpeed,
		MoveSpeed:   moveSpeed,
		Basic: &general.Basic{
			EntityScale: general.DrawScale,
		},
	}
}

// NewPlayer creates a new blank player to be passed around the game.
func NewPlayer() (*Player, error) {
	p := setupPlayer()

	p.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{150, 150, 0, 255},
	}

	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    20,
		Height:   20,
	}

	p.CollisionComponent = common.CollisionComponent{
		Main: general.Solid,
	}

	// Set the 3 dimensional drawing index.
	// p.SetZIndex(general.Player)

	return p, nil
}

func (p *Player) Prepare(pos *engo.Point) {
	p.SetPosition(pos)
	// other steps
}

// SetPosition takes a reference to a point and sets the position to it.
func (p *Player) SetPosition(pos *engo.Point) {
	p.Position.X = pos.X
	p.Position.Y = pos.Y
}
