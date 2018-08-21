package scenes

import "engo.io/engo"

type Online struct{}

// Preload is used for loading in sprites and other assets used
// within the scene.
func (*Online) Preload() {

}

// Setup creates all the essential objects and functionality
// for the player to traverse the scene.
func (*Online) Setup(u engo.Updater) {

}

// Type returns the type of scene this is
func (*Online) Type() string {
	return "Online Scene"
}
