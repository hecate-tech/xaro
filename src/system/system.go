package system

import (
	"fmt"
	"os"

	"engo.io/engo"
	"github.com/spf13/viper"
)

// Init initializes controls and other configurations when starting the window
func Init() {
	_, config := LoadViperConfig("/config/")

	//// Registering Buttons /////
	engo.Input.RegisterButton("left", engo.Key(config.Controls.Left))
	engo.Input.RegisterButton("right", engo.Key(config.Controls.Right))
	engo.Input.RegisterButton("up", engo.Key(config.Controls.Up))
	engo.Input.RegisterButton("down", engo.Key(config.Controls.Down))
	engo.Input.RegisterButton("menu", engo.Key(config.Controls.Menu))

	engo.Input.RegisterButton("quit", engo.KeyEscape)
}

// LoadViperConfig loads up the configuration TOML file and returns a viper object
func LoadViperConfig(file string) (*viper.Viper, Configuration) {
	v := viper.New()
	wd, _ := os.Getwd()

	v.SetConfigName("config.development")
	v.AddConfigPath("$HOME/.go-xaro")
	v.AddConfigPath("$GOPATH/src/github.com/damienfamed75/engo-xaro/config/")
	v.AddConfigPath("$GOPATH/src/github.com/engo-xaro/config/")
	v.AddConfigPath("$GOPATH/src/damienfamed75/engo-xaro/config/")
	v.AddConfigPath("$GOPATH/src/engo-xaro/config/")
	v.AddConfigPath(wd + file)
	v.AddConfigPath(".")
	v.AddConfigPath(file)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return v, c
}

// ChangeConfig updates the current config file's value
func ChangeConfig(v *viper.Viper, key string, value interface{}) {
	v.Set(key, value)
	if err := v.WriteConfig(); err != nil {
		fmt.Printf("couldn't write config: %s", err)
	}
}
