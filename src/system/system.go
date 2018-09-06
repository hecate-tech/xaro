package system

import (
	"os"

	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/damienfamed75/engo-xaro/src/report"
	"github.com/spf13/viper"
)

// Init initializes controls and other configurations when starting the window
func Init() {
	_, config := LoadViperConfig()

	// Setting Volume Settings...
	common.SetMasterVolume(config.Settings.SoundVolume)

	// Registering Buttons...
	engo.Input.RegisterButton("left", engo.Key(config.Controls.Left))
	engo.Input.RegisterButton("right", engo.Key(config.Controls.Right))
	engo.Input.RegisterButton("up", engo.Key(config.Controls.Up))
	engo.Input.RegisterButton("down", engo.Key(config.Controls.Down))
	engo.Input.RegisterButton("menu", engo.Key(config.Controls.Menu))
	engo.Input.RegisterButton("quit", engo.KeyEscape)
}

// LoadViperConfig loads up the configuration TOML file and returns a viper object
func LoadViperConfig() (*viper.Viper, Configuration) {
	v := viper.New()
	var c Configuration

	wd, err := os.Getwd()
	report.Error("cannot find working directory:", err)

	// Adding config paths...
	v.SetConfigName("config.development")
	v.AddConfigPath("../../../config/")
	v.AddConfigPath("$HOME/.go-xaro")
	v.AddConfigPath(wd + "/config/")
	v.AddConfigPath(".")
	v.AddConfigPath(wd)

	err = v.ReadInConfig()
	report.Error("unable to read in config file from selected paths:", err)

	err = v.Unmarshal(&c)
	report.Error("unable to unmarshal config file:", err)

	return v, c
}

// ChangeConfig updates the current config file's value
func ChangeConfig(v *viper.Viper, key string, value interface{}) {
	v.Set(key, value)

	err := v.WriteConfig()
	report.Error("unable to write to config:", err)
}
