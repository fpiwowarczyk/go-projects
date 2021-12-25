package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func handleKeyboard(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyW && action == glfw.Press {
		direction = 0
	} else if key == glfw.KeyD && action == glfw.Press {
		direction = 1
	} else if key == glfw.KeyS && action == glfw.Press {
		direction = 2
	} else if key == glfw.KeyA && action == glfw.Press {
		direction = 3
	}
}
