package resource

import (
	"fmt"
	"image/color"

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
			Color:    color.RGBA{255, 0, 0, 255},
			Drawable: common.Rectangle{},
			Scale:    engo.Point{X: 1, Y: 1},
		},
		CollisionComponent: common.CollisionComponent{
			Group: 1,
		},
	}
}

func NewWall(pos engo.Point, size engo.Point) *Wall {
	w := setupWall()

	w.SpaceComponent = common.SpaceComponent{
		Position: pos,
		Width:    size.X,
		Height:   size.Y,
	}

	fmt.Printf("extra[%v][%v]\n", w.Extra.X, w.Extra.Y)
	// w.Extra = size.Y

	return w
}
