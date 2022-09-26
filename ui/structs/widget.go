package structs

import (
	"image"
)

type Widget interface {
	Buffer() *image.RGBA
	SetBuffer(rgba *image.RGBA)
	SetNeedsRepaint(bool)
	NeedsRepaint() bool
	Widgets() []Widget
	ComputedBox() *Box
	SetWindow(window *Window)
	BaseWidget() *BaseWidget

	draw()
}
