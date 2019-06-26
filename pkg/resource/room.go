package resource

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Room struct {
	// walls is the exterior of the room
	walls [4]*Wall
	// obstacles is a temporary object for the sake of the maze demo.
	obstacles []*Wall
	goal      *Goal

	// in the future
	// exits will be the room's doors to exit into another room.
	// exits []*Exit
}

// setupRoom sets up the base four walls of the room.
func setupRoom() *Room {
	return &Room{
		walls: [4]*Wall{
			NewWall(engo.Point{X: 0, Y: 0}, engo.Point{X: 50, Y: engo.GameHeight()}),
			NewWall(engo.Point{X: engo.GameWidth() - 50, Y: 0}, engo.Point{X: 50, Y: engo.GameHeight()}),
			NewWall(engo.Point{X: 0, Y: 0}, engo.Point{X: engo.GameWidth(), Y: 50}),
			NewWall(engo.Point{X: 0, Y: engo.GameHeight() - 50}, engo.Point{X: engo.GameWidth(), Y: 50}),
		},
	}
}

// NewRoom is a debug function for generating a room with obstacles inside of it.
// This function will not be the same in the final game, but for the sake
// of testing the ability to switch levels this function will return a choice
// of three different room types with their own associated obstacles and
// goal position for the player to progress through the next room.
func NewRoom(w *ecs.World, num int) *Room {
	r := setupRoom()

	for _, wall := range r.walls {
		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&wall.BasicEntity, &wall.RenderComponent, &wall.SpaceComponent)
			case *common.CollisionSystem:
				sys.Add(&wall.BasicEntity, &wall.CollisionComponent, &wall.SpaceComponent)
			}
		}
	}

	switch num {
	case 1:
		r.obstacles = []*Wall{
			NewWall(engo.Point{X: 100, Y: 50}, engo.Point{X: 25, Y: 100}),
			NewWall(engo.Point{X: 100, Y: 150}, engo.Point{X: 100, Y: 25}),
			NewWall(engo.Point{X: 200, Y: 150}, engo.Point{X: 25, Y: 100}),
			NewWall(engo.Point{X: 300, Y: 50}, engo.Point{X: 25, Y: 300}),
			NewWall(engo.Point{X: 400, Y: 150}, engo.Point{X: 25, Y: 150}),
			NewWall(engo.Point{X: 400, Y: 150}, engo.Point{X: 100, Y: 25}),
			NewWall(engo.Point{X: 500, Y: 50}, engo.Point{X: 25, Y: 125}),

			NewWall(engo.Point{X: 625, Y: 150}, engo.Point{X: 125, Y: 25}),
			NewWall(engo.Point{X: 625, Y: 175}, engo.Point{X: 25, Y: 125}),
			NewWall(engo.Point{X: 400, Y: 275}, engo.Point{X: 250, Y: 25}),

			NewWall(engo.Point{X: 50, Y: 250}, engo.Point{X: 125, Y: 25}),
			NewWall(engo.Point{X: 100, Y: 400}, engo.Point{X: 75, Y: 25}),
			NewWall(engo.Point{X: 175, Y: 325}, engo.Point{X: 25, Y: 150}),
			NewWall(engo.Point{X: 200, Y: 450}, engo.Point{X: 200, Y: 25}),
			NewWall(engo.Point{X: 400, Y: 300}, engo.Point{X: 25, Y: 175}),

			NewWall(engo.Point{X: 500, Y: 375}, engo.Point{X: 25, Y: 100}),
			NewWall(engo.Point{X: 650, Y: 375}, engo.Point{X: 25, Y: 100}),
			NewWall(engo.Point{X: 575, Y: 450}, engo.Point{X: 25, Y: 100}),
			NewWall(engo.Point{X: 575, Y: 450}, engo.Point{X: 100, Y: 25}),
			NewWall(engo.Point{X: 525, Y: 375}, engo.Point{X: 125, Y: 25}),
		}
		r.goal = NewGoal(engo.Point{X: 614.5, Y: 487.5})
	case 2:
	default:
	}

	for _, ob := range r.obstacles {
		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&ob.BasicEntity, &ob.RenderComponent, &ob.SpaceComponent)
			case *common.CollisionSystem:
				sys.Add(&ob.BasicEntity, &ob.CollisionComponent, &ob.SpaceComponent)
			}
		}
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&r.goal.BasicEntity, &r.goal.RenderComponent, &r.goal.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&r.goal.BasicEntity, &r.goal.CollisionComponent, &r.goal.SpaceComponent)
		}
	}

	return r
}
