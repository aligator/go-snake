package scenes

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
	"snake"
	"snake/entities"
	"snake/systems"
)

type Game struct {
	snake.Properties
}

func (g *Game) Preload() {
	entities.PreloadSnake()
	entities.PreloadCooky()
	entities.PreloadText()
}

func (g *Game) Setup(u engo.Updater) {
	world, ok := u.(*ecs.World)
	if !ok {
		panic("updater is no world")
	}
	common.SetBackground(color.Black)

	systems.SetupSnake()

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.Snake{Props: g.Properties})
	world.AddSystem(&systems.CookyEater{})
	world.AddSystem(&systems.CookySpawner{})
	world.AddSystem(&systems.Hud{})
}

func (g *Game) Type() string {
	return "gameScene"
}
