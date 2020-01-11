package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"log"
)

type Object struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func loadTexture(path string) *common.Texture {
	texture, err := common.LoadedSprite(path)
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	return texture
}
