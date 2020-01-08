package main

import (
	"github.com/EngoEngine/engo"
	"snake/scenes"
)

func main() {
	opts := engo.RunOptions{
		Title:  "Snake",
		Width:  600,
		Height: 400,
	}

	engo.Run(opts, &scenes.Game{})
}
