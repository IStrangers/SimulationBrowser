package structs

import "image"

type ButtonWidget struct {
	BaseWidget
	content string

	icon      image.Image
	fontSize  float64
	fontColor string
	selected  bool
	padding   float64
	onClick   func()
}

func (button *ButtonWidget) draw() {

}
