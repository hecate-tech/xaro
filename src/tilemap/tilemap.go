package tilemap

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/gameobject/constant"
	"github.com/damienfamed75/engo-xaro/src/report"
)

// Tile will contain data required
// to draw a full tilemap
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Load will take the given world and path
// To load a resource previously loaded and create
// a new tilemap from it using engo's inbuilt Tiled
// support and then set it beneath the player.
func Load(w *ecs.World, path string) {
	report.Status("Loading Tilemap")

	r, err := engo.Files.Resource(path)
	report.Error("Couldn't load resource:", err)

	tmxResource := r.(common.TMXResource)
	levelData := tmxResource.Level

	tiles := make([]*Tile, 0)
	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{X: 4, Y: 4},
				}
				tile.SetZIndex(constant.BACKGROUND) // Drawn in background

				pos := tileElement.Point
				pos.MultiplyScalar(tile.Scale.X)
				tile.SpaceComponent = common.SpaceComponent{
					Position: pos,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	// add the tiles to the RenderSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	common.CameraBounds = levelData.Bounds()
}
