package scenes

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"
)

// Nexus is the main collection for players
type Nexus struct{}

var (
	spawnPosition = engo.Point{X: 0, Y: 0}

	setBackground = common.SetBackground
	playerNew     = player.New
)

// Preload loads in essential graphics and assets
func (*Nexus) Preload() {
	engo.Files.Load("graphics/player.png")
}

// Setup creates and instantiates everything in the world
func (*Nexus) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	//// Button Setup /////
	engo.Input.RegisterButton("Attack", engo.KeySpace)

	//// Making Scene /////
	setBackground(color.White)
	w.AddSystem(&common.RenderSystem{})

	//// Player Setup /////
	player := playerNew(w)
	player.Position = spawnPosition

	//// System Setup ////
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}

}

// Type returns the type of the world
func (*Nexus) Type() string {
	return "GameWorld"
}
