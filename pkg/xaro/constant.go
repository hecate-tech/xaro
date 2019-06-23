package xaro

//go:generate stringer -type=SceneNumber

// SceneNumber is used to index the scenes in the map in an efficient manner
// without using strings. This type also has a stringer type associated with it.
// This is located in scenenumber_stringer.go
type SceneNumber int8

const (
	// Game represents the main scene of Xaro that the player will use.
	Game SceneNumber = iota
	// MainMenu is the first scene the player will see when they start Xaro.
	MainMenu
)
