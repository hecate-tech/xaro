package general

import (
	"github.com/EngoEngine/engo/common"
	"github.com/solarlune/goaseprite"
)

// Basic contains essential pieces
// to render an object in the game.
type Basic struct {
	// EntityScale controls the size at which
	// the object's entity is rendered at. (multiplier)
	EntityScale float32
	// Spritesheet contains all the drawables used by
	// the engo RenderComponent.
	Spritesheet *common.Spritesheet
	// Ase is used to keep track of what frame
	// is being played animation by animation using jsons.
	Ase goaseprite.File
}
