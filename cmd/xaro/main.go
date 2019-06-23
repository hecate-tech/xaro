package main

import (
	"flag"
	"log"

	"github.com/hecatetech/xaro/pkg/xaro"

	"github.com/EngoEngine/engo"
)

var (
	debug bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()

	// Init the xaro manager object
	m, err := xaro.NewManager(debug)
	if err != nil {
		log.Fatalln(err)
	}

	config := m.GetConfig()

	// Load the configurations using the Xaro object.
	opts := engo.RunOptions{
		Title:         "Xaro",
		ScaleOnResize: true,
		MSAA:          0,
		Width:         config.Window.Width,
		Height:        config.Window.Height,
		VSync:         config.Window.VSync,
		Fullscreen:    config.Window.FullScreen,
	}

	// Run the Xaro object's Game scene.
	engo.Run(opts, m.GetScene(xaro.Game))
}
