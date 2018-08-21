package main

import (
	"engo.io/engo"
	"github.com/damienfamed75/engo-xaro/src/scenes"
)

func main() {
	opts := engo.RunOptions{
		Title:         "Xaro",
		Width:         800,
		Height:        600,
		ScaleOnResize: true,
		VSync:         false,
		MSAA:          0,
	}
	engo.Run(opts, &scenes.Nexus{})
}
