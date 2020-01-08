package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

const snakeBodyTexture = "textures/snake.png"
const snakeFrontTexture = "textures/snake_front.png"
const snakeBackTexture = "textures/snake_back.png"

type SnakePart struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func PreloadSnake() {
	err := engo.Files.Load(snakeBodyTexture, snakeFrontTexture, snakeBackTexture)
	if err != nil {
		panic(err)
	}
}

func loadTexture(path string) *common.Texture {
	texture, err := common.LoadedSprite(path)
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	return texture
}

func newSnakePart(pos engo.Point, texture *common.Texture) *SnakePart {
	return &SnakePart{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Scale:    engo.Point{X: 1, Y: 1},
			Drawable: texture,
		},
		SpaceComponent: common.SpaceComponent{
			Position: pos,
			Width:    16,
			Height:   16,
		},
	}
}

func (s *SnakePart) TransformToBodyPart() {
	s.Drawable = loadTexture(snakeBodyTexture)
}

func (s *SnakePart) TransformToTailPart() {
	s.Drawable = loadTexture(snakeBackTexture)
}

func NewSnakeFront(pos engo.Point) *SnakePart {
	return newSnakePart(pos, loadTexture(snakeFrontTexture))
}

func NewSnakeBody(pos engo.Point) *SnakePart {
	return newSnakePart(pos, loadTexture(snakeBodyTexture))
}

func NewSnakeBack(pos engo.Point) *SnakePart {
	return newSnakePart(pos, loadTexture(snakeBackTexture))
}
