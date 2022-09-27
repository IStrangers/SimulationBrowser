package structs

import (
	"image"
	"image/draw"
)

type Widget interface {
	Buffer() *image.RGBA
	SetBuffer(rgba *image.RGBA)
	SetNeedsRepaint(bool)
	NeedsRepaint() bool
	Widgets() []Widget
	ComputedBox() *Box
	SetWindow(window *Window)
	GetBaseWidget() *BaseWidget

	draw()
}

func CopyWidgetToBuffer(widget Widget, src image.Image) {
	computedBox := widget.ComputedBox()
	top, left, width, height := int(computedBox.top), int(computedBox.left), int(computedBox.width), int(computedBox.height)

	buffer := widget.Buffer()
	if buffer == nil || buffer.Bounds().Max.X != width && buffer.Bounds().Max.Y != height {
		widget.GetBaseWidget().SetBuffer(image.NewRGBA(image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: width,
				Y: height,
			},
		}))
	}

	draw.Draw(widget.GetBaseWidget().buffer, image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: width, Y: height},
	}, src, image.Point{X: left, Y: top}, draw.Over)
}
