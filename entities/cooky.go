package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const cookyTexture = "textures/cooky.png"

type Cooky Object

func PreloadCooky() {
	err := engo.Files.Load(cookyTexture)
	if err != nil {
		panic(err)
	}
}

func NewCooky(pos engo.Point) *Cooky {
	return &Cooky{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Scale:    engo.Point{X: 1, Y: 1},
			Drawable: loadTexture(cookyTexture),
		},
		SpaceComponent: common.SpaceComponent{
			Position: pos,
			Width:    16,
			Height:   16,
		},
	}
}
