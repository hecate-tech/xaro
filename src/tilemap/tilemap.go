package tilemap

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/hecategames/xaro/src/gameobject/constant"
	"github.com/hecategames/xaro/src/report"
)

// Load will take the given world and path
// To load a resource previously loaded and create
// a new tilemap from it using engo's inbuilt Tiled
// support and then set it beneath the player.
func Load(w *ecs.World, path string) {
	report.Status("Loading Tilemap")

	r, err := engo.Files.Resource(path) // Loads a resource from the tilemap's path.
	report.Error("Couldn't load resource:", err)

	tmxResource := r.(common.TMXResource) // type assert the Resource into a TMXResource.
	levelData := tmxResource.Level

	tiles := make([]*Tile, 0)
	for _, tileLayer := range levelData.TileLayers { // For every layer in the loaded tilemap.
		for _, tileElement := range tileLayer.Tiles { // For every tile within the layer.
			if tileElement.Image != nil { // If the image for the tile isn't nil then add it.
				tile := &Tile{BasicEntity: ecs.NewBasic()} // Create a new tile object with a unique ID.

				tile.RenderComponent = common.RenderComponent{ // Render Component with the drawable object and scale.
					Drawable: tileElement,
					Scale:    engo.Point{X: SCALE, Y: SCALE},
				}
				tile.SetZIndex(constant.BACKGROUND) // Set the drawing layer to the background.

				pos := tileElement.Point         // Set the position to the default position
				pos.MultiplyScalar(tile.Scale.X) // Multiply in the Scale for position correction.
				tile.SpaceComponent = common.SpaceComponent{
					Position: pos,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}

	// add the loaded tilemap to the world's systems.
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	common.CameraBounds = levelData.Bounds() // Set the camera's boundaries to match tilemap.
}
