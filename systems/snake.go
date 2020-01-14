package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"snake"
	"snake/entities"
)

type direction byte

const (
	// start with one so that 0 (at the beginning) does nothing
	LEFT direction = 1 + iota
	RIGHT
	UP
	DOWN
)

type Snake struct {
	Props         snake.Properties
	body          []*entities.SnakePart
	world         *ecs.World
	lastDirection direction
	lastMove      float32
	lost          bool
}

const leftButton = "left"
const rightButton = "right"
const upButton = "up"
const downButton = "down"

func SetupSnake() {
	engo.Input.RegisterButton(leftButton, engo.KeyArrowLeft)
	engo.Input.RegisterButton(rightButton, engo.KeyArrowRight)
	engo.Input.RegisterButton(upButton, engo.KeyArrowUp)
	engo.Input.RegisterButton(downButton, engo.KeyArrowDown)
}

func (s *Snake) Update(dt float32) {
	if s.lost {
		return
	}

	// update last direction if any button was pressed
	if engo.Input.Button(rightButton).JustPressed() {
		s.lastDirection = RIGHT
	}

	if engo.Input.Button(leftButton).JustPressed() {
		s.lastDirection = LEFT
	}

	if engo.Input.Button(upButton).JustPressed() {
		s.lastDirection = UP
	}

	if engo.Input.Button(downButton).JustPressed() {
		s.lastDirection = DOWN
	}

	if s.lastMove < 0.2 {
		s.lastMove = s.lastMove + dt
		return
	}
	s.lastMove = 0

	// get new head pos
	newPos := s.body[len(s.body)-1].Position
	switch s.lastDirection {
	case RIGHT:
		newPos.X = newPos.X + s.Props.TileSizeFloat32()
	case LEFT:
		newPos.X = newPos.X - s.Props.TileSizeFloat32()
	case UP:
		newPos.Y = newPos.Y - s.Props.TileSizeFloat32()
	case DOWN:
		newPos.Y = newPos.Y + s.Props.TileSizeFloat32()
	default:
		return
	}

	// check for collision on whole snake
	for _, part := range s.body {
		if part.Position == newPos {
			// todo: send lost message which prints "lost" to the screen by using a HUD-system
			s.lost = true
			return
		}
	}

	// check if snake comes out of the screen
	if newPos.X < 0 || newPos.Y < 0 || newPos.X >= engo.GameWidth() || newPos.Y >= engo.GameHeight() {
		// todo: send lost message which prints "lost" to the screen by using a HUD-system
		s.lost = true
		return
	}

	// add new head
	s.addParts([]*entities.SnakePart{entities.NewSnakeFront(newPos)})

	// transform former head to body part
	s.body[len(s.body)-2].TransformToBodyPart()

	// remove old tail
	// if old tail has the cooky, grow snake
	tail := s.body[0]
	if !tail.HasCooky {
		s.world.RemoveEntity(tail.BasicEntity)

		// transform last body part to tail
		s.body[0].TransformToTailPart()
	} else {
		tail.HasCooky = false
	}
}

func (s *Snake) Remove(e ecs.BasicEntity) {
	for i, part := range s.body {
		if part.ID() == e.ID() {
			s.body = append(s.body[:i], s.body[i+1:]...)
			break
		}
	}
}

// New implements the Initializer interface and
// will be called on creation automatically.
// It sets up the starting snake.
func (s *Snake) New(world *ecs.World) {
	s.world = world

	center := engo.Point{
		X: float32(s.Props.Width/2) * s.Props.TileSizeFloat32(),
		Y: float32(s.Props.Height/2) * s.Props.TileSizeFloat32(),
	}

	// set starting snake
	parts := []*entities.SnakePart{
		entities.NewSnakeBack(engo.Point{
			X: center.X - s.Props.TileSizeFloat32(),
			Y: center.Y,
		}),
		entities.NewSnakeBody(engo.Point{
			X: center.X,
			Y: center.Y,
		}),
		entities.NewSnakeFront(engo.Point{
			X: center.X + s.Props.TileSizeFloat32(),
			Y: center.Y,
		}),
	}

	s.addParts(parts)
}

func (s *Snake) addParts(parts []*entities.SnakePart) {
	for _, system := range s.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			// we assume there is only one render system
			// Otherwise the append has to be in an extra loop outside of the systems-loop
			for _, part := range parts {
				s.body = append(s.body, part)
				sys.Add(&part.BasicEntity, &part.RenderComponent, &part.SpaceComponent)
			}
		case *CookyEater:
			sys.SetCurrentHead(parts[len(parts)-1])
		}
	}
}
