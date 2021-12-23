package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width = 500
	heigh = 500
)

func main() {
	runtime.LockOSThread()

	window := initGlwf()
	defer glfw.Terminate()
	for !window.ShouldClose() {
		//TODO
	}

}

// initGlfw initiializes glfw and returns a Window to use
func initGlwf() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, heigh, "SnakeGO", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	return window

}
