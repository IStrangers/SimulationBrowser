package structs

import "image"

type Overlay struct {
	ref string

	active bool

	top  float64
	left float64

	width  float64
	height float64

	position image.Point

	backgroundBuffer *image.RGBA
	buffer           *image.RGBA
}

func extractOverlay(buffer *image.RGBA, position image.Point) *Overlay {
	return &Overlay{
		ref:    "contextMenu",
		active: true,

		top:  float64(position.Y),
		left: float64(position.X),

		width:  float64(buffer.Rect.Max.X),
		height: float64(buffer.Rect.Max.Y),

		position: position,
		buffer:   buffer,
	}
}
