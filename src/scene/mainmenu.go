package scene

import "engo.io/engo"

// MainMenu is the starting scene
// and way to access the game scene.
type MainMenu struct{}

// Preload loads in all essential
// graphics and assets for the scene.
func (*MainMenu) Preload() {
	
}

// Setup creates and instantiates all
// the gameobjects in the world.
func (*MainMenu) Setup(u engo.Updater) {

}

// Type returns the name of
// the given scene.
func (*MainMenu) Type() string {
	return "Main Menu"
}
