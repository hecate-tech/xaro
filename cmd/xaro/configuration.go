package main

import (
	"os"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/spf13/viper"
)

// Init initializes controls for the game
// based on the configuration file.
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

// LoadViperConfig loads a configuration file and returns a new
// copy of the data via a Configuration struct and viper object.
func LoadViperConfig() (*viper.Viper, Configuration) {
	v := viper.New()
	var c Configuration

	// wd, err := os.Getwd()
	wd, _ := os.Getwd()
	// report.Error("cannot find working directory:", err)

	// Adding config paths...
	v.SetConfigName("config.development")
	v.AddConfigPath("../../../config/")
	v.AddConfigPath("$HOME/.go-xaro")
	v.AddConfigPath(wd + "/config/")
	v.AddConfigPath(".")
	v.AddConfigPath(wd)

	// err = v.ReadInConfig()
	_ = v.ReadInConfig()
	// report.Error("unable to read in config file from selected paths:", err)

	// err = v.Unmarshal(&c)
	_ = v.Unmarshal(&c)
	// report.Error("unable to unmarshal config file:", err)

	return v, c
}

// ChangeConfig takes your viper to the configuration file along with a key
// you want to change and its new value then will be saved to your config file.
func ChangeConfig(v *viper.Viper, key string, value interface{}) {
	v.Set(key, value) // Sets the viper's key and value

	// err := v.WriteConfig() // Writes the new configuration to your config file.
	_ = v.WriteConfig() // Writes the new configuration to your config file.
	// report.Error("unable to write to config:", err)
}
