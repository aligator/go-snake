package entities

import (
	"bytes"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"image/color"
)

type Text struct {
	Object
	prefix string
	font   *common.Font
}

func (t *Text) Set(newText string) {
	t.Drawable = common.Text{
		Font: t.font,
		Text: t.prefix + newText,
	}
}

func PreloadText() {
	err := engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
	if err != nil {
		panic(err)
	}

	err = engo.Files.Load(snakeBodyTexture, snakeFrontTexture, snakeBackTexture)
	if err != nil {
		panic(err)
	}
}

func NewText(prefix string, initialText string, position engo.Point) *Text {
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: 24,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	text := &Text{
		Object: Object{
			BasicEntity: ecs.NewBasic(),
			RenderComponent: common.RenderComponent{
				Scale: engo.Point{X: 1, Y: 1},
				Drawable: common.Text{
					Font: fnt,
					Text: prefix + initialText,
				},
			},
			SpaceComponent: common.SpaceComponent{
				Position: position,
				Width:    200,
				Height:   200,
			},
		},
		prefix: prefix,
		font:   fnt,
	}

	text.SetShader(common.TextHUDShader)

	return text
}
