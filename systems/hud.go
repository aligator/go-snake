package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"snake/entities"
	"snake/messages"
	"strconv"
)

type Hud struct {
	world      *ecs.World
	pointsText *entities.Text
}

func (h *Hud) Update(float32) {

}

func (h *Hud) Remove(ecs.BasicEntity) {

}

// New implements the Initializer interface and
// will be called on creation automatically.
// It injects the world into the system
func (h *Hud) New(world *ecs.World) {
	h.world = world

	h.pointsText = entities.NewText("Score: ", "0", engo.Point{1, 1})

	engo.Mailbox.Listen(messages.HUDTextMessageType, func(m engo.Message) {
		msg, ok := m.(messages.UpdatePoints)
		if !ok {
			return
		}

		h.pointsText.Set(strconv.Itoa(msg.Points()))
	})

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.pointsText.BasicEntity, &h.pointsText.RenderComponent, &h.pointsText.SpaceComponent)
		}
	}
}
