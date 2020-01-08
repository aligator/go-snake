package scenes

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
	"snake/entities"
	"snake/systems"
)

type Game struct{}

func (g *Game) Preload() {
	entities.PreloadSnake()
}

func (g *Game) Setup(u engo.Updater) {
	world, ok := u.(*ecs.World)
	if !ok {
		panic("updater is no world")
	}
	common.SetBackground(color.Black)

	systems.SetupSnake()

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.Snake{})
}

func (g *Game) Type() string {
	return "gameScene"
}
