package xaro

import (
	general "github.com/hecatetech/xaro/general"
	"github.com/hecatetech/xaro/pkg/logging"
	p "github.com/hecatetech/xaro/pkg/player"
	"github.com/hecatetech/xaro/pkg/scene"

	"github.com/EngoEngine/engo"
)

// Manager is used to manage the whole Xaro game including all its scenes
// its configurations, and the player itself.
type Manager struct {
	player *p.Player
	cfg    *general.Configuration
	scenes map[SceneNumber]engo.Scene
	logger logging.Logger
}

func setupManager() *Manager {
	return &Manager{
		cfg:    general.LoadViperConfig(),
		scenes: make(map[SceneNumber]engo.Scene),
		logger: logging.NewNilLogger(),
	}
}

// NewManager returns a game manager with all Xaro scenes and player in it.
func NewManager(debug bool) (*Manager, error) {
	m := setupManager()

	// Perhaps this could be done more elegantly?
	p, err := p.NewPlayer()
	if err != nil {
		return nil, err
	}
	m.player = p

	// This should be a config function instead an if to prevent branching.
	if debug {
		m.logger = logging.NewDebugLogger()
	}

	// This is going to get sloppy if there are more than three assignments.
	m.scenes[Game] = scene.NewGame(p, Game.String(), m.logger)
	m.scenes[MainMenu] = &scene.MainMenu{}

	return m, nil
}

// GetConfig returns a reference to the configuration loaded by Viper.
func (m *Manager) GetConfig() *general.Configuration {
	return m.cfg
}

// GetScene returns the scene from the given SceneNumber key.
func (m *Manager) GetScene(n SceneNumber) engo.Scene {
	return m.scenes[n]
}
