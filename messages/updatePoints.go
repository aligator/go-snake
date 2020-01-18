package messages

type UpdatePoints struct {
	points int
}

func (up UpdatePoints) Points() int {
	return up.points
}

func NewUpdatePoints(points int) UpdatePoints {
	return UpdatePoints{
		points: points,
	}
}

const HUDTextMessageType string = "HUDTextMessage"

// Type implements the engo.Message Interface
func (UpdatePoints) Type() string {
	return HUDTextMessageType
}
