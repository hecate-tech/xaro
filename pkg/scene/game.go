package scene

import (
	"image/color"

	"github.com/hecatetech/xaro/general"
	"github.com/hecatetech/xaro/pkg/logging"
	p "github.com/hecatetech/xaro/pkg/player"
	"github.com/hecatetech/xaro/pkg/resource"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Game is the main game scene for Xaro.
type Game struct {
	// Given by the manager
	player    *p.Player
	sceneName string
	logger    logging.Logger

	// Made here in the scene package.
	spawnPos *engo.Point
	bkgPath  string
}

func setupGame() *Game {
	return &Game{
		spawnPos: &engo.Point{X: 100, Y: 100},
		bkgPath:  "/maps/Nexus.tmx",
	}
}

// NewGame stores the player and scene name and sets up the rest of the scene
// fields in the Game struct.
func NewGame(player *p.Player, sceneName string, logger logging.Logger) *Game {
	s := setupGame()

	s.player = player
	s.sceneName = sceneName
	s.logger = logger

	return s
}

// Preload loads in all essential
// graphics and assets for the scene.
func (g *Game) Preload() {
	g.logger.Debug("Loading Scene " + g.Type())
	// Apply settings from into the engo engine.
	general.Apply()
	// g.player.SetPosition(g.spawnPos)
	g.player.Prepare(g.spawnPos)
}

// Setup creates and instantiates all
// the gameobjects in the world.
func (g *Game) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(g.player)

	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.CollisionSystem{Solids: 1})

	wall := resource.NewWall(engo.Point{X: 0, Y: 0},
		engo.Point{X: 50, Y: engo.WindowHeight()})

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&g.player.BasicEntity, &g.player.RenderComponent, &g.player.SpaceComponent)
			sys.Add(&wall.BasicEntity, &wall.RenderComponent, &wall.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&g.player.BasicEntity, &g.player.CollisionComponent, &g.player.SpaceComponent)
			sys.Add(&wall.BasicEntity, &wall.CollisionComponent, &wall.SpaceComponent)
		}
	}

	g.logger.Debug("Finished Loading Scene " + g.Type())
}

// Type returns the name of
// the given scene.
func (g *Game) Type() string {
	return g.sceneName
}
