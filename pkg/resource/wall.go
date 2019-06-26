package resource

import (
	"image/color"

	"github.com/hecatetech/xaro/general"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Wall struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

func setupWall() *Wall {
	return &Wall{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Color:    color.RGBA{160, 160, 160, 255},
			Drawable: common.Rectangle{},
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{X: 0, Y: 0},
			Width:    10,
			Height:   10,
		},
		CollisionComponent: common.CollisionComponent{
			Group: general.Solid,
		},
	}
}

func NewWall(pos engo.Point, size engo.Point) *Wall {
	w := setupWall()

	w.Position.X = pos.X
	w.Position.Y = pos.Y

	w.Width = size.X
	w.Height = size.Y
	// w.Extra = size.Y

	return w
}
