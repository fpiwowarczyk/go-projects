package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Border struct {
	*tl.Entity
	width, height int
	coords        map[Coord]int
}

func NewBorder(width, height int) *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)

	b.width, b.height = width-1, height-1

	b.coords = make(map[Coord]int)

	for x := 0; x < b.width; x++ {
		b.coords[Coord{x, 0}] = 1
		b.coords[Coord{x, b.height}] = 1
	}

	for y := 0; y < b.height+1; y++ {
		b.coords[Coord{0, y}] = 1
		b.coords[Coord{b.width, y}] = 1
	}

	return b
}

func (b *Border) Contains(coord Coord) bool {
	_, exists := b.coords[coord]
	return exists
}

func (b *Border) Draw(screen *tl.Screen) {
	if b == nil {
		return
	}

	for c := range b.coords {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}
}
