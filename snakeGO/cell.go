package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	squarePoints = []float32{
		-1, 1, 0,
		1, -1, 0,
		-1, -1, 0,

		-1, 1, 0,
		1, 1, 0,
		1, -1, 0,
	}

	squarePointsCount = int32(len(squarePoints) / 3)
)

type cell struct {
	drawable uint32

	x int
	y int

	head bool
	body bool
	back bool
}

func (cell *cell) updateState(cells [][]*cell) {
	if !cell.head && cell.x > 0 {
		if cells[cell.x-1][cell.y].head {
		}
	} else {
		if cells[len(cells)-1][cell.y].head {
		}
	}

}

func (c *cell) draw() {
	if !c.head {
		return
	}
	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, squarePointsCount)

}

func newCell(x, y int) *cell {
	points := make([]float32, len(squarePoints))
	copy(points, squarePoints)

	for i := 0; i < len(points); i++ {
		var factor float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(columns)
			factor = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			factor = float32(y) * size
		default:
			continue
		}
		if points[i] < 0 {
			points[i] = (factor * 2) - 1

		} else {
			points[i] = ((factor + size) * 2) - 1
		}

	}
	return &cell{
		drawable: makeVao(points),

		x:    x,
		y:    y,
		head: false,
		body: false,
		back: false,
	}
}

func makeCells() [][]*cell {

	cells := make([][]*cell, rows, columns)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)

			if c.x == 10 && c.y == 10 {
				c.head = true
			}
			if c.x == 9 && c.y == 10 {
				c.body = true
			}
			if c.x == 8 && c.y == 10 {
				c.back = true
			}

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}
