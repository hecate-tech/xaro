package scenes

import (
	"image/color"

	"github.com/damienfamed75/engo-xaro/src/gameobjects/manager"
	"github.com/damienfamed75/engo-xaro/src/gameobjects/player"
	"github.com/damienfamed75/engo-xaro/src/system"

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
)

// Preload loads in essential graphics and assets
func (*Nexus) Preload() {
	system.Init()
}

// Setup creates and instantiates everything in the world
func (*Nexus) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	//// Making Scene /////
	setBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})

	//// Player Setup /////
	m := manager.New(w)
	m.Player.Position = spawnPosition

	//// System Setup ////
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.RenderComponent, &m.Player.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.MouseComponent, &m.Player.SpaceComponent, &m.Player.RenderComponent)
		}
	}

}

// Type returns the type of the world
func (*Nexus) Type() string {
	return "GameWorld"
}
