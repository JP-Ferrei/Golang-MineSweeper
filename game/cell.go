package game

import (
	"fmt"
)

type Cell struct {
	X           int
	Y           int
	hasMine     bool
	IsVisible   bool
	MinesAround int
	IsMarked    bool
}

func NewCell(x, y int, hasMine bool) *Cell {
	return &Cell{
		X:         x,
		Y:         y,
		hasMine:   hasMine,
		IsVisible: false,
	}
}

func (c *Cell) MarkCell() {
	c.IsMarked = !c.IsMarked
}

func (c *Cell) ShowCell() {
	c.IsVisible = true
}

func (c *Cell) String() string {
	return fmt.Sprintf("x: %d, y: %d, mine: %t, visible: %t, minesAround: %d, marked: %t",
		c.X, c.Y, c.hasMine, c.IsVisible, c.MinesAround, c.IsMarked)
}
