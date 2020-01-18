package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"snake/entities"
	"snake/messages"
)

type CookyEater struct {
	world        *ecs.World
	currentCooky *entities.Cooky
	currentHead  *entities.SnakePart
	points       int
}

func (c *CookyEater) Update(float32) {
	if c.currentCooky == nil || c.currentHead == nil {
		return
	}

	if c.currentCooky.Position == c.currentHead.Position {
		c.points++
		c.world.RemoveEntity(c.currentCooky.BasicEntity)
		c.currentHead.HasCooky = true
		engo.Mailbox.Dispatch(messages.NewUpdatePoints(c.points))
	}
}

func (c *CookyEater) Remove(e ecs.BasicEntity) {
	if c.currentCooky != nil && c.currentCooky.ID() == e.ID() {
		c.currentCooky = nil
	}
	if c.currentHead != nil && c.currentHead.ID() == e.ID() {
		c.currentHead = nil
	}
}

// New implements the Initializer interface and
// will be called on creation automatically.
// It injects the world into the system
func (c *CookyEater) New(world *ecs.World) {
	c.world = world
}

func (c *CookyEater) SetCurrentCooky(cooky *entities.Cooky) {
	c.currentCooky = cooky
}

func (c *CookyEater) SetCurrentHead(head *entities.SnakePart) {
	c.currentHead = head
}
