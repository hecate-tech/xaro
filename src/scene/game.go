package scene

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/hecategames/xaro/src/gameobject/manager"
	"github.com/hecategames/xaro/src/report"
	"github.com/hecategames/xaro/src/system"
	"github.com/hecategames/xaro/src/tilemap"
)

// Game is the testing level
// for Xaro at the moment.
type Game struct{}

var (
	gameSpawn = engo.Point{X: 0, Y: 0}
	gamePath  = "/maps/Nexus.tmx"
)

// Preload loads in all essential
// graphics and assets for the scene.
func (*Game) Preload() {
	engo.Files.Load(gamePath)
	system.Init()
}

// Setup creates and instantiates all
// the gameobjects in the world.
func (*Game) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	report.Status("Loading Game")

	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})

	tilemap.Load(w, gamePath)

	m := manager.New(w)
	m.Player.Position = gameSpawn

	report.Status("Adding Systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.RenderComponent, &m.Player.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&m.Player.BasicEntity, &m.Player.MouseComponent, &m.Player.SpaceComponent, &m.Player.RenderComponent)
		}
	}

	report.Success("Game successfully loaded")
}

// Type returns the name of
// the given scene.
func (*Game) Type() string {
	return "Game"
}
