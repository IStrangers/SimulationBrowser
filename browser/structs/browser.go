package structs

import (
	"fmt"
	"image"
	renderer_structs "renderer/structs"
	ui_structs "ui/structs"
)

func loadDocumentByUrl(webBrowser *WebBrowser)  {

}

func treeNodeFromDOM(node *renderer_structs.NodeDOM) *ui_structs.TreeWidgetNode {
	nodeString := fmt.Sprintf(node.NodeName)
	xPath := node.GetXPath()
	treeNode := ui_structs.CreateTreeWidgetNode(nodeString, xPath)
	treeNode.Open()
	for _, childNode := range node.Children {
		treeNode.AddNode(treeNodeFromDOM(childNode))
	}
	return treeNode
}

func showDebugOverlay(webBrowser *WebBrowser) {
	webBrowser.Window.RemoveStaticOverlay("debugOverlay")

	debugEl := webBrowser.CurrentDocument.SelectedElement
	top, left, _, height := debugEl.RenderBox.GetRect()
	ctx := renderer_structs.CreateContext(int(webBrowser.CurrentDocument.DOM.RenderBox.Width), int(height+20))
	paintDebugRect(ctx, debugEl)

	overlay := ui_structs.CreateStaticOverlay("debugOverlay", ctx, image.Point{
		X: int(left), Y: int(top+webBrowser.UI.Viewport.View.GetTop()) + webBrowser.UI.Viewport.View.GetOffset(),
	})

	webBrowser.Window.AddStaticOverlay(overlay)
}

func paintDebugRect(ctx *renderer_structs.Context, node *renderer_structs.NodeDOM) {
	debugString := node.NodeName + " {" + fmt.Sprint(node.RenderBox.Top, node.RenderBox.Left, node.RenderBox.Width, node.RenderBox.Height) + "}"
	ctx.DrawRectangle(0, 0, node.RenderBox.Width, node.RenderBox.Height)
	ctx.SetRGBA(.2, .8, .4, .3)
	ctx.Fill()

	w, h := ctx.MeasureString(debugString)

	if node.RenderBox.Width < w {
		ctx.DrawRectangle(0, node.RenderBox.Height, w+4, h+4)
		ctx.SetRGB(1, 1, 0)
		ctx.Fill()

		ctx.SetRGB(0, 0, 0)
		ctx.DrawString(debugString, 2, node.RenderBox.Height+h)
		ctx.Fill()
	} else {
		ctx.DrawRectangle(node.RenderBox.Width-w-2, node.RenderBox.Height, w+4, h+4)
		ctx.SetRGB(1, 1, 0)
		ctx.Fill()

		ctx.SetRGB(0, 0, 0)
		ctx.DrawString(debugString, node.RenderBox.Width-w, node.RenderBox.Height+h)
		ctx.Fill()
	}
}