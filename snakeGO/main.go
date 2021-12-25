package main

import (
	"flag"
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
	columns   = 100
	rows      = 100
	seed      = time.Now().UnixNano()
	threshold = 0.06
	fps       = 1
	direction = 4
)

func init() {
	flag.IntVar(&columns, "columns", columns, "Sets the number of columns.")
	flag.IntVar(&rows, "rows", rows, "Sets the number of rows.")
	flag.Int64Var(&seed, "seed", seed, "Sets the starting seed of the game, used to randomize the initial state.")
	flag.Float64Var(&threshold, "threshold", threshold, "A percentage between 0 and 1 used in conjunction with the -seed to deermine if a cell starts alive. For example, 0.15 means each cell has 15% chance of starting alvie.")
	flag.IntVar(&fps, "fps", fps, "Sets the frames-per-second, used set the spped of the simulation.")
	flag.Parse()
}

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	cells := makeCells(seed, threshold)
	t := time.Now()
	for !window.ShouldClose() {
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
		for _, c := range cells[x] {
			c.checkState(cells)
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
