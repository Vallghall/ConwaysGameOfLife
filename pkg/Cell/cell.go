package cell

import (
	"image/color"
)

type Cell struct {
	Alive bool
	Size  int
}

func (c *Cell) Color() color.Color {
	if c.Alive {
		return color.Black
	}

	return color.White
}

func New(size int) *Cell {
	cell := &Cell{
		Alive: false,
		Size:  size,
	}

	return cell
}
