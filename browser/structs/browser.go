package structs

import (
	"filesystem"
	"fmt"
	"image"
	"renderer"
	renderer_structs "renderer/structs"
	"strings"
	ui_structs "ui/structs"
)

func loadDocument(webBrowser *WebBrowser, url string) {
	URL := filesystem.ParseURL(url)
	window := webBrowser.Window
	history := webBrowser.History
	currentDocument := webBrowser.CurrentDocument
	ui := webBrowser.UI
	headBar := ui.HeadBar
	urlInput := headBar.UrlInput

	if URL.Scheme == "" && URL.Host == "" {
		if !strings.HasPrefix(URL.Path, "/") {
			URL.Path = "/" + URL.Path
		}

		if strings.HasSuffix(currentDocument.URL.String(), "/") {
			URL.Path = strings.TrimSuffix(currentDocument.URL.Path, "/") + URL.Path
		}

		URL = filesystem.ParseURL(currentDocument.URL.Scheme + "://" + currentDocument.URL.Host + URL.Path)
	}

	resource := filesystem.GetResourceByURL(URL)
	rawDocument := string(resource.Body)

	switch strings.Split(resource.ContentType, ";")[0] {
	case "text/plain", "text/xml", "application/json":
		//ParsePlainText(rawDocument)
	default:
		webBrowser.CurrentDocument = renderer.ParseHTMLDocument(rawDocument)
	}
	currentDocument = webBrowser.CurrentDocument

	currentDocument.URL = resource.URL
	currentDocument.ContentType = resource.ContentType
	currentDocument.Title = currentDocument.GetDocumentTitle()

	urlInput.SetValue(webBrowser.CurrentDocument.URL.String())

	window.SetTitle(currentDocument.Title)

	window.RemoveStaticOverlay("debugOverlay")

	if history.PageCount() == 0 || history.Last().String() != resource.URL.String() {
		history.Push(resource.URL)
	}
}

func loadDocumentByUrl(webBrowser *WebBrowser) {
	ui := webBrowser.UI
	headBar := ui.HeadBar
	urlInput := headBar.UrlInput
	statusLabel := headBar.StatusLabel

	viewport := ui.Viewport
	view := viewport.View

	statusLabel.SetContent("Loading: " + urlInput.GetValue())
	statusLabel.RequestRepaint()

	loadDocument(webBrowser, urlInput.GetValue())

	view.SetOffset(0)
	view.SetDrawingRepaint(true)
	view.RequestRepaint()
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

	debugger := webBrowser.DebuggerMap[webBrowser.CurrentDocument]
	debugEl := debugger.SelectedElement
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
