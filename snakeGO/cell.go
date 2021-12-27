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

	alive         bool
	nextState     bool
	foundNextHead bool
}

func (c *cell) checkState(cells [][]*cell) {

	c.alive = c.nextState
	c.nextState = c.alive

	if c.nextHeadPosition(cells) && !c.alive {
		c.nextState = true
		c.foundNextHead = true
	} else {
		c.nextState = false
	}

}

func (c *cell) nextHeadPosition(cells [][]*cell) bool {
	var liveCount int
	check := func(x, y int) {
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].alive {
			liveCount++
		}
	}

	switch direction {
	case 0:
		check(c.x+1, c.y)
	case 1:
		check(c.x-1, c.y)
	case 2:
		check(c.x, c.y+1)
	case 3:
		check(c.x, c.y-1)

	}
	return liveCount > 0
}

func (c *cell) draw() {
	if !c.alive {
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
			factor = float32(x) * (1.0 / float32(columns))
		case 1:
			size = 1.0 / float32(rows)
			factor = float32(y) * (1.0 / float32(rows))
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

		x: x,
		y: y,
	}
}

func makeCells() [][]*cell {

	cells := make([][]*cell, rows, columns)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)

			c.alive = false
			if x == 20 && y == 20 {
				c.alive = true
			}
			c.nextState = c.alive

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}
