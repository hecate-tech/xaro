package general

const (
	// Background is the background image of the stage.
	// In a top down perspective this would be the floor.
	Background float32 = iota - 1
	// Scene contains all scene objects for decoration.
	// This would include a collidable barrel for instance.
	// Or the walls of the scene.
	Scene
	// Enemies stores the enemy characters in the scene
	// that the player can interact with directly.
	Enemies
	// Player contains the playable character itself.
	// This is a high priority render.
	Player
	// Particles holds all the special effect particles.
	// This could include explosions and bullets.
	Particles
	// MenuBackground is for the background of a menu. This could be
	// a flat image for the main menu or a blurry dark half opacity for options.
	MenuBackground
	// MenuGraphics is for the buttons and misc. graphics to go over the menu.
	MenuGraphics
)

// DrawScale is the default multiplier when drawing
// tiles and objects in the game.
const DrawScale float32 = 4.0
