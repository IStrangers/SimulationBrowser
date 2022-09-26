package structs

import (
	browser "browser/structs"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	"image"
)

type BaseWidget struct {
	box            Box
	computedBox    Box
	widgetPosition WidgetPosition

	font *truetype.Font

	needsRepaint bool
	fixedWidth   bool
	fixedHeight  bool

	widgets []Widget

	backgroundColor string

	widgetType WidgetType
	cursor     *glfw.Cursor

	focusable  bool
	selectable bool

	focused  bool
	selected bool

	buffer *image.RGBA
	window *browser.Window
}

func (widget *BaseWidget) RequestReflow() {
	if widget.window != nil {
		widget.window.SetNeedsReflow(true)
	}
}

func (widget *BaseWidget) SetBuffer(buffer *image.RGBA) {
	widget.buffer = buffer
}
