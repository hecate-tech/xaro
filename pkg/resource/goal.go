package resource

import (
	"fmt"
	"image/color"

	"github.com/hecatetech/xaro/general"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Goal is used for switching between rooms. For instance in the final game
// there should be four of these objects at most on each side of the room.
type Goal struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
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
		CollisionComponent: common.CollisionComponent{
			Group: general.Solid,
		},
	}
}

func NewGoal(pos engo.Point) *Goal {
	g := setupGoal()

	g.Position.X = pos.X
	g.Position.Y = pos.Y

	engo.Mailbox.Listen("CollisionMessage", func(msg engo.Message) {
		collision, isCollision := msg.(common.CollisionMessage)
		if isCollision {
			if collision.To.ID() == g.ID() {
				fmt.Println("Switch levels")
				// engo.SetSceneByName("MainMenu", false)
			}
		}
	})

	return g
}

func (g *Goal) Update(dt float32) {

}

// Remove deletes the player and player systems
func (g *Goal) Remove(ecs.BasicEntity) {
	g.Drawable.Close()
}
