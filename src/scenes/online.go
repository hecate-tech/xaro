package scenes

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/system"
)

// Online is a testing scene for online play
type Online struct{}

// Preload is used for loading in sprites and other assets used
// within the scene.
func (*Online) Preload() {
	system.Init()
}

// Setup creates all the essential objects and functionality
// for the player to traverse the scene.
func (*Online) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	setBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})

	p := playerNew(w)
	p.Position = spawnPosition

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&p.BasicEntity, &p.MouseComponent, &p.SpaceComponent, &p.RenderComponent)
		}
	}
}

// Type returns the type of scene this is
func (*Online) Type() string {
	return "Online Scene"
}
