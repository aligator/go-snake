package main

import (
	"github.com/EngoEngine/engo"
	"snake"
	"snake/scenes"
)

func main() {
	width := 37
	height := 25
	tileSize := 16

	opts := engo.RunOptions{
		Title:  "Snake",
		Width:  width * tileSize,
		Height: height * tileSize,
	}

	engo.Run(opts, &scenes.Game{
		Properties: snake.Properties{
			Width:    width,
			Height:   height,
			TileSize: tileSize,
		},
	})
}
