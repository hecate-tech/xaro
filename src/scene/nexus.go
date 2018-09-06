package scene

import (
	"image/color"

	"github.com/damienfamed75/engo-xaro/src/gameobject/manager"
	"github.com/damienfamed75/engo-xaro/src/gameobject/player"
	"github.com/damienfamed75/engo-xaro/src/report"
	"github.com/damienfamed75/engo-xaro/src/system"
	"github.com/damienfamed75/engo-xaro/src/tilemap"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Nexus is the main collection for players
type Nexus struct{}

var (
	spawnPosition = engo.Point{X: 0, Y: 0}
	mapPath       = "/maps/Nexus.tmx"

	setBackground = common.SetBackground
	playerNew     = player.New
)

// Preload loads in essential graphics and assets
func (*Nexus) Preload() {
	engo.Files.Load(mapPath)
	system.Init()
}

// Setup creates and instantiates everything in the world
func (*Nexus) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	report.Status("Loading Nexus")

	//-- Making Scene --//
	setBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})

	tilemap.Load(w, mapPath)

	//-- Player Setup --//
	m := manager.New(w)
	m.Player.Position = spawnPosition

	//-- System Setup --//
	report.Status("Adding Systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.RenderComponent, &m.Player.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.MouseComponent, &m.Player.SpaceComponent, &m.Player.RenderComponent)
		}
	}

	report.Success("Nexus successfully loaded")
}

// Type returns the type of the world
func (*Nexus) Type() string {
	return "Nexus"
}
