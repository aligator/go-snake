package snake

type Properties struct {
	Width, Height, TileSize int
}

func (p Properties) TileSizeFloat32() float32 {
	return float32(p.TileSize)
}

func (p Properties) WidthFloat32() float32 {
	return float32(p.Width)
}

func (p Properties) HeightFloat32() float32 {
	return float32(p.Height)
}
