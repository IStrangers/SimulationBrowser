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
