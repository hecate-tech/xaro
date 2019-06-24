package resource

import (
	// "github.com/hecatetech/xaro/general"

	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Goal struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
}

func setupGoal() *Goal {
	return &Goal{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Color:    color.RGBA{0, 255, 25, 255},
			Drawable: common.Circle{},
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{X: 0, Y: 0},
			Width:    50,
			Height:   50,
		},
	}
}

func NewGoal(pos engo.Point) *Goal {
	g := setupGoal()

	g.Position.X = pos.X
	g.Position.Y = pos.Y

	return g
}
