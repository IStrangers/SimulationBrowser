package structs

import renderer_structs "renderer/structs"

type InputWidget struct {
	BaseWidget

	value           string
	selected        bool
	active          bool
	padding         float64
	fontSize        float64
	context         *renderer_structs.Context
	fontColor       string
	cursorFloat     bool
	cursorPosition  int
	cursorDirection bool
	returnCallback  func()
}

func (button *InputWidget) draw() {

}
