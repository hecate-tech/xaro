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
	Ase *goaseprite.File
}

type window struct {
	Width      int
	Height     int
	VSync      bool
	FullScreen bool
}

type settings struct {
	SoundVolume float64
	MusicVolume float64
}

type controls struct {
	Left  int
	Right int
	Up    int
	Down  int
	Menu  int
}

// Configuration is the global object used to hold game settings
type Configuration struct {
	Window   window   `mapstructure:"window"`
	Settings settings `mapstructure:"settings"`
	Controls controls `mapstructure:"controls"`
}
