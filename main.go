package main

import (
	"engo.io/engo"
	"github.com/damienfamed75/engo-xaro/src/scenes"
	"github.com/damienfamed75/engo-xaro/src/system"
)

func main() {
	_, config := system.LoadViperConfig("/config/")

	opts := engo.RunOptions{
		Title:         "Xaro",
		ScaleOnResize: true,
		MSAA:          0,
		Width:         config.Window.Width,
		Height:        config.Window.Height,
		VSync:         config.Window.VSync,
		Fullscreen:    config.Window.FullScreen,
	}

	//// Run the Scene /////
	engo.Run(opts, &scenes.Nexus{})
}
