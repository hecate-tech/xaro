package tilemap

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

// Tile will contain data required
// to draw a full tilemap
type Tile struct {
	// ecs.BasicEntity to save the unique ID
	// of the specific tile.
	ecs.BasicEntity
	// common.RenderComponent to store the
	// drawable for the tile object.
	common.RenderComponent
	// common.SpaceComponent to save the
	// tile's position and rotation.
	common.SpaceComponent
}

// SCALE is the default scale for each
// drawn tile within each tilemap.
const SCALE = 4
