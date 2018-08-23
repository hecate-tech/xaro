package system

type connection struct {
	IP   string
	Port string
}

func (c *connection) GetAddress() string {
	return c.IP + ":" + c.Port
}

type window struct {
	Width      int
	Height     int
	VSync      bool
	FullScreen bool
}

type settings struct {
	SoundVolume float32
	MusicVolume float32
}

type controls struct {
	Left  int
	Right int
	Up    int
	Down  int
	Menu  int
}

type playerdata struct {
	Username string
}

// Configuration is the global object used to hold game settings
type Configuration struct {
	Window     window     `mapstructure:"window"`
	Connection connection `mapstructure:"connection"`
	Settings   settings   `mapstructure:"settings"`
	Controls   controls   `mapstructure:"controls"`
	PlayerData playerdata `mapstructure:"playerdata"`
}
