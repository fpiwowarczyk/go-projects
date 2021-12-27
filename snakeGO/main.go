package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 1000.0
	height = 1000.0
)

var (
	columns   = 30
	rows      = 30
	seed      = time.Now().UnixNano()
	fps       = 3
	direction = 3
)

func init() {
	flag.IntVar(&columns, "columns", columns, "Sets the number of columns.")
	flag.IntVar(&rows, "rows", rows, "Sets the number of rows.")
	flag.Int64Var(&seed, "seed", seed, "Sets the starting seed of the game, used to randomize the initial state.")
	flag.IntVar(&fps, "fps", fps, "Sets the frames-per-second, used set the spped of the simulation.")
	flag.Parse()
}

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	cells := makeCells()
	t := time.Now()
	for !window.ShouldClose() {
		fmt.Print("Tick")
		tick(cells)
		if err := draw(program, window, cells); err != nil {
			panic(err)
		}

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
		t = time.Now()
	}
}

func tick(cells [][]*cell) {
	for x := range cells {
		for _, cell := range cells[x] {
			if cell.x > 0 && isBeforeHead(cell.x, cell.y, cells) {
				cells[cell.x-1][cell.y].head = false
				cell.head = true
			} else {
				cells[29][cell.y].head = true
				cell.head = true
			}

		}
	}
}

func draw(program uint32, window *glfw.Window, cells [][]*cell) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range cells {
		for _, c := range cells[x] {
			c.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
	return nil
}
