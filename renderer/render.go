package renderer

import (
	"layout"
	renderer_structs "renderer/structs"
)

func RenderDocument(context *renderer_structs.Context, document *renderer_structs.Document) error {
	body, _ := document.DOM.FindChildByName("body")

	document.DOM.RenderBox.Width = float64(context.Width())
	document.DOM.RenderBox.Height = float64(context.Height())

	context.SetRGB(1, 1, 1)
	context.Clear()

	layout.LayoutDOM(context, body, 0)
	return nil
}
