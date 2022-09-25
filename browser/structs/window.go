package structs

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

type Window struct {
	title       string
	width       int
	height      int
	hiDPI       bool
	needsReflow bool
	visible     bool
	glw         *glfw.Window

	defaultCursor *glfw.Cursor
	pointerCursor *glfw.Cursor
}

func CreateWindow(title string, width int, height int, hiDPI bool) *Window {
	glw, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	glw.SetSizeLimits(300, 200, glfw.DontCare, glfw.DontCare)

	xScale, yScale := float32(1), float32(1)
	if hiDPI {
		xScale, yScale = glw.GetContentScale()
	}

	window := &Window{
		title:  title,
		width:  int(float32(width) / xScale),
		height: int(float32(height) / yScale),
		hiDPI:  hiDPI,
		glw:    glw,

		defaultCursor: glfw.CreateStandardCursor(glfw.ArrowCursor),
		pointerCursor: glfw.CreateStandardCursor(glfw.HandCursor),
	}

	window.RecreateContext()
	glw.MakeContextCurrent()

	return window
}

func (window *Window) Show() {
	window.needsReflow = true
	window.visible = true
	window.glw.Show()
}

func (window *Window) destroy() {
	//window.visible = false
	window.glw.Destroy()

	window.glw = nil
	//window.context = nil
	//window.backend = nil
	//window.frameBuffer = nil

	window.defaultCursor = nil
	window.pointerCursor = nil

	//window.registeredButtons = nil
	//window.registeredInputs = nil
	//window.activeInput = nil
	//
	//window.rootFrame = nil

	window = nil
}

func (window *Window) RecreateContext() {
}
