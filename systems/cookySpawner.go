package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"math/rand"
	"snake/entities"
	"time"
)

type CookySpawner struct {
	world *ecs.World
	cooky *entities.Cooky
}

func (c *CookySpawner) Update(dt float32) {

}

func (c *CookySpawner) Remove(e ecs.BasicEntity) {
	if c.cooky.ID() == e.ID() {
		c.cooky = nil
	}
}

// New implements the Initializer interface and
// will be called on creation automatically.
// It spawns the first cooky.
func (c *CookySpawner) New(world *ecs.World) {
	c.world = world
	rand.Seed(time.Now().UnixNano())
	c.setNewCooky()
}

func (c *CookySpawner) setNewCooky() {
	if c.cooky != nil {
		c.world.RemoveEntity(c.cooky.BasicEntity)
	}

	maxX := int(float64(engo.GameWidth()) / 16.0)
	maxY := int(float64(engo.GameHeight()) / 16.0)

	x := float32(rand.Intn(maxX))
	y := float32(rand.Intn(maxY))

	c.cooky = entities.NewCooky(engo.Point{X: x * 16, Y: y * 16})

	for _, system := range c.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&c.cooky.BasicEntity, &c.cooky.RenderComponent, &c.cooky.SpaceComponent)
		}
	}
}
