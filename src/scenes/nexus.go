package scenes

import (
	"image/color"

	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Nexus is the main collection for players
type Nexus struct{}

var (
	spawnPosition = engo.Point{X: 0, Y: 0}

	setBackground = common.SetBackground
	playerNew     = player.New

	//p = &player.Player{}
)

// Preload loads in essential graphics and assets
func (*Nexus) Preload() {
	//engo.Files.Load("graphics/player.png")

}

// Setup creates and instantiates everything in the world
func (*Nexus) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	//// Making Scene /////
	setBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})

	//// Player Setup /////
	p := playerNew(w)
	p.Position = spawnPosition

	//// System Setup ////
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&p.BasicEntity, &p.MouseComponent, &p.SpaceComponent, &p.RenderComponent)
		}
	}

}

// Type returns the type of the world
func (*Nexus) Type() string {
	return "GameWorld"
}
