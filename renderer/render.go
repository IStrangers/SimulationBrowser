package renderer

import (
	"SimulationBrowser/layout"
	renderer_structs "SimulationBrowser/renderer/structs"
)

func RenderDocument(context *renderer_structs.Context, document *renderer_structs.Document) error {

	html, _ := document.DOM.FindChildByName("html")
	html.RenderBox.Width = float64(context.Width())
	html.RenderBox.Height = float64(context.Height())

	body, _ := html.FindChildByName("body")

	context.SetRGB(1, 1, 1)
	context.Clear()

	layout.LayoutDOM(context, body, 0)
	return nil
}
